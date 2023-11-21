// This package is used to interact with all Riot Games endpoints.
//   - AccountV1
//
// Note: this package is automatically generated.
package riot

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 9e564de22543560b3ce444e320b260b7e759455a

import "github.com/Kyagara/equinox/internal"

// Note: this struct is automatically generated.
type RiotClient struct {
	internalClient *internal.InternalClient
	AccountV1      *AccountV1
}

// Creates a new RiotClient using the InternalClient provided.
func NewRiotClient(client *internal.InternalClient) *RiotClient {
	return &RiotClient{
		internalClient: client,
		AccountV1:      &AccountV1{internalClient: client},
	}
}
