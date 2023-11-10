// Automatically generated package.
package riot

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 95a5cf31a385d91b952e19190af5a828d2e60ed8

// AccountDTO data object.
type AccountV1DTO struct {
    PUUID string `json:"puuid"`
    // This field may be excluded from the response if the account doesn't have a gameName.
    GameName string `json:"gameName"`
    // This field may be excluded from the response if the account doesn't have a tagLine.
    TagLine string `json:"tagLine"`
}

// ActiveShardDTO data object.
type ActiveShardV1DTO struct {
    PUUID string `json:"puuid"`
    Game string `json:"game"`
    ActiveShard string `json:"activeShard"`
}