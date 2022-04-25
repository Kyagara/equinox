package lol

import (
	"github.com/Kyagara/equinox/internal"
)

// League of Legends endpoints
const (
	ChampionEndpointURL = "/lol/platform/v3/champion-rotations"
	StatusEndpointURL   = "/lol/status/v4/platform-data"
)

type LOLClient struct {
	internalClient *internal.InternalClient
	Champion       *ChampionEndpoint
	Status         *StatusEndpoint
}

// Creates a new LOLClient using a InternalClient
func NewLOLClient(client *internal.InternalClient) *LOLClient {
	return &LOLClient{
		internalClient: client,
		Champion:       &ChampionEndpoint{internalClient: client},
		Status: &StatusEndpoint{
			internalClient: client,
		},
	}
}
