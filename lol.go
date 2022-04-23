package equinox

import (
	"fmt"
	"net/http"
)

// URLs
const (
	LOLBaseURL          = "/lol/platform"
	ChampionEndpointURL = "/v3/champion-rotations"
)

type Region string

// Region enums
const (
	BR1 Region = "br1"
	NA1 Region = "na1"
)

type LOLClient struct {
	*Equinox
	Champion
}

func NewLOLClient(base *Equinox) *LOLClient {
	return &LOLClient{
		Equinox: base,
		Champion: Champion{
			Equinox: base,
		},
	}
}

type Champion struct {
	*Equinox
}

type FreeChampionsRotation struct {
	FreeChampionIds              []int `json:"freeChampionIds"`
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

// Free Champions Rotation
func (c *Champion) FreeRotation(region Region) (*FreeChampionsRotation, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s.api.riotgames.com%s%s", region, LOLBaseURL, ChampionEndpointURL), nil)

	if err != nil {
		return nil, err
	}

	res := FreeChampionsRotation{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
