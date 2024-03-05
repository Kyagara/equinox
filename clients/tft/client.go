// This package is used to interact with all TFT endpoints.
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

// Spec version = 031d3e7fc343bd86d82c45559fc79d3a87fa1b82

import "github.com/Kyagara/equinox/internal"

type Client struct {
	LeagueV1   LeagueV1
	MatchV1    MatchV1
	StatusV1   StatusV1
	SummonerV1 SummonerV1
}

// Creates a new TFT Client using the internal.Client provided.
func NewTFTClient(client *internal.Client) *Client {
	return &Client{
		LeagueV1:   LeagueV1{internal: client},
		MatchV1:    MatchV1{internal: client},
		StatusV1:   StatusV1{internal: client},
		SummonerV1: SummonerV1{internal: client},
	}
}
