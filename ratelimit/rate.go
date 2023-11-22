package ratelimit

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

const (
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

func (r *RateLimit) Check(route any, method string, headers *http.Header) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.Buckets[route] == nil {
		r.Buckets[route] = NewBuckets()
	}
	if r.Buckets[route].Methods == nil {
		r.Buckets[route].Methods = make(map[string][]*rate.Limiter)
	}
	if len(r.Buckets[route].App) == 0 {
		r.Buckets[route].App = parseHeaders(headers.Get(APP_RATE_LIMIT_HEADER), headers.Get(APP_RATE_LIMIT_COUNT_HEADER))
	}
	if r.Buckets[route].Methods[method] == nil {
		r.Buckets[route].Methods[method] = parseHeaders(headers.Get(METHOD_RATE_LIMIT_HEADER), headers.Get(METHOD_RATE_LIMIT_COUNT_HEADER))
	}
	for _, rate := range r.Buckets[route].App {
		if !rate.Allow() {
			return fmt.Errorf("app rate limit exceeded")
		}
	}
	for _, rate := range r.Buckets[route].Methods[method] {
		if !rate.Allow() {
			return fmt.Errorf("method rate limit exceeded")
		}
	}
	return nil
}

func parseHeaders(limitHeader string, countHeader string) []*rate.Limiter {
	if limitHeader == "" || countHeader == "" {
		return []*rate.Limiter{}
	}
	limits := strings.Split(limitHeader, ",")
	counts := strings.Split(countHeader, ",")
	size := len(limits)
	rates := make([]*rate.Limiter, size)
	for i := 0; i < size; i++ {
		limit, seconds := getNumbersFromPair(limits[i])
		count, _ := getNumbersFromPair(counts[i])
		rates[i] = rate.NewLimiter(rate.Every(time.Second*time.Duration(seconds)), limit)
		rates[i].AllowN(time.Now(), count)
	}
	return rates
}

func getNumbersFromPair(pair string) (int, int) {
	numbers := strings.Split(pair, ":")
	seconds, _ := strconv.Atoi(numbers[1])
	limitOrCount, _ := strconv.Atoi(numbers[0])
	return limitOrCount, seconds
}
