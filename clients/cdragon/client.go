// This package is used to interact with CDragon endpoints.
package cdragon

import (
	"github.com/Kyagara/equinox/internal"
)

// CDragon endpoint URLs.
const (
	VersionsURL = "/api/versions.json"
	ChampionURL = "/%s/champion/%v/data"
)

type CDragonClient struct {
	internalClient *internal.InternalClient
	Version        *VersionEndpoint
	Champion       *ChampionEndpoint
}

// Returns a new CDragonClient using the InternalClient provided.
func NewCDragonClient(client *internal.InternalClient) *CDragonClient {
	return &CDragonClient{
		internalClient: client,
		Version:        &VersionEndpoint{client},
		Champion:       &ChampionEndpoint{client},
	}
}
