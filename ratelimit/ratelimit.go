package ratelimit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Kyagara/equinox/api"
)

const (
	RATE_LIMIT_TYPE_HEADER         = "X-Rate-Limit-Type"
	RETRY_AFTER_HEADER             = "Retry-After"
	APP_RATE_LIMIT_HEADER          = "X-App-Rate-Limit"
	APP_RATE_LIMIT_COUNT_HEADER    = "X-App-Rate-Limit-Count"
	METHOD_RATE_LIMIT_HEADER       = "X-Method-Rate-Limit"
	METHOD_RATE_LIMIT_COUNT_HEADER = "X-Method-Rate-Limit-Count"
)

var (
	errDeadlineExceeded  = errors.New("waiting would exceed context deadline")
	errRateLimitExceeded = errors.New("rate limit exceeded")
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
	methodBucket := bucket.Methods[equinoxReq.MethodID]
	if methodBucket == nil {
		bucket.Methods[equinoxReq.MethodID] = make([]*Bucket, 0)
	}
	for _, rate := range bucket.App {
		err := rate.IsRateLimited(ctx)
		if err != nil {
			if errors.Is(err, errRateLimitExceeded) {
				equinoxReq.Logger.Warn(fmt.Sprintf("App rate limit reached on '%v' route for method '%s'. %v", equinoxReq.Route, equinoxReq.MethodID, err))
				err = rate.wait(ctx)
				if err != nil {
					select {
					case <-ctx.Done():
						return ctx.Err()
					default:
						return err
					}
				}
			}
			return err
		}
	}
	for _, rate := range methodBucket {
		err := rate.IsRateLimited(ctx)
		if err != nil {
			if errors.Is(err, errRateLimitExceeded) {
				equinoxReq.Logger.Warn(fmt.Sprintf("Method rate limit reached on '%v' route for method '%s'. %v", equinoxReq.Route, equinoxReq.MethodID, err))
				err = rate.wait(ctx)
				if err != nil {
					select {
					case <-ctx.Done():
						return ctx.Err()
					default:
						return err
					}
				}
			}
			return err
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