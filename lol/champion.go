package lol

import (
	"net/http"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type ChampionEndpoint struct {
	internalClient *internal.InternalClient
}

type ChampionRotationsDTO struct {
	FreeChampionIds              []int `json:"freeChampionIds"`
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

// Get champion rotations, including free-to-play and low-level free-to-play rotations.
func (c *ChampionEndpoint) Rotations(region api.LOLRegion) (*ChampionRotationsDTO, error) {
	res := ChampionRotationsDTO{}

	err := c.internalClient.Do(http.MethodGet, region, ChampionURL, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
