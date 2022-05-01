package lol

import (
	"fmt"
	"net/http"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type SummonerEndpoint struct {
	internalClient *internal.InternalClient
}

type SummonerDTO struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

// Get a summoner by summoner name.
func (c *SummonerEndpoint) ByName(region api.LOLRegion, summonerName string) (*SummonerDTO, error) {
	url := fmt.Sprintf(SummonerByNameURL, summonerName)

	res := SummonerDTO{}

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get a summoner by summoner account ID.
func (c *SummonerEndpoint) ByAccountID(region api.LOLRegion, accountID string) (*SummonerDTO, error) {
	url := fmt.Sprintf(SummonerByAccountIDURL, accountID)

	res := SummonerDTO{}

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get a summoner by summoner PUUID.
func (c *SummonerEndpoint) ByPUUID(region api.LOLRegion, PUUID string) (*SummonerDTO, error) {
	url := fmt.Sprintf(SummonerByPUUIDURL, PUUID)

	res := SummonerDTO{}

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get a summoner by summoner ID.
func (c *SummonerEndpoint) ByID(region api.LOLRegion, summonerID string) (*SummonerDTO, error) {
	url := fmt.Sprintf(SummonerByID, summonerID)

	res := SummonerDTO{}

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
