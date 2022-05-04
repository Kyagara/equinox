package tft

import "github.com/Kyagara/equinox/internal"

// Teamfight Tactics endpoints
const (
	SummonerByIDURL          = "/tft/summoner/v1/summoners/%s"
	SummonerByNameURL        = "/tft/summoner/v1/summoners/by-name/%s"
	SummonerByPUUIDURL       = "/tft/summoner/v1/summoners/by-puuid/%s"
	SummonerByAccountIDURL   = "/tft/summoner/v1/summoners/by-account/%s"
	SummonerByAccessTokenURL = "/tft/summoner/v1/summoners/me"
)

type TFTClient struct {
	internalClient *internal.InternalClient
	Summoner       *SummonerEndpoint
}

// Returns a new client using the API key provided.
func NewTFTClient(client *internal.InternalClient) *TFTClient {
	return &TFTClient{
		internalClient: client,
		Summoner:       &SummonerEndpoint{internalClient: client},
	}
}
