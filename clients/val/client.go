// This package is used to interact with VAL endpoints.
//
// Automatically generated package.
package val

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
type VALClient struct {
	internalClient  *internal.InternalClient
	ContentV1  *ContentV1
	MatchV1  *MatchV1
	RankedV1  *RankedV1
	StatusV1  *StatusV1
}

// Creates a new VALClient using the InternalClient provided.
func NewVALClient(client *internal.InternalClient) *VALClient {
	return &VALClient{
        internalClient: client,
        ContentV1: &ContentV1{internalClient: client},
        MatchV1: &MatchV1{internalClient: client},
        RankedV1: &RankedV1{internalClient: client},
        StatusV1: &StatusV1{internalClient: client},
	}
}
