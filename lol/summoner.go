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
func (c *SummonerEndpoint) ByName(region Region, summonerName string) (*SummonerDTO, error) {
	return c.getSummoner(SummonerByNameURL, region, summonerName)
}

// Get a summoner by summoner account ID.
func (c *SummonerEndpoint) ByAccountID(region Region, accountID string) (*SummonerDTO, error) {
	return c.getSummoner(SummonerByAccountIDURL, region, accountID)
}

// Get a summoner by summoner PUUID.
func (c *SummonerEndpoint) ByPUUID(region Region, PUUID string) (*SummonerDTO, error) {
	return c.getSummoner(SummonerByPUUIDURL, region, PUUID)
}

// Get a summoner by summoner ID.
func (c *SummonerEndpoint) ByID(region Region, PUUID string) (*SummonerDTO, error) {
	return c.getSummoner(SummonerByID, region, PUUID)
}

func (c *SummonerEndpoint) getSummoner(method string, region Region, summoner string) (*SummonerDTO, error) {
	url := fmt.Sprintf(method, summoner)

	res := SummonerDTO{}

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
