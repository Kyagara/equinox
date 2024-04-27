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
	// Riot API Key.
	Key string
	// http.Client used internally.
	HTTPClient *http.Client
	// The Cache store, storing all GET requests done by the client, optional.
	Cache *cache.Cache
	// The RateLimit store, only disable it if you know what you're doing.
	RateLimit *ratelimit.RateLimit
	// Configuration object for the Logger.
	Logger Logger
	// Configuration object for Retry.
	Retry Retry
}

func (c EquinoxConfig) MarshalZerologObject(encoder *zerolog.Event) {
	if c.Retry.MaxRetries > 0 {
		encoder.Object("retry", c.Retry)
	}
	if c.Cache.TTL != 0 {
		encoder.Object("cache", c.Cache)
	}
	if c.RateLimit.Enabled {
		encoder.Object("ratelimit", c.RateLimit)
	}
}

// Retry configuration object.
//
// Retries have a exponential backoff mechanism.
type Retry struct {
	// Maximum number of retries, 0 disables retries.
	MaxRetries int
	// Jitter, in milliseconds, added to the retry interval.
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

// Logger configuration object.
type Logger struct {
	TimeFieldFormat string
	Level           zerolog.Level
	// Enables prettified logging.
	Pretty bool
	// Prints the timestamp.
	EnableTimestamp bool
	// Logs configurations objects from the client, includes Cache and RateLimit.
	EnableConfigLogging bool
}
