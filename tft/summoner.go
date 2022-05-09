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
func (e *SummonerEndpoint) ByName(region lol.Region, summonerName string) (*lol.SummonerDTO, error) {
	return e.getSummoner(fmt.Sprintf(SummonerByNameURL, summonerName), region, "", "ByName")
}

// Get a summoner by summoner account ID.
func (e *SummonerEndpoint) ByAccountID(region lol.Region, accountID string) (*lol.SummonerDTO, error) {
	return e.getSummoner(fmt.Sprintf(SummonerByAccountIDURL, accountID), region, "", "ByAccountID")
}

// Get a summoner by summoner PUUID.
func (e *SummonerEndpoint) ByPUUID(region lol.Region, PUUID string) (*lol.SummonerDTO, error) {
	return e.getSummoner(fmt.Sprintf(SummonerByPUUIDURL, PUUID), region, "", "ByPUUID")
}

// Get a summoner by summoner ID.
func (e *SummonerEndpoint) ByID(region lol.Region, PUUID string) (*lol.SummonerDTO, error) {
	return e.getSummoner(fmt.Sprintf(SummonerByIDURL, PUUID), region, "", "ByID")
}

// Get a summoner by access token.
func (e *SummonerEndpoint) ByAccessToken(region lol.Region, accessToken string) (*lol.SummonerDTO, error) {
	return e.getSummoner(SummonerByAccessTokenURL, region, accessToken, "ByAccessToken")
}

func (e *SummonerEndpoint) getSummoner(url string, region lol.Region, accessToken string, methodName string) (*lol.SummonerDTO, error) {
	logger := e.internalClient.Logger("TFT", "summoner", methodName)

	var summoner *lol.SummonerDTO

	err := e.internalClient.Do(http.MethodGet, region, url, nil, &summoner, accessToken)

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return summoner, nil
}
