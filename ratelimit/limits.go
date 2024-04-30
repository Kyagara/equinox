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
	Type       string
	Buckets    []*Bucket
	RetryAfter time.Duration
	mutex      sync.Mutex
}

func (l *Limit) MarshalZerologObject(encoder *zerolog.Event) {
	buckets := zerolog.Arr()
	for _, bucket := range l.Buckets {
		buckets.Object(bucket)
	}
	encoder.Str("type", l.Type).Array("buckets", buckets).Dur("retry_after", l.RetryAfter)
}

func NewLimit(limitType string) *Limit {
	return &Limit{
		Type:       limitType,
		Buckets:    make([]*Bucket, 0),
		RetryAfter: 0,
		mutex:      sync.Mutex{},
	}
}

// Checks if any of the buckets provided are rate limited, and if so, blocks until the next reset.
func (l *Limit) CheckBuckets(ctx context.Context, logger zerolog.Logger, route string, methodID string) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.RetryAfter > 0 {
		err := WaitN(ctx, time.Now().Add(l.RetryAfter), l.RetryAfter)
		if err != nil {
			logger.Warn().Err(err).Msg("Failed to wait for retry after")
			return err
		}

		l.RetryAfter = 0
	}

	// Reverse loop, buckets with higher limits will be checked first
	for i := len(l.Buckets) - 1; i >= 0; i-- {
		bucket := l.Buckets[i]
		bucket.mutex.Lock()

		if bucket.IsRateLimited() {
			logger.Warn().
				Str("route", route).
				Str("method_id", methodID).
				Str("limit_type", l.Type).
				Object("bucket", bucket).
				Msg("Rate limited")

			err := WaitN(ctx, bucket.Next, time.Until(bucket.Next))
			if err != nil {
				bucket.mutex.Unlock()
				logger.Warn().Err(err).Msg("Failed to wait for reset")
				return err
			}

			// 'Next' is now in the past, the bucket should reset now
			bucket.Check()
			bucket.Tokens++
		}

		bucket.mutex.Unlock()
	}

	return nil
}

// Checks if the limits given in the header match the current buckets.
func (l *Limit) LimitsMatch(limitHeader string) bool {
	if limitHeader == "" {
		return false
	}

	limits := strings.Split(limitHeader, ",")
	if len(l.Buckets) != len(limits) {
		return false
	}

	for i, pair := range limits {
		bucket := l.Buckets[i]
		if bucket == nil {
			return false
		}

		limit, interval := GetNumbersFromPair(pair)
		if bucket.BaseLimit != limit || bucket.Interval != interval {
			return false
		}
	}

	return true
}

func (l *Limit) SetRetryAfter(delay time.Duration) {
	l.mutex.Lock()
	l.RetryAfter = delay
	l.mutex.Unlock()
}
