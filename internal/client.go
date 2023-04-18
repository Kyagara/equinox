package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/rate_limit"
	"go.uber.org/zap"
)

type InternalClient struct {
	key                string
	Cluster            api.Cluster
	http               *http.Client
	logger             *zap.Logger
	cache              *cache.Cache
	rateLimit          *rate_limit.RateLimit
	IsCacheEnabled     bool
	IsRateLimitEnabled bool
	IsRetryEnabled     bool
}

// Creates an EquinoxConfig for tests.
func NewTestEquinoxConfig() *api.EquinoxConfig {
	return &api.EquinoxConfig{
		Key:       "RGAPI-TEST",
		Cluster:   api.AmericasCluster,
		LogLevel:  api.DebugLevel,
		Timeout:   15,
		Retry:     false,
		Cache:     &cache.Cache{TTL: 0},
		RateLimit: &rate_limit.RateLimit{Enabled: false},
	}
}

// Returns a new InternalClient using the configuration provided.
func NewInternalClient(config *api.EquinoxConfig) (*InternalClient, error) {
	var cacheEnabled bool

	if config.Cache == nil {
		config.Cache = &cache.Cache{TTL: 0}
	}

	var rateEnabled bool

	if config.RateLimit == nil {
		config.RateLimit = &rate_limit.RateLimit{Enabled: false}
	}

	logger := NewLogger(config)

	if logger == nil {
		return nil, fmt.Errorf("error initializing logger")
	}

	rateEnabled = config.RateLimit.Enabled

	client := &InternalClient{
		key:                config.Key,
		Cluster:            config.Cluster,
		http:               &http.Client{Timeout: time.Duration(config.Timeout * int(time.Second))},
		logger:             logger,
		cache:              config.Cache,
		rateLimit:          config.RateLimit,
		IsCacheEnabled:     cacheEnabled,
		IsRateLimitEnabled: rateEnabled,
		IsRetryEnabled:     config.Retry,
	}

	if config.Cache.TTL > 0 {
		client.IsCacheEnabled = true
		client.cache = config.Cache
	}

	return client, nil
}

// Performs a GET request to the Riot API.
func (c *InternalClient) Get(route interface{}, endpointPath string, target interface{}, endpointName string, methodName string, authorizationHeader string) error {
	baseUrl := fmt.Sprintf(api.BaseURLFormat, route)

	url := fmt.Sprintf("%s%s", baseUrl, endpointPath)

	logger := c.logger.With(zap.String("httpMethod", http.MethodGet), zap.String("url", url))

	return c.get(logger, url, route, target, endpointName, methodName, authorizationHeader)
}

// Performs a GET request to the Data Dragon API.
func (c *InternalClient) DataDragonGet(endpointPath string, target interface{}, endpointName string, methodName string) error {
	url := fmt.Sprintf(api.DataDragonURLFormat, endpointPath)

	logger := c.logger.With(zap.String("httpMethod", http.MethodGet), zap.String("url", url))

	return c.get(logger, url, "", target, endpointName, methodName, "")
}

func (c *InternalClient) get(logger *zap.Logger, url string, route interface{}, target interface{}, endpointName string, methodName string, authorizationHeader string) error {
	// Creating a new HTTP Request
	req, err := c.newRequest(http.MethodGet, url, nil)

	if err != nil {
		return err
	}

	if authorizationHeader != "" {
		req.Header.Set("Authorization", authorizationHeader)
	}

	if c.IsCacheEnabled {
		item, err := c.cache.Get(url)

		// If there was an error with retrieving the cached response, only log the error
		if err != nil {
			logger.Error("Method failed", zap.Error(err))
		}

		if item != nil {
			logger.Debug("Cache hit")

			// Decoding the cached body into the target
			err = json.Unmarshal(item, &target)

			if err != nil {
				return err
			}

			return nil
		}
	}

	// Sending HTTP request and returning the response
	body, err := c.sendRequest(logger, req, url, route, endpointName, methodName, false)

	if err != nil {
		return err
	}

	if c.IsCacheEnabled {
		err := c.cache.Set(url, body)

		if err != nil {
			logger.Error("Error caching item", zap.Error(err))
		} else {
			logger.Debug("Cache set")
		}
	}

	// Decoding the body into the target
	err = json.Unmarshal(body, &target)

	if err != nil {
		return err
	}

	return nil
}

// Performs a POST request, authorizationHeader can be blank.
func (c *InternalClient) Post(route interface{}, endpointPath string, requestBody interface{}, target interface{}, endpointName string, methodName string, authorizationHeader string) error {
	baseUrl := fmt.Sprintf(api.BaseURLFormat, route)

	url := fmt.Sprintf("%s%s", baseUrl, endpointPath)

	logger := c.logger.With(zap.String("httpMethod", http.MethodPost), zap.String("url", url))

	// Creating a new HTTP Request
	req, err := c.newRequest(http.MethodPost, url, requestBody)

	if err != nil {
		return err
	}

	if authorizationHeader != "" {
		req.Header.Set("Authorization", authorizationHeader)
	}

	// Sending HTTP request and returning the response
	body, err := c.sendRequest(logger, req, url, route, endpointName, methodName, false)

	if err != nil {
		return err
	}

	// Decoding the body into the target
	err = json.Unmarshal(body, &target)

	if err != nil {
		return err
	}

	return nil
}

// Performs a PUT request.
func (c *InternalClient) Put(route interface{}, endpointPath string, requestBody interface{}, endpointName string, methodName string) error {
	baseUrl := fmt.Sprintf(api.BaseURLFormat, route)

	url := fmt.Sprintf("%s%s", baseUrl, endpointPath)

	logger := c.logger.With(zap.String("httpMethod", http.MethodPut), zap.String("url", url))

	// Creating a new HTTP Request
	req, err := c.newRequest(http.MethodPut, url, requestBody)

	if err != nil {
		return err
	}

	// Sending HTTP request and returning the response
	_, err = c.sendRequest(logger, req, url, route, endpointName, methodName, false)

	if err != nil {
		return err
	}

	return nil
}

// Creates a new HTTP Request and sets headers.
func (c *InternalClient) newRequest(httpMethod string, url string, body interface{}) (*http.Request, error) {
	var buffer io.ReadWriter

	if body != nil {
		buffer = &bytes.Buffer{}

		encoder := json.NewEncoder(buffer)

		err := encoder.Encode(body)

		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(httpMethod, url, buffer)

	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Riot-Token", c.key)
	req.Header.Set("User-Agent", "equinox")

	return req, nil
}

// Sends a HTTP request.
func (c *InternalClient) sendRequest(logger *zap.Logger, req *http.Request, url string, route interface{}, endpointName string, methodName string, retrying bool) ([]byte, error) {
	// If rate limiting is enabled
	if c.IsRateLimitEnabled {
		err := c.checkRateLimit(route, endpointName, methodName)

		if err != nil {
			return nil, err
		}
	}

	logger.Info("Sending request")

	// Sending request
	res, err := c.http.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Update rate limits
	if c.IsRateLimitEnabled {
		// Updating app rate limit
		err := c.rateLimit.SetAppRate(route, &res.Header)

		if err != nil {
			return nil, err
		}

		// Updating method rate limit
		err = c.rateLimit.Set(route, endpointName, methodName, &res.Header)

		if err != nil {
			return nil, err
		}
	}

	// Checking the response
	err = c.checkResponse(logger, url, res)

	// The body is defined here so if we retry the request we can later return the value
	// without having to read the body again, causing an error
	var body []byte

	// If retry is enabled and c.checkResponse() returns an error, retry the request
	if c.IsRetryEnabled && !retrying && errors.Is(err, api.ErrTooManyRequests) {
		logger.Debug("Retrying request")

		// If this retry is successful, the body var will be the res.Body
		body, err = c.sendRequest(logger, req, url, route, endpointName, methodName, true)
	}

	// Returns the error from c.checkResponse() if any
	// If retry is enabled, this error will be the error from the retried request if it failed again
	if err != nil {
		return nil, err
	}

	logger.Info("Request successful")

	// If the retry was successful, the body won't be nil, so return the result here to avoid reading the body again, causing an error
	if body != nil {
		return body, nil
	}

	body, err = io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	// In case of a post request returning just a single, non JSON response
	// This requires the endpoint method to handle the response as a api.PlainTextResponse and do type assertion
	if res.Header.Get("Content-Type") == "" {
		body = []byte(fmt.Sprintf(`{"response":"%s"}`, string(body)))

		return body, nil
	}

	return body, nil
}

func (c *InternalClient) checkResponse(logger *zap.Logger, url string, res *http.Response) error {
	// If the API returns a 429 code
	if c.IsRetryEnabled && res.StatusCode == http.StatusTooManyRequests {
		retryAfter := res.Header.Get("Retry-After")

		// If the header isn't found, don't retry and return error
		if retryAfter == "" {
			logger.Error("Request failed", zap.Error(api.ErrRetryAfterHeaderNotFound))
			return api.ErrRetryAfterHeaderNotFound
		}

		seconds, err := strconv.Atoi(retryAfter)

		if err != nil {
			logger.Error("Error converting Retry-After header", zap.Error(err))
			return err
		}

		logger.Warn(fmt.Sprintf("Too Many Requests, retrying request in %ds", seconds))

		time.Sleep(time.Duration(seconds) * time.Second)

		return api.ErrTooManyRequests
	}

	// If the status code is lower than 200 or higher than 299, return an error
	if res.StatusCode < http.StatusOK || res.StatusCode > 299 {
		logger.Error("Request failed", zap.Error(fmt.Errorf("endpoint method returned an error code: %v", res.Status)))

		// Handling errors documented in the Riot API docs
		// This StatusCodeToError solution is from https://github.com/KnutZuidema/golio
		err, ok := api.StatusCodeToError[res.StatusCode]

		if !ok {
			return api.ErrorResponse{
				Status: api.Status{
					Message:    "Unknown error",
					StatusCode: res.StatusCode,
				},
			}
		}

		return err
	}

	return nil
}

// Checks the app and method rate limit, returns true if rate limited.
func (c *InternalClient) checkRateLimit(route interface{}, endpointName string, methodName string) error {
	if route == "" {
		return nil
	}

	// Checking rate limits for the app
	rate, err := c.rateLimit.GetAppRate(route)

	if err != nil {
		return err
	}

	isRateLimited, err := c.rateLimit.IsRateLimited(rate)

	if err != nil {
		return err
	}

	if isRateLimited {
		return api.ErrTooManyRequests
	}

	// Checking rate limits for the endpoint method
	rate, err = c.rateLimit.Get(route, endpointName, methodName)

	if err != nil {
		return err
	}

	isRateLimited, err = c.rateLimit.IsRateLimited(rate)

	if err != nil {
		return err
	}

	if isRateLimited {
		return api.ErrTooManyRequests
	}

	return nil
}
