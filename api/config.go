package api

import (
	"net/http"
	"time"

	"github.com/Kyagara/equinox/cache"
	"github.com/rs/zerolog"
)

// Configuration for the Equinox client.
type EquinoxConfig struct {
	// Riot API Key.
	Key string
	// Zerolog log level.
	LogLevel zerolog.Level
	// http.Client used internally.
	HTTPClient *http.Client
	// Retry a request `n` times if it returns a 429 status code.
	Retry int
	// The cache used to store all GET requests done by the client.
	Cache *cache.Cache
}

func (c EquinoxConfig) MarshalZerologObject(encoder *zerolog.Event) {
	encoder.Int("retry-if-429", c.Retry).Dur("http-client-timeout", c.HTTPClient.Timeout)
	if c.Cache.TTL > 0 {
		cache := CacheConfig{Store: string(c.Cache.StoreType), TTL: c.Cache.TTL}
		encoder.Object("cache", cache)
	}
}

type CacheConfig struct {
	TTL   time.Duration
	Store string
}

func (c CacheConfig) MarshalZerologObject(encoder *zerolog.Event) {
	encoder.Str("store", c.Store).Dur("cache-ttl", c.TTL)
}
