// This package is used to interact with all Legends of Runeterra endpoints.
//   - DeckV1
//   - InventoryV1
//   - MatchV1
//   - RankedV1
//   - StatusV1
//
// Note: this package is automatically generated.
package lor

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = a5a3a5f5d5f2a617a56302a0afac77c745e4fd56

import "github.com/Kyagara/equinox/internal"

// Note: this struct is automatically generated.
type LORClient struct {
	internalClient *internal.InternalClient
	DeckV1         *DeckV1
	InventoryV1    *InventoryV1
	MatchV1        *MatchV1
	RankedV1       *RankedV1
	StatusV1       *StatusV1
}

// Creates a new LORClient using the InternalClient provided.
func NewLORClient(client *internal.InternalClient) *LORClient {
	return &LORClient{
		internalClient: client,
		DeckV1:         &DeckV1{internalClient: client},
		InventoryV1:    &InventoryV1{internalClient: client},
		MatchV1:        &MatchV1{internalClient: client},
		RankedV1:       &RankedV1{internalClient: client},
		StatusV1:       &StatusV1{internalClient: client},
	}
}
