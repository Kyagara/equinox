package ratelimit

import (
	"context"
	"errors"
	"net/http"
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
	Reserve(ctx context.Context, logger zerolog.Logger, route string, methodID string) error

	// Creates new buckets in a route with the limits provided in the response headers.
	Update(ctx context.Context, logger zerolog.Logger, route string, methodID string, headers http.Header, delay time.Duration) error
}

type RateLimit struct {
	store     Store
	StoreType StoreType
	// Factor to be applied to the limit. E.g. if set to 0.5, the limit will be reduced by 50%.
	LimitUsageFactor float64
	// Delay, in milliseconds, added to reset intervals.
	IntervalOverhead time.Duration
	Enabled          bool
}

func (r RateLimit) MarshalZerologObject(encoder *zerolog.Event) {
	if r.Enabled {
		encoder.Str("store", string(r.StoreType)).Float64("limit_usage_factor", r.LimitUsageFactor).Dur("interval_overhead", r.IntervalOverhead)
	}
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

func (r *RateLimit) Reserve(ctx context.Context, logger zerolog.Logger, route string, methodID string) error {
	if !r.Enabled {
		return ErrRateLimitIsDisabled
	}
	return r.store.Reserve(ctx, logger, route, methodID)
}

func (r *RateLimit) Update(ctx context.Context, logger zerolog.Logger, route string, methodID string, headers http.Header, retryAfter time.Duration) error {
	if !r.Enabled {
		return ErrRateLimitIsDisabled
	}
	return r.store.Update(ctx, logger, route, methodID, headers, retryAfter)
}
