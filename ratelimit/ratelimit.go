package ratelimit

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Kyagara/equinox/api"
	"go.uber.org/zap"
)

const (
	RATE_LIMIT_TYPE_HEADER = "X-Rate-Limit-Type"
	RETRY_AFTER_HEADER     = "Retry-After"

	APP_RATE_LIMIT_HEADER          = "X-App-Rate-Limit"
	APP_RATE_LIMIT_COUNT_HEADER    = "X-App-Rate-Limit-Count"
	METHOD_RATE_LIMIT_HEADER       = "X-Method-Rate-Limit"
	METHOD_RATE_LIMIT_COUNT_HEADER = "X-Method-Rate-Limit-Count"
)

var (
	Err429ButNoRetryAfterHeader = errors.New("received 429 but no Retry-After header was found")
	ErrContextDeadlineExceeded  = errors.New("waiting would exceed context deadline")
)

type RateLimit struct {
	// Map of limits, keyed by route
	Limits map[any]*Limits
	mutex  sync.Mutex
}

type Limits struct {
	App     []*Bucket
	Methods map[string][]*Bucket
}

func NewLimits() *Limits {
	return &Limits{
		App:     []*Bucket{},
		Methods: make(map[string][]*Bucket),
	}
}

// Take decreases tokens for the App and Method rate limit buckets in a route by one.
//
// If rate limited, will block until the next bucket reset.
func (r *RateLimit) Take(ctx context.Context, equinoxReq *api.EquinoxRequest) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	limits := r.Limits[equinoxReq.Route]
	if limits == nil {
		limits = NewLimits()
		r.Limits[equinoxReq.Route] = limits
	}
	methods := limits.Methods[equinoxReq.MethodID]
	if methods == nil {
		limits.Methods[equinoxReq.MethodID] = make([]*Bucket, 0)
	}
	err := r.checkBuckets(ctx, equinoxReq, limits.App, "application")
	if err != nil {
		return err
	}
	err = r.checkBuckets(ctx, equinoxReq, methods, "method")
	if err != nil {
		return err
	}
	return nil
}

// Update creates new buckets in a route with the limits provided in the response headers.
func (r *RateLimit) Update(equinoxReq *api.EquinoxRequest, responseHeaders *http.Header) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	limits := r.Limits[equinoxReq.Route]
	if limitsDontMatch(limits.App, responseHeaders.Get(APP_RATE_LIMIT_HEADER)) {
		limits.App = parseHeaders(responseHeaders.Get(APP_RATE_LIMIT_HEADER), responseHeaders.Get(APP_RATE_LIMIT_COUNT_HEADER))
		equinoxReq.Logger.Debug("New Application buckets", zap.Objects("buckets", limits.App))
	}
	if limitsDontMatch(limits.Methods[equinoxReq.MethodID], responseHeaders.Get(METHOD_RATE_LIMIT_HEADER)) {
		limits.Methods[equinoxReq.MethodID] = parseHeaders(responseHeaders.Get(METHOD_RATE_LIMIT_HEADER), responseHeaders.Get(METHOD_RATE_LIMIT_COUNT_HEADER))
		equinoxReq.Logger.Debug("New Method buckets", zap.Objects("buckets", limits.Methods[equinoxReq.MethodID]))
	}
}

// Checks if the limits given in the header match the current buckets
//
// Doesn't look good
func limitsDontMatch(buckets []*Bucket, limitHeader string) bool {
	if limitHeader == "" {
		return false
	}
	limits := strings.Split(limitHeader, ",")
	if len(buckets) != len(limits) {
		return true
	}
	for i, pair := range limits {
		if buckets[i] == nil {
			return true
		}
		limit, interval := getNumbersFromPair(pair)
		if buckets[i].limit != limit || buckets[i].interval != time.Duration(interval)*time.Second {
			return true
		}
	}
	return false
}

// checkBuckets checks if any of the buckets provided are rate limited, and if so, blocks until the next reset.
//
// It loops from the end of the slice to ensure that the bigger limits are checked first.
func (r *RateLimit) checkBuckets(ctx context.Context, equinoxReq *api.EquinoxRequest, buckets []*Bucket, bucket_type string) error {
	var limited []*Bucket
	for _, bucket := range buckets {
		if bucket.isRateLimited(ctx) {
			limited = append(limited, bucket)
		}
	}
	for i := len(limited) - 1; i >= 0; i-- {
		equinoxReq.Logger.Warn("Rate limited", zap.String("bucket_type", bucket_type), zap.Any("route", equinoxReq.Route), zap.Object("bucket", limited[i]))
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

func parseHeaders(limitHeader string, countHeader string) []*Bucket {
	if limitHeader == "" || countHeader == "" {
		return []*Bucket{}
	}
	limits := strings.Split(limitHeader, ",")
	counts := strings.Split(countHeader, ",")
	rates := make([]*Bucket, len(limits))
	for i := range limits {
		limit, seconds := getNumbersFromPair(limits[i])
		count, _ := getNumbersFromPair(counts[i])
		rates[i] = NewBucket(time.Duration(seconds), limit, limit-count)
	}
	return rates
}

func getNumbersFromPair(pair string) (int, int) {
	numbers := strings.Split(pair, ":")
	interval, _ := strconv.Atoi(numbers[1])
	limitOrCount, _ := strconv.Atoi(numbers[0])
	return limitOrCount, interval
}
