package riot

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = cd204d7d764a025c280943766bc498278e439a6c

// AccountDTO data object.
type AccountV1DTO struct {
	PUUID string `json:"puuid,omitempty"`
	// This field may be excluded from the response if the account doesn't have a gameName.
	GameName string `json:"gameName,omitempty"`
	// This field may be excluded from the response if the account doesn't have a tagLine.
	TagLine string `json:"tagLine,omitempty"`
}

// ActiveShardDTO data object.
type ActiveShardV1DTO struct {
	PUUID       string `json:"puuid,omitempty"`
	Game        string `json:"game,omitempty"`
	ActiveShard string `json:"activeShard,omitempty"`
}
