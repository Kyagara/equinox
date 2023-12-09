// This package is used to interact with Community Dragon endpoints.
package cdragon

import (
	"github.com/Kyagara/equinox/internal"
)

// CDragon endpoint URLs.
const (
	VersionsURL = "/api/versions.json"
	ChampionURL = "/%s/champion/%v/data"
)

type Client struct {
	Version  VersionEndpoint
	Champion ChampionEndpoint
}

// Returns a new CDragon Client using the internal.Client provided.
func NewCDragonClient(client *internal.Client) *Client {
	return &Client{
		Version:  VersionEndpoint{client},
		Champion: ChampionEndpoint{client},
	}
}
