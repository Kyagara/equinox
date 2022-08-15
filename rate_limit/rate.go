package rate_limit

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type RateLimitStoreType string

const (
	InternalRateLimiter RateLimitStoreType = "Internal"
	RedisRateLimiter    RateLimitStoreType = "Redis"
)

type RateLimit struct {
	store     Store
	Enabled   bool
	StoreType RateLimitStoreType
}

type Store interface {
	Get(route interface{}, endpointName string, methodName string) (*Rate, error)
	GetAppRate(route interface{}) (*Rate, error)
	Set(route interface{}, endpointName string, methodName string, headers *http.Header) error
	SetAppRate(route interface{}, headers *http.Header) error
	IsRateLimited(rate *Rate) (bool, error)
}

type Rate struct {
	Seconds RateTiming
	Minutes RateTiming
}

type RateTiming struct {
	// The amount in seconds the rate limit should reset.
	Time int
	// Maximum amount of requests in n seconds.
	Limit int
	// Current count of requests made in n seconds.
	Count int

	Expire time.Time
	Access time.Time
}

const (
	AppRateLimitHeader         = "X-App-Rate-Limit"
	AppRateLimitCountHeader    = "X-App-Rate-Limit-Count"
	MethodRateLimitHeader      = "X-Method-Rate-Limit"
	MethodRateLimitCountHeader = "X-Method-Rate-Limit-Count"
)

var (
	ErrRateLimitingIsDisabled = errors.New("Rate limiting is disabled")
)

func NewInternalRateLimit() (*RateLimit, error) {
	rate := &RateLimit{
		store: &InternalRateStore{
			client: nil,
			Route:  map[interface{}]*Enpoints{},
		},
		Enabled:   true,
		StoreType: InternalRateLimiter,
	}

	return rate, nil
}

func (r *RateLimit) Get(route interface{}, endpointName string, methodName string) (*Rate, error) {
	if !r.Enabled {
		return nil, ErrRateLimitingIsDisabled
	}

	return r.store.Get(route, endpointName, methodName)
}

func (r *RateLimit) GetAppRate(route interface{}) (*Rate, error) {
	if !r.Enabled {
		return nil, ErrRateLimitingIsDisabled
	}

	return r.store.GetAppRate(route)
}

func (r *RateLimit) Set(route interface{}, endpointName string, methodName string, headers *http.Header) error {
	if !r.Enabled {
		return ErrRateLimitingIsDisabled
	}

	return r.store.Set(route, endpointName, methodName, headers)
}

func (r *RateLimit) SetAppRate(route interface{}, headers *http.Header) error {
	if !r.Enabled {
		return ErrRateLimitingIsDisabled
	}

	return r.store.SetAppRate(route, headers)
}

func (r *RateLimit) IsRateLimited(rate *Rate) (bool, error) {
	if !r.Enabled {
		return false, ErrRateLimitingIsDisabled
	}

	if rate == nil {
		return false, nil
	}

	return r.store.IsRateLimited(rate)
}

func ParseHeaders(headers *http.Header, limitHeader string, countHeader string) (*Rate, error) {
	rateLimitHeader := headers.Get(limitHeader)

	if rateLimitHeader == "" {
		return nil, nil
	}

	// Obtaining the rate limit header values
	rates := strings.Split(rateLimitHeader, ",")

	// Obtaining the rate limit for seconds
	limit, seconds, err := getNumberPairs(rates[0])

	if err != nil {
		return nil, err
	}

	now := time.Now()

	rate := &Rate{
		Seconds: RateTiming{
			Time:   seconds,
			Limit:  limit,
			Count:  0,
			Expire: now.Add(time.Duration(seconds) * time.Second),
			Access: now,
		},
		Minutes: RateTiming{
			Time:   0,
			Limit:  0,
			Count:  0,
			Expire: time.Time{},
			Access: time.Time{},
		},
	}

	if len(rates) == 2 {
		// Obtaining the rate limit for minutes
		// Minutes are in seconds
		limit, seconds, err := getNumberPairs(rates[1])

		if err != nil {
			return nil, err
		}

		rate.Minutes.Expire = now.Add(time.Duration(seconds) * time.Second)
		rate.Minutes.Access = now
		rate.Minutes.Limit = limit
		rate.Minutes.Time = seconds
	}

	rateCountHeader := headers.Get(countHeader)

	if rateCountHeader == "" {
		return nil, nil
	}

	// Obtaining rate counts
	counts := strings.Split(rateCountHeader, ",")

	// Obtaining rate count for seconds
	// Discarding the limit from this header since we already have it
	current, _, err := getNumberPairs(counts[0])

	if err != nil {
		return nil, err
	}

	rate.Seconds.Count = current

	if len(counts) == 2 {
		// Obtaining rate count for minutes
		// Minutes are in seconds
		// Discarding the limit from this header since we already have it
		current, _, err := getNumberPairs(counts[1])

		if err != nil {
			return nil, err
		}

		rate.Minutes.Count = current
	}

	return rate, nil
}

func getNumberPairs(str string) (int, int, error) {
	numbers := strings.Split(str, ":")

	number, err := strconv.Atoi(numbers[0])

	if err != nil {
		return 0, 0, err
	}

	limit, err := strconv.Atoi(numbers[1])

	if err != nil {
		return 0, 0, err
	}

	return number, limit, nil
}
