// This package is used to interact with LOR endpoints.
//
// Automatically generated package.
package lor

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = ed83574d1b85ef4c52f267ee5558e3c1c3ffb412

import (
	"github.com/Kyagara/equinox/internal"
)

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
