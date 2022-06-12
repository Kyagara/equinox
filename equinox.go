package equinox

import (
	"fmt"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
	"github.com/Kyagara/equinox/lor"
	"github.com/Kyagara/equinox/riot"
	"github.com/Kyagara/equinox/tft"
	"github.com/Kyagara/equinox/val"
)

type Equinox struct {
	internalClient *internal.InternalClient
	Riot           *riot.RiotClient
	LOL            *lol.LOLClient
	TFT            *tft.TFTClient
	LOR            *lor.LORClient
	VAL            *val.VALClient
}

//	Creates a new Equinox client with a default configuration
//
//		- `Cluster`    : api.AmericasCluster
//		- `LogLevel`   : api.FatalLevel
//		- `Timeout`    : 10
//		- `TTL`        : 120
//		- `Retry`      : true
//		- `RateLimit`  : true
func NewClient(key string) (*Equinox, error) {
	if key == "" {
		return nil, fmt.Errorf("API Key not provided")
	}

	config := &api.EquinoxConfig{
		Cluster:   api.AmericasCluster,
		Key:       key,
		LogLevel:  api.FatalLevel,
		Timeout:   10,
		TTL:       120,
		Retry:     true,
		RateLimit: true,
	}

	client := internal.NewInternalClient(config)

	equinox := &Equinox{
		internalClient: client,
		Riot:           riot.NewRiotClient(client),
		LOL:            lol.NewLOLClient(client),
		TFT:            tft.NewTFTClient(client),
		LOR:            lor.NewLORClient(client),
		VAL:            val.NewVALClient(client),
	}

	return equinox, nil
}

//	Creates a new Equinox client using a custom configuration.
//
//	If you don't specify a Timeout this will disable the timeout for the http.Client.
func NewClientWithConfig(config *api.EquinoxConfig) (*Equinox, error) {
	if config == nil {
		return nil, fmt.Errorf("equinox configuration not provided")
	}

	if config.Key == "" {
		return nil, fmt.Errorf("API Key not provided")
	}

	if config.Cluster == "" {
		return nil, fmt.Errorf("cluster not provided")
	}

	client := internal.NewInternalClient(config)

	equinox := &Equinox{
		internalClient: client,
		Riot:           riot.NewRiotClient(client),
		LOL:            lol.NewLOLClient(client),
		TFT:            tft.NewTFTClient(client),
		LOR:            lor.NewLORClient(client),
		VAL:            val.NewVALClient(client),
	}

	return equinox, nil
}

// Clears the Cache
func (c *Equinox) ClearCache() {
	c.internalClient.ClearInternalClientCache()
}
