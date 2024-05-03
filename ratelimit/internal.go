package ratelimit

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

type InternalRateLimitStore struct {
	Route            map[string]*Limits
	limitUsageFactor float64
	intervalOverhead time.Duration
	mutex            sync.Mutex
}

func (r *InternalRateLimitStore) Reserve(ctx context.Context, logger zerolog.Logger, route string, methodID string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	limits, ok := r.Route[route]
	if !ok {
		limits = NewLimits()
		r.Route[route] = limits
	}

	methods, ok := limits.Methods[methodID]
	if !ok {
		methods = NewLimit(METHOD_RATE_LIMIT_TYPE)
		limits.Methods[methodID] = methods
	}

	if err := limits.App.CheckBuckets(ctx, logger, route, methodID); err != nil {
		return err
	}

	return methods.CheckBuckets(ctx, logger, route, methodID)
}

func (r *InternalRateLimitStore) Update(ctx context.Context, logger zerolog.Logger, route string, methodID string, headers http.Header, retryAfter time.Duration) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	limits := r.Route[route]

	// If rate limited, set RetryAfter delay based on the rate limit type
	limitType := headers.Get(RATE_LIMIT_TYPE_HEADER)
	if limitType != "" {
		if limitType == APP_RATE_LIMIT_TYPE {
			limits.App.SetRetryAfter(retryAfter)
		} else {
			limits.Methods[methodID].SetRetryAfter(retryAfter)
		}
	}

	appLimitHeader := headers.Get(APP_RATE_LIMIT_HEADER)
	methodLimitHeader := headers.Get(METHOD_RATE_LIMIT_HEADER)

	if !limits.App.LimitsMatch(appLimitHeader) {
		countHeader := headers.Get(APP_RATE_LIMIT_COUNT_HEADER)
		newLimit := ParseHeaders(APP_RATE_LIMIT_TYPE, appLimitHeader, countHeader, r.limitUsageFactor, r.intervalOverhead)
		limits.App = newLimit
		logger.Debug().Object("limit", newLimit).Msg("New application limit")
	}

	if !limits.Methods[methodID].LimitsMatch(methodLimitHeader) {
		countHeader := headers.Get(METHOD_RATE_LIMIT_COUNT_HEADER)
		newLimit := ParseHeaders(METHOD_RATE_LIMIT_TYPE, methodLimitHeader, countHeader, r.limitUsageFactor, r.intervalOverhead)
		limits.Methods[methodID] = newLimit
		logger.Debug().Object("limit", newLimit).Msg("New method limit")
	}

	return nil
}
