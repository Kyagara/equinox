package lol

import (
	"fmt"
	"net/http"

	"github.com/Kyagara/equinox/internal"
)

type ChampionEndpoint struct {
	*internal.InternalClient
}

type FreeChampionsRotation struct {
	FreeChampionIds              []int `json:"freeChampionIds"`
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

// Get Free Champions Rotation
func (c *ChampionEndpoint) FreeRotation(region Region) (*FreeChampionsRotation, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s.api.riotgames.com%s%s", region, LOLBaseURL, ChampionEndpointURL), nil)

	if err != nil {
		return nil, err
	}

	res := FreeChampionsRotation{}

	if err := c.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
