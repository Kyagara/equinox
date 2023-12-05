package ratelimit

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/Kyagara/equinox/api"
	"go.uber.org/zap"
)

// Limit represents a collection of buckets and the type of limit (application or method).
type Limit struct {
	buckets    []*Bucket
	limitType  string
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
func (l *Limit) checkBuckets(ctx context.Context, equinoxReq *api.EquinoxRequest) error {
	if l.retryAfter > 0 {
		deadline, ok := ctx.Deadline()
		if ok && deadline.Before(time.Now().Add(l.retryAfter)) {
			return ErrContextDeadlineExceeded
		}
		select {
		case <-time.After(l.retryAfter):
		case <-ctx.Done():
			return ctx.Err()
		}
		l.retryAfter = 0
	}
	var limited []*Bucket
	for _, bucket := range l.buckets {
		if bucket.isRateLimited() {
			limited = append(limited, bucket)
		}
	}
	for i := len(limited) - 1; i >= 0; i-- {
		equinoxReq.Logger.Warn("Rate limited", zap.Any("route", equinoxReq.Route), zap.String("method_id", equinoxReq.MethodID), zap.String("limit_type", string(l.limitType)))
		err := limited[i].wait(ctx)
		if err != nil {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				return err
			}
		}
	}
	return nil
}

// Checks if the limits given in the header match the current buckets.
func (l *Limit) limitsDontMatch(limitHeader string) bool {
	if limitHeader == "" {
		return false
	}
	limits := strings.Split(limitHeader, ",")
	if len(l.buckets) != len(limits) {
		return true
	}
	for i, pair := range limits {
		if l.buckets[i] == nil {
			return true
		}
		limit, interval := getNumbersFromPair(pair)
		if l.buckets[i].limit != limit || l.buckets[i].interval != time.Duration(interval)*time.Second {
			return true
		}
	}
	return false
}

func (l *Limit) setDelay(delay time.Duration) {
	l.mutex.Lock()
	l.retryAfter = delay
	l.mutex.Unlock()
}
