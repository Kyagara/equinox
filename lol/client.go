package lol

import (
	"github.com/Kyagara/equinox/internal"
)

// League of Legends endpoints
const (
	LOLBaseURLFormat    = "https://%s.api.riotgames.com/lol/platform%s"
	ChampionEndpointURL = "/v3/champion-rotations"
)

type LOLClient struct {
	internalClient *internal.InternalClient
	Champion       *ChampionEndpoint
}

// Creates a new LOLClient using a InternalClient
func NewLOLClient(client *internal.InternalClient) *LOLClient {
	return &LOLClient{
		internalClient: client,
		Champion: &ChampionEndpoint{
			internalClient: client,
		},
	}
}
