package ratelimit

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Kyagara/equinox/api"
	"golang.org/x/time/rate"
)

const (
	RATE_LIMIT_TYPE_HEADER         = "X-Rate-Limit-Type"
	RETRY_AFTER_HEADER             = "Retry-After"
	APP_RATE_LIMIT_HEADER          = "X-App-Rate-Limit"
	APP_RATE_LIMIT_COUNT_HEADER    = "X-App-Rate-Limit-Count"
	METHOD_RATE_LIMIT_HEADER       = "X-Method-Rate-Limit"
	METHOD_RATE_LIMIT_COUNT_HEADER = "X-Method-Rate-Limit-Count"
)

type RateLimit struct {
	Buckets map[any]*Buckets
	mu      sync.Mutex
}

type Buckets struct {
	App     []*rate.Limiter
	Methods map[string][]*rate.Limiter
}

func NewBuckets() *Buckets {
	return &Buckets{
		App:     []*rate.Limiter{},
		Methods: make(map[string][]*rate.Limiter),
	}
}

// Take increases the count for the App and Method rate limit buckets by one in a route.
func (r *RateLimit) Take(equinoxReq *api.EquinoxRequest) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	bucket := r.Buckets[equinoxReq.Route]
	if bucket == nil {
		bucket = NewBuckets()
		r.Buckets[equinoxReq.Route] = bucket
	}
	if bucket.Methods == nil {
		bucket.Methods = make(map[string][]*rate.Limiter)
	}
	methodBucket := bucket.Methods[equinoxReq.MethodID]
	if methodBucket == nil {
		bucket.Methods[equinoxReq.MethodID] = make([]*rate.Limiter, 0)
	}
	// For now, just return an error if the rate limit is reached.
	for _, rate := range bucket.App {
		if !rate.Allow() {
			return fmt.Errorf("app rate limit reached on '%v' route for method '%s'", equinoxReq.Route, equinoxReq.MethodID)
		}
	}
	for _, rate := range methodBucket {
		if !rate.Allow() {
			return fmt.Errorf("method rate limit reached on '%v' route for method '%s'", equinoxReq.Route, equinoxReq.MethodID)
		}
	}
	return nil
}

// Update creates new buckets in a route with the limits provided in the response headers.
// TODO: Maybe add a way to dinamically update the buckets with new rates?
// Currently this only runs one time, when it is known that the buckets are empty.
func (r *RateLimit) Update(equinoxReq *api.EquinoxRequest, responseHeaders *http.Header) {
	r.mu.Lock()
	defer r.mu.Unlock()
	bucket := r.Buckets[equinoxReq.Route]
	if bucket == nil {
		bucket = NewBuckets()
		r.Buckets[equinoxReq.Route] = bucket
	}
	if bucket.Methods == nil {
		bucket.Methods = make(map[string][]*rate.Limiter)
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

func parseHeaders(limitHeader string, countHeader string) []*rate.Limiter {
	if limitHeader == "" || countHeader == "" {
		return []*rate.Limiter{}
	}
	limits := strings.Split(limitHeader, ",")
	counts := strings.Split(countHeader, ",")
	rates := make([]*rate.Limiter, len(limits))
	now := time.Now()
	for i := range limits {
		limit, seconds := getNumbersFromPair(limits[i])
		count, _ := getNumbersFromPair(counts[i])
		rates[i] = rate.NewLimiter(rate.Every(time.Second*time.Duration(seconds)), limit)
		rates[i].AllowN(now, count)
	}
	return rates
}

func getNumbersFromPair(pair string) (int, int) {
	numbers := strings.Split(pair, ":")
	seconds, _ := strconv.Atoi(numbers[1])
	limitOrCount, _ := strconv.Atoi(numbers[0])
	return limitOrCount, seconds
}
