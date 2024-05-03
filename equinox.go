package equinox

//go:generate go run ./codegen

import (
	"context"
	"net/http"
	"time"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/clients/lor"
	"github.com/Kyagara/equinox/clients/riot"
	"github.com/Kyagara/equinox/clients/tft"
	"github.com/Kyagara/equinox/clients/val"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/allegro/bigcache/v3"
	"github.com/rs/zerolog"
)

type Equinox struct {
	Internal  *internal.Client
	Cache     *cache.Cache
	RateLimit *ratelimit.RateLimit
	Riot      *riot.Client
	LOL       *lol.Client
	TFT       *tft.Client
	VAL       *val.Client
	LOR       *lor.Client
}

// Creates a new equinox client with the default configuration.
//
//   - Key  	  : The provided Riot API key.
//   - HTTPClient : Default http.Client.
//   - Cache	  : BigCache with an eviction time of 4 minutes.
//   - RateLimit  : InternalRateLimit with a limit usage factor of 0.99 and interval overhead of 1 second.
//   - Logger	  : Logger with zerolog.WarnLevel. Will log if rate limited or when retrying a request.
//   - Retry	  : Will retry a request a maximum of 3 times and jitter of 500 milliseconds.
func NewClient(key string) (*Equinox, error) {
	config := DefaultConfig(key)
	cache, err := DefaultCache()
	if err != nil {
		return nil, err
	}
	rateLimit := DefaultRateLimit()
	return NewCustomClient(config, nil, cache, rateLimit)
}

// Creates a new equinox client using a custom configuration.
//
//   - Config     : The equinox config.
//   - HTTPClient : Can be nil, will default to a default http.Client.
//   - Cache      : Can be nil.
//   - RateLimit  : Can be nil, only disable it if you know what you're doing.
func NewCustomClient(config api.EquinoxConfig, httpClient *http.Client, cache *cache.Cache, rateLimit *ratelimit.RateLimit) (*Equinox, error) {
	client, err := internal.NewInternalClient(config, httpClient, cache, rateLimit)
	if err != nil {
		return nil, err
	}
	equinox := &Equinox{
		Internal:  client,
		Cache:     cache,
		RateLimit: rateLimit,
		Riot:      riot.NewRiotClient(client),
		LOL:       lol.NewLOLClient(client),
		TFT:       tft.NewTFTClient(client),
		VAL:       val.NewVALClient(client),
		LOR:       lor.NewLORClient(client),
	}
	return equinox, nil
}

// Returns the default equinox config.
//
// Logger with zerolog.WarnLevel. Retry with a limit of 3 and jitter of 500 milliseconds.
func DefaultConfig(key string) api.EquinoxConfig {
	return api.EquinoxConfig{
		Key:    key,
		Retry:  DefaultRetry(),
		Logger: DefaultLogger(),
	}
}

// Returns the default Cache.
//
// BigCache with an eviction time of 4 minutes.
func DefaultCache() (*cache.Cache, error) {
	ctx := context.Background()
	cacheConfig := bigcache.DefaultConfig(4 * time.Minute)
	cacheConfig.Verbose = false
	return cache.NewBigCache(ctx, cacheConfig)
}

// Returns the default RateLimit.
//
//   - StoreType        : InternalRateLimit
//   - LimitUsageFactor : 0.99
//   - IntervalOverhead : 1 second
func DefaultRateLimit() *ratelimit.RateLimit {
	return ratelimit.NewInternalRateLimit(0.99, time.Second)
}

// Returns the default Retry config.
//
//   - MaxRetries : 3
//   - Jitter     : 500 milliseconds
func DefaultRetry() api.Retry {
	return api.Retry{MaxRetries: 3, Jitter: 500 * time.Millisecond}
}

// Returns the default Logger config.
//
//   - Level		              : zerolog.WarnLevel
//   - Pretty		              : false
//   - TimeFieldFormat            : zerolog.TimeFormatUnix
//   - EnableTimestamp            : true
//   - EnableConfigurationLogging : false
func DefaultLogger() api.Logger {
	return api.Logger{
		Level:                      zerolog.WarnLevel,
		Pretty:                     false,
		TimeFieldFormat:            zerolog.TimeFormatUnix,
		EnableTimestamp:            true,
		EnableConfigurationLogging: false,
	}
}
