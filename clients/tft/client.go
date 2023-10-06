package tft

import "github.com/Kyagara/equinox/internal"

// Teamfight Tactics endpoint URLs.
const (
	MatchListURL = "/tft/match/v1/matches/by-puuid/%s/ids"
	MatchByIDURL = "/tft/match/v1/matches/%s"

	LeagueRatedLaddersURL      = "/tft/league/v1/rated-ladders/%s/top "
	LeagueEntriesURL           = "/tft/league/v1/entries/%s/%s"
	LeagueEntriesBySummonerURL = "/tft/league/v1/entries/by-summoner/%s"
	LeagueByIDURL              = "/tft/league/v1/leagues/%s"
	LeagueChallengerURL        = "/tft/league/v1/challenger"
	LeagueGrandmasterURL       = "/tft/league/v1/grandmaster"
	LeagueMasterURL            = "/tft/league/v1/master"

	SummonerByIDURL          = "/tft/summoner/v1/summoners/%s"
	SummonerByNameURL        = "/tft/summoner/v1/summoners/by-name/%s"
	SummonerByPUUIDURL       = "/tft/summoner/v1/summoners/by-puuid/%s"
	SummonerByAccountIDURL   = "/tft/summoner/v1/summoners/by-account/%s"
	SummonerByAccessTokenURL = "/tft/summoner/v1/summoners/me"
)

type TFTClient struct {
	internalClient *internal.InternalClient
	Match          *MatchEndpoint
	League         *LeagueEndpoint
	Summoner       *SummonerEndpoint
}

// Returns a new TFTClient using the InternalClient provided.
func NewTFTClient(client *internal.InternalClient) *TFTClient {
	if client.IsDataDragonOnly {
		return nil
	}

	return &TFTClient{
		internalClient: client,
		Match:          &MatchEndpoint{internalClient: client},
		League:         &LeagueEndpoint{internalClient: client},
		Summoner:       &SummonerEndpoint{internalClient: client},
	}
}
