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
	"go.uber.org/zap"
)

type InternalClient struct {
	key                string
	Cluster            api.Cluster
	http               *http.Client
	logger             *zap.Logger
	logLevel           api.LogLevel
	cache              *cache.Cache
	isCacheEnabled     bool
	rateLimit          map[interface{}]*RateLimit
	isRateLimitEnabled bool
	isRetryEnabled     bool
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
		RateLimit: false,
	}
}

// Returns a new InternalClient using the configuration provided.
func NewInternalClient(config *api.EquinoxConfig) (*InternalClient, error) {
	var cacheEnabled bool

	if config.Cache == nil {
		config.Cache = &cache.Cache{TTL: 0}
	}

	client := &InternalClient{
		key:                config.Key,
		Cluster:            config.Cluster,
		http:               &http.Client{Timeout: time.Duration(config.Timeout * int(time.Second))},
		logger:             NewLogger(config),
		logLevel:           config.LogLevel,
		isCacheEnabled:     cacheEnabled,
		cache:              config.Cache,
		rateLimit:          map[interface{}]*RateLimit{},
		isRateLimitEnabled: config.RateLimit,
		isRetryEnabled:     config.Retry,
	}

	if config.Cache.TTL > 0 {
		client.isCacheEnabled = true
		client.cache = config.Cache
	}

	return client, nil
}

func (c *InternalClient) ClearInternalClientCache() error {
	if c.isCacheEnabled {
		err := c.cache.Clear()

		return err
	}

	return fmt.Errorf("cache is disabled")
}

// Performs a GET request to the Riot API
func (c *InternalClient) Get(route interface{}, endpointPath string, target interface{}, endpointName string, methodName string, authorizationHeader string) error {
	baseUrl := fmt.Sprintf(api.BaseURLFormat, route)

	url := fmt.Sprintf("%s%s", baseUrl, endpointPath)

	logger := c.logger.With(zap.String("httpMethod", http.MethodGet), zap.String("url", url))

	return c.get(logger, url, route, target, endpointName, methodName, authorizationHeader)
}

// Performs a GET request to the Data Dragon API
func (c *InternalClient) DataDragonGet(endpointPath string, target interface{}, endpointName string, methodName string) error {
	url := fmt.Sprintf(api.DataDragonURLFormat, endpointPath)

	logger := c.logger.With(zap.String("httpMethod", http.MethodGet), zap.String("url", url))

	return c.get(logger, url, "", target, endpointName, methodName, "")
}

func (c *InternalClient) get(logger *zap.Logger, url string, route interface{}, target interface{}, endpointName string, methodName string, authorizationHeader string) error {
	// Creating a new HTTP Request.
	req, err := c.newRequest(http.MethodGet, url, nil)

	if err != nil {
		return err
	}

	if authorizationHeader != "" {
		req.Header.Set("Authorization", authorizationHeader)
	}

	if c.isCacheEnabled {
		item, err := c.cache.Get(fmt.Sprintf("%s/%s", url, authorizationHeader))

		// If there was an error with retrieving the cached response, only log the error
		if err != nil {
			logger.Error("Method failed", zap.Error(err))
		}

		if item != nil {
			logger.Debug("Cache hit")

			// Decoding the cached body into the target.
			err = json.Unmarshal(item, &target)

			if err != nil {
				return err
			}

			return nil
		}
	}

	// Sending HTTP request and returning the response.
	_, body, err := c.sendRequest(logger, url, req, false, endpointName, methodName, route)

	if err != nil {
		return err
	}

	if c.isCacheEnabled {
		err := c.cache.Set(fmt.Sprintf("%s/%s", url, authorizationHeader), body)

		if err != nil {
			logger.Error("Method failed", zap.Error(err))
		} else {
			logger.Debug("Cache set")
		}
	}

	// Decoding the body into the target.
	err = json.Unmarshal(body, &target)

	if err != nil {
		return err
	}

	return nil
}

// Performs a POST request, authorizationHeader can be blank
func (c *InternalClient) Post(route interface{}, endpointPath string, requestBody interface{}, target interface{}, endpointName string, method string, authorizationHeader string) error {
	baseUrl := fmt.Sprintf(api.BaseURLFormat, route)

	url := fmt.Sprintf("%s%s", baseUrl, endpointPath)

	logger := c.logger.With(zap.String("httpMethod", http.MethodPost), zap.String("url", url))

	// Creating a new HTTP Request.
	req, err := c.newRequest(http.MethodPost, url, requestBody)

	if err != nil {
		return err
	}

	if authorizationHeader != "" {
		req.Header.Set("Authorization", authorizationHeader)
	}

	// Sending HTTP request and returning the response.
	res, body, err := c.sendRequest(logger, url, req, false, endpointName, method, route)

	if err != nil {
		return err
	}

	// In case of a post request returning just a single, non JSON response.
	// This requires the endpoint method to handle the response as a api.PlainTextResponse and do type assertion.
	// This implementation looks horrible, I don't know another way of decoding any non JSON value to the &target.
	if res.Header.Get("Content-Type") == "" {
		body := []byte(fmt.Sprintf(`{"response":"%s"}`, string(body)))

		err = json.Unmarshal(body, &target)

		if err != nil {
			return err
		}

		return nil
	}

	// Decoding the body into the target.
	err = json.Unmarshal(body, &target)

	if err != nil {
		return err
	}

	return nil
}

// Performs a PUT request
func (c *InternalClient) Put(route interface{}, endpointPath string, requestBody interface{}, endpointName string, methodName string) error {
	baseUrl := fmt.Sprintf(api.BaseURLFormat, route)

	url := fmt.Sprintf("%s%s", baseUrl, endpointPath)

	logger := c.logger.With(zap.String("httpMethod", http.MethodPut), zap.String("url", url))

	// Creating a new HTTP Request.
	req, err := c.newRequest(http.MethodPut, url, requestBody)

	if err != nil {
		return err
	}

	// Sending HTTP request and returning the response.
	_, _, err = c.sendRequest(logger, url, req, false, endpointName, methodName, route)

	if err != nil {
		return err
	}

	return nil
}

// Sends a HTTP request.
func (c *InternalClient) sendRequest(logger *zap.Logger, url string, req *http.Request, retried bool, endpointName string, methodName string, route interface{}) (*http.Response, []byte, error) {
	// If rate limiting is enabled
	if c.isRateLimitEnabled {
		isRateLimited := c.checkRateLimit(route, endpointName, methodName)

		if isRateLimited {
			return nil, nil, api.TooManyRequestsError
		}
	}

	logger.Info("Sending request")

	// Sending request.
	res, err := c.http.Do(req)

	if err != nil {
		return nil, nil, err
	}

	defer res.Body.Close()

	// Update rate limits
	if c.isRateLimitEnabled && res.Header.Get("X-App-Rate-Limit") != "" {
		// Updating app rate limit
		rate := ParseHeaders(res.Header, "X-App-Rate-Limit", "X-App-Rate-Limit-Count")

		c.rateLimit[route].SetAppRate(rate)

		// Updating method rate limit
		rate = ParseHeaders(res.Header, "X-Method-Rate-Limit", "X-Method-Rate-Limit-Count")

		c.rateLimit[route].Set(endpointName, methodName, rate)
	}

	// Checking the response
	err = c.checkResponse(logger, url, res)

	// The body is defined here so if we retry the request we can later return the value
	// without having to read the body again, causing an error
	var body []byte

	// If retry is enabled and c.checkResponse() returns an api.RateLimitedError, retry the request
	if c.isRetryEnabled && errors.Is(err, api.TooManyRequestsError) && !retried {
		// If this retry is successful, the body var will be the res.Body
		res, body, err = c.sendRequest(logger, url, req, true, endpointName, methodName, route)
	}

	// Returns the error from c.checkResponse() if any
	// If retry is enabled, this error could also be the error from the retried request if it failed again
	if err != nil {
		return nil, nil, err
	}

	// If the retry was successful, the body won't be nil, so return the result here to avoid reading the body again
	if body != nil {
		return res, body, nil
	}

	logger.Info("Request successful")

	body, err = io.ReadAll(res.Body)

	if err != nil {
		return nil, nil, err
	}

	return res, body, nil
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

func (c *InternalClient) checkResponse(logger *zap.Logger, url string, res *http.Response) error {
	// If the API returns a 429 code.
	if res.StatusCode == http.StatusTooManyRequests && c.isRetryEnabled {
		retryAfter := res.Header.Get("Retry-After")

		// If the header isn't found, don't retry and return error.
		if retryAfter == "" {
			return fmt.Errorf("rate limited but no Retry-After header was found, stopping")
		}

		seconds, err := strconv.Atoi(retryAfter)

		if err != nil {
			logger.Error("Error converting retry after header", zap.Error(err))
			return err
		}

		logger.Warn(fmt.Sprintf("Too Many Requests, retrying request in %ds", seconds))

		time.Sleep(time.Duration(seconds) * time.Second)

		return api.TooManyRequestsError
	}

	// If the status code is lower than 200 or higher than 299, return an error.
	if res.StatusCode < http.StatusOK || res.StatusCode > 299 {
		logger.Error("Request failed", zap.Error(fmt.Errorf("endpoint method returned an error code: %v", res.Status)))

		// Handling errors documented in the Riot API docs
		// This StatusCodeToError solution is from KnutZuidema/golio
		err, ok := api.StatusCodeToError[res.StatusCode]

		if !ok {
			err = api.ErrorResponse{
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

// Checks the app and method rate limit, returns true if rate limited
func (c *InternalClient) checkRateLimit(route interface{}, endpointName string, methodName string) bool {
	if route == "" {
		return false
	}

	if c.rateLimit[route] == nil {
		c.rateLimit[route] = NewRateLimit()
	}

	// Checking rate limits for the app
	isRateLimited := c.rateLimit[route].IsRateLimited(c.rateLimit[route].appRate)

	if isRateLimited {
		return true
	}

	// Checking rate limits for the endpoint method
	rate := c.rateLimit[route].Get(endpointName, methodName)

	if rate != nil {
		isRateLimited := c.rateLimit[route].IsRateLimited(rate)

		if isRateLimited {
			return true
		}
	}

	return false
}
