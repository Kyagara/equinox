package api

import (
	"net/http"
	"time"

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
	// Configuration for retries.
	Retry Retry
	// Zerolog log level.
	LogLevel zerolog.Level
}

type Retry struct {
	// 0 disables retries
	MaxRetries int
	// In milliseconds
	Jitter time.Duration
}

func (c EquinoxConfig) MarshalZerologObject(encoder *zerolog.Event) {
	if c.Retry.MaxRetries > 0 {
		encoder.Object("retry", c.Retry)
	}
	if c.HTTPClient.Timeout > 0 {
		encoder.Dur("http_client_timeout", c.HTTPClient.Timeout)
	}
	if c.Cache.TTL != 0 {
		encoder.Str("cache_store", string(c.Cache.StoreType)).Dur("cache_ttl", c.Cache.TTL)
	}
}

func (r Retry) MarshalZerologObject(encoder *zerolog.Event) {
	if r.MaxRetries > 0 {
		encoder.Int("max_retries", r.MaxRetries)
	}
	if r.Jitter > 0 {
		encoder.Dur("jitter", r.Jitter)
	}
}
