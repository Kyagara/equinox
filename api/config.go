package api

import (
	"net/http"

	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/rs/zerolog"
)

// Configuration for the Equinox client.
type EquinoxConfig struct {
	// http.Client used internally.
	HTTPClient *http.Client
	// The cache used to store all GET requests done by the client.
	Cache *cache.Cache
	// The type of rate limiter to use, only disable it if you know what you're doing.
	RateLimit *ratelimit.RateLimit
	// Riot API Key.
	Key string
	// Maximum amount of times to retry a request on error.
	Retries int
	// Zerolog log level.
	LogLevel zerolog.Level
}

func (c EquinoxConfig) MarshalZerologObject(encoder *zerolog.Event) {
	if c.Retries > 0 {
		encoder.Int("max_retries", c.Retries)
	}
	if c.HTTPClient.Timeout > 0 {
		encoder.Dur("http_client_timeout", c.HTTPClient.Timeout)
	}
	if c.Cache.TTL != 0 {
		encoder.Str("cache_store", string(c.Cache.StoreType)).Dur("cache_ttl", c.Cache.TTL)
	}
}
