package lol

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type ChampionMasteryEndpoint struct {
	internalClient *internal.InternalClient
}

type ChampionMasteryDTO struct {
	ChampionID                   int       `json:"championId"`
	ChampionLevel                int       `json:"championLevel"`
	ChampionPoints               int       `json:"championPoints"`
	LastPlayTime                 time.Time `json:"lastPlayTime"`
	ChampionPointsSinceLastLevel int       `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel int       `json:"championPointsUntilNextLevel"`
	ChestGranted                 bool      `json:"chestGranted"`
	TokensEarned                 int       `json:"tokensEarned"`
	SummonerID                   string    `json:"summonerId"`
}

// Get all champion mastery entries sorted by number of champion points descending.
func (c *ChampionMasteryEndpoint) BySummonerID(region api.Region, summonerID string) (*[]ChampionMasteryDTO, error) {
	res := []ChampionMasteryDTO{}

	err := c.internalClient.Do(http.MethodGet, region, ChampionMasteriesURL, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get a champion mastery by player ID and champion ID.
func (c *ChampionMasteryEndpoint) ChampionScore(region api.Region, summonerID string, championID int) (*ChampionMasteryDTO, error) {
	url := fmt.Sprintf(ChampionMasteriesByChampionURL, summonerID, championID)

	res := ChampionMasteryDTO{}

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get a player's total champion mastery score, which is the sum of individual champion mastery levels.
func (c *ChampionMasteryEndpoint) MasteryScore(region api.Region, summonerID string) (int, error) {
	url := fmt.Sprintf(ChampionMasteriesScoresURL, summonerID)

	res := 0

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &res)

	if err != nil {
		return -1, err
	}

	return res, nil
}
