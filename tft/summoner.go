package tft

import (
	"fmt"
	"net/http"

	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
)

type SummonerEndpoint struct {
	internalClient *internal.InternalClient
}

// Get a summoner by summoner name.
func (s *SummonerEndpoint) ByName(region lol.Region, summonerName string) (*lol.SummonerDTO, error) {
	return s.getSummoner(fmt.Sprintf(SummonerByNameURL, summonerName), region, "", "ByName")
}

// Get a summoner by summoner account ID.
func (s *SummonerEndpoint) ByAccountID(region lol.Region, accountID string) (*lol.SummonerDTO, error) {
	return s.getSummoner(fmt.Sprintf(SummonerByAccountIDURL, accountID), region, "", "ByAccountID")
}

// Get a summoner by summoner PUUID.
func (s *SummonerEndpoint) ByPUUID(region lol.Region, PUUID string) (*lol.SummonerDTO, error) {
	return s.getSummoner(fmt.Sprintf(SummonerByPUUIDURL, PUUID), region, "", "ByPUUID")
}

// Get a summoner by summoner ID.
func (s *SummonerEndpoint) ByID(region lol.Region, PUUID string) (*lol.SummonerDTO, error) {
	return s.getSummoner(fmt.Sprintf(SummonerByIDURL, PUUID), region, "", "ByID")
}

// Get a summoner by access token.
func (s *SummonerEndpoint) ByAccessToken(region lol.Region, accessToken string) (*lol.SummonerDTO, error) {
	return s.getSummoner(SummonerByAccessTokenURL, region, accessToken, "ByAccessToken")
}

func (s *SummonerEndpoint) getSummoner(url string, region lol.Region, accessToken string, methodName string) (*lol.SummonerDTO, error) {
	logger := s.internalClient.Logger("tft").With("endpoint", "summoner", "method", methodName)

	var summoner *lol.SummonerDTO

	err := s.internalClient.Do(http.MethodGet, region, url, nil, &summoner, accessToken)

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return summoner, nil
}
