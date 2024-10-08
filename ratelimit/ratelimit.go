package ratelimit

import (
	"context"
	"errors"
	"math"
	"net/http"
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

	DEFAULT_RETRY_AFTER = 1 * time.Second
)

var (
	ErrContextDeadlineExceeded = errors.New("waiting would exceed context deadline")

	ErrRateLimitIsDisabled = errors.New("rate limit is disabled")
)

type StoreType string

const (
	InternalRateLimit StoreType = "Internal"
)

type Store interface {
	// Reserves one request for the App and Method buckets in a route.
	//
	// If rate limited, will block until the next bucket reset.
	Reserve(ctx context.Context, logger zerolog.Logger, route string, methodID string, isRSO bool) error

	// Creates new buckets in a route with the limits provided in the response headers.
	Update(ctx context.Context, logger zerolog.Logger, route string, methodID string, headers http.Header, retryAfter time.Duration) error
}

type RateLimit struct {
	store     Store
	StoreType StoreType
	// Factor to be applied to any Limit. E.g. If set to 0.5, the limit will be reduced by 50%.
	LimitUsageFactor float64
	// Delay, in milliseconds, added to reset intervals.
	IntervalOverhead time.Duration
	Enabled          bool
}

func NewInternalRateLimit(limitUsageFactor float64, intervalOverhead time.Duration) *RateLimit {
	limitUsageFactor, intervalOverhead = ValidateRateLimitOptions(limitUsageFactor, intervalOverhead)
	return &RateLimit{
		store: &InternalRateLimitStore{
			Route:            map[string]*Limits{},
			limitUsageFactor: limitUsageFactor,
			intervalOverhead: intervalOverhead,
			mutex:            sync.Mutex{},
		},
		StoreType:        InternalRateLimit,
		LimitUsageFactor: limitUsageFactor,
		IntervalOverhead: intervalOverhead,
		Enabled:          true,
	}
}

func (r *RateLimit) Reserve(ctx context.Context, logger zerolog.Logger, route string, methodID string, isRSO bool) error {
	if !r.Enabled {
		return ErrRateLimitIsDisabled
	}
	return r.store.Reserve(ctx, logger, route, methodID, isRSO)
}

func (r *RateLimit) Update(ctx context.Context, logger zerolog.Logger, route string, methodID string, headers http.Header, retryAfter time.Duration) error {
	if !r.Enabled {
		return ErrRateLimitIsDisabled
	}
	return r.store.Update(ctx, logger, route, methodID, headers, retryAfter)
}

// Parses the headers and returns a new Limit with its buckets.
func ParseHeaders(limitType string, limitHeader string, countHeader string, limitUsageFactor float64, intervalOverhead time.Duration) *Limit {
	if limitHeader == "" || countHeader == "" {
		return NewLimit(limitType)
	}

	limits := strings.Split(limitHeader, ",")
	counts := strings.Split(countHeader, ",")

	limit := &Limit{
		Type:       limitType,
		Buckets:    make([]*Bucket, len(limits)),
		RetryAfter: 0,
		mutex:      sync.Mutex{},
	}

	for i, limitString := range limits {
		baseLimit, interval := GetNumbersFromPair(limitString)
		newLimit := int(math.Max(1, float64(baseLimit)*limitUsageFactor))
		count, _ := GetNumbersFromPair(counts[i])
		limit.Buckets[i] = NewBucket(interval, intervalOverhead, baseLimit, newLimit, count)
	}

	return limit
}
