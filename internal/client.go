package internal

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"net/http"
	"slices"
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
	isCacheEnabled     bool
	isRateLimitEnabled bool
}

var (
	errContextIsNil   = errors.New("context must be non-nil")
	errKeyNotProvided = errors.New("api key not provided")
	errServerError    = errors.New("server error")

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
			methods: make(map[string]zerolog.Logger),
			mutex:   sync.Mutex{},
		},
		cache:              config.Cache,
		ratelimit:          config.RateLimit,
		maxRetries:         config.Retries,
		isCacheEnabled:     config.Cache.TTL != 0,
		isRateLimitEnabled: config.RateLimit.Enabled,
	}
	apiHeaders.Set("X-Riot-Token", config.Key)
	return client
}

func (c *Client) Request(ctx context.Context, logger zerolog.Logger, baseURL string, httpMethod string, route any, path string, methodID string, body any) (api.EquinoxRequest, error) {
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

	url := fmt.Sprintf(baseURL, route, path)

	request, err := http.NewRequestWithContext(ctx, httpMethod, url, buffer)
	if err != nil {
		return api.EquinoxRequest{}, err
	}

	equinoxReq := api.EquinoxRequest{
		Logger:   logger,
		MethodID: methodID,
		Route:    route,
		URL:      url,
		Request:  request,
		Retries:  0,
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
	if ctx == nil {
		return errContextIsNil
	}

	url := GetURLWithAuthorizationHash(equinoxReq)

	if c.isCacheEnabled && equinoxReq.Request.Method == http.MethodGet {
		if item, err := c.cache.Get(ctx, url); err != nil {
			equinoxReq.Logger.Error().Err(err).Msg("Error retrieving cached response")
		} else if item != nil {
			equinoxReq.Logger.Debug().Msg("Cache hit")
			return jsonv2.Unmarshal(item, target)
		}
	}

	if c.isRateLimitEnabled && !equinoxReq.IsCDN {
		err := c.ratelimit.Take(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		if err != nil {
			return err
		}
	}

	equinoxReq.Logger.Info().Msg("Sending request")
	response, err := c.http.Do(equinoxReq.Request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	delay, err := c.checkResponse(equinoxReq, response)
	if errors.Is(err, errServerError) {
		equinoxReq.Retries++
		return c.Execute(ctx, equinoxReq, target)
	}

	if err != nil {
		return err
	}

	if delay > 0 {
		equinoxReq.Logger.Info().Dur("sleep", delay).Msg("Retrying request after sleep")
		err := ratelimit.WaitN(ctx, time.Now().Add(delay), delay)
		if err != nil {
			return err
		}
		equinoxReq.Retries++
		return c.Execute(ctx, equinoxReq, target)
	}

	equinoxReq.Logger.Info().Msg("Request successful")
	if !c.isCacheEnabled {
		return jsonv2.UnmarshalRead(response.Body, target)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if equinoxReq.Request.Method == http.MethodGet {
		if err := c.cache.Set(ctx, url, body); err != nil {
			equinoxReq.Logger.Error().Err(err).Msg("Error caching item")
		} else {
			equinoxReq.Logger.Debug().Msg("Cache set")
		}
	}
	return jsonv2.Unmarshal(body, target)
}

// ExecuteRaw executes a request without checking cache and returns []byte
func (c *Client) ExecuteRaw(ctx context.Context, equinoxReq api.EquinoxRequest) ([]byte, error) {
	if ctx == nil {
		return nil, errContextIsNil
	}

	if c.isRateLimitEnabled && !equinoxReq.IsCDN {
		err := c.ratelimit.Take(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		if err != nil {
			return nil, err
		}
	}

	equinoxReq.Logger.Info().Msg("Sending request")
	response, err := c.http.Do(equinoxReq.Request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	delay, err := c.checkResponse(equinoxReq, response)
	if errors.Is(err, errServerError) {
		equinoxReq.Retries++
		return c.ExecuteRaw(ctx, equinoxReq)
	}

	if err != nil {
		return nil, err
	}

	if delay > 0 {
		equinoxReq.Logger.Info().Dur("sleep", delay).Msg("Retrying request after sleep")
		err := ratelimit.WaitN(ctx, time.Now().Add(delay), delay)
		if err != nil {
			return nil, err
		}
		equinoxReq.Retries++
		return c.ExecuteRaw(ctx, equinoxReq)
	}

	equinoxReq.Logger.Info().Msg("Request successful")
	return io.ReadAll(response.Body)
}

func (c *Client) checkResponse(equinoxReq api.EquinoxRequest, response *http.Response) (time.Duration, error) {
	if c.isRateLimitEnabled && !equinoxReq.IsCDN {
		c.ratelimit.Update(equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, response.Header)
	}

	// 2xx responses
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return 0, nil
	}

	// 4xx and 5xx responses will be retried
	if equinoxReq.Retries < c.maxRetries {
		if response.StatusCode == http.StatusTooManyRequests {
			equinoxReq.Logger.Warn().Msg("Received 429 response")
			return c.ratelimit.CheckRetryAfter(equinoxReq.Route, equinoxReq.MethodID, response.Header), nil
		}

		if response.StatusCode >= 500 && response.StatusCode < 600 {
			equinoxReq.Logger.Warn().Str("status_code", response.Status).Msg("Retrying request")
			return 0, errServerError
		}
	}

	if err, ok := api.StatusCodeToError[response.StatusCode]; ok {
		return 0, err
	}

	return 0, api.HTTPErrorResponse{
		Status: api.Status{
			Message:    "Unknown error",
			StatusCode: response.StatusCode,
		},
	}
}

func (c *Client) GetDDragonLOLVersions(ctx context.Context, id string) ([]string, error) {
	logger := c.Logger(id)
	logger.Debug().Msg("Method started execution")
	equinoxReq, err := c.Request(ctx, logger, api.D_DRAGON_BASE_URL_FORMAT, http.MethodGet, "", "/api/versions.json", "", nil)
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
func GetURLWithAuthorizationHash(req api.EquinoxRequest) string {
	auth := req.Request.Header.Get("Authorization")
	if auth == "" {
		return req.URL
	}

	hash := sha256.New()
	hash.Write([]byte(auth))
	hashedAuth := hash.Sum(nil)

	return fmt.Sprintf("%s-%x", req.URL, hashedAuth)
}
