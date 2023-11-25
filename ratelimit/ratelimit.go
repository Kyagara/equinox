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
	ErrRateLimitedButNoRetryAfterHeader = errors.New("rate limited but no Retry-After header was found, stopping")
	ErrContextDeadlineExceeded          = errors.New("waiting would exceed context deadline")
)

type RateLimit struct {
	Buckets map[any]*Limits
	mutex   sync.Mutex
}

func NewLimits() *Limits {
	return &Limits{
		App:     []*Bucket{},
		Methods: make(map[string][]*Bucket),
	}
}

type Limits struct {
	App     []*Bucket
	Methods map[string][]*Bucket
}

// Take decreases tokens for the App and Method rate limit buckets in a route by one.
func (r *RateLimit) Take(ctx context.Context, equinoxReq *api.EquinoxRequest) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	bucket := r.Buckets[equinoxReq.Route]
	if bucket == nil {
		bucket = NewLimits()
		r.Buckets[equinoxReq.Route] = bucket
	}
	if bucket.Methods == nil {
		bucket.Methods = make(map[string][]*Bucket)
	}
	err := r.checkBuckets(ctx, equinoxReq, bucket.App, "app")
	if err != nil {
		return err
	}
	methodBucket := bucket.Methods[equinoxReq.MethodID]
	if methodBucket == nil {
		bucket.Methods[equinoxReq.MethodID] = make([]*Bucket, 0)
	}
	err = r.checkBuckets(ctx, equinoxReq, methodBucket, "method")
	if err != nil {
		return err
	}
	return nil
}

func (r *RateLimit) checkBuckets(ctx context.Context, equinoxReq *api.EquinoxRequest, buckets []*Bucket, bucket_type string) error {
	var limited []*Bucket
	for _, bucket := range buckets {
		if bucket.isRateLimited(ctx) {
			limited = append(limited, bucket)
		}
	}
	for i := len(limited) - 1; i >= 0; i-- {
		equinoxReq.Logger.Warn("Rate limited", zap.String("bucket", bucket_type), zap.Any("route", equinoxReq.Route), zap.Object("bucket", limited[i]))
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

// Update creates new buckets in a route with the limits provided in the response headers.
// TODO: Maybe add a way to dinamically update the buckets with new rates?
// Currently this only runs one time, when it is known that the buckets are empty.
func (r *RateLimit) Update(equinoxReq *api.EquinoxRequest, responseHeaders *http.Header) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	bucket := r.Buckets[equinoxReq.Route]
	if bucket == nil {
		bucket = NewLimits()
		r.Buckets[equinoxReq.Route] = bucket
	}
	if bucket.Methods == nil {
		bucket.Methods = make(map[string][]*Bucket)
	}
	if len(bucket.App) == 0 {
		bucket.App = parseHeaders(responseHeaders.Get(APP_RATE_LIMIT_HEADER), responseHeaders.Get(APP_RATE_LIMIT_COUNT_HEADER))
	}
	methodBucket := bucket.Methods[equinoxReq.MethodID]
	if methodBucket == nil {
		methodBucket = parseHeaders(responseHeaders.Get(METHOD_RATE_LIMIT_HEADER), responseHeaders.Get(METHOD_RATE_LIMIT_COUNT_HEADER))
		bucket.Methods[equinoxReq.MethodID] = methodBucket
	}
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
	seconds, _ := strconv.Atoi(numbers[1])
	limitOrCount, _ := strconv.Atoi(numbers[0])
	return limitOrCount, seconds
}
