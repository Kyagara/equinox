package lol

import (
	"fmt"

	"github.com/Kyagara/equinox/internal"
	"go.uber.org/zap"
)

type SummonerEndpoint struct {
	internalClient *internal.InternalClient
}

type SummonerDTO struct {
	// Encrypted summoner ID. Max length 63 characters.
	ID string `json:"id"`
	// Encrypted account ID. Max length 56 characters.
	AccountID string `json:"accountId"`
	// Encrypted PUUID. Exact length of 78 characters.
	PUUID string `json:"puuid"`
	// Summoner name.
	Name string `json:"name"`
	// ID of the summoner icon associated with the summoner.
	ProfileIconID int `json:"profileIconId"`
	// Date summoner was last modified specified as epoch milliseconds. The following events will update this timestamp: summoner name change, summoner level change, or profile icon change.
	RevisionDate int64 `json:"revisionDate"`
	// Summoner level associated with the summoner.
	SummonerLevel int `json:"summonerLevel"`
}

// Get a summoner by summoner name.
func (e *SummonerEndpoint) ByName(region Region, summonerName string) (*SummonerDTO, error) {
	return e.getSummoner(fmt.Sprintf(SummonerByNameURL, summonerName), region, "", "ByName")
}

// Get a summoner by summoner account ID.
func (e *SummonerEndpoint) ByAccountID(region Region, accountID string) (*SummonerDTO, error) {
	return e.getSummoner(fmt.Sprintf(SummonerByAccountIDURL, accountID), region, "", "ByAccountID")
}

// Get a summoner by summoner PUUID.
func (e *SummonerEndpoint) ByPUUID(region Region, puuid string) (*SummonerDTO, error) {
	return e.getSummoner(fmt.Sprintf(SummonerByPUUIDURL, puuid), region, "", "ByPUUID")
}

// Get a summoner by summoner ID.
func (e *SummonerEndpoint) ByID(region Region, summonerID string) (*SummonerDTO, error) {
	return e.getSummoner(fmt.Sprintf(SummonerByIDURL, summonerID), region, "", "ByID")
}

// Get a summoner by access token.
func (e *SummonerEndpoint) ByAccessToken(region Region, accessToken string) (*SummonerDTO, error) {
	return e.getSummoner(SummonerByAccessTokenURL, region, accessToken, "ByAccessToken")
}

func (e *SummonerEndpoint) getSummoner(url string, region Region, accessToken string, methodName string) (*SummonerDTO, error) {
	logger := e.internalClient.Logger("LOL", "summoner-v4", methodName)
	logger.Debug("Method executed")

	if methodName == "ByAccessToken" && accessToken == "" {
		return nil, fmt.Errorf("accessToken is required")
	}

	var summoner *SummonerDTO

	err := e.internalClient.Get(region, url, &summoner, "summoner-v4", methodName, accessToken)
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return summoner, nil
}
