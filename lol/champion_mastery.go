package lol

import (
	"fmt"
	"net/http"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
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
func (c *ChampionMasteryEndpoint) SummonerMasteries(region api.LOLRegion, summonerID string) (*[]ChampionMasteryDTO, error) {
	url := fmt.Sprintf(ChampionMasteriesURL, summonerID)

	res := []ChampionMasteryDTO{}

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get a champion mastery by player ID and champion ID.
func (c *ChampionMasteryEndpoint) ChampionScore(region api.LOLRegion, summonerID string, championID int) (*ChampionMasteryDTO, error) {
	url := fmt.Sprintf(ChampionMasteriesByChampionURL, summonerID, championID)

	res := ChampionMasteryDTO{}

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get a player's total champion mastery score, which is the sum of individual champion mastery levels.
func (c *ChampionMasteryEndpoint) MasteryScoreSum(region api.LOLRegion, summonerID string) (int, error) {
	url := fmt.Sprintf(ChampionMasteriesScoresURL, summonerID)

	res := 0

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &res)

	if err != nil {
		return -1, err
	}

	return res, nil
}
