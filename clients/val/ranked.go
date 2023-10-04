package val

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/Kyagara/equinox/internal"
	"go.uber.org/zap"
)

type RankedEndpoint struct {
	internalClient *internal.InternalClient
}

type LeaderboardDTO struct {
	// The act id for the given leaderboard. Act ids can be found using the val-content api.
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
	Shard Shard `json:"shard"`
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
// Size defaults to 200. Valid values: 1 to 200
//
// Start defaults to 0
func (e *RankedEndpoint) LeaderboardsByActID(shard Shard, actID string, size uint8, start int) (*LeaderboardDTO, error) {
	logger := e.internalClient.Logger("VAL", "val-ranked-v1", "LeaderboardsByActID")
	logger.Debug("Method executed")

	if shard == ESPORTS {
		return nil, fmt.Errorf("the region ESPORTS is not available for this method")
	}

	if size < 1 || size > 200 {
		size = 200
	}

	if start < 0 {
		start = 0
	}

	query := url.Values{}
	query.Set("size", strconv.Itoa(int(size)))
	query.Set("start", strconv.Itoa(int(start)))

	method := fmt.Sprintf(RankedURL, actID)
	url := fmt.Sprintf("%s?%s", method, query.Encode())

	var leaderboard LeaderboardDTO

	err := e.internalClient.Get(shard, url, &leaderboard, "val-ranked-v1", "LeaderboardsByActID", "")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return &leaderboard, nil
}
