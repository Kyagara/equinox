package ratelimit

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

// Limits in a route.
type Limits struct {
	App     *Limit
	Methods map[string]*Limit
}

func NewLimits() *Limits {
	return &Limits{
		App:     NewLimit(APP_RATE_LIMIT_TYPE),
		Methods: make(map[string]*Limit, 1),
	}
}

// Represents a collection of buckets and the type of limit (application or method).
type Limit struct {
	limitType  string
	buckets    []*Bucket
	retryAfter time.Duration
	mutex      sync.Mutex
}

func NewLimit(limitType string) *Limit {
	return &Limit{
		buckets:    make([]*Bucket, 0),
		limitType:  limitType,
		retryAfter: 0,
		mutex:      sync.Mutex{},
	}
}

// Checks if any of the buckets provided are rate limited, and if so, blocks until the next reset.
func (l *Limit) checkBuckets(ctx context.Context, logger zerolog.Logger, route string, methodID string) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.retryAfter > 0 {
		err := WaitN(ctx, time.Now().Add(l.retryAfter), l.retryAfter)
		if err != nil {
			logger.Warn().Err(err).Msg("Failed to wait for retry after")
			return err
		}

		l.retryAfter = 0
	}

	// Reverse loop, buckets with higher limits will be checked first
	for i := len(l.buckets) - 1; i >= 0; i-- {
		bucket := l.buckets[i]
		bucket.mutex.Lock()

		if bucket.IsRateLimited() {
			logger.Warn().
				Str("route", route).
				Str("method_id", methodID).
				Str("limit_type", l.limitType).
				Object("bucket", bucket).
				Msg("Rate limited")

			err := WaitN(ctx, bucket.Next, time.Until(bucket.Next))
			if err != nil {
				bucket.mutex.Unlock()
				logger.Warn().Err(err).Msg("Failed to wait for reset")
				return err
			}

			// next reset is now in the past, so reset the bucket
			bucket.Check()
			bucket.Tokens++
		}

		bucket.mutex.Unlock()
	}

	return nil
}

// Checks if the limits given in the header match the current buckets.
func (l *Limit) limitsMatch(limitHeader string) bool {
	if limitHeader == "" {
		return false
	}

	limits := strings.Split(limitHeader, ",")
	if len(l.buckets) != len(limits) {
		return false
	}

	for i, pair := range limits {
		bucket := l.buckets[i]
		if bucket == nil {
			return false
		}

		limit, interval := getNumbersFromPair(pair)
		if bucket.BaseLimit != limit || bucket.Interval != interval {
			return false
		}
	}

	return true
}

func (l *Limit) setRetryAfter(delay time.Duration) {
	l.mutex.Lock()
	l.retryAfter = delay
	l.mutex.Unlock()
}
