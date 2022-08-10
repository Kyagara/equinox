package rate_limit

import (
	"net/http"
	"time"
)

type InternalRateLimitClient interface {
	Get(route interface{}, endpointName string, methodName string) (*Rate, error)
	GetAppRate(route interface{}) (*Rate, error)
	Set(route interface{}, endpointName string, methodName string, headers *http.Header) error
	SetAppRate(route interface{}, headers *http.Header) error
	ParseHeaders(headers *http.Header, limitHeader string, countHeader string) (*Rate, error)
	IsRateLimited(rate *Rate) (bool, error)
}

type InternalRateStore struct {
	client InternalRateLimitClient
	route  map[interface{}]*Enpoints
}

type Enpoints struct {
	endpoints map[string]*Methods
	appRate   *Rate
}

type Methods struct {
	methods map[string]*Rate
}

func (s *InternalRateStore) Get(route interface{}, endpointName string, methodName string) (*Rate, error) {
	if s.route[route] == nil {
		s.route[route] = &Enpoints{
			endpoints: map[string]*Methods{},
			appRate: &Rate{
				Seconds: RateTiming{},
				Minutes: RateTiming{},
			},
		}

		return nil, nil
	}

	if s.route[route].endpoints[endpointName] == nil {
		s.route[route].endpoints[endpointName] = &Methods{
			methods: map[string]*Rate{},
		}

		return nil, nil
	}

	return s.route[route].endpoints[endpointName].methods[methodName], nil
}

func (s *InternalRateStore) GetAppRate(route interface{}) (*Rate, error) {
	if s.route[route] == nil {
		s.route[route] = &Enpoints{
			endpoints: map[string]*Methods{},
			appRate: &Rate{
				Seconds: RateTiming{},
				Minutes: RateTiming{},
			},
		}

		return nil, nil
	}

	return s.route[route].appRate, nil
}

func (s *InternalRateStore) Set(route interface{}, endpointName string, methodName string, headers *http.Header) error {
	if s.route[route] == nil {
		s.route[route] = &Enpoints{
			endpoints: map[string]*Methods{},
			appRate: &Rate{
				Seconds: RateTiming{},
				Minutes: RateTiming{},
			},
		}
	}

	if s.route[route].endpoints[endpointName] == nil {
		s.route[route].endpoints[endpointName] = &Methods{
			methods: map[string]*Rate{},
		}
	}

	rate, err := ParseHeaders(headers, MethodRateLimitHeader, MethodRateLimitCountHeader)

	if err != nil {
		return err
	}

	if rate == nil {
		return nil
	}

	if s.route[route].endpoints[endpointName].methods[methodName] == nil {
		s.route[route].endpoints[endpointName].methods[methodName] = rate

		return nil
	}

	s.updateInternalRateCount(s.route[route].endpoints[endpointName].methods[methodName], rate)

	return nil
}

func (s *InternalRateStore) SetAppRate(route interface{}, headers *http.Header) error {
	if s.route[route] == nil {
		s.route[route] = &Enpoints{
			endpoints: map[string]*Methods{},
			appRate: &Rate{
				Seconds: RateTiming{},
				Minutes: RateTiming{},
			},
		}
	}

	rate, err := ParseHeaders(headers, AppRateLimitHeader, AppRateLimitCountHeader)

	if err != nil {
		return err
	}

	if rate == nil {
		return nil
	}

	if s.route[route].appRate.Seconds.Limit == 0 {
		s.route[route].appRate = rate

		return nil
	}

	s.updateInternalRateCount(s.route[route].appRate, rate)

	return nil
}

func (s *InternalRateStore) IsRateLimited(rate *Rate) (bool, error) {
	now := time.Now()

	rate.Seconds.Access = now

	if rate.Seconds.Access.After(rate.Seconds.Expire) {
		rate.Seconds.Count = 0
		rate.Seconds.Expire = now.Add(time.Duration(rate.Seconds.Time) * time.Second)
		return false, nil
	}

	if rate.Seconds.Limit == 0 {
		return false, nil
	}

	if rate.Seconds.Count >= rate.Seconds.Limit {
		return true, nil
	}

	rate.Minutes.Access = now

	if rate.Minutes.Access.After(rate.Minutes.Expire) {
		rate.Minutes.Count = 0
		rate.Minutes.Expire = now.Add(time.Duration(rate.Minutes.Time) * time.Second)
		return false, nil
	}

	if rate.Minutes.Limit == 0 {
		return false, nil
	}

	if rate.Minutes.Time == 0 {
		return false, nil
	}

	return rate.Minutes.Count >= rate.Minutes.Limit, nil
}

func (s *InternalRateStore) updateInternalRateCount(old *Rate, new *Rate) {
	now := time.Now()

	old.Seconds.Count = new.Seconds.Count
	old.Seconds.Access = now
	old.Seconds.Expire = now.Add(time.Duration(new.Seconds.Time) * time.Second)

	old.Minutes.Count = new.Minutes.Count
	old.Minutes.Access = now
	old.Minutes.Expire = now.Add(time.Duration(new.Minutes.Time) * time.Second)
}
