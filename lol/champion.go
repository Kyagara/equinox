package lol

import (
	"net/http"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type ChampionEndpoint struct {
	internalClient *internal.InternalClient
}

type ChampionInfoDTO struct {
	FreeChampionIds              []int `json:"freeChampionIds"`
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

// Get Free Champions Rotation
func (c *ChampionEndpoint) FreeRotation(region api.Region) (*ChampionInfoDTO, error) {
	res := ChampionInfoDTO{}

	err := c.internalClient.Do(http.MethodGet, region, ChampionEndpointURL, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
