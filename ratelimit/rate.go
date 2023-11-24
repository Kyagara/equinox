package ratelimit

import (
	"context"
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

type RateLimit struct {
	Buckets map[any]*Limits
	mu      sync.Mutex
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

type Bucket struct {
	// Time interval in seconds
	interval time.Duration
	// Maximum number of tokens
	limit int
	// Current number of tokens
	tokens int
	// Updates every time its checked
	updated time.Time
	// Next reset
	next time.Time
	mu   sync.Mutex
}

func NewBucket(interval time.Duration, limit int, tokens int) *Bucket {
	now := time.Now()
	return &Bucket{
		interval: interval * time.Second,
		limit:    limit,
		tokens:   tokens,
		updated:  now,
		next:     now.Add(interval * time.Second),
		mu:       sync.Mutex{},
	}
}

// Responsible for updating the bucket, resets the tokens if necessary.
func (b *Bucket) check() {
	now := time.Now()
	if now.Sub(b.updated) >= b.interval {
		b.tokens = b.limit
		b.next = now.Add(b.interval)
	}
	b.updated = now
}

// TODO: Wait should block if the rate limit is reached until the bucket resets.
func (b *Bucket) Wait(ctx context.Context) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.check()
	if b.limit == 0 {
		return nil
	}
	if b.tokens-1 <= 0 {
		return fmt.Errorf("exceeded the bucket's limit %d", b.limit)
	}
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(b.next) {
		return fmt.Errorf("waiting would exceed context deadline")
	}
	b.tokens--
	return nil
}

// Take decreases tokens for the App and Method rate limit buckets in a route by one.
func (r *RateLimit) Take(ctx context.Context, equinoxReq *api.EquinoxRequest) error {
	r.mu.Lock()
	defer r.mu.Unlock()
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
		err := rate.Wait(ctx)
		if err != nil {
			return fmt.Errorf("app rate limit reached on '%v' route for method '%s'. %v", equinoxReq.Route, equinoxReq.MethodID, err)
		}
	}
	for _, rate := range methodBucket {
		err := rate.Wait(ctx)
		if err != nil {
			return fmt.Errorf("method rate limit reached on '%v' route for method '%s'. %v", equinoxReq.Route, equinoxReq.MethodID, err)
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
