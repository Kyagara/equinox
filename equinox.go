package equinox

//go:generate go run ./codegen

import (
	"context"
	"net/http"
	"time"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/clients/cdragon"
	"github.com/Kyagara/equinox/clients/ddragon"
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
	Internal *internal.Client
	Cache    *cache.Cache
	DDragon  *ddragon.Client
	CDragon  *cdragon.Client
	Riot     *riot.Client
	LOL      *lol.Client
	TFT      *tft.Client
	VAL      *val.Client
	LOR      *lor.Client
}

// Creates a new equinox client with the default configuration
func NewClient(key string) (*Equinox, error) {
	config, err := DefaultConfig(key)
	if err != nil {
		return nil, err
	}
	return NewClientWithConfig(config), nil
}

// Creates a new equinox client using a custom configuration.
func NewClientWithConfig(config api.EquinoxConfig) *Equinox {
	client := internal.NewInternalClient(config)
	equinox := &Equinox{
		Internal: client,
		Cache:    config.Cache,
		DDragon:  ddragon.NewDDragonClient(client),
		CDragon:  cdragon.NewCDragonClient(client),
		Riot:     riot.NewRiotClient(client),
		LOL:      lol.NewLOLClient(client),
		TFT:      tft.NewTFTClient(client),
		VAL:      val.NewVALClient(client),
		LOR:      lor.NewLORClient(client),
	}
	return equinox
}

// Returns the default equinox config with a provided key.
//
//   - Key  	  : The provided Riot API key.
//   - HTTPClient : http.Client with a timeout of 15 seconds.
//   - Cache	  : BigCache with an eviction time of 4 minutes.
//   - RateLimit  : Internal rate limiter with a limit usage factor of 1.0 and interval overhead of 1 second.
//   - Logger	  : api.Logger object with zerolog.WarnLevel. Will log if rate limited or when retrying a request (before waiting).
//   - Retry	  : api.Retry object with a limit of 3 and jitter of 500 milliseconds.
func DefaultConfig(key string) (api.EquinoxConfig, error) {
	ctx := context.Background()
	cacheConfig := bigcache.DefaultConfig(4 * time.Minute)
	cacheConfig.Verbose = false
	cache, err := cache.NewBigCache(ctx, cacheConfig)
	if err != nil {
		return api.EquinoxConfig{}, err
	}
	config := api.EquinoxConfig{
		Key: key,
		HTTPClient: &http.Client{
			Timeout: 15 * time.Second,
		},
		Cache:     cache,
		RateLimit: ratelimit.NewInternalRateLimit(0.99, time.Second),
		Retry:     DefaultRetry(),
		Logger:    DefaultLogger(),
	}
	return config, nil
}

// Returns the default retry config
//
//   - MaxRetries : 3
//   - Jitter     : 500 milliseconds
func DefaultRetry() api.Retry {
	return api.Retry{MaxRetries: 3, Jitter: 500 * time.Millisecond}
}

// Returns the default logger config
//
//   - Level			   : zerolog.WarnLevel
//   - Pretty			   : false
//   - TimeFieldFormat	   : zerolog.TimeFormatUnix
//   - EnableConfigLogging : true
//   - EnableTimestamp	   : true
func DefaultLogger() api.Logger {
	return api.Logger{
		Level:               zerolog.WarnLevel,
		Pretty:              false,
		TimeFieldFormat:     zerolog.TimeFormatUnix,
		EnableConfigLogging: true,
		EnableTimestamp:     true,
	}
}
