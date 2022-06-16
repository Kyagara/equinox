package internal

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type RateLimit struct {
	endpoints map[string]*Methods
	appRate   *Rate
}

type Methods struct {
	methods map[string]*Rate
}

type Rate struct {
	Seconds *RateTiming
	Minutes *RateTiming
}

func NewRateLimit() *RateLimit {
	return &RateLimit{
		endpoints: map[string]*Methods{},
		appRate: &Rate{
			Seconds: &RateTiming{},
			Minutes: &RateTiming{},
		},
	}
}

type RateTiming struct {
	// The amount in seconds the rate limit should reset
	Time int
	// Maximum amount of requests in n seconds
	Limit int
	// Current count of requests made in n seconds
	Count int

	Ticker *time.Ticker
	Mutex  *sync.Mutex
}

func (r *RateTiming) tick() {
	for range r.Ticker.C {
		r.Mutex.Lock()

		r.Count = 0

		r.Mutex.Unlock()
	}
}

// Checks if the *Rate is currently rate limited
func (r *RateLimit) IsRateLimited(rate *Rate) bool {
	if rate.Seconds.Limit == 0 {
		return false
	}

	rate.Seconds.Mutex.Lock()
	defer rate.Seconds.Mutex.Unlock()

	if rate.Seconds.Count >= rate.Seconds.Limit {
		return true
	}

	rate.Minutes.Mutex.Lock()
	defer rate.Minutes.Mutex.Unlock()

	return rate.Minutes.Count >= rate.Minutes.Limit
}

func (r *RateLimit) Get(endpointName string, methodName string) *Rate {
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

func (r *RateLimit) Set(endpointName string, methodName string, rate *Rate) {
	endpoint := r.endpoints[endpointName]

	if endpoint == nil {
		r.endpoints[endpointName] = &Methods{
			methods: map[string]*Rate{},
		}

		if r.endpoints[endpointName].methods[methodName] == nil {
			r.endpoints[endpointName].methods[methodName] = rate

			go rate.Seconds.tick()

			if rate.Minutes.Limit != 0 {
				go rate.Minutes.tick()
			}
		}

		return
	}

	methodRate := endpoint.methods[methodName]

	// Update seconds count
	updateRateCount(methodRate.Seconds, rate.Seconds)

	// Update minutes count
	updateRateCount(methodRate.Minutes, rate.Minutes)
}

func (r *RateLimit) SetAppRate(rate *Rate) {
	if r.appRate.Seconds.Limit == 0 {
		r.appRate = rate

		go rate.Seconds.tick()

		if rate.Minutes.Limit != 0 {
			go rate.Minutes.tick()
		}

		return
	}

	// Update seconds count
	updateRateCount(r.appRate.Seconds, rate.Seconds)

	// Update minutes count
	updateRateCount(r.appRate.Minutes, rate.Minutes)
}

func ParseHeaders(headers http.Header, limitHeader string, countHeader string) *Rate {
	rate := &Rate{
		Seconds: &RateTiming{},
		Minutes: &RateTiming{},
	}

	rateLimit := headers.Get(limitHeader)

	if rateLimit == "" {
		return nil
	}

	// Obtaining Rate-Limit header values
	rates := strings.Split(rateLimit, ",")

	// Obtaining rate limit for seconds
	limit, seconds := getNumberPairs(rates[0])

	rate.Seconds = getRateTiming(limit, seconds)

	if len(rates) == 2 {
		// Obtaining rate limit for minutes
		limit, seconds = getNumberPairs(rates[1])

		rate.Minutes = getRateTiming(limit, seconds)
	}

	rateCount := headers.Get(countHeader)

	if rateCount == "" {
		return nil
	}

	// Obtaining rate counts
	counts := strings.Split(rateCount, ",")

	// Obtaining rate count for seconds
	current, _ := getNumberPairs(counts[0])

	rate.Seconds.Count = current

	if len(counts) == 2 {
		// Obtaining rate count for seconds
		current, _ = getNumberPairs(counts[1])

		rate.Minutes.Count = current
	}

	return rate
}

func getNumberPairs(str string) (int, int) {
	numbers := strings.Split(str, ":")

	number, _ := strconv.Atoi(numbers[0])

	limit, _ := strconv.Atoi(numbers[1])

	return number, limit
}

func getRateTiming(limit int, seconds int) *RateTiming {
	timing := &RateTiming{
		Time:   seconds,
		Limit:  limit,
		Count:  0,
		Ticker: &time.Ticker{},
		Mutex:  &sync.Mutex{},
	}

	timing.Ticker = time.NewTicker(time.Duration(seconds) * time.Second)

	return timing
}

func updateRateCount(old *RateTiming, new *RateTiming) {
	old.Mutex.Lock()
	defer old.Mutex.Unlock()

	old.Count = new.Count
}
