package ratelimit

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

// Limit represents a collection of buckets and the type of limit (application or method).
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
func (l *Limit) checkBuckets(ctx context.Context, logger zerolog.Logger, route any, methodID string) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.retryAfter > 0 {
		err := WaitN(ctx, time.Now().Add(l.retryAfter), l.retryAfter)
		if err != nil {
			return err
		}
		l.retryAfter = 0
	}

	for i := len(l.buckets) - 1; i >= 0; i-- {
		bucket := l.buckets[i]
		bucket.mutex.Lock()
		if bucket.isRateLimited() {
			logger.Warn().
				Any("route", route).
				Str("method_id", methodID).
				Str("limit_type", l.limitType).
				Object("bucket", bucket).
				Msg("Rate limited")
			err := WaitN(ctx, bucket.next, time.Until(bucket.next))
			if err != nil {
				bucket.mutex.Unlock()
				return err
			}
			bucket.check()
			bucket.tokens--
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
		if bucket.baseLimit != limit || bucket.interval != interval {
			return false
		}
	}
	return true
}

func (l *Limit) setDelay(delay time.Duration) {
	l.mutex.Lock()
	l.retryAfter = delay
	l.mutex.Unlock()
}
