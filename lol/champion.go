package lol

import (
	"net/http"

	"github.com/Kyagara/equinox/internal"
)

type ChampionEndpoint struct {
	internalClient *internal.InternalClient
}

type ChampionRotationsDTO struct {
	// List of free champions IDs
	FreeChampionIDs []int `json:"freeChampionIds"`
	// List of free champions IDs for new players
	FreeChampionIDsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	// Max new player level
	MaxNewPlayerLevel int `json:"maxNewPlayerLevel"`
}

// Get champion rotations, including free-to-play and low-level free-to-play rotations.
func (c *ChampionEndpoint) Rotations(region Region) (*ChampionRotationsDTO, error) {
	logger := c.internalClient.Logger().With("endpoint", "champion", "method", "Rotations")

	var rotations *ChampionRotationsDTO

	err := c.internalClient.Do(http.MethodGet, region, ChampionURL, nil, &rotations)

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return rotations, nil
}
