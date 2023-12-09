package ratelimit

import (
	"context"
	"errors"
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
	Err429ButNoRetryAfterHeader = errors.New("received 429 but no Retry-After header was found")
	ErrContextDeadlineExceeded  = errors.New("waiting would exceed context deadline")
)

type RateLimit struct {
	Region  map[any]*Limits
	Enabled bool
	mutex   sync.Mutex
}

func NewInternalRateLimit() *RateLimit {
	return &RateLimit{Region: make(map[any]*Limits), Enabled: true}
}

// Limits in a region.
type Limits struct {
	App     *Limit
	Methods map[string]*Limit
}

func NewLimits() *Limits {
	return &Limits{
		App:     NewLimit(APP_RATE_LIMIT_TYPE),
		Methods: make(map[string]*Limit),
	}
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
	if err := methods.checkBuckets(ctx, logger, route, methodID); err != nil {
		return err
	}
	return nil
}

// Update creates new buckets in a route with the limits provided in the response headers.
func (r *RateLimit) Update(logger zerolog.Logger, route any, methodID string, headers *http.Header) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	limits := r.Region[route]
	if limits.App.limitsDontMatch(headers.Get(APP_RATE_LIMIT_HEADER)) {
		limits.App = parseHeaders(headers.Get(APP_RATE_LIMIT_HEADER), headers.Get(APP_RATE_LIMIT_COUNT_HEADER), APP_RATE_LIMIT_TYPE)
		logger.Debug().Msg("New Application buckets")
	}
	if limits.Methods[methodID].limitsDontMatch(headers.Get(METHOD_RATE_LIMIT_HEADER)) {
		limits.Methods[methodID] = parseHeaders(headers.Get(METHOD_RATE_LIMIT_HEADER), headers.Get(METHOD_RATE_LIMIT_COUNT_HEADER), METHOD_RATE_LIMIT_TYPE)
		logger.Debug().Msg("New Method buckets")
	}
}

func (r *RateLimit) CheckRetryAfter(route any, methodID string, headers *http.Header) (time.Duration, error) {
	retryAfter := headers.Get(RETRY_AFTER_HEADER)
	if retryAfter == "" {
		return 0, Err429ButNoRetryAfterHeader
	}

	delayF, _ := strconv.ParseFloat(retryAfter, 32)
	delay := time.Duration(delayF+0.5) * time.Second

	r.mutex.Lock()
	defer r.mutex.Unlock()

	limits := r.Region[route]
	limitType := headers.Get(RATE_LIMIT_TYPE_HEADER)

	if limitType == APP_RATE_LIMIT_TYPE {
		limits.App.setDelay(delay)
	} else {
		limits.Methods[methodID].setDelay(delay)
	}
	return delay, nil
}

// WaitN waits for the given duration after checking if the context deadline will be exceeded.
func WaitN(ctx context.Context, estimated time.Time, duration time.Duration) error {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(estimated) {
		return ErrContextDeadlineExceeded
	}
	select {
	case <-time.After(duration):
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func parseHeaders(limitHeader string, countHeader string, limitType string) *Limit {
	if limitHeader == "" || countHeader == "" {
		return NewLimit(limitType)
	}
	limits := strings.Split(limitHeader, ",")
	counts := strings.Split(countHeader, ",")
	limit := NewLimit(limitType)
	rates := make([]*Bucket, len(limits))
	for i := range limits {
		limit, seconds := getNumbersFromPair(limits[i])
		count, _ := getNumbersFromPair(counts[i])
		rates[i] = NewBucket(time.Duration(seconds), limit, limit-count)
	}
	limit.buckets = rates
	return limit
}

func getNumbersFromPair(pair string) (int, int) {
	numbers := strings.Split(pair, ":")
	interval, _ := strconv.Atoi(numbers[1])
	limitOrCount, _ := strconv.Atoi(numbers[0])
	return limitOrCount, interval
}
