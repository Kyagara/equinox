package riot

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = a70746fcf353ba0ad0aceceafcc70d4ba8de4431

// account-v1.ActiveShardDto
type AccountActiveShardV1DTO struct {
	ActiveShard string `json:"activeShard,omitempty"`
	Game        string `json:"game,omitempty"`
	PUUID       string `json:"puuid,omitempty"`
}

// account-v1.AccountDto
type AccountV1DTO struct {
	// This field may be excluded from the response if the account doesn't have a gameName.
	GameName string `json:"gameName,omitempty"`
	PUUID    string `json:"puuid,omitempty"`
	// This field may be excluded from the response if the account doesn't have a tagLine.
	TagLine string `json:"tagLine,omitempty"`
}
