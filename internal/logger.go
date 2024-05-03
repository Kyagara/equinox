package internal

import (
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/rs/zerolog"
)

type loggers struct {
	// The main logger the others will be created from.
	main    zerolog.Logger
	methods map[string]zerolog.Logger
	mutex   sync.Mutex
}

type configuration struct {
	cache     cacheConfiguration
	rateLimit rateLimitConfiguration
	retry     retryConfiguration
}

type cacheConfiguration struct {
	store string
	ttl   time.Duration
}

type rateLimitConfiguration struct {
	storeType        string
	intervalOverhead time.Duration
	usageFactor      float64
	enabled          bool
}

type retryConfiguration struct {
	maxRetries int
	jitter     time.Duration
}

func (c configuration) MarshalZerologObject(encoder *zerolog.Event) {
	if c.cache.ttl > 0 {
		encoder.Object("cache", c.cache)
	}
	if c.rateLimit.enabled {
		encoder.Object("rate_limit", c.rateLimit)
	}
	if c.retry.maxRetries > 0 {
		encoder.Object("retry", c.retry)
	}
}

func (c cacheConfiguration) MarshalZerologObject(encoder *zerolog.Event) {
	encoder.Dur("ttl", c.ttl).Str("store", c.store)
}

func (r rateLimitConfiguration) MarshalZerologObject(encoder *zerolog.Event) {
	encoder.Str("store", r.storeType).Dur("interval_overhead", r.intervalOverhead).Float64("limit_usage_factor", r.usageFactor)
}

func (r retryConfiguration) MarshalZerologObject(encoder *zerolog.Event) {
	encoder.Int("max_retries", r.maxRetries).Dur("jitter", r.jitter)
}

// Creates a new zerolog.Logger from an EquinoxConfig.
func NewLogger(config api.EquinoxConfig, cache *cache.Cache, ratelimit *ratelimit.RateLimit) zerolog.Logger {
	if config == (api.EquinoxConfig{}) || config.Logger.Level == zerolog.Disabled {
		return zerolog.Nop()
	}

	zerolog.TimeFieldFormat = config.Logger.TimeFieldFormat

	var logger zerolog.Logger

	if config.Logger.Pretty {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).Level(config.Logger.Level)
	} else {
		logger = zerolog.New(os.Stderr).Level(config.Logger.Level)
	}

	if config.Logger.EnableTimestamp {
		logger = logger.With().Timestamp().Logger()
	}

	var equinoxConfig configuration
	emptyConfig := true

	if cache != nil {
		equinoxConfig.cache = cacheConfiguration{
			ttl:   cache.TTL,
			store: string(cache.StoreType),
		}
		emptyConfig = false
	}

	if ratelimit != nil {
		equinoxConfig.rateLimit = rateLimitConfiguration{
			storeType:        string(ratelimit.StoreType),
			intervalOverhead: ratelimit.IntervalOverhead,
			usageFactor:      ratelimit.LimitUsageFactor,
			enabled:          ratelimit.Enabled,
		}
		emptyConfig = false
	}

	if config.Retry.MaxRetries > 0 {
		equinoxConfig.retry = retryConfiguration{
			maxRetries: config.Retry.MaxRetries,
			jitter:     config.Retry.Jitter,
		}
		emptyConfig = false
	}

	if config.Logger.EnableConfigurationLogging && !emptyConfig {
		logger = logger.With().Object("equinox", equinoxConfig).Logger()
	}

	return logger
}

// Used to create/retrieve the zerolog.Logger for the specified endpoint method.
func (c *Client) Logger(id string) zerolog.Logger {
	c.loggers.mutex.Lock()
	defer c.loggers.mutex.Unlock()

	if logger, ok := c.loggers.methods[id]; ok {
		return logger
	}

	names := strings.Split(id, "_")
	logger := c.loggers.main.With().Str("client", names[0]).Str("endpoint", names[1]).Str("method", names[2]).Logger()
	c.loggers.methods[id] = logger
	return logger
}
