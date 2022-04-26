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
	FreeChampionIds              []int16 `json:"freeChampionIds"`
	FreeChampionIdsForNewPlayers []int16 `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            uint8   `json:"maxNewPlayerLevel"`
}

// Get Free Champions Rotation
func (c *ChampionEndpoint) FreeRotation(region api.Region) (*ChampionInfoDTO, error) {
	res := ChampionInfoDTO{}

	err := c.internalClient.Do(http.MethodGet, region, ChampionEndpointURL, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
