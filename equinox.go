package equinox

import (
	"context"
	"fmt"
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
	"github.com/allegro/bigcache/v3"
)

type Equinox struct {
	Cache   *cache.Cache
	DDragon *ddragon.DDragonClient
	CDragon *cdragon.CDragonClient
	Riot    *riot.RiotClient
	LOL     *lol.LOLClient
	TFT     *tft.TFTClient
	LOR     *lor.LORClient
	VAL     *val.VALClient
}

// Returns the default Equinox config with a provided key.
//
//   - `Cluster`    : api.AmericasCluster
//   - `LogLevel`   : api.NopLevel
//   - `Timeout`    : 15 Seconds
//   - `Retry`      : true
//   - `Cache`      : BigCache with TTL of 4 minutes
func DefaultConfig(key string) (*api.EquinoxConfig, error) {
	ctx := context.Background()
	cache, err := cache.NewBigCache(ctx, bigcache.DefaultConfig(4*time.Minute))
	if err != nil {
		return nil, err
	}

	config := &api.EquinoxConfig{
		Key:      key,
		Cluster:  api.AmericasCluster,
		LogLevel: api.NopLevel,
		Timeout:  15,
		Retry:    true,
		Cache:    cache,
	}

	return config, nil
}

// Creates a new Equinox client with a default configuration
//
//   - `Cluster`    : api.AmericasCluster
//   - `LogLevel`   : api.NopLevel
//   - `Timeout`    : 15 Seconds
//   - `Retry`      : true
//   - `Cache`      : BigCache with TTL of 4 minutes
func NewClient(key string) (*Equinox, error) {
	config, err := DefaultConfig(key)
	if err != nil {
		return nil, err
	}

	client, err := internal.NewInternalClient(config)
	if err != nil {
		return nil, err
	}

	if key == "" {
		client.GetInternalLogger().Warn("API key was not provided, only Data Dragon methods will be enabled.")
		client.IsDataDragonOnly = true
	}

	equinox := &Equinox{
		Cache:   config.Cache,
		DDragon: ddragon.NewDDragonClient(client),
		CDragon: cdragon.NewCDragonClient(client),
		Riot:    riot.NewRiotClient(client),
		LOL:     lol.NewLOLClient(client),
		TFT:     tft.NewTFTClient(client),
		LOR:     lor.NewLORClient(client),
		VAL:     val.NewVALClient(client),
	}

	return equinox, nil
}

// Creates a new Equinox client using a custom configuration.
//
// If you don't specify a Timeout this will disable the timeout for the http.Client.
func NewClientWithConfig(config *api.EquinoxConfig) (*Equinox, error) {
	if config == nil {
		return nil, fmt.Errorf("equinox configuration not provided")
	}

	if config.Cluster == "" {
		return nil, fmt.Errorf("cluster not provided")
	}

	client, err := internal.NewInternalClient(config)
	if err != nil {
		return nil, err
	}

	if config.Key == "" {
		client.GetInternalLogger().Warn("API key was not provided, only Data Dragon methods will be enabled.")
		client.IsDataDragonOnly = true
	}

	if config.Cache == nil {
		config.Cache = &cache.Cache{TTL: 0}
	}

	equinox := &Equinox{
		Cache:   config.Cache,
		DDragon: ddragon.NewDDragonClient(client),
		CDragon: cdragon.NewCDragonClient(client),
		Riot:    riot.NewRiotClient(client),
		LOL:     lol.NewLOLClient(client),
		TFT:     tft.NewTFTClient(client),
		LOR:     lor.NewLORClient(client),
		VAL:     val.NewVALClient(client),
	}

	return equinox, nil
}
