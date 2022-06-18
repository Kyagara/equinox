package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/Kyagara/equinox/api"
	"go.uber.org/zap"
)

type InternalClient struct {
	Cluster    api.Cluster
	cache      *Cache
	rates      map[interface{}]*RateLimit
	defaultTTL int64
	key        string
	logLevel   api.LogLevel
	retry      bool
	http       *http.Client
	logger     *zap.SugaredLogger
	rateLimit  bool
}

// Creates an EquinoxConfig for tests.
func NewTestEquinoxConfig() *api.EquinoxConfig {
	return &api.EquinoxConfig{
		Key:       "RGAPI-KEY",
		Cluster:   api.AmericasCluster,
		LogLevel:  api.DebugLevel,
		Timeout:   10,
		TTL:       0,
		Retry:     false,
		RateLimit: false,
	}
}

// Returns a new InternalClient using configuration object provided.
func NewInternalClient(config *api.EquinoxConfig) *InternalClient {
	return &InternalClient{
		Cluster:    config.Cluster,
		cache:      NewCache(),
		rates:      map[interface{}]*RateLimit{},
		defaultTTL: config.TTL * int64(time.Second),
		key:        config.Key,
		logLevel:   config.LogLevel,
		retry:      config.Retry,
		http:       &http.Client{Timeout: time.Duration(config.Timeout * int(time.Second))},
		logger:     NewLogger(config),
		rateLimit:  config.RateLimit,
	}
}

func (c *InternalClient) ClearInternalClientCache() {
	if c.defaultTTL > 0 {
		c.cache.Clear()
	}
}

// Performs a GET request, authorizationHeader can be blank
func (c *InternalClient) Get(route interface{}, endpoint string, object interface{}, endpointName string, method string, authorizationHeader string) error {
	baseUrl := fmt.Sprintf(api.BaseURLFormat, route)

	// Creating a new HTTP Request.
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s%s", baseUrl, endpoint), nil)

	if err != nil {
		return err
	}

	if authorizationHeader != "" {
		req.Header.Set("Authorization", authorizationHeader)
	}

	// If caching is enabled
	if c.defaultTTL > 0 {
		cacheItem, err := c.cache.Get(req.URL.String())

		if err != nil {
			return err
		}

		if cacheItem != nil {
			logger := c.logger.With("httpMethod", http.MethodGet, "path", req.URL.Path)

			logger.Info("Cache hit")

			if err != nil {
				logger.Error(err)
				return err
			}

			// Decoding the cached body into the endpoint method response object.
			err = json.Unmarshal(cacheItem.response, &object)

			if err != nil {
				return err
			}

			return nil
		}
	}

	// Sending HTTP request and returning the response.
	_, body, err := c.sendRequest(req, false, endpointName, method, route)

	if err != nil {
		return err
	}

	if c.defaultTTL > 0 {
		c.cache.Set(req.URL.String(), body, c.defaultTTL)
	}

	// Decoding the body into the endpoint method response object.
	err = json.Unmarshal(body, &object)

	if err != nil {
		return err
	}

	return nil
}

// Performs a POST request, authorizationHeader can be blank
func (c *InternalClient) Post(route interface{}, endpoint string, requestBody interface{}, object interface{}, endpointName string, method string, authorizationHeader string) error {
	baseUrl := fmt.Sprintf(api.BaseURLFormat, route)

	// Creating a new HTTP Request.
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s%s", baseUrl, endpoint), requestBody)

	if err != nil {
		return err
	}

	if authorizationHeader != "" {
		req.Header.Set("Authorization", authorizationHeader)
	}

	// Sending HTTP request and returning the response.
	res, body, err := c.sendRequest(req, false, endpointName, method, route)

	if err != nil {
		return err
	}

	// In case of a post request returning just a single, non JSON response.
	// This requires the endpoint method to handle the response as a api.PlainTextResponse and do type assertion.
	// This implementation looks horrible, I don't know another way of decoding any non JSON value to the &object.
	if res.Header.Get("Content-Type") == "" {
		body := []byte(fmt.Sprintf(`{"response":"%s"}`, string(body)))

		err = json.Unmarshal(body, &object)

		if err != nil {
			return err
		}

		return nil
	}

	// Decoding the body into the endpoint method response object.
	err = json.Unmarshal(body, &object)

	if err != nil {
		return err
	}

	return nil
}

// Performs a PUT request
func (c *InternalClient) Put(route interface{}, endpoint string, requestBody interface{}, endpointName string, method string) error {
	baseUrl := fmt.Sprintf(api.BaseURLFormat, route)

	// Creating a new HTTP Request.
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("%s%s", baseUrl, endpoint), requestBody)

	if err != nil {
		return err
	}

	// Sending HTTP request and returning the response.
	_, _, err = c.sendRequest(req, false, endpointName, method, route)

	if err != nil {
		return err
	}

	return nil
}

// Sends a HTTP request.
func (c *InternalClient) sendRequest(req *http.Request, retried bool, endpoint string, method string, route interface{}) (*http.Response, []byte, error) {
	logger := c.logger.With("httpMethod", req.Method, "path", req.URL.Path)

	// If rate limiting is enabled
	if c.rateLimit {
		isRateLimited := c.checkRates(route, endpoint, method)

		if isRateLimited {
			return nil, nil, api.TooManyRequestsError
		}
	}

	logger.Debug("Making request")

	// Sending request.
	res, err := c.http.Do(req)

	if err != nil {
		logger.Error("Request failed")
		return nil, nil, err
	}

	defer res.Body.Close()

	// Update rate limits
	if c.rateLimit && res.Header.Get("X-App-Rate-Limit") != "" {
		// Updating app rate limit
		rate := ParseHeaders(res.Header, "X-App-Rate-Limit", "X-App-Rate-Limit-Count")

		c.rates[route].SetAppRate(rate)

		// Updating method rate limit
		rate = ParseHeaders(res.Header, "X-Method-Rate-Limit", "X-Method-Rate-Limit-Count")

		c.rates[route].Set(endpoint, method, rate)
	}

	// Checking the response
	err = c.checkResponse(res)

	// The body is defined here so if we retry the request we can later return the value
	// without having to read the body again, causing an error
	var body []byte

	// If retry is enabled and c.checkResponse() returns an api.RateLimitedError, retry the request
	if c.retry && errors.Is(err, api.TooManyRequestsError) && !retried {
		// If this retry is successful, the body var will be the res.Body
		res, body, err = c.sendRequest(req, true, endpoint, method, route)
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

	body, err = ioutil.ReadAll(res.Body)

	if err != nil {
		logger.Error(err)
		return nil, nil, err
	}

	return res, body, nil
}

// Creates a new HTTP Request and sets headers.
func (c *InternalClient) newRequest(method string, url string, body interface{}) (*http.Request, error) {
	var buffer io.ReadWriter

	if body != nil {
		buffer = &bytes.Buffer{}

		enc := json.NewEncoder(buffer)

		err := enc.Encode(body)

		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, buffer)

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

func (c *InternalClient) checkResponse(res *http.Response) error {
	logger := c.logger.With("httpMethod", res.Request.Method, "path", res.Request.URL.Path)

	// If the API returns a 429 code.
	if res.StatusCode == http.StatusTooManyRequests && c.retry {
		retryAfter := res.Header.Get("Retry-After")

		// If the header isn't found, don't retry and return error.
		if retryAfter == "" {
			return fmt.Errorf("rate limited but no Retry-After header was found, stopping")
		}

		seconds, _ := strconv.Atoi(retryAfter)

		logger.Warn(fmt.Sprintf("Too Many Requests, retrying request in %ds", seconds))

		time.Sleep(time.Duration(seconds) * time.Second)

		return api.TooManyRequestsError
	}

	// If the status code is lower than 200 or higher than 299, return an error.
	if res.StatusCode < http.StatusOK || res.StatusCode > 299 {
		logger.Errorf("Endpoint method returned an error response: %v", res.Status)

		// Handling errors documented in the Riot API docs
		// This StatusCodeToError solution is from KnutZuidema/golio
		// https://github.com/KnutZuidema/golio/blob/master/api/error.go
		// https://github.com/KnutZuidema/golio/blob/master/internal/client.go
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
func (c *InternalClient) checkRates(route interface{}, endpoint string, method string) bool {
	if c.rates[route] == nil {
		c.rates[route] = NewRateLimit()
	}

	// Checking rate limits for the app
	isRateLimited := c.rates[route].IsRateLimited(c.rates[route].appRate)

	if isRateLimited {
		return true
	}

	// Checking rate limits for the endpoint method
	rate := c.rates[route].Get(endpoint, method)

	if rate != nil {
		isRateLimited := c.rates[route].IsRateLimited(rate)

		if isRateLimited {
			return true
		}
	}

	return false
}
