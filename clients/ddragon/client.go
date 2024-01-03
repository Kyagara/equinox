// This package is used to interact with Data Dragon endpoints.
package ddragon

import (
	"github.com/Kyagara/equinox/internal"
)

type Client struct {
	Version  VersionEndpoint
	Realm    RealmEndpoint
	Champion ChampionEndpoint
}

// Returns a new DDragon Client using the internal.Client provided.
func NewDDragonClient(client *internal.Client) *Client {
	return &Client{
		Version:  VersionEndpoint{client},
		Realm:    RealmEndpoint{client},
		Champion: ChampionEndpoint{client},
	}
}
