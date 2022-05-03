package equinox

import (
	"fmt"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
)

type Equinox struct {
	internalClient *internal.InternalClient
	LOL            *lol.LOLClient
}

//	Creates a new Equinox client with a default configuration
//
//		- `LogLevel`   : api.FatalLevel
//		- `Timeout`    : 10
//		- `Retry`      : true
func NewClient(key string) (*Equinox, error) {
	if key == "" {
		return nil, fmt.Errorf("API Key not provided")
	}

	config := &api.EquinoxConfig{
		Key:      key,
		LogLevel: api.FatalLevel,
		Timeout:  10,
		Retry:    true,
	}

	internalClient := internal.NewInternalClient(config)

	client := &Equinox{
		internalClient: internalClient,
		LOL:            lol.NewLOLClient(internalClient),
	}

	return client, nil
}

//	Creates a new Equinox client using a custom configuration.
//
//	If you don't specify a Timeout this will disable the timeout for the http.Client.
func NewClientWithConfig(config *api.EquinoxConfig) (*Equinox, error) {
	if config.Key == "" {
		return nil, fmt.Errorf("API Key not provided")
	}

	internalClient := internal.NewInternalClient(config)

	client := &Equinox{
		internalClient: internalClient,
		LOL:            lol.NewLOLClient(internalClient),
	}

	return client, nil
}
