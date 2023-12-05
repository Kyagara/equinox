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

	APP_RATE_LIMIT_TYPE     = "application"
	METHOD_RATE_LIMIT_TYPE  = "method"
	SERVICE_RATE_LIMIT_TYPE = "service"
)

var (
	Err429ButNoRetryAfterHeader = errors.New("received 429 but no Retry-After header was found")
	ErrContextDeadlineExceeded  = errors.New("waiting would exceed context deadline")
)

type RateLimit struct {
	Region map[any]*Limits
	mutex  sync.Mutex
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
func (r *RateLimit) Take(ctx context.Context, equinoxReq *api.EquinoxRequest) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	limits, ok := r.Region[equinoxReq.Route]
	if !ok {
		limits = NewLimits()
		r.Region[equinoxReq.Route] = limits
	}
	methods, ok := limits.Methods[equinoxReq.MethodID]
	if !ok {
		methods = NewLimit(METHOD_RATE_LIMIT_TYPE)
		limits.Methods[equinoxReq.MethodID] = methods
	}
	if err := limits.App.checkBuckets(ctx, equinoxReq); err != nil {
		return err
	}
	if err := methods.checkBuckets(ctx, equinoxReq); err != nil {
		return err
	}
	return nil
}

// Update creates new buckets in a route with the limits provided in the response headers.
func (r *RateLimit) Update(equinoxReq *api.EquinoxRequest, headers *http.Header) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	limits := r.Region[equinoxReq.Route]
	if limits.App.limitsDontMatch(headers.Get(APP_RATE_LIMIT_HEADER)) {
		limits.App = parseHeaders(headers.Get(APP_RATE_LIMIT_HEADER), headers.Get(APP_RATE_LIMIT_COUNT_HEADER), APP_RATE_LIMIT_TYPE)
		equinoxReq.Logger.Debug("New Application buckets", zap.Objects("buckets", limits.App.buckets))
	}
	if limits.Methods[equinoxReq.MethodID].limitsDontMatch(headers.Get(METHOD_RATE_LIMIT_HEADER)) {
		limits.Methods[equinoxReq.MethodID] = parseHeaders(headers.Get(METHOD_RATE_LIMIT_HEADER), headers.Get(METHOD_RATE_LIMIT_COUNT_HEADER), METHOD_RATE_LIMIT_TYPE)
		equinoxReq.Logger.Debug("New Method buckets", zap.Objects("buckets", limits.Methods[equinoxReq.MethodID].buckets))
	}
}

func (r *RateLimit) CheckRetryAfter(equinoxReq *api.EquinoxRequest, headers *http.Header) (time.Duration, error) {
	retryAfter := headers.Get(RETRY_AFTER_HEADER)
	if retryAfter == "" {
		return 0, Err429ButNoRetryAfterHeader
	}

	delayF, _ := strconv.ParseFloat(retryAfter, 32)
	delay := time.Duration(delayF+0.5) * time.Second

	r.mutex.Lock()
	defer r.mutex.Unlock()

	limits := r.Region[equinoxReq.Route]
	limitType := headers.Get(RATE_LIMIT_TYPE_HEADER)

	if limitType == APP_RATE_LIMIT_TYPE {
		limits.App.setDelay(delay)
	} else {
		limits.Methods[equinoxReq.MethodID].setDelay(delay)
	}
	return delay, nil
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
