package internal

import (
	"bytes"
	"context"
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

	"github.com/Kyagara/equinox/v2/api"
	"github.com/Kyagara/equinox/v2/cache"
	"github.com/Kyagara/equinox/v2/ratelimit"
)

var (
	ErrMaxRetries     = errors.New("max retries reached")
	ErrContextIsNil   = errors.New("context must be non-nil")
	ErrKeyNotProvided = errors.New("api key not provided")
)

var (
	apiHeaders = http.Header{
		"X-Riot-Token": {""},
	}
)

type Client struct {
	http               *http.Client
	cache              *cache.Cache
	ratelimit          *ratelimit.RateLimit
	loggers            loggers
	key                string
	maxRetries         int
	jitter             time.Duration
	IsCacheEnabled     bool
	IsRateLimitEnabled bool
	IsRetryEnabled     bool
}

func NewInternalClient(config api.EquinoxConfig, h *http.Client, c *cache.Cache, r *ratelimit.RateLimit) (*Client, error) {
	if config.Key == "" {
		return nil, ErrKeyNotProvided
	}

	if h == nil {
		h = http.DefaultClient
	}
	if c == nil {
		c = &cache.Cache{TTL: 0}
	}
	if r == nil {
		r = &ratelimit.RateLimit{Enabled: false}
	}

	client := &Client{
		key:  config.Key,
		http: h,
		loggers: loggers{
			main:    NewLogger(config, c, r),
			methods: make(map[string]zerolog.Logger, 1),
			mutex:   sync.Mutex{},
		},
		cache:              c,
		ratelimit:          r,
		maxRetries:         config.Retry.MaxRetries,
		jitter:             config.Retry.Jitter,
		IsCacheEnabled:     c.TTL > 0,
		IsRateLimitEnabled: r.Enabled,
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

	request.Header = apiHeaders

	equinoxReq := api.EquinoxRequest{
		Logger:   logger,
		Request:  request,
		URL:      url,
		Route:    urlComponents[1],
		MethodID: methodID,
	}

	return equinoxReq, nil
}

// Executes a 'EquinoxRequest', checks cache and unmarshals the response into 'target'.
//
// ctx accepts 'api.ExecuteOptions', 'api.Revalidate' for example can be used to revalidate the cache, forcing an update to it.
func (c *Client) Execute(ctx context.Context, equinoxReq api.EquinoxRequest, target any) error {
	equinoxReq.Logger.Trace().Msg("Execute")

	if ctx == nil {
		return ErrContextIsNil
	}

	key, isRSO := cache.GetCacheKey(equinoxReq.URL, equinoxReq.Request.Header.Get("Authorization"))

	if c.IsCacheEnabled && equinoxReq.Request.Method == http.MethodGet {
		revalidate := ctx.Value(api.Revalidate)
		if revalidate == nil {
			item, err := c.cache.Get(ctx, key)
			if err != nil {
				equinoxReq.Logger.Error().Err(err).Msg("Error retrieving cached response")
				return err
			}

			if item != nil {
				equinoxReq.Logger.Debug().Str("route", equinoxReq.Route).Msg("Cache hit")

				// Only valid json is cached, so unmarshal shouldn't fail
				_ = jsonv2.Unmarshal(item, target)
				return nil
			}
		}
	}

	if c.IsRateLimitEnabled {
		err := c.ratelimit.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, isRSO)
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

	// There is only one endpoint method that is a Put and it also returns nothing.
	// For now it should be okay to just return.
	if equinoxReq.Request.Method == http.MethodPut {
		return nil
	}

	// If cache is not enabled, the method does not matter, unmarshal and return
	if !c.IsCacheEnabled {
		err := jsonv2.UnmarshalRead(response.Body, target)
		if err != nil {
			equinoxReq.Logger.Error().Err(err).Msg("Error unmarshalling response")
			return err
		}

		return nil
	}

	// Cache is enabled

	// Unmarshal and return any method that isn't Get
	if equinoxReq.Request.Method != http.MethodGet {
		err = jsonv2.UnmarshalRead(response.Body, target)
		if err != nil {
			equinoxReq.Logger.Error().Err(err).Msg("Error unmarshalling response")
			return err
		}

		return nil
	}

	// Method is a get, unmarshal and cache only valid json

	body, err := io.ReadAll(response.Body)
	if err != nil {
		equinoxReq.Logger.Error().Err(err).Msg("Error reading response")
		return err
	}

	err = jsonv2.Unmarshal(body, target)
	if err != nil {
		equinoxReq.Logger.Error().Err(err).Msg("Error unmarshalling body")
		return err
	}

	err = c.cache.Set(ctx, key, body)
	if err != nil {
		equinoxReq.Logger.Error().Err(err).Msg("Error caching item")
		return err
	}

	equinoxReq.Logger.Debug().Str("route", equinoxReq.Route).Msg("Cache set")
	return nil
}

// Executes a 'EquinoxRequest', skips checking cache and returns []byte.
func (c *Client) ExecuteBytes(ctx context.Context, equinoxReq api.EquinoxRequest) ([]byte, error) {
	equinoxReq.Logger.Trace().Msg("ExecuteBytes")

	if ctx == nil {
		return nil, ErrContextIsNil
	}

	if c.IsRateLimitEnabled {
		isRSO := equinoxReq.Request.Header.Get("Authorization") != ""
		err := c.ratelimit.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, isRSO)
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

// Sends the request using the internal http.Client, retries if enabled.
func (c *Client) Do(ctx context.Context, equinoxReq api.EquinoxRequest) (*http.Response, error) {
	equinoxReq.Logger.Trace().Msg("Do")

	var httpErr error

	// MaxRetries+1 to run this loop at least once
	for i := 0; i < c.maxRetries+1; i++ {
		response, err := c.http.Do(equinoxReq.Request)
		if err != nil {
			// Stop if the http.Client itself returns any error
			return nil, err
		}

		delay, retryable, err := c.checkResponse(ctx, equinoxReq, response)
		if err == nil && delay == 0 {
			equinoxReq.Logger.Trace().Str("route", equinoxReq.Route).Msg("Success")
			return response, nil
		}

		if !c.IsRetryEnabled || !retryable {
			return nil, err
		}

		if i < c.maxRetries {
			// Exponential backoff with jitter
			wait := delay*time.Duration(math.Pow(2, float64(i))) + c.jitter
			equinoxReq.Logger.Warn().Str("route", equinoxReq.Route).Str("status_code", response.Status).Dur("wait", wait).Int("retries", i).Msg("Retrying request")
			err := ratelimit.WaitN(ctx, time.Now().Add(wait), wait)
			if err != nil {
				return nil, err
			}
		}

		httpErr = err
	}

	return nil, fmt.Errorf("%w: %w", ErrMaxRetries, httpErr)
}

// Checks if the response contains a 'Retry-After' header and if it should be retried (StatusCode within range 429-599).
func (c *Client) checkResponse(ctx context.Context, equinoxReq api.EquinoxRequest, response *http.Response) (time.Duration, bool, error) {
	// Delay in milliseconds
	var retryAfter time.Duration

	if response.StatusCode == http.StatusTooManyRequests {
		retryAfter = ratelimit.GetRetryAfterHeader(ratelimit.RETRY_AFTER_HEADER)
	}

	if c.IsRateLimitEnabled {
		err := c.ratelimit.Update(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, response.Header, retryAfter)
		if err != nil {
			return 0, false, err
		}
	}

	// 2xx responses
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return 0, false, nil
	}

	err := api.StatusCodeToError(response.StatusCode)
	if err != nil {
		// 429 and 5xx responses will be retried
		if response.StatusCode == http.StatusTooManyRequests || (response.StatusCode > 499 && response.StatusCode < 600) {
			return retryAfter, true, err
		}

		return 0, false, err
	}

	return 0, false, fmt.Errorf("unexpected status code: %d", response.StatusCode)
}
