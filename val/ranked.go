package val

import (
	"fmt"
	"net/http"

	"github.com/Kyagara/equinox/internal"
)

type RankedEndpoint struct {
	internalClient *internal.InternalClient
}

type LeaderboardDTO struct {
	// The act id for the given leaderboard. Act ids can be found using the val-content API.
	ActID   string       `json:"actId"`
	Players []PlayersDTO `json:"players"`
	// The total number of players in the leaderboard.
	TotalPlayers          int            `json:"totalPlayers"`
	ImmortalStartingPage  int            `json:"immortalStartingPage"`
	ImmortalStartingIndex int            `json:"immortalStartingIndex"`
	TopTierRRThreshold    int            `json:"topTierRRThreshold"`
	TierDetails           TierDetailsDTO `json:"tierDetails"`
	StartIndex            int            `json:"startIndex"`
	Query                 string         `json:"query"`
	// The shard for the given leaderboard.
	Shard Region `json:"shard"`
}

type PlayersDTO struct {
	// This field may be omitted if the player has been anonymized.
	PUUID string `json:"puuid"`
	// This field may be omitted if the player has been anonymized.
	GameName string `json:"gameName"`
	// This field may be omitted if the player has been anonymized.
	TagLine         string `json:"tagLine"`
	LeaderboardRank int    `json:"leaderboardRank"`
	RankedRating    int    `json:"rankedRating"`
	NumberOfWins    int    `json:"numberOfWins"`
	CompetitiveTier int    `json:"competitiveTier"`
}

type AdditionalPropDTO struct {
	RankedRatingThreshold int `json:"rankedRatingThreshold"`
	StartingPage          int `json:"startingPage"`
	StartingIndex         int `json:"startingIndex"`
}

type TierDetailsDTO struct {
	Num21 AdditionalPropDTO `json:"21"`
	Num22 AdditionalPropDTO `json:"22"`
	Num23 AdditionalPropDTO `json:"23"`
	Num24 AdditionalPropDTO `json:"24"`
}

// Get leaderboard for the competitive queue.
//
// Size defaults to 200. Valid values: 1 to 200.
//
// Start defaults to 0.
func (r *RankedEndpoint) LeaderboardsByActID(region Region, actID string, size uint8, start int) (*LeaderboardDTO, error) {
	logger := r.internalClient.Logger("val").With("endpoint", "ranked", "method", "LeaderboardsByActID")

	if region == ESPORTS {
		return nil, fmt.Errorf("the region ESPORTS is not available for this method")
	}

	if size < 1 {
		size = 200
	}

	if start < 0 {
		start = 0
	}

	url := fmt.Sprintf(RankedURL, actID)

	var leaderboard *LeaderboardDTO

	err := r.internalClient.Do(http.MethodGet, region, url, nil, &leaderboard, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return leaderboard, nil
}
