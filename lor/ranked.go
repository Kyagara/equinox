package lor

import (
	"net/http"

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
	logger := e.internalClient.Logger("lor").With("endpoint", "ranked", "method", "Leaderboards")

	var leaderboard *LeaderboardDTO

	err := e.internalClient.Do(http.MethodGet, region, RankedURL, nil, &leaderboard, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return leaderboard, nil
}
