package internal

import (
	"encoding/json"
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
	Cluster  api.Cluster
	key      string
	logLevel api.LogLevel
	retry    bool
	http     *http.Client
	logger   *zap.SugaredLogger
}

// Creates an EquinoxConfig for tests.
func NewTestEquinoxConfig() *api.EquinoxConfig {
	return &api.EquinoxConfig{
		Key:      "RIOT_API_KEY",
		Cluster:  api.AmericasCluster,
		LogLevel: api.DebugLevel,
		Timeout:  10,
		Retry:    true,
	}
}

// Returns a new InternalClient using configuration object provided.
func NewInternalClient(config *api.EquinoxConfig) *InternalClient {
	return &InternalClient{
		Cluster:  config.Cluster,
		key:      config.Key,
		logLevel: config.LogLevel,
		retry:    config.Retry,
		http:     &http.Client{Timeout: config.Timeout * time.Second},
		logger:   NewLogger(config.Retry, config.Timeout, config.LogLevel),
	}
}

// Creates, sends and decodes a HTTP request.
func (c *InternalClient) Do(method string, route interface{}, endpoint string, requestBody io.Reader, object interface{}, authorizationHeader string) error {
	if route == "" {
		return fmt.Errorf("region is required")
	}

	baseUrl := fmt.Sprintf(api.BaseURLFormat, route)

	// Creating a new HTTP Request.
	req, err := c.newRequest(method, fmt.Sprintf("%s%s", baseUrl, endpoint), requestBody)

	if err != nil {
		return err
	}

	if authorizationHeader != "" {
		req.Header.Set("Authorization", authorizationHeader)
	}

	// Sending HTTP request and returning the response.
	res, err := c.sendRequest(req, 0)

	if err != nil {
		return err
	}

	// In case of a PUT request return nil.
	if res.Request.Method == http.MethodPut {
		return nil
	}

	// In case of a post request returning just a single, non JSON response.
	// This has a Post requirement because at the moment only one post request returns a plain text response.
	// This requires the endpoint method to handle the response as a api.PlainTextResponse and do type assertion.
	// This implementation looks horrible, I don't know another way of decoding any non JSON value to the &object.
	if res.Request.Method == http.MethodPost && res.Header.Get("Content-Type") == "" {
		value, err := ioutil.ReadAll(res.Body)

		if err != nil {
			return err
		}

		body := []byte(fmt.Sprintf(`{"response":"%s"}`, string(value)))

		err = json.Unmarshal(body, &object)

		if err != nil {
			return err
		}

		return nil
	}

	// Decoding the body into the endpoint method response object.
	if err := json.NewDecoder(res.Body).Decode(&object); err != nil {
		return err
	}

	return nil
}

// Sends a HTTP request.
func (c *InternalClient) sendRequest(req *http.Request, retryCount int8) (*http.Response, error) {
	logger := c.logger.With("httpMethod", req.Method, "path", req.URL.Path)

	if c.retry && retryCount > 1 {
		logger.Debug("Retried 2 times, stopping")

		return nil, fmt.Errorf("retried 2 times, stopping")
	}

	logger.Debug("Making request")

	// Sending request.
	res, err := c.http.Do(req)

	if err != nil {
		return nil, err
	}

	// Handling errors documented in the Riot API docs
	switch res.StatusCode {
	case http.StatusBadRequest:
		return nil, api.BadRequestError

	case http.StatusUnauthorized:
		return nil, api.UnauthorizedError

	case http.StatusForbidden:
		return nil, api.ForbiddenError

	case http.StatusNotFound:
		return nil, api.NotFoundError
	}

	// If the API returns a 429 code.
	if res.StatusCode == http.StatusTooManyRequests {
		logger.Debug("Too many requests")

		// If Retry is disabled just return an error.
		if !c.retry {
			return nil, api.RateLimitedError
		}

		retryAfter := res.Header.Get("Retry-After")

		// If the header isn't found, don't retry and return error.
		if retryAfter == "" {
			logger.Debug("Retry-After header not found, not retrying")

			return nil, fmt.Errorf("rate limited but no Retry-After header was found, stopping")
		}

		seconds, err := strconv.Atoi(retryAfter)

		if err != nil {
			return nil, err
		}

		logger.Debug(fmt.Sprintf("Too Many Requests, retrying request in %ds", seconds))

		time.Sleep(time.Duration(seconds) * time.Second)

		return c.sendRequest(req, retryCount+1)
	}

	// If the status code is lower than 200 or higher than 300, return an error.
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		logger.Error("Endpoint method returned an error response")

		err := api.ErrorResponse{
			Status: api.Status{
				Message:    "Unknown error",
				StatusCode: res.StatusCode,
			},
		}

		return nil, err
	}

	logger.Debug("Request successful")

	return res, nil
}

// Creates a new HTTP Request and sets headers.
func (c *InternalClient) newRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Riot-Token", c.key)

	return req, nil
}
