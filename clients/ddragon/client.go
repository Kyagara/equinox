// This package is used to interact with Data Dragon endpoints.
package ddragon

import (
	"github.com/Kyagara/equinox/internal"
)

// Data Dragon endpoint URLs.
const (
	RealmURL = "/realms/%s.json"

	ChampionURL  = "/cdn/%s/data/%s/champion/%s.json"
	ChampionsURL = "/cdn/%s/data/%s/champion.json"
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
