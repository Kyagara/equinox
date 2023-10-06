package lol

import (
	"github.com/Kyagara/equinox/internal"
)

// League of Legends endpoint URLs.
const (
	ChallengesConfigurationsURL      = "/lol/challenges/v1/challenges/config"
	ChallengesPercentilesURL         = "/lol/challenges/v1/challenges/percentiles"
	ChallengesConfigurationByIDURL   = "/lol/challenges/v1/challenges/%d/config"
	ChallengesLeaderboardsByLevelURL = "/lol/challenges/v1/challenges/%d/leaderboards/by-level/%s"
	ChallengesPercentileByIDURL      = "/lol/challenges/v1/challenges/%d/percentiles"
	ChallengesByPUUIDURL             = "/lol/challenges/v1/player-data/%s"

	ChampionURL = "/lol/platform/v3/champion-rotations"

	ChampionMasteriesURL           = "/lol/champion-mastery/v4/champion-masteries/by-summoner/%s"
	ChampionMasteriesByChampionURL = "/lol/champion-mastery/v4/champion-masteries/by-summoner/%s/by-champion/%d"
	ChampionMasteriesScoresURL     = "/lol/champion-mastery/v4/scores/by-summoner/%s"

	ClashURL                   = "/lol/clash/v1/tournaments"
	ClashTournamentTeamByIDURL = "/lol/clash/v1/teams/%s"
	ClashSummonerEntriesURL    = "/lol/clash/v1/players/by-summoner/%s"
	ClashByTeamIDURL           = "/lol/clash/v1/tournaments/by-team/%s"
	ClashByIDURL               = "/lol/clash/v1/tournaments/%s"

	StatusURL = "/lol/status/v4/platform-data"

	MatchListURL     = "/lol/match/v5/matches/by-puuid/%s/ids"
	MatchByIDURL     = "/lol/match/v5/matches/%s"
	MatchTimelineURL = "/lol/match/v5/matches/%s/timeline"

	SpectatorFeaturedGamesURL = "/lol/spectator/v4/featured-games"
	SpectatorCurrentGameURL   = "/lol/spectator/v4/active-games/by-summoner/%s"

	SummonerByIDURL          = "/lol/summoner/v4/summoners/%s"
	SummonerByNameURL        = "/lol/summoner/v4/summoners/by-name/%s"
	SummonerByPUUIDURL       = "/lol/summoner/v4/summoners/by-puuid/%s"
	SummonerByAccountIDURL   = "/lol/summoner/v4/summoners/by-account/%s"
	SummonerByAccessTokenURL = "/lol/summoner/v4/summoners/me"

	LeagueEntriesURL           = "/lol/league/v4/entries/%s/%s/%s"
	LeagueEntriesBySummonerURL = "/lol/league/v4/entries/by-summoner/%s"
	LeagueByIDURL              = "/lol/league/v4/leagues/%s"
	LeagueChallengerURL        = "/lol/league/v4/challengerleagues/by-queue/%s"
	LeagueGrandmasterURL       = "/lol/league/v4/grandmasterleagues/by-queue/%s"
	LeagueMasterURL            = "/lol/league/v4/masterleagues/by-queue/%s"

	TournamentLobbyEventsURL = "/lol/tournament/v5/lobby-events/by-code/%s"
	TournamentCodesURL       = "/lol/tournament/v5/codes"
	TournamentByCodeURL      = "/lol/tournament/v5/codes/%s"
	TournamentProvidersURL   = "/lol/tournament/v5/providers"
	TournamentURL            = "/lol/tournament/v5/tournaments"

	TournamentStubLobbyEventsURL = "/lol/tournament-stub/v5/lobby-events/by-code/%s"
	TournamentStubCodesURL       = "/lol/tournament-stub/v5/codes"
	TournamentStubProvidersURL   = "/lol/tournament-stub/v5/providers"
	TournamentStubURL            = "/lol/tournament-stub/v5/tournaments"
)

type LOLClient struct {
	internalClient    *internal.InternalClient
	Challenges        *ChallengesEndpoint
	Champion          *ChampionEndpoint
	ChampionMasteries *ChampionMasteryEndpoint
	Clash             *ClashEndpoint
	Match             *MatchEndpoint
	Status            *StatusEndpoint
	Spectator         *SpectatorEndpoint
	Summoner          *SummonerEndpoint
	League            *LeagueEndpoint
	Tournament        *TournamentEndpoint
	TournamentStub    *TournamentStubEndpoint
}

// Creates a new LOLClient using the InternalClient provided.
func NewLOLClient(client *internal.InternalClient) *LOLClient {
	if client.IsDataDragonOnly {
		return nil
	}

	return &LOLClient{
		internalClient:    client,
		Challenges:        &ChallengesEndpoint{internalClient: client},
		Champion:          &ChampionEndpoint{internalClient: client},
		ChampionMasteries: &ChampionMasteryEndpoint{internalClient: client},
		Clash:             &ClashEndpoint{internalClient: client},
		Match:             &MatchEndpoint{internalClient: client},
		Status:            &StatusEndpoint{internalClient: client},
		Spectator:         &SpectatorEndpoint{internalClient: client},
		Summoner:          &SummonerEndpoint{internalClient: client},
		League:            &LeagueEndpoint{internalClient: client},
		Tournament:        &TournamentEndpoint{internalClient: client},
		TournamentStub:    &TournamentStubEndpoint{internalClient: client},
	}
}
