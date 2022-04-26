package lol

import (
	"github.com/Kyagara/equinox/internal"
)

// League of Legends endpoints
const (
	ChampionEndpointURL            = "/lol/platform/v3/champion-rotations"
	StatusEndpointURL              = "/lol/status/v4/platform-data"
	SpectatorEndpointURL           = "/lol/spectator/v4/featured-games"
	SpectatorBySummonerEndpointURL = "/lol/spectator/v4/active-games/by-summoner/%s"
)

type LOLClient struct {
	internalClient *internal.InternalClient
	Champion       *ChampionEndpoint
	Status         *StatusEndpoint
	Spectator      *SpectatorEndpoint
}

// Creates a new LOLClient using an InternalClient provided
func NewLOLClient(client *internal.InternalClient) *LOLClient {
	return &LOLClient{
		internalClient: client,
		Champion:       &ChampionEndpoint{internalClient: client},
		Status:         &StatusEndpoint{internalClient: client},
		Spectator:      &SpectatorEndpoint{internalClient: client},
	}
}
