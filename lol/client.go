package lol

import (
	"github.com/Kyagara/equinox/internal"
)

// URLs
const (
	LOLBaseURL          = "/lol/platform"
	ChampionEndpointURL = "/v3/champion-rotations"
)

type Region string

// Region enums
const (
	BR1 Region = "br1"
	NA1 Region = "na1"
)

type LOLClient struct {
	*internal.InternalClient
	Champion *ChampionEndpoint
}

// Creates a new LOLClient using a internal client
func NewLOLClient(client *internal.InternalClient) *LOLClient {
	return &LOLClient{
		InternalClient: client,
		Champion: &ChampionEndpoint{
			InternalClient: client,
		},
	}
}
