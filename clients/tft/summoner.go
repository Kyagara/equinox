package tft

import (
	"fmt"

	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/internal"
	"go.uber.org/zap"
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
func (e *SummonerEndpoint) ByPUUID(region lol.Region, puuid string) (*lol.SummonerDTO, error) {
	return e.getSummoner(fmt.Sprintf(SummonerByPUUIDURL, puuid), region, "", "ByPUUID")
}

// Get a summoner by summoner ID.
func (e *SummonerEndpoint) ByID(region lol.Region, puuid string) (*lol.SummonerDTO, error) {
	return e.getSummoner(fmt.Sprintf(SummonerByIDURL, puuid), region, "", "ByID")
}

// Get a summoner by access token.
func (e *SummonerEndpoint) ByAccessToken(region lol.Region, accessToken string) (*lol.SummonerDTO, error) {
	return e.getSummoner(SummonerByAccessTokenURL, region, accessToken, "ByAccessToken")
}

func (e *SummonerEndpoint) getSummoner(url string, region lol.Region, accessToken string, methodName string) (*lol.SummonerDTO, error) {
	logger := e.internalClient.Logger("TFT", "tft-summoner-v1", methodName)
	logger.Debug("Method executed")

	var summoner lol.SummonerDTO

	err := e.internalClient.Get(region, url, &summoner, "tft-summoner-v1", methodName, accessToken)
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return &summoner, nil
}
