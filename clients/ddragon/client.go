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

type DDragonClient struct {
	internalClient *internal.InternalClient
	Version        *VersionEndpoint
	Realm          *RealmEndpoint
	Champion       *ChampionEndpoint
}

// Returns a new DDragonClient using the InternalClient provided.
func NewDDragonClient(client *internal.InternalClient) *DDragonClient {
	return &DDragonClient{
		internalClient: client,
		Version:        &VersionEndpoint{client},
		Realm:          &RealmEndpoint{client},
		Champion:       &ChampionEndpoint{client},
	}
}
