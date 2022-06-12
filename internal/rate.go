package internal

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type RateLimiter struct {
	endpoints map[string]*EndpointMethods
	appRate   *Rate
}

type EndpointMethods struct {
	methods map[string]*Rate
}

type Rate struct {
	Seconds int
	// Maximum amount of requests in n seconds
	SecondsLimit int
	// Current count of requests in n seconds
	SecondsCount int
	Ticker       *time.Ticker
	Mutex        *sync.Mutex
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		endpoints: map[string]*EndpointMethods{},
		appRate: &Rate{
			Seconds:      0,
			SecondsLimit: 0,
			SecondsCount: 0,
			Ticker:       &time.Ticker{},
			Mutex:        &sync.Mutex{},
		},
	}
}

func (r *Rate) tick() {
	for range r.Ticker.C {
		r.Mutex.Lock()

		r.SecondsCount = 0

		r.Mutex.Unlock()
	}
}

func (r *RateLimiter) Get(endpointName string, methodName string) *Rate {
	endpoint := r.endpoints[endpointName]

	if endpoint != nil {
		rate := endpoint.methods[methodName]

		if rate != nil {
			return rate
		}

		return nil
	}

	return nil
}

func (r *RateLimiter) Set(endpointName string, methodName string, rate *Rate) {
	endpoint := r.endpoints[endpointName]

	if endpoint == nil {
		r.endpoints[endpointName] = &EndpointMethods{
			methods: map[string]*Rate{},
		}

		if r.endpoints[endpointName].methods[methodName] == nil {
			r.endpoints[endpointName].methods[methodName] = rate

			go rate.tick()
		}

		return
	}

	methodRate := endpoint.methods[methodName]

	methodRate.Mutex.Lock()

	methodRate.SecondsCount = rate.SecondsCount

	methodRate.Mutex.Unlock()
}

func (r *RateLimiter) SetAppRate(rate *Rate) {
	if r.appRate.Seconds == 0 {
		r.appRate = rate

		go rate.tick()

		return
	}

	r.appRate.Mutex.Lock()

	r.appRate.SecondsCount = rate.SecondsCount

	r.appRate.Mutex.Unlock()
}

func (r *RateLimiter) ParseHeaders(headers http.Header, limitHeader string, countHeader string) *Rate {
	rate := &Rate{
		Seconds:      0,
		SecondsLimit: 0,
		SecondsCount: 0,
		Ticker:       &time.Ticker{},
		Mutex:        &sync.Mutex{},
	}

	rateLimit := headers.Get(limitHeader)

	if rateLimit == "" {
		return nil
	}

	// Obtaining Rate-Limit header values
	rates := strings.Split(rateLimit, ",")

	// Obtaining rate limit for seconds
	max, inSeconds := getNumberPairs(rates[0])

	rate.Seconds = inSeconds

	rate.SecondsLimit = max

	rateCount := headers.Get(countHeader)

	if rateCount == "" {
		return nil
	}

	// Obtaining rate counts
	counts := strings.Split(rateCount, ",")

	// Obtaining a specific rate count
	current, _ := getNumberPairs(counts[0])

	rate.SecondsCount = current

	rate.Ticker = time.NewTicker(time.Duration(rate.Seconds) * time.Second)

	return rate
}

func getNumberPairs(str string) (int, int) {
	numbers := strings.Split(str, ":")

	number, _ := strconv.Atoi(numbers[0])

	limit, _ := strconv.Atoi(numbers[1])

	return number, limit
}
