package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/cache"
	"go.uber.org/zap"
)

type InternalClient struct {
	key    string
	http   *http.Client
	logger *zap.Logger
	// Used by endpoint methods to retrieve their respective logger
	endpointLoggers map[string]*zap.Logger
	cache           *cache.Cache
	IsCacheEnabled  bool
	IsRetryEnabled  bool
}

var (
	headers = http.Header{
		"X-Riot-Token": {""},
		"Accept":       {"application/json"},
		"User-Agent":   {"equinox - https://github.com/Kyagara/equinox"},
	}
	cdns = []string{"ddragon.leagueoflegends.com", "cdn.communitydragon.org"}
)

// Creates an EquinoxConfig for tests.
func NewTestEquinoxConfig() *api.EquinoxConfig {
	return &api.EquinoxConfig{
		Key:      "RGAPI-TEST",
		LogLevel: api.DEBUG_LOG_LEVEL,
		Timeout:  15,
		Retry:    false,
		Cache:    &cache.Cache{TTL: 0},
	}
}

// Returns a new InternalClient using the configuration provided.
func NewInternalClient(config *api.EquinoxConfig) (*InternalClient, error) {
	if config == nil {
		return nil, fmt.Errorf("equinox configuration not provided")
	}
	logger, err := NewLogger(config)
	if err != nil {
		return nil, err
	}
	if config.Cache == nil {
		config.Cache = &cache.Cache{TTL: 0}
	}
	client := &InternalClient{
		key:             config.Key,
		http:            &http.Client{Timeout: time.Second * time.Duration(config.Timeout)},
		logger:          logger,
		endpointLoggers: make(map[string]*zap.Logger),
		cache:           config.Cache,
		IsCacheEnabled:  config.Cache.TTL > 0,
		IsRetryEnabled:  config.Retry,
	}
	headers.Set("X-Riot-Token", config.Key)
	return client, nil
}

// Creates a request to the provided route and URL.
func (c *InternalClient) Request(base string, method string, route any, path string, body any) (*http.Request, error) {
	baseURL := fmt.Sprintf(base, route)
	url := fmt.Sprintf("%s%s", baseURL, path)
	var buffer io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buffer = bytes.NewReader(bodyBytes)
	}
	request, err := http.NewRequest(method, url, buffer)
	if err != nil {
		return nil, err
	}
	request.Header = headers
	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}
	for _, cdn := range cdns {
		if strings.Contains(request.URL.Host, cdn) {
			request.Header.Del("X-Riot-Token")
		}
	}
	return request, nil
}

// Performs a GET request to the Riot API.
func (c *InternalClient) Execute(logger *zap.Logger, request *http.Request, target any) error {
	url := request.URL.String()
	if c.IsCacheEnabled {
		if item, err := c.cache.Get(url); err != nil {
			logger.Error("Error retrieving cached response", zap.Error(err))
		} else if item != nil {
			logger.Debug("Cache hit")
			return json.Unmarshal(item, &target)
		}
	}
	body, err := c.sendRequest(logger, request, false)
	if err != nil {
		return err
	}
	if c.IsCacheEnabled {
		if err := c.cache.Set(url, body); err != nil {
			logger.Error("Error caching item", zap.Error(err))
		} else {
			logger.Debug("Cache set")
		}
	}
	return json.Unmarshal(body, &target)
}

// sendRequest sends an HTTP request and returns the response body as a byte array.
func (c *InternalClient) sendRequest(logger *zap.Logger, request *http.Request, retrying bool) ([]byte, error) {
	logger.Info("Sending request")
	response, err := c.http.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	err = c.checkResponse(logger, response)
	if err != nil && c.IsRetryEnabled && !retrying && errors.Is(err, api.ErrTooManyRequests) {
		return c.sendRequest(logger, request, true)
	} else if err != nil {
		return nil, err
	}
	logger.Info("Request successful")
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.Header.Get("Content-Type") == "" {
		jsonResponse := map[string]interface{}{"response": body}
		return json.Marshal(jsonResponse)
	}
	return body, nil
}

func (c *InternalClient) checkResponse(logger *zap.Logger, response *http.Response) error {
	if response.StatusCode == http.StatusTooManyRequests {
		limitType := response.Header.Get(api.X_RATE_LIMIT_TYPE_HEADER)
		if limitType != "" {
			logger.Warn("Rate limited, type:", zap.String("limit_type", limitType))
		} else {
			logger.Warn("Rate limited but no service was specified")
		}
		if c.IsRetryEnabled {
			retryAfter := response.Header.Get(api.RETRY_AFTER_HEADER)
			if retryAfter == "" {
				err := api.ErrRetryAfterHeaderNotFound
				logger.Error("Request failed", zap.Error(err))
				return err
			}
			// Convert the value of Retry-After header to seconds
			seconds, err := strconv.Atoi(retryAfter)
			if err != nil {
				logger.Error("Error converting Retry-After header", zap.Error(err))
				return err
			}
			logger.Warn("Retrying request in", zap.Duration("retry_after", time.Second*time.Duration(seconds)))
			// Sleep for the retry duration
			time.Sleep(time.Second * time.Duration(seconds))
			return api.ErrTooManyRequests
		}
	}
	// Check if the response status code is not within the range of 200-299 (success codes)
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

func (c *InternalClient) GetDDragonLOLVersions(id string) ([]string, error) {
	logger := c.Logger(id)
	logger.Debug("Method started execution")
	request, err := c.Request(api.D_DRAGON_BASE_URL_FORMAT, http.MethodGet, "", "/api/versions.json", nil)
	if err != nil {
		logger.Error("Error creating request", zap.Error(err))
		return nil, err
	}
	var data []string
	err = c.Execute(logger, request, &data)
	if err != nil {
		logger.Error("Error executing request", zap.Error(err))
		return nil, err
	}
	return data, nil
}
