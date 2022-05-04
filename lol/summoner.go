package lol

import (
	"fmt"
	"net/http"

	"github.com/Kyagara/equinox/internal"
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
func (s *SummonerEndpoint) ByName(region Region, summonerName string) (*SummonerDTO, error) {
	return s.getSummoner(fmt.Sprintf(SummonerByNameURL, summonerName), region, "", "ByName")
}

// Get a summoner by summoner account ID.
func (s *SummonerEndpoint) ByAccountID(region Region, accountID string) (*SummonerDTO, error) {
	return s.getSummoner(fmt.Sprintf(SummonerByAccountIDURL, accountID), region, "", "ByAccountID")
}

// Get a summoner by summoner PUUID.
func (s *SummonerEndpoint) ByPUUID(region Region, PUUID string) (*SummonerDTO, error) {
	return s.getSummoner(fmt.Sprintf(SummonerByPUUIDURL, PUUID), region, "", "ByPUUID")
}

// Get a summoner by summoner ID.
func (s *SummonerEndpoint) ByID(region Region, summonerID string) (*SummonerDTO, error) {
	return s.getSummoner(fmt.Sprintf(SummonerByIDURL, summonerID), region, "", "ByID")
}

// Get a summoner by access token.
func (s *SummonerEndpoint) ByAccessToken(region Region, accessToken string) (*SummonerDTO, error) {
	return s.getSummoner(SummonerByAccessTokenURL, region, accessToken, "ByAccessToken")
}

func (s *SummonerEndpoint) getSummoner(url string, region Region, accessToken string, methodName string) (*SummonerDTO, error) {
	logger := s.internalClient.Logger("lol").With("endpoint", "summoner", "method", methodName)

	if methodName == "ByAccessToken" && accessToken == "" {
		return nil, fmt.Errorf("accessToken is required")
	}

	var summoner *SummonerDTO

	err := s.internalClient.Do(http.MethodGet, region, url, nil, &summoner, accessToken)

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return summoner, nil
}
