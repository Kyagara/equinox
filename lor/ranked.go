package lor

import (
	"github.com/Kyagara/equinox/internal"
)

type RankedEndpoint struct {
	internalClient *internal.InternalClient
}

type LeaderboardDTO struct {
	// A list of players in Master tier.
	Players []LeaderboardPlayersDTO `json:"players"`
}

type LeaderboardPlayersDTO struct {
	Name string `json:"name"`
	Rank int    `json:"rank"`
	// League points.
	Lp int `json:"lp"`
}

// Get the players in Master tier.
//
// The leaderboard is updated once an hour.
func (e *RankedEndpoint) Leaderboards(region Region) (*LeaderboardDTO, error) {
	logger := e.internalClient.Logger("LOR", "ranked", "Leaderboards")

	var leaderboard *LeaderboardDTO

	err := e.internalClient.Get(region, RankedURL, &leaderboard, "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return leaderboard, nil
}
