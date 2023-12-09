package equinox

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
	Cache   *cache.Cache
	DDragon *ddragon.Client
	CDragon *cdragon.Client
	Riot    *riot.Client
	LOL     *lol.Client
	TFT     *tft.Client
	VAL     *val.Client
	LOR     *lor.Client
}

// Creates a new Equinox client with the default configuration
func NewClient(key string) (*Equinox, error) {
	config, err := DefaultConfig(key)
	if err != nil {
		return nil, err
	}
	return NewClientWithConfig(config)
}

// Creates a new Equinox client using a custom configuration.
func NewClientWithConfig(config api.EquinoxConfig) (*Equinox, error) {
	client, err := internal.NewInternalClient(config)
	if err != nil {
		return nil, err
	}
	equinox := &Equinox{
		Cache:   config.Cache,
		DDragon: ddragon.NewDDragonClient(client),
		CDragon: cdragon.NewCDragonClient(client),
		Riot:    riot.NewRiotClient(client),
		LOL:     lol.NewLOLClient(client),
		TFT:     tft.NewTFTClient(client),
		VAL:     val.NewVALClient(client),
		LOR:     lor.NewLORClient(client),
	}
	return equinox, nil
}

// Returns the default Equinox config with a provided key.
//
//   - `LogLevel`   : zerolog.WarnLevel
//   - `Retries`    : Retry 3 times
//   - `HTTPClient` : http.Client with timeout of 15 seconds
//   - `Cache`      : BigCache with TTL of 4 minutes
//   - `RateLimit`  : Internal rate limit
func DefaultConfig(key string) (api.EquinoxConfig, error) {
	ctx := context.Background()
	cache, err := cache.NewBigCache(ctx, bigcache.DefaultConfig(4*time.Minute))
	if err != nil {
		return api.EquinoxConfig{}, err
	}
	config := api.EquinoxConfig{
		Key:      key,
		LogLevel: zerolog.WarnLevel,
		HTTPClient: &http.Client{
			Timeout: 15 * time.Second,
		},
		Retries:   3,
		Cache:     cache,
		RateLimit: ratelimit.NewInternalRateLimit(),
	}
	return config, nil
}
