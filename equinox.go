package equinox

import (
	"fmt"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
	"github.com/Kyagara/equinox/riot"
)

type Equinox struct {
	internalClient *internal.InternalClient
	Riot           *riot.RiotClient
	LOL            *lol.LOLClient
}

//	Creates a new Equinox client with a default configuration
//
//		- `Cluster`    : api.Americas
//		- `LogLevel`   : api.FatalLevel
//		- `Timeout`    : 10
//		- `Retry`      : true
func NewClient(key string) (*Equinox, error) {
	if key == "" {
		return nil, fmt.Errorf("API Key not provided")
	}

	config := &api.EquinoxConfig{
		Cluster:  api.Americas,
		Key:      key,
		LogLevel: api.FatalLevel,
		Timeout:  10,
		Retry:    true,
	}

	client := internal.NewInternalClient(config)

	equinox := &Equinox{
		internalClient: client,
		Riot:           riot.NewRiotClient(client),
		LOL:            lol.NewLOLClient(client),
	}

	return equinox, nil
}

//	Creates a new Equinox client using a custom configuration.
//
//	If you don't specify a Timeout this will disable the timeout for the http.Client.
func NewClientWithConfig(config *api.EquinoxConfig) (*Equinox, error) {
	if config.Key == "" {
		return nil, fmt.Errorf("API Key not provided")
	}

	client := internal.NewInternalClient(config)

	equinox := &Equinox{
		internalClient: client,
		Riot:           riot.NewRiotClient(client),
		LOL:            lol.NewLOLClient(client),
	}

	return equinox, nil
}
