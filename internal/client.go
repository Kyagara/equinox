package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
	"sync"
	"time"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/ratelimit"
	"go.uber.org/zap"
)

type InternalClient struct {
	key            string
	http           *http.Client
	loggers        *Loggers
	cache          *cache.Cache
	ratelimit      *ratelimit.RateLimit
	retry          int
	isCacheEnabled bool
}

var (
	staticHeaders = http.Header{
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

// Returns a new InternalClient using the configuration provided.
func NewInternalClient(config *api.EquinoxConfig) (*InternalClient, error) {
	if config == nil {
		return nil, fmt.Errorf("equinox configuration not provided")
	}
	logger, err := NewLogger(config)
	if err != nil {
		return nil, err
	}
	if config.Key == "" {
		logger.Warn("API key was not provided, requests using other clients will result in errors.")
	}
	if config.Cache == nil {
		config.Cache = &cache.Cache{TTL: 0}
	}
	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{Timeout: 15 * time.Second}
	}
	client := &InternalClient{
		key:  config.Key,
		http: config.HTTPClient,
		loggers: &Loggers{
			main:    logger,
			methods: make(map[string]*zap.Logger),
			mu:      sync.Mutex{},
		},
		cache:          config.Cache,
		ratelimit:      &ratelimit.RateLimit{Buckets: make(map[any]*ratelimit.Buckets)},
		retry:          config.Retry,
		isCacheEnabled: config.Cache.TTL > 0,
	}
	apiHeaders.Set("X-Riot-Token", config.Key)
	return client, nil
}

// Creates a request to the provided route and URL.
func (c *InternalClient) Request(logger *zap.Logger, baseURL string, httpMethod string, route any, path string, methodID string, body any) (*api.EquinoxRequest, error) {
	equinoxReq := &api.EquinoxRequest{
		Logger:   logger,
		MethodID: methodID,
		Route:    route,
		BaseURL:  baseURL,
		Method:   httpMethod,
		Path:     path,
		Body:     body,
		Request:  nil,
	}
	host := fmt.Sprintf(baseURL, route)
	url := fmt.Sprintf("%s%s", host, path)
	var buffer io.Reader
	if equinoxReq.Body != nil {
		bodyBytes, err := json.Marshal(equinoxReq.Body)
		if err != nil {
			return nil, err
		}
		buffer = bytes.NewReader(bodyBytes)
	}
	request, err := http.NewRequest(httpMethod, url, buffer)
	if err != nil {
		return nil, err
	}
	equinoxReq.Request = request
	if slices.Contains(cdns, request.URL.Host) {
		request.Header = staticHeaders
	} else {
		if c.key == "" {
			return nil, fmt.Errorf("api key not provided")
		}
		request.Header = apiHeaders
		// The request is going to the Riot API, so check the rate limit.
		err = c.ratelimit.Take(equinoxReq)
		if err != nil {
			return nil, err
		}
	}
	return equinoxReq, nil
}

// Performs a request to the Riot API.
func (c *InternalClient) Execute(equinoxReq *api.EquinoxRequest, target any) error {
	url := equinoxReq.Request.URL.String()
	if c.isCacheEnabled && equinoxReq.Method == http.MethodGet {
		if item, err := c.cache.Get(url); err != nil {
			equinoxReq.Logger.Error("Error retrieving cached response", zap.Error(err))
		} else if item != nil {
			equinoxReq.Logger.Debug("Cache hit")
			return json.Unmarshal(item, &target)
		}
	}
	body, err := c.sendRequest(equinoxReq, 0)
	if err != nil {
		return err
	}
	if c.isCacheEnabled && equinoxReq.Method == http.MethodGet {
		if err := c.cache.Set(url, body); err != nil {
			equinoxReq.Logger.Error("Error caching item", zap.Error(err))
		} else {
			equinoxReq.Logger.Debug("Cache set")
		}
	}
	return json.Unmarshal(body, &target)
}

// sendRequest sends an HTTP request and returns the response body as a byte array.
func (c *InternalClient) sendRequest(equinoxReq *api.EquinoxRequest, retriedCount int) ([]byte, error) {
	equinoxReq.Logger.Info("Sending request")
	response, err := c.http.Do(equinoxReq.Request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	err = c.checkResponse(equinoxReq, response)
	if err != nil && retriedCount < c.retry && errors.Is(err, api.ErrTooManyRequests) {
		return c.sendRequest(equinoxReq, retriedCount+1)
	} else if err != nil {
		return nil, err
	}
	equinoxReq.Logger.Info("Request successful")
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

func (c *InternalClient) checkResponse(equinoxReq *api.EquinoxRequest, response *http.Response) error {
	if !slices.Contains(cdns, response.Request.Host) {
		err := c.ratelimit.Update(equinoxReq, &response.Header)
		if err != nil {
			return err
		}
	}
	if response.StatusCode == http.StatusTooManyRequests {
		limitType := response.Header.Get(ratelimit.RATE_LIMIT_TYPE_HEADER)
		if limitType != "" {
			equinoxReq.Logger.Warn("Rate limited, type:", zap.String("limit_type", limitType))
		} else {
			equinoxReq.Logger.Warn("Rate limited but no service was specified")
		}
		if c.retry > 0 {
			retryAfter := response.Header.Get(ratelimit.RETRY_AFTER_HEADER)
			if retryAfter == "" {
				err := api.ErrRetryAfterHeaderNotFound
				equinoxReq.Logger.Error("Request failed", zap.Error(err))
				return err
			}
			// Convert the value of Retry-After header to seconds
			seconds, err := strconv.Atoi(retryAfter)
			if err != nil {
				equinoxReq.Logger.Error("Error converting Retry-After header", zap.Error(err))
				return err
			}
			equinoxReq.Logger.Warn("Retrying request in", zap.Duration("retry_after", time.Second*time.Duration(seconds)))
			// Sleep for the retry duration
			time.Sleep(time.Second * time.Duration(seconds))
			return api.ErrTooManyRequests
		}
	}
	// Check if the response status code is not within the range of 200-299 (success codes)
	if response.StatusCode < http.StatusOK || response.StatusCode > 299 {
		equinoxReq.Logger.Error("Request failed", zap.Error(fmt.Errorf("endpoint method returned an error code: %v", response.Status)))
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
	equinoxReq, err := c.Request(logger, api.D_DRAGON_BASE_URL_FORMAT, http.MethodGet, "", "/api/versions.json", "", nil)
	if err != nil {
		logger.Error("Error creating request", zap.Error(err))
		return nil, err
	}
	var data []string
	err = c.Execute(equinoxReq, &data)
	if err != nil {
		logger.Error("Error executing request", zap.Error(err))
		return nil, err
	}
	return data, nil
}
