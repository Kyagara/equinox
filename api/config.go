package api

import (
	"net/http"
	"time"

	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/rs/zerolog"
)

// Configuration for the equinox client.
type EquinoxConfig struct {
	// http.Client used internally.
	HTTPClient *http.Client
	// The cache used to store all GET requests done by the client.
	Cache *cache.Cache
	// The type of rate limiter to use, only disable it if you know what you're doing.
	RateLimit *ratelimit.RateLimit
	// Riot API Key.
	Key string
	// Configuration for the logger.
	Logger Logger
	// Configuration for retries.
	Retry Retry
}

func (c EquinoxConfig) MarshalZerologObject(encoder *zerolog.Event) {
	if c.Retry.MaxRetries > 0 {
		encoder.Object("retry", c.Retry)
	}
	if c.Cache.Enabled {
		encoder.Object("cache", c.Cache)
	}
	if c.RateLimit.Enabled {
		encoder.Object("ratelimit", c.RateLimit)
	}
}

// Retry configuration.
type Retry struct {
	// Retries are exponential, 0 disables retries
	MaxRetries int
	// In milliseconds
	Jitter time.Duration
}

func (r Retry) MarshalZerologObject(encoder *zerolog.Event) {
	if r.MaxRetries > 0 {
		encoder.Int("max_retries", r.MaxRetries)
	}
	if r.Jitter > 0 {
		encoder.Dur("jitter", r.Jitter)
	}
}

// Logger configuration.
type Logger struct {
	TimeFieldFormat string
	Level           zerolog.Level
	// Enables prettified logging
	Pretty bool
	// Prints the timestamp
	EnableTimestamp bool
	// Adds the equinox configuration to logs
	EnableConfigLogging bool
}
