package ratelimit

import (
	"context"
	"errors"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

const (
	RATE_LIMIT_TYPE_HEADER = "X-Rate-Limit-Type"
	RETRY_AFTER_HEADER     = "Retry-After"

	APP_RATE_LIMIT_HEADER          = "X-App-Rate-Limit"
	APP_RATE_LIMIT_COUNT_HEADER    = "X-App-Rate-Limit-Count"
	METHOD_RATE_LIMIT_HEADER       = "X-Method-Rate-Limit"
	METHOD_RATE_LIMIT_COUNT_HEADER = "X-Method-Rate-Limit-Count"

	APP_RATE_LIMIT_TYPE     = "application"
	METHOD_RATE_LIMIT_TYPE  = "method"
	SERVICE_RATE_LIMIT_TYPE = "service"
)

var (
	ErrContextDeadlineExceeded = errors.New("waiting would exceed context deadline")
)

type RateLimit struct {
	// 'any' is used here because routes can be PlatformRoute, RegionalRoute...
	Region  map[any]*Limits
	Enabled bool
	// Factor to be applied to the limit. E.g. if set to 0.5, the limit will be reduced by 50%.
	LimitUsageFactor float32
	// Delay in milliseconds to be add to reset intervals.
	IntervalOverhead time.Duration
	mutex            sync.Mutex
}

func NewInternalRateLimit(limitUsageFactor float32, intervalOverhead time.Duration) *RateLimit {
	if limitUsageFactor < 0.0 || limitUsageFactor > 1.0 {
		limitUsageFactor = 1.0
	}
	if intervalOverhead < 0 {
		intervalOverhead = 1 * time.Second
	}
	return &RateLimit{Region: make(map[any]*Limits), LimitUsageFactor: limitUsageFactor, IntervalOverhead: intervalOverhead, Enabled: true}
}

// Take decreases tokens for the App and Method rate limit buckets in a route by one.
//
// If rate limited, will block until the next bucket reset.
func (r *RateLimit) Take(ctx context.Context, logger zerolog.Logger, route any, methodID string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	limits, ok := r.Region[route]
	if !ok {
		limits = NewLimits()
		r.Region[route] = limits
	}

	methods, ok := limits.Methods[methodID]
	if !ok {
		methods = NewLimit(METHOD_RATE_LIMIT_TYPE)
		limits.Methods[methodID] = methods
	}

	if err := limits.App.checkBuckets(ctx, logger, route, methodID); err != nil {
		return err
	}

	return methods.checkBuckets(ctx, logger, route, methodID)
}

// Update creates new buckets in a route with the limits provided in the response headers.
func (r *RateLimit) Update(logger zerolog.Logger, route any, methodID string, headers http.Header) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	limits := r.Region[route]

	appRateLimitHeader := headers.Get(APP_RATE_LIMIT_HEADER)
	appRateLimitCountHeader := headers.Get(APP_RATE_LIMIT_COUNT_HEADER)
	methodRateLimitHeader := headers.Get(METHOD_RATE_LIMIT_HEADER)
	methodRateLimitCountHeader := headers.Get(METHOD_RATE_LIMIT_COUNT_HEADER)

	if !limits.App.limitsMatch(appRateLimitHeader) {
		limits.App = r.parseHeaders(appRateLimitHeader, appRateLimitCountHeader, APP_RATE_LIMIT_TYPE)
		logger.Debug().Any("route", route).Msg("New Application buckets")
	}

	if !limits.Methods[methodID].limitsMatch(methodRateLimitHeader) {
		limits.Methods[methodID] = r.parseHeaders(methodRateLimitHeader, methodRateLimitCountHeader, METHOD_RATE_LIMIT_TYPE)
		logger.Debug().Any("route", route).Str("method", methodID).Msg("New Method buckets")
	}
}

// CheckRetryAfter returns the number of seconds to wait before retrying from the Retry-After header, or 2 seconds if not found.
func (r *RateLimit) CheckRetryAfter(route any, methodID string, headers http.Header) time.Duration {
	retryAfter := headers.Get(RETRY_AFTER_HEADER)
	if retryAfter == "" {
		return 2 * time.Second
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	delayF, _ := strconv.ParseFloat(retryAfter, 32)
	delay := time.Duration(delayF+0.5) * time.Second

	limits := r.Region[route]
	limitType := headers.Get(RATE_LIMIT_TYPE_HEADER)

	if limitType == APP_RATE_LIMIT_TYPE {
		limits.App.setDelay(delay)
	} else {
		limits.Methods[methodID].setDelay(delay)
	}

	return delay
}

func (r *RateLimit) parseHeaders(limitHeader string, countHeader string, limitType string) *Limit {
	if limitHeader == "" || countHeader == "" {
		return NewLimit(limitType)
	}

	limit := NewLimit(limitType)

	limits := strings.Split(limitHeader, ",")
	counts := strings.Split(countHeader, ",")
	rates := make([]*Bucket, len(limits))

	for i := range limits {
		baseLimit, interval := getNumbersFromPair(limits[i])
		count, _ := getNumbersFromPair(counts[i])
		limit := int(math.Floor(math.Max(1, float64(baseLimit)*float64(r.LimitUsageFactor))))
		rates[i] = NewBucket(interval, r.IntervalOverhead, baseLimit, limit, limit-count)
	}

	limit.buckets = rates
	return limit
}
