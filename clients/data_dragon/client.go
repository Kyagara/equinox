package data_dragon

import (
	"github.com/Kyagara/equinox/internal"
)

// Data Dragon endpoint URLs.
const (
	VersionsURL = "/api/versions.json"
	RealmURL    = "/realms/%s.json"
	ChampionURL = "/cdn/%s/data/%s/champion/%s.json"
)

type DataDragonClient struct {
	internalClient *internal.InternalClient
	Version        *VersionEndpoint
	Realm          *RealmEndpoint
	Champion       *ChampionEndpoint
}

// Returns a new DataDragonClient using the InternalClient provided.
func NewDataDragonClient(client *internal.InternalClient) *DataDragonClient {
	return &DataDragonClient{
		internalClient: client,
		Version:        &VersionEndpoint{client},
		Realm:          &RealmEndpoint{client},
		Champion:       &ChampionEndpoint{client},
	}
}
