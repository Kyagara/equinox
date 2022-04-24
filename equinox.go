package equinox

import (
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
)

type Equinox struct {
	*internal.InternalClient
	LOL *lol.LOLClient
}

// Creates a new Equinox client
func NewClient(key string) *Equinox {
	internalClient := internal.NewClient(key)

	client := &Equinox{
		InternalClient: internalClient,
		LOL:            lol.NewLOLClient(internalClient),
	}

	return client
}
