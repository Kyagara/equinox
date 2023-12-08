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
	http           *http.Client
	loggers        *Loggers
	cache          *cache.Cache
	ratelimit      *ratelimit.RateLimit
	key            string
	maxRetries     int
	isCacheEnabled bool
}

var (
	ErrConfigurationNotProvided = errors.New("configuration not provided")
)

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

func NewInternalClient(config *api.EquinoxConfig) (*Client, error) {
	if config == nil {
		return nil, ErrConfigurationNotProvided
	}
	if config.Cache == nil {
		config.Cache = &cache.Cache{TTL: 0}
	}
	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{Timeout: 15 * time.Second}
	}
	client := &Client{
		key:  config.Key,
		http: config.HTTPClient,
		loggers: &Loggers{
			main:    NewLogger(config),
			methods: make(map[string]zerolog.Logger),
			mutex:   sync.Mutex{},
		},
		cache:          config.Cache,
		ratelimit:      &ratelimit.RateLimit{Region: make(map[any]*ratelimit.Limits)},
		maxRetries:     config.Retries,
		isCacheEnabled: config.Cache.TTL > 0,
	}
	apiHeaders.Set("X-Riot-Token", config.Key)
	return client, nil
}

func (c *Client) Request(ctx context.Context, logger zerolog.Logger, baseURL string, httpMethod string, route any, path string, methodID string, body any) (*api.EquinoxRequest, error) {
	if ctx == nil {
		return nil, errContextIsNil
	}

	equinoxReq := &api.EquinoxRequest{
		Logger:   logger,
		MethodID: methodID,
		Route:    route,
		BaseURL:  baseURL,
		Method:   httpMethod,
		Path:     path,
		Body:     body,
		Request:  nil,
		Retries:  0,
	}

	url := fmt.Sprintf(baseURL, route, path)

	var buffer io.Reader
	if equinoxReq.Body != nil {
		bodyBytes, err := jsonv2.Marshal(equinoxReq.Body)
		if err != nil {
			return nil, err
		}
		buffer = bytes.NewReader(bodyBytes)
	}

	request, err := http.NewRequestWithContext(ctx, httpMethod, url, buffer)
	if err != nil {
		return nil, err
	}
	equinoxReq.IsCDN = slices.Contains(cdns, request.URL.Host)

	if equinoxReq.IsCDN {
		request.Header = cdnHeaders.Clone()
	} else {
		if c.key == "" {
			return nil, errKeyNotProvided
		}
		request.Header = apiHeaders.Clone()
	}
	equinoxReq.Request = request
	return equinoxReq, nil
}

func (c *Client) Execute(ctx context.Context, equinoxReq *api.EquinoxRequest, target any) error {
	if ctx == nil {
		return errContextIsNil
	}

	url := getURLWithAuthorizationHash(equinoxReq.Request)

	if c.isCacheEnabled && equinoxReq.Method == http.MethodGet {
		if item, err := c.cache.Get(url); err != nil {
			equinoxReq.Logger.Error().Err(err).Msg("Error retrieving cached response")
		} else if item != nil {
			equinoxReq.Logger.Debug().Msg("Cache hit")
			return jsonv2.Unmarshal(item, &target)
		}
	}

	if !equinoxReq.IsCDN {
		err := c.ratelimit.Take(ctx, equinoxReq)
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
	if err != nil && !errors.Is(err, errServerError) {
		return err
	}

	if delay > 0 || errors.Is(err, errServerError) {
		equinoxReq.Logger.Info().Dur("sleep", delay).Msg("Retrying request after sleep")
		err := ratelimit.WaitN(ctx, time.Now().Add(delay), delay)
		if err != nil {
			return err
		}
		equinoxReq.Retries++
		return c.Execute(ctx, equinoxReq, target)
	}

	equinoxReq.Logger.Info().Msg("Request successful")
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if c.isCacheEnabled && equinoxReq.Method == http.MethodGet {
		if err := c.cache.Set(url, body); err != nil {
			equinoxReq.Logger.Error().Err(err).Msg("Error caching item")
		} else {
			equinoxReq.Logger.Debug().Msg("Cache set")
		}
	}
	return jsonv2.Unmarshal(body, &target)
}

func (c *Client) checkResponse(equinoxReq *api.EquinoxRequest, response *http.Response) (time.Duration, error) {
	if !equinoxReq.IsCDN {
		c.ratelimit.Update(equinoxReq, &response.Header)
	}

	// 2xx responses
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return 0, nil
	}

	// 4xx and 5xx responses will be retried
	if equinoxReq.Retries < c.maxRetries {
		if response.StatusCode == http.StatusTooManyRequests {
			equinoxReq.Logger.Warn().Msg("Received 429 response, checking Retry-After header")
			return c.ratelimit.CheckRetryAfter(equinoxReq, &response.Header)
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
func getURLWithAuthorizationHash(req *http.Request) string {
	url := req.URL.String()
	auth := req.Header.Get("Authorization")
	if auth == "" {
		return url
	}
	return fmt.Sprintf("%s%x", url, sha256.Sum256([]byte(auth)))
}
