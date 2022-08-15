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
	Route  map[interface{}]*Enpoints
}

type Enpoints struct {
	Endpoints map[string]*Methods
	AppRate   *Rate
}

type Methods struct {
	Methods map[string]*Rate
}

func (s *InternalRateStore) Get(route interface{}, endpointName string, methodName string) (*Rate, error) {
	if s.Route[route] == nil {
		s.Route[route] = &Enpoints{
			Endpoints: map[string]*Methods{},
			AppRate: &Rate{
				Seconds: RateTiming{},
				Minutes: RateTiming{},
			},
		}

		return nil, nil
	}

	if s.Route[route].Endpoints[endpointName] == nil {
		s.Route[route].Endpoints[endpointName] = &Methods{
			Methods: map[string]*Rate{},
		}

		return nil, nil
	}

	return s.Route[route].Endpoints[endpointName].Methods[methodName], nil
}

func (s *InternalRateStore) GetAppRate(route interface{}) (*Rate, error) {
	if s.Route[route] == nil {
		s.Route[route] = &Enpoints{
			Endpoints: map[string]*Methods{},
			AppRate: &Rate{
				Seconds: RateTiming{},
				Minutes: RateTiming{},
			},
		}

		return nil, nil
	}

	return s.Route[route].AppRate, nil
}

func (s *InternalRateStore) Set(route interface{}, endpointName string, methodName string, headers *http.Header) error {
	if s.Route[route] == nil {
		s.Route[route] = &Enpoints{
			Endpoints: map[string]*Methods{},
			AppRate: &Rate{
				Seconds: RateTiming{},
				Minutes: RateTiming{},
			},
		}
	}

	if s.Route[route].Endpoints[endpointName] == nil {
		s.Route[route].Endpoints[endpointName] = &Methods{
			Methods: map[string]*Rate{},
		}
	}

	rate, err := ParseHeaders(headers, MethodRateLimitHeader, MethodRateLimitCountHeader)

	if err != nil {
		return err
	}

	if rate == nil {
		return nil
	}

	if s.Route[route].Endpoints[endpointName].Methods[methodName] == nil {
		s.Route[route].Endpoints[endpointName].Methods[methodName] = rate

		return nil
	}

	s.UpdateInternalRateCount(s.Route[route].Endpoints[endpointName].Methods[methodName], rate)

	return nil
}

func (s *InternalRateStore) SetAppRate(route interface{}, headers *http.Header) error {
	if s.Route[route] == nil {
		s.Route[route] = &Enpoints{
			Endpoints: map[string]*Methods{},
			AppRate: &Rate{
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

	if s.Route[route].AppRate.Seconds.Limit == 0 {
		s.Route[route].AppRate = rate

		return nil
	}

	s.UpdateInternalRateCount(s.Route[route].AppRate, rate)

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

func (s *InternalRateStore) UpdateInternalRateCount(old *Rate, new *Rate) {
	old.Seconds.Count = new.Seconds.Count
	old.Seconds.Access = new.Seconds.Access

	if old.Seconds.Access.After(old.Seconds.Expire) {
		old.Seconds.Expire = new.Seconds.Expire
	}

	old.Minutes.Count = new.Minutes.Count
	old.Minutes.Access = new.Minutes.Access

	if old.Minutes.Access.After(old.Minutes.Expire) {
		old.Minutes.Expire = new.Minutes.Expire
	}
}
