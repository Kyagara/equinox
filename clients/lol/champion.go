package lol

import (
	"fmt"

	"github.com/Kyagara/equinox/internal"
	"go.uber.org/zap"
)

type ChampionEndpoint struct {
	internalClient *internal.InternalClient
}

type ChampionRotationsDTO struct {
	// List of free champions IDs.
	FreeChampionIDs []int `json:"freeChampionIds"`
	// List of free champions IDs for new players.
	FreeChampionIDsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	// Max new player level.
	MaxNewPlayerLevel int `json:"maxNewPlayerLevel"`
}

// Get champion rotations, including free-to-play and low-level free-to-play rotations.
func (e *ChampionEndpoint) Rotations(region Region) (*ChampionRotationsDTO, error) {
	logger := e.internalClient.Logger("LOL", "champion-v3", "Rotations")
	logger.Debug("Method executed")

	if region == PBE1 {
		return nil, fmt.Errorf("the region PBE1 is not available for this method")
	}

	var rotations ChampionRotationsDTO

	err := e.internalClient.Get(region, ChampionURL, &rotations, "champion-v3", "Rotations", "")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return &rotations, nil
}
