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
	ErrMaxRetries     = errors.New("max retries reached")
	ErrContextIsNil   = errors.New("context must be non-nil")
	ErrKeyNotProvided = errors.New("api key not provided")
)

var (
	apiHeaders = http.Header{
		"X-Riot-Token": {""},
		"Accept":       {"application/json"},
		"Content-Type": {"application/json"},
	}
)

func NewInternalClient(config api.EquinoxConfig) (*Client, error) {
	if config.Key == "" {
		return nil, ErrKeyNotProvided
	}

	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{Timeout: 15 * time.Second}
	}
	if config.Cache == nil {
		config.Cache = &cache.Cache{}
	}
	if config.RateLimit == nil {
		config.RateLimit = &ratelimit.RateLimit{}
	}

	client := &Client{
		key:  config.Key,
		http: config.HTTPClient,
		loggers: Loggers{
			main:    NewLogger(config),
			methods: make(map[string]zerolog.Logger, 1),
			mutex:   sync.Mutex{},
		},
		cache:              config.Cache,
		ratelimit:          config.RateLimit,
		maxRetries:         config.Retry.MaxRetries,
		jitter:             config.Retry.Jitter,
		IsCacheEnabled:     config.Cache.TTL > 0,
		IsRateLimitEnabled: config.RateLimit.Enabled,
		IsRetryEnabled:     config.Retry.MaxRetries > 0,
	}

	apiHeaders.Set("X-Riot-Token", config.Key)
	return client, nil
}

// Creates a new 'EquinoxRequest' object for the 'Execute' and 'ExecuteBytes' methods.
func (c *Client) Request(ctx context.Context, logger zerolog.Logger, httpMethod string, urlComponents []string, methodID string, body any) (api.EquinoxRequest, error) {
	logger.Trace().Msg("Request")

	if ctx == nil {
		return api.EquinoxRequest{}, ErrContextIsNil
	}

	var bodyReader io.Reader
	if body != nil {
		json, err := jsonv2.Marshal(body)
		if err != nil {
			logger.Error().Err(err).Msg("Error marshalling body")
			return api.EquinoxRequest{}, err
		}
		bodyReader = bytes.NewReader(json)
	}

	url := strings.Join(urlComponents, "")

	request, err := http.NewRequestWithContext(ctx, httpMethod, url, bodyReader)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating HTTP request")
		return api.EquinoxRequest{}, err
	}

	equinoxReq := api.EquinoxRequest{
		Logger:   logger,
		MethodID: methodID,
		Route:    urlComponents[1],
		URL:      url,
		Request:  request,
	}

	request.Header = apiHeaders

	return equinoxReq, nil
}

// Executes a 'EquinoxRequest', checks cache and unmarshals the response into 'target'.
func (c *Client) Execute(ctx context.Context, equinoxReq api.EquinoxRequest, target any) error {
	equinoxReq.Logger.Trace().Msg("Execute")

	if ctx == nil {
		return ErrContextIsNil
	}

	url := GetURLWithAuthorizationHash(equinoxReq)

	if c.IsCacheEnabled && equinoxReq.Request.Method == http.MethodGet {
		item, err := c.cache.Get(ctx, url)
		if err != nil {
			equinoxReq.Logger.Error().Err(err).Msg("Error retrieving cached response")
		}

		if item != nil {
			equinoxReq.Logger.Debug().Msg("Cache hit")

			err := jsonv2.Unmarshal(item, target)
			if err != nil {
				equinoxReq.Logger.Error().Err(err).Msg("Error unmarshalling cached response")
				return err
			}

			return nil
		}
	}

	if c.IsRateLimitEnabled {
		err := c.ratelimit.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		if err != nil {
			return err
		}
	}

	response, err := c.Do(ctx, equinoxReq)
	if err != nil {
		equinoxReq.Logger.Error().Err(err).Msg("Do failed")
		return err
	}
	defer response.Body.Close()

	if !c.IsCacheEnabled {
		err := jsonv2.UnmarshalRead(response.Body, target)
		if err != nil {
			equinoxReq.Logger.Error().Err(err).Msg("Error unmarshalling response")
			return err
		}

		return nil
	}

	if equinoxReq.Request.Method == http.MethodGet {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			equinoxReq.Logger.Error().Err(err).Msg("Error reading response")
			return err
		}

		err = c.cache.Set(ctx, url, body)
		if err != nil {
			equinoxReq.Logger.Error().Err(err).Msg("Error caching item")
		} else {
			equinoxReq.Logger.Debug().Msg("Cache set")
		}

		err = jsonv2.Unmarshal(body, target)
		if err != nil {
			equinoxReq.Logger.Error().Err(err).Msg("Error unmarshalling body")
		}

		return nil
	}

	err = jsonv2.UnmarshalRead(response.Body, target)
	if err != nil {
		equinoxReq.Logger.Error().Err(err).Msg("Error unmarshalling response")
		return err
	}

	return nil
}

// Executes a 'EquinoxRequest', skips cache check and returns []byte.
func (c *Client) ExecuteBytes(ctx context.Context, equinoxReq api.EquinoxRequest) ([]byte, error) {
	equinoxReq.Logger.Trace().Msg("ExecuteBytes")

	if ctx == nil {
		return nil, ErrContextIsNil
	}

	if c.IsRateLimitEnabled {
		err := c.ratelimit.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		if err != nil {
			return nil, err
		}
	}

	response, err := c.Do(ctx, equinoxReq)
	if err != nil {
		equinoxReq.Logger.Error().Err(err).Msg("Do failed")
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		equinoxReq.Logger.Error().Err(err).Msg("Error reading response body")
		return nil, err
	}

	return body, nil
}

// Do sends the request and retry if necessary and enabled.
func (c *Client) Do(ctx context.Context, equinoxReq api.EquinoxRequest) (*http.Response, error) {
	equinoxReq.Logger.Trace().Msg("Do")

	var httpErr error

	// MaxRetries+1 to run this loop at least once.
	for i := 0; i < c.maxRetries+1; i++ {
		response, err := c.http.Do(equinoxReq.Request)
		if err != nil {
			// Stop retrying if the http.Client itself returns any error.
			return nil, err
		}

		delay, retryable, err := c.checkResponse(equinoxReq, response)
		if err == nil && delay == 0 {
			return response, nil
		}

		if !c.IsRetryEnabled || !retryable {
			return nil, err
		}

		if i < c.maxRetries {
			// Exponential backoff with jitter.
			sleep := delay*time.Duration(math.Pow(2, float64(i))) + c.jitter
			equinoxReq.Logger.Warn().Str("status_code", response.Status).Dur("sleep", sleep).Msg("Retrying request")
			err := ratelimit.WaitN(ctx, time.Now().Add(sleep), sleep)
			if err != nil {
				return nil, err
			}
		}

		httpErr = err
	}

	return nil, fmt.Errorf("%w: %w", ErrMaxRetries, httpErr)
}

func (c *Client) checkResponse(equinoxReq api.EquinoxRequest, response *http.Response) (time.Duration, bool, error) {
	if c.IsRateLimitEnabled {
		c.ratelimit.Update(equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, response.Header)
	}

	// 2xx responses.
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return 0, false, nil
	}

	err := api.StatusCodeToError(response.StatusCode)
	if err != nil {
		// 429 and 5xx responses will be retried.
		if response.StatusCode == http.StatusTooManyRequests || (response.StatusCode >= 500 && response.StatusCode < 600) {
			return c.ratelimit.CheckRetryAfter(equinoxReq.Route, equinoxReq.MethodID, response.Header), true, err
		}

		return 0, false, err
	}

	return 0, false, fmt.Errorf("unexpected status code: %d", response.StatusCode)
}

// Returns an URL with a hashed Authorization key if it exists.
func GetURLWithAuthorizationHash(req api.EquinoxRequest) string {
	auth := req.Request.Header.Get("Authorization")
	if auth == "" {
		return req.URL
	}

	hash := sha256.New()
	_, _ = hash.Write([]byte(auth))
	hashedAuth := hash.Sum(nil)

	return fmt.Sprintf("%s-%x", req.URL, hashedAuth)
}
