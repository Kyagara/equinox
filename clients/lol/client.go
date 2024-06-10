// This package is used to interact with all LOL endpoints.
//   - ChampionMasteryV4
//   - ChampionV3
//   - ClashV1
//   - LeagueExpV4
//   - LeagueV4
//   - ChallengesV1
//   - StatusV4
//   - MatchV5
//   - SpectatorV5
//   - SummonerV4
//   - TournamentStubV5
//   - TournamentV5
//
// Note: this package is automatically generated.
package lol

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 9fef246d3ece1da9515c8941f7a3c7cd57e330fc

import "github.com/Kyagara/equinox/v2/internal"

type Client struct {
	ChampionMasteryV4 ChampionMasteryV4
	ChampionV3        ChampionV3
	ClashV1           ClashV1
	LeagueExpV4       LeagueExpV4
	LeagueV4          LeagueV4
	ChallengesV1      ChallengesV1
	StatusV4          StatusV4
	MatchV5           MatchV5
	SpectatorV5       SpectatorV5
	SummonerV4        SummonerV4
	TournamentStubV5  TournamentStubV5
	TournamentV5      TournamentV5
}

// Creates a new LOL Client using the internal.Client provided.
func NewLOLClient(client *internal.Client) *Client {
	return &Client{
		ChampionMasteryV4: ChampionMasteryV4{internal: client},
		ChampionV3:        ChampionV3{internal: client},
		ClashV1:           ClashV1{internal: client},
		LeagueExpV4:       LeagueExpV4{internal: client},
		LeagueV4:          LeagueV4{internal: client},
		ChallengesV1:      ChallengesV1{internal: client},
		StatusV4:          StatusV4{internal: client},
		MatchV5:           MatchV5{internal: client},
		SpectatorV5:       SpectatorV5{internal: client},
		SummonerV4:        SummonerV4{internal: client},
		TournamentStubV5:  TournamentStubV5{internal: client},
		TournamentV5:      TournamentV5{internal: client},
	}
}
