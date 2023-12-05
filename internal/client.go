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
	maxRetries     int
	isCacheEnabled bool
}

var (
	errContextIsNil   = errors.New("context must be non-nil")
	errKeyNotProvided = errors.New("api key not provided")

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
			mutex:   sync.Mutex{},
		},
		cache:          config.Cache,
		ratelimit:      &ratelimit.RateLimit{Region: make(map[any]*ratelimit.Limits)},
		maxRetries:     config.Retry,
		isCacheEnabled: config.Cache.TTL > 0,
	}
	apiHeaders.Set("X-Riot-Token", config.Key)
	return client, nil
}

// Creates a request to the provided route and URL.
func (c *InternalClient) Request(ctx context.Context, logger *zap.Logger, baseURL string, httpMethod string, route any, path string, methodID string, body any) (*api.EquinoxRequest, error) {
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

	host := fmt.Sprintf(baseURL, route)
	url := fmt.Sprintf("%s%s", host, path)

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
	equinoxReq.Request = request
	equinoxReq.IsCDN = slices.Contains(cdns, request.URL.Host)

	if equinoxReq.IsCDN {
		request.Header = cdnHeaders
	} else {
		if c.key == "" {
			return nil, errKeyNotProvided
		}
		request.Header = apiHeaders
	}
	return equinoxReq, nil
}

// Performs a request to the Riot API.
func (c *InternalClient) Execute(ctx context.Context, equinoxReq *api.EquinoxRequest, target any) error {
	if ctx == nil {
		return errContextIsNil
	}

	url := equinoxReq.Request.URL.String() + getAuthorizationHeaderHash(equinoxReq.Request.Header.Get("Authorization"))

	if c.isCacheEnabled && equinoxReq.Method == http.MethodGet {
		if item, err := c.cache.Get(url); err != nil {
			equinoxReq.Logger.Error("Error retrieving cached response", zap.Error(err))
		} else if item != nil {
			equinoxReq.Logger.Debug("Cache hit")
			return jsonv2.Unmarshal(item, &target)
		}
	}

	// Request not cached/cache disabled, so take from the bucket
	if !equinoxReq.IsCDN {
		err := c.ratelimit.Take(ctx, equinoxReq)
		if err != nil {
			return err
		}
	}

	equinoxReq.Logger.Info("Sending request")
	response, err := c.http.Do(equinoxReq.Request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	delay, err := c.checkResponse(equinoxReq, response)
	if err != nil {
		equinoxReq.Logger.Error("Request failed", zap.Error(err))
		return err
	}

	if delay > 0 {
		deadline, ok := ctx.Deadline()
		if ok && deadline.Before(time.Now().Add(delay)) {
			return ratelimit.ErrContextDeadlineExceeded
		}
		equinoxReq.Logger.Info("Retrying request after sleep", zap.Duration("sleep", delay))
		select {
		case <-time.After(delay):
			equinoxReq.Retries++
			return c.Execute(ctx, equinoxReq, target)
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	equinoxReq.Logger.Info("Request successful")
	body, err := io.ReadAll(response.Body)
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
	return jsonv2.Unmarshal(body, &target)
}

// Checks the response from the Riot API and returns a Retry-After value if present.
func (c *InternalClient) checkResponse(equinoxReq *api.EquinoxRequest, response *http.Response) (time.Duration, error) {
	if !equinoxReq.IsCDN {
		c.ratelimit.Update(equinoxReq, &response.Header)
		if response.StatusCode == http.StatusTooManyRequests && equinoxReq.Retries < c.maxRetries {
			return c.ratelimit.CheckRetryAfter(equinoxReq, &response.Header)
		}
	}

	if response.StatusCode < http.StatusOK || response.StatusCode > 299 {
		equinoxReq.Logger.Error("Response with error code", zap.String("code", response.Status))
		err, ok := api.StatusCodeToError[response.StatusCode]
		if !ok {
			return 0, api.ErrorResponse{
				Status: api.Status{
					Message:    "Unknown error",
					StatusCode: response.StatusCode,
				},
			}
		}
		return 0, err
	}
	return 0, nil
}

// Generates a hash for the Authorization header. Don't want to store the Authorization header as key in plaintext.
func getAuthorizationHeaderHash(key string) string {
	if key == "" {
		return ""
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(key)))
}

func (c *InternalClient) GetDDragonLOLVersions(ctx context.Context, id string) ([]string, error) {
	logger := c.Logger(id)
	logger.Debug("Method started execution")
	equinoxReq, err := c.Request(ctx, logger, api.D_DRAGON_BASE_URL_FORMAT, http.MethodGet, "", "/api/versions.json", "", nil)
	if err != nil {
		logger.Error("Error creating request", zap.Error(err))
		return nil, err
	}
	var data []string
	err = c.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error("Error executing request", zap.Error(err))
		return nil, err
	}
	return data, nil
}
