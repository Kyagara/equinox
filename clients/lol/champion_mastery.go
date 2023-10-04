package lol

import (
	"fmt"

	"github.com/Kyagara/equinox/internal"
	"go.uber.org/zap"
)

type ChampionMasteryEndpoint struct {
	internalClient *internal.InternalClient
}

type ChampionMasteryDTO struct {
	// Champion ID for this entry.
	ChampionID int `json:"championId"`
	// Champion level for specified player and champion combination.
	ChampionLevel int `json:"championLevel"`
	// Total number of champion points for this player and champion combination - they are used to determine championLevel.
	ChampionPoints int `json:"championPoints"`
	// Last time this champion was played by this player - in Unix milliseconds time format.
	LastPlayTime int64 `json:"lastPlayTime"`
	// Number of points earned since current level has been achieved.
	ChampionPointsSinceLastLevel int `json:"championPointsSinceLastLevel"`
	// Number of points needed to achieve next level. Zero if player reached maximum champion level for this champion.
	ChampionPointsUntilNextLevel int `json:"championPointsUntilNextLevel"`
	// Is chest granted for this champion or not in current season.
	ChestGranted bool `json:"chestGranted"`
	// The token earned for this champion at the current championLevel. When the championLevel is advanced the tokensEarned resets to 0.
	TokensEarned int `json:"tokensEarned"`
	// Summoner ID for this entry. (Encrypted)
	SummonerID string `json:"summonerId"`
}

// Get all champion mastery entries sorted by number of champion points descending.
func (e *ChampionMasteryEndpoint) SummonerMasteries(region Region, summonerID string) (*[]ChampionMasteryDTO, error) {
	logger := e.internalClient.Logger("LOL", "champion-mastery-v4", "SummonerMasteries")

	logger.Debug("Method executed")

	if region == PBE1 {
		return nil, fmt.Errorf("the region PBE1 is not available for this method")
	}

	url := fmt.Sprintf(ChampionMasteriesURL, summonerID)

	var masteries *[]ChampionMasteryDTO

	err := e.internalClient.Get(region, url, &masteries, "champion-mastery-v4", "SummonerMasteries", "")

	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return masteries, nil
}

// Get a champion mastery by player ID and champion ID.
func (e *ChampionMasteryEndpoint) ChampionScore(region Region, summonerID string, championID int) (*ChampionMasteryDTO, error) {
	logger := e.internalClient.Logger("LOL", "champion-mastery-v4", "ChampionScore")
	logger.Debug("Method executed")

	if region == PBE1 {
		return nil, fmt.Errorf("the region PBE1 is not available for this method")
	}

	url := fmt.Sprintf(ChampionMasteriesByChampionURL, summonerID, championID)

	var score *ChampionMasteryDTO

	err := e.internalClient.Get(region, url, &score, "champion-mastery-v4", "ChampionScore", "")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return score, nil
}

// Get a player's total champion mastery score, which is the sum of individual champion mastery levels.
func (e *ChampionMasteryEndpoint) MasteryScoreSum(region Region, summonerID string) (*int, error) {
	logger := e.internalClient.Logger("LOL", "champion-mastery-v4", "MasteryScoreSum")
	logger.Debug("Method executed")

	if region == PBE1 {
		return nil, fmt.Errorf("the region PBE1 is not available for this method")
	}

	url := fmt.Sprintf(ChampionMasteriesScoresURL, summonerID)

	var sum *int

	err := e.internalClient.Get(region, url, &sum, "champion-mastery-v4", "MasteryScoreSum", "")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return sum, nil
}
