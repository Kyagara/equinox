// This package is used to interact with all LOR endpoints.
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

// Spec version = 965a512f06ece005d0879258a6234ff0765d9a35

import "github.com/Kyagara/equinox/v2/internal"

type Client struct {
	DeckV1      DeckV1
	InventoryV1 InventoryV1
	MatchV1     MatchV1
	RankedV1    RankedV1
	StatusV1    StatusV1
}

// Creates a new LOR Client using the internal.Client provided.
func NewLORClient(client *internal.Client) *Client {
	return &Client{
		DeckV1:      DeckV1{internal: client},
		InventoryV1: InventoryV1{internal: client},
		MatchV1:     MatchV1{internal: client},
		RankedV1:    RankedV1{internal: client},
		StatusV1:    StatusV1{internal: client},
	}
}
