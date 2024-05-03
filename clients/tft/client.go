// This package is used to interact with all TFT endpoints.
//   - SpectatorV5
//   - LeagueV1
//   - MatchV1
//   - StatusV1
//   - SummonerV1
//
// Note: this package is automatically generated.
package tft

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = a70746fcf353ba0ad0aceceafcc70d4ba8de4431

import "github.com/Kyagara/equinox/internal"

type Client struct {
	SpectatorV5 SpectatorV5
	LeagueV1    LeagueV1
	MatchV1     MatchV1
	StatusV1    StatusV1
	SummonerV1  SummonerV1
}

// Creates a new TFT Client using the internal.Client provided.
func NewTFTClient(client *internal.Client) *Client {
	return &Client{
		SpectatorV5: SpectatorV5{internal: client},
		LeagueV1:    LeagueV1{internal: client},
		MatchV1:     MatchV1{internal: client},
		StatusV1:    StatusV1{internal: client},
		SummonerV1:  SummonerV1{internal: client},
	}
}
