package lol

import (
	"github.com/Kyagara/equinox/internal"
)

// League of Legends endpoints URLs
const (
	ChampionURL                    = "/lol/platform/v3/champion-rotations"
	ChampionMasteriesURL           = "/lol/champion-mastery/v4/champion-masteries/by-summoner/%s"
	ChampionMasteriesByChampionURL = "/lol/champion-mastery/v4/champion-masteries/by-summoner/%s/by-champion/%d"
	ChampionMasteriesScoresURL     = "/lol/champion-mastery/v4/scores/by-summoner/%s"
	StatusURL                      = "/lol/status/v4/platform-data"
	MatchlistURL                   = "/lol/match/v5/matches/by-puuid/%s/ids"
	MatchURL                       = "/lol/match/v5/matches/%s"
	MatchTimelineURL               = "/lol/match/v5/matches/%s/timeline"
	SpectatorURL                   = "/lol/spectator/v4/featured-games"
	SpectatorCurrentGameURL        = "/lol/spectator/v4/active-games/by-summoner/%s"
	SummonerByAccountIDURL         = "/lol/summoner/v4/summoners/by-account/%s"
	SummonerByNameURL              = "/lol/summoner/v4/summoners/by-name/%s"
	SummonerByPUUIDURL             = "/lol/summoner/v4/summoners/by-puuid/%s"
	SummonerByID                   = "/lol/summoner/v4/summoners/%s"
)

type LOLClient struct {
	internalClient    *internal.InternalClient
	Champion          *ChampionEndpoint
	ChampionMasteries *ChampionMasteryEndpoint
	Match             *MatchEndpoint
	Status            *StatusEndpoint
	Spectator         *SpectatorEndpoint
	Summoner          *SummonerEndpoint
}

// Creates a new LOLClient using an InternalClient provided.
func NewLOLClient(client *internal.InternalClient) *LOLClient {
	return &LOLClient{
		internalClient:    client,
		Champion:          &ChampionEndpoint{internalClient: client},
		ChampionMasteries: &ChampionMasteryEndpoint{internalClient: client},
		Match:             &MatchEndpoint{internalClient: client},
		Status:            &StatusEndpoint{internalClient: client},
		Spectator:         &SpectatorEndpoint{internalClient: client},
		Summoner:          &SummonerEndpoint{internalClient: client},
	}
}
