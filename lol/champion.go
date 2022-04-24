package lol

import (
	"fmt"
	"net/http"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type ChampionEndpoint struct {
	internalClient *internal.InternalClient
}

type FreeChampionsRotation struct {
	FreeChampionIds              []int `json:"freeChampionIds"`
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

// Get Free Champions Rotation
func (c *ChampionEndpoint) FreeRotation(region api.Region) (*FreeChampionsRotation, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(LOLBaseURLFormat, region, ChampionEndpointURL), nil)

	if err != nil {
		return nil, err
	}

	res := FreeChampionsRotation{}

	if err := c.internalClient.SendRequest(req, ChampionEndpointURL, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
