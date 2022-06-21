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
	logger := e.internalClient.Logger("LOR", "lor-ranked-v1", "Leaderboards")

	logger.Debug("Method executed")

	var leaderboard *LeaderboardDTO

	err := e.internalClient.Get(region, RankedURL, &leaderboard, "lor-ranked-v1", "Leaderboards", "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return leaderboard, nil
}
