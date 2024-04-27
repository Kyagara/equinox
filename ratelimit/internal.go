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

	if err := limits.App.checkBuckets(ctx, logger, route, methodID); err != nil {
		return err
	}

	return methods.checkBuckets(ctx, logger, route, methodID)
}

func (r *InternalRateLimitStore) Update(ctx context.Context, logger zerolog.Logger, route string, methodID string, headers http.Header, delay time.Duration) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	limits := r.Route[route]

	// If rate limited, set retry after delay based on the rate limit type
	limitType := headers.Get(RATE_LIMIT_TYPE_HEADER)
	if limitType != "" {
		if limitType == APP_RATE_LIMIT_TYPE {
			limits.App.setRetryAfter(delay)
		} else {
			limits.Methods[methodID].setRetryAfter(delay)
		}
	}

	appRateLimitHeader := headers.Get(APP_RATE_LIMIT_HEADER)
	methodRateLimitHeader := headers.Get(METHOD_RATE_LIMIT_HEADER)

	if !limits.App.limitsMatch(appRateLimitHeader) {
		appRateLimitCountHeader := headers.Get(APP_RATE_LIMIT_COUNT_HEADER)
		limits.App = ParseHeaders(appRateLimitHeader, appRateLimitCountHeader, APP_RATE_LIMIT_TYPE, r.limitUsageFactor, r.intervalOverhead)
		logger.Debug().Str("route", route).Msg("New Application buckets")
	}

	if !limits.Methods[methodID].limitsMatch(methodRateLimitHeader) {
		methodRateLimitCountHeader := headers.Get(METHOD_RATE_LIMIT_COUNT_HEADER)
		limits.Methods[methodID] = ParseHeaders(methodRateLimitHeader, methodRateLimitCountHeader, METHOD_RATE_LIMIT_TYPE, r.limitUsageFactor, r.intervalOverhead)
		logger.Debug().Str("route", route).Str("method", methodID).Msg("New Method buckets")
	}

	return nil
}
