package lol

import (
	"fmt"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type ChampionEndpoint struct {
	internalClient *internal.InternalClient
}

type ChampionInfo struct {
	FreeChampionIds              []int `json:"freeChampionIds"`
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

// Get Free Champions Rotation
func (c *ChampionEndpoint) FreeRotation(region api.Region) (*ChampionInfo, error) {
	res := ChampionInfo{}

	if err := c.internalClient.SendRequest("GET", fmt.Sprintf(api.BaseURLFormat, region), ChampionEndpointURL, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
