package api

import (
	"time"

	"github.com/Kyagara/equinox/cache"
	"go.uber.org/zap/zapcore"
)

// Configuration for the Equinox client.
type EquinoxConfig struct {
	// Riot API Key.
	Key string
	// Log level, api.NOP_LOG_LEVEL disables logging.
	LogLevel LogLevel
	// Timeout for the internal http.Client in seconds, 0 disables the timeout.
	Timeout int
	// Allows retrying a request if it returns a 429 status code.
	Retry bool
	// The cache used to store all GET requests done by the client.
	Cache *cache.Cache
}

func (c *EquinoxConfig) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddBool("retry-if-429", c.Retry)
	encoder.AddInt("http-client-timeout", c.Timeout)
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
