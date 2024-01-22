// This package is used to interact with all Riot endpoints.
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

// Spec version = 82c5b64c16bd63688a0d19f471a19301bae8be4a

import "github.com/Kyagara/equinox/internal"

type Client struct {
	AccountV1 AccountV1
}

// Creates a new Riot Client using the internal.Client provided.
func NewRiotClient(client *internal.Client) *Client {
	return &Client{
		AccountV1: AccountV1{internal: client},
	}
}
