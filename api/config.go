package api

import (
	"time"

	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/rate_limit"
	"go.uber.org/zap/zapcore"
)

// Configuration for the Equinox client.
type EquinoxConfig struct {
	// Riot API Key.
	Key string
	// Cluster name, using the nearest cluster to you is recommended.
	Cluster Cluster
	// Log level, api.NopLevel disables logging.
	LogLevel LogLevel
	// Timeout for the internal http.Client in seconds, 0 disables the timeout.
	Timeout int
	// Allows retrying a request if it returns a 429 status code.
	Retry bool
	// The cache used to store all GET requests done by the client.
	Cache *cache.Cache
	// The rate limit store.
	RateLimit *rate_limit.RateLimit
}

func (c *EquinoxConfig) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddBool("retry-if-429", c.Retry)
	encoder.AddInt("http-client-timeout", c.Timeout)

	if c.RateLimit.Enabled {
		rate := RateConfig{Store: string(c.RateLimit.StoreType)}

		err := encoder.AddObject("rate-limit", rate)

		if err != nil {
			return err
		}
	}

	if c.Cache.TTL > 0 {
		cache := CacheConfig{Store: string(c.Cache.StoreType), TTL: c.Cache.TTL}

		err := encoder.AddObject("cache", cache)

		if err != nil {
			return err
		}
	}

	return nil
}

type CacheConfig struct {
	TTL   time.Duration
	Store string
}

func (c CacheConfig) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddString("store", c.Store)
	encoder.AddDuration("cache-ttl", c.TTL)

	return nil
}

type RateConfig struct {
	Store string
}

func (c RateConfig) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddString("store", c.Store)

	return nil
}
