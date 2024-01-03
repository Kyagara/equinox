package internal

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"slices"
	"strings"
	"sync"
	"time"

	jsonv2 "github.com/go-json-experiment/json"
	"github.com/rs/zerolog"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/ratelimit"
)

type Client struct {
	http               *http.Client
	loggers            Loggers
	cache              *cache.Cache
	ratelimit          *ratelimit.RateLimit
	key                string
	maxRetries         int
	jitter             time.Duration
	IsCacheEnabled     bool
	IsRateLimitEnabled bool
	IsRetryEnabled     bool
}

var (
	ErrMaxRetries = errors.New("max retries reached")
)

var (
	errContextIsNil   = errors.New("context must be non-nil")
	errKeyNotProvided = errors.New("api key not provided")

	cdnHeaders = http.Header{
		"Accept":     {"application/json"},
		"User-Agent": {"equinox - https://github.com/Kyagara/equinox"},
	}
	apiHeaders = http.Header{
		"X-Riot-Token": {""},
		"Accept":       {"application/json"},
		"Content-Type": {"application/json"},
		"User-Agent":   {"equinox - https://github.com/Kyagara/equinox"},
	}
	cdns = []string{"ddragon.leagueoflegends.com", "cdn.communitydragon.org"}
)

func NewInternalClient(config api.EquinoxConfig) *Client {
	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{Timeout: 15 * time.Second}
	}
	if config.Cache == nil {
		config.Cache = &cache.Cache{}
	}
	if config.RateLimit == nil {
		config.RateLimit = &ratelimit.RateLimit{Enabled: false}
	}
	client := &Client{
		key:  config.Key,
		http: config.HTTPClient,
		loggers: Loggers{
			main:    NewLogger(config),
			methods: make(map[string]zerolog.Logger, 5),
			mutex:   sync.Mutex{},
		},
		cache:              config.Cache,
		ratelimit:          config.RateLimit,
		maxRetries:         config.Retry.MaxRetries,
		jitter:             config.Retry.Jitter,
		IsCacheEnabled:     config.Cache.TTL != 0,
		IsRateLimitEnabled: config.RateLimit.Enabled,
		IsRetryEnabled:     config.Retry.MaxRetries > 0,
	}
	if config.Retry.MaxRetries == 0 {
		config.Retry.MaxRetries = 1
	}
	apiHeaders.Set("X-Riot-Token", config.Key)
	return client
}

func (c *Client) Request(ctx context.Context, logger zerolog.Logger, httpMethod string, urlComponents []string, methodID string, body any) (api.EquinoxRequest, error) {
	logger.Debug().Msg("Creating request")

	if ctx == nil {
		return api.EquinoxRequest{}, errContextIsNil
	}

	var buffer io.Reader
	if body != nil {
		bodyBytes, err := jsonv2.Marshal(body)
		if err != nil {
			return api.EquinoxRequest{}, err
		}
		buffer = bytes.NewReader(bodyBytes)
	}

	url := strings.Join(urlComponents, "")

	request, err := http.NewRequestWithContext(ctx, httpMethod, url, buffer)
	if err != nil {
		return api.EquinoxRequest{}, err
	}

	equinoxReq := api.EquinoxRequest{
		Logger:   logger,
		MethodID: methodID,
		Route:    urlComponents[1],
		URL:      url,
		Request:  request,
		IsCDN:    slices.Contains(cdns, request.URL.Host),
	}

	if equinoxReq.IsCDN {
		request.Header = cdnHeaders
	} else {
		if c.key == "" {
			return api.EquinoxRequest{}, errKeyNotProvided
		}
		request.Header = apiHeaders
	}

	return equinoxReq, nil
}

func (c *Client) Execute(ctx context.Context, equinoxReq api.EquinoxRequest, target any) error {
	equinoxReq.Logger.Debug().Msg("Execute")

	if ctx == nil {
		return errContextIsNil
	}

	url, err := GetURLWithAuthorizationHash(equinoxReq)
	if err != nil {
		return err
	}

	if c.IsCacheEnabled && equinoxReq.Request.Method == http.MethodGet {
		if item, err := c.cache.Get(ctx, url); err != nil {
			equinoxReq.Logger.Error().Err(err).Msg("Error retrieving cached response")
		} else if item != nil {
			equinoxReq.Logger.Trace().Msg("Cache hit")
			return jsonv2.Unmarshal(item, target)
		}
	}

	if c.IsRateLimitEnabled && !equinoxReq.IsCDN {
		err := c.ratelimit.Take(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		if err != nil {
			return err
		}
	}

	response, err := c.Do(ctx, equinoxReq)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if !c.IsCacheEnabled {
		return jsonv2.UnmarshalRead(response.Body, target)
	}

	if equinoxReq.Request.Method == http.MethodGet {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		err = c.cache.Set(ctx, url, body)
		if err != nil {
			equinoxReq.Logger.Error().Err(err).Msg("Error caching item")
		} else {
			equinoxReq.Logger.Trace().Msg("Cache set")
		}

		return jsonv2.Unmarshal(body, target)
	}

	return jsonv2.UnmarshalRead(response.Body, target)
}

// ExecuteRaw executes a request without checking cache and returns []byte
func (c *Client) ExecuteRaw(ctx context.Context, equinoxReq api.EquinoxRequest) ([]byte, error) {
	equinoxReq.Logger.Debug().Msg("ExecuteRaw")

	if ctx == nil {
		return nil, errContextIsNil
	}

	if c.IsRateLimitEnabled && !equinoxReq.IsCDN {
		err := c.ratelimit.Take(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		if err != nil {
			return nil, err
		}
	}

	response, err := c.Do(ctx, equinoxReq)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return io.ReadAll(response.Body)
}

// Do sends the request n times from c.maxRetries and returns the response
//
// The loop will always run at least once
func (c *Client) Do(ctx context.Context, equinoxReq api.EquinoxRequest) (*http.Response, error) {
	equinoxReq.Logger.Debug().Msg("Sending request")
	for i := 0; i < c.maxRetries+1; i++ {
		response, err := c.http.Do(equinoxReq.Request)
		if err != nil {
			return nil, err
		}

		delay, retryable, err := c.checkResponse(equinoxReq, response)
		if err == nil && delay == 0 {
			equinoxReq.Logger.Debug().Msg("Request successful")
			return response, nil
		}

		if !c.IsRetryEnabled || !retryable {
			equinoxReq.Logger.Error().Err(err).Msg("Request failed")
			return nil, err
		}

		if i < c.maxRetries {
			sleep := delay*time.Duration(math.Pow(2, float64(i+1))) + c.jitter
			equinoxReq.Logger.Warn().Str("status_code", response.Status).Dur("sleep", sleep).Msg("Retrying request")
			err := ratelimit.WaitN(ctx, time.Now().Add(sleep), sleep)
			if err != nil {
				return nil, err
			}
		}
	}

	equinoxReq.Logger.Error().Err(ErrMaxRetries).Msg("Request failed")
	return nil, ErrMaxRetries
}

func (c *Client) checkResponse(equinoxReq api.EquinoxRequest, response *http.Response) (time.Duration, bool, error) {
	if c.IsRateLimitEnabled && !equinoxReq.IsCDN {
		c.ratelimit.Update(equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, response.Header)
	}

	// 2xx responses
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return 0, false, nil
	}

	err, ok := api.StatusCodeToError[response.StatusCode]
	if ok {
		// 429 and 5xx responses will be retried
		if response.StatusCode == http.StatusTooManyRequests || (response.StatusCode >= 500 && response.StatusCode < 600) {
			return c.ratelimit.CheckRetryAfter(equinoxReq.Route, equinoxReq.MethodID, response.Header), true, err
		}

		return 0, false, err
	}

	return 0, false, fmt.Errorf("unexpected status code: %d", response.StatusCode)
}

func (c *Client) GetDDragonLOLVersions(ctx context.Context, id string) ([]string, error) {
	logger := c.Logger(id)
	logger.Trace().Msg("Method started execution")
	urlComponents := []string{"https://", "", api.D_DRAGON_BASE_URL_FORMAT, "/api/versions.json"}
	equinoxReq, err := c.Request(ctx, logger, http.MethodGet, urlComponents, "", nil)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating request")
		return nil, err
	}
	var data []string
	err = c.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error().Err(err).Msg("Error executing request")
		return nil, err
	}
	return data, nil
}

// Generates an URL with the Authorization header if it exists. Don't want to store the Authorization header as key in plaintext.
func GetURLWithAuthorizationHash(req api.EquinoxRequest) (string, error) {
	auth := req.Request.Header.Get("Authorization")
	if auth == "" {
		return req.URL, nil
	}

	hash := sha256.New()
	_, err := hash.Write([]byte(auth))
	if err != nil {
		return req.URL, err
	}
	hashedAuth := hash.Sum(nil)

	return fmt.Sprintf("%s-%x", req.URL, hashedAuth), nil
}
