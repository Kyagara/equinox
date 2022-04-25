package equinox

import (
	"fmt"

	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
)

type Equinox struct {
	internalClient *internal.InternalClient
	LOL            *lol.LOLClient
}

// Creates a new Equinox client
func NewClient(key string) (*Equinox, error) {
	if key == "" {
		return nil, fmt.Errorf("API Key not provided")
	}

	internalClient := internal.NewInternalClient(key, false)

	client := &Equinox{
		internalClient: internalClient,
		LOL:            lol.NewLOLClient(internalClient),
	}

	return client, nil
}

// Creates a new Equinox client with debug
func NewClientWithDebug(key string) (*Equinox, error) {
	if key == "" {
		return nil, fmt.Errorf("API Key not provided")
	}

	internalClient := internal.NewInternalClient(key, true)

	client := &Equinox{
		internalClient: internalClient,
		LOL:            lol.NewLOLClient(internalClient),
	}

	return client, nil
}
