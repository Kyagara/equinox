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
	key            string
	http           *http.Client
	logger         *zap.Logger
	cache          *cache.Cache
	IsCacheEnabled bool
	IsRetryEnabled bool
}

// Creates an EquinoxConfig for tests.
func NewTestEquinoxConfig() *api.EquinoxConfig {
	return &api.EquinoxConfig{
		Key:      "RGAPI-TEST",
		LogLevel: api.DebugLevel,
		Timeout:  15,
		Retry:    false,
		Cache:    &cache.Cache{TTL: 0},
	}
}

// Returns a new InternalClient using the configuration provided.
func NewInternalClient(config *api.EquinoxConfig) (*InternalClient, error) {
	logger, err := NewLogger(config)
	if err != nil {
		return nil, fmt.Errorf("error initializing logger, %w", err)
	}
	if config.Cache == nil {
		config.Cache = &cache.Cache{TTL: 0}
	}
	client := &InternalClient{
		key:            config.Key,
		http:           &http.Client{Timeout: time.Duration(config.Timeout * int(time.Second))},
		logger:         logger,
		cache:          config.Cache,
		IsRetryEnabled: config.Retry,
	}
	client.IsCacheEnabled = config.Cache.TTL > 0
	return client, nil
}

// Creates a request to the provided route and url.
func (c *InternalClient) Request(base string, method string, route any, url string, body any) (*http.Request, error) {
	baseUrl := fmt.Sprintf(base, route)
	url = fmt.Sprintf("%s%s", baseUrl, url)
	var buffer io.ReadWriter
	if body != nil {
		buffer = &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		err := encoder.Encode(body)
		if err != nil {
			return nil, err
		}
	}
	request, err := http.NewRequest(method, url, buffer)
	if err != nil {
		return nil, err
	}
	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("X-Riot-Token", c.key)
	request.Header.Set("User-Agent", "equinox - https://github.com/Kyagara/equinox")
	return request, nil
}

// Performs a GET request to the Riot API.
func (c *InternalClient) Execute(request *http.Request, target any) error {
	url := request.URL.String()
	logger := c.logger.With(zap.String("httpMethod", request.Method), zap.String("url", url))
	if c.IsCacheEnabled {
		item, err := c.cache.Get(url)
		// If there was an error with retrieving the cached response, only log the error
		if err != nil {
			logger.Error("Error retrieving cached response", zap.Error(err))
		}
		if item != nil {
			logger.Debug("Cache hit")
			// Decoding the cached body into the target
			return json.Unmarshal(item, &target)
		}
	}
	// Sending HTTP request and returning the response
	body, err := c.sendRequest(logger, request, false)
	if err != nil {
		return err
	}
	if c.IsCacheEnabled {
		err := c.cache.Set(url, body)
		if err == nil {
			logger.Debug("Cache set")
		} else {
			logger.Error("Error caching item", zap.Error(err))
		}
	}
	// Decoding the body into the target
	return json.Unmarshal(body, &target)
}

// Sends a HTTP request.
func (c *InternalClient) sendRequest(logger *zap.Logger, request *http.Request, retrying bool) ([]byte, error) {
	logger.Info("Sending request")
	// Sending request
	response, err := c.http.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Checking the response
	err = c.checkResponse(logger, response)

	// The body is defined here so if we retry the request we can later return the value
	// without having to read the body again, causing an error
	var body []byte

	// If retry is enabled and c.checkResponse() returns an error, retry the request
	if c.IsRetryEnabled && !retrying && errors.Is(err, api.ErrTooManyRequests) {
		logger.Info("Retrying request")
		// If this retry is successful, the body var will be the response.Body
		body, err = c.sendRequest(logger, request, true)
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

	body, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// In case of a post request returning just a single, non JSON response
	// This requires the endpoint method to handle the response as a api.PlainTextResponse and do type requireion
	if response.Header.Get("Content-Type") == "" {
		body = []byte(fmt.Sprintf(`{"response":"%s"}`, string(body)))
		return body, nil
	}
	return body, nil
}

func (c *InternalClient) checkResponse(logger *zap.Logger, response *http.Response) error {
	// If the API returns a 429 code
	if response.StatusCode == http.StatusTooManyRequests {
		limit_type := response.Header.Get(api.RateLimitTypeHeader)

		if limit_type != "" {
			logger.Warn(fmt.Sprintf("Rate limited, type: %s", limit_type))
		} else {
			logger.Warn("Rate limited but no service was specified")
		}

		if c.IsRetryEnabled {
			retryAfter := response.Header.Get(api.RetryAfterHeader)

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

			logger.Warn(fmt.Sprintf("Retrying request in %ds", seconds))
			time.Sleep(time.Duration(seconds) * time.Second)
			return api.ErrTooManyRequests
		}
	}

	// If the status code is lower than 200 or higher than 299, return an error
	if response.StatusCode < http.StatusOK || response.StatusCode > 299 {
		logger.Error("Request failed", zap.Error(fmt.Errorf("endpoint method returned an error code: %v", response.Status)))
		// Handling errors documented in the Riot API docs
		// This StatusCodeToError solution is from https://github.com/KnutZuidema/golio
		err, ok := api.StatusCodeToError[response.StatusCode]
		if !ok {
			return api.ErrorResponse{
				Status: api.Status{
					Message:    "Unknown error",
					StatusCode: response.StatusCode,
				},
			}
		}
		return err
	}
	return nil
}

func (c *InternalClient) GetDDragonLOLVersions(client string, endpoint string, method string) ([]string, error) {
	logger := c.Logger(client, endpoint, method)
	logger.Debug("Method started execution")
	request, err := c.Request(api.DataDragonURLFormat, http.MethodGet, "", api.DataDragonLOLVersionURL, nil)
	if err != nil {
		logger.Error("Error creating request", zap.Error(err))
		return nil, err
	}
	var data []string
	err = c.Execute(request, &data)
	if err != nil {
		logger.Error("Error executing request", zap.Error(err))
		return nil, err
	}
	return data, nil
}
