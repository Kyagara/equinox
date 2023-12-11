package api

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = cd204d7d764a025c280943766bc498278e439a6c

// Base API URLs formats.
const (
	RIOT_API_BASE_URL_FORMAT = "https://%s.api.riotgames.com%s"
	D_DRAGON_BASE_URL_FORMAT = "https://ddragon.leagueoflegends.com%s%s"
	C_DRAGON_BASE_URL_FORMAT = "https://cdn.communitydragon.org%s%s"
)

// Regional routes, used in tournament services, Legends of Runeterra (LoR), and some other endpoints.
type RegionalRoute string

const (
	// North and South America.
	AMERICAS RegionalRoute = "americas"
	// Asia, used for LoL matches (`match-v5`) and TFT matches (`tft-match-v1`).
	ASIA RegionalRoute = "asia"
	// Europe.
	EUROPE RegionalRoute = "europe"
	// South East Asia, used for LoR, LoL matches (`match-v5`), and TFT matches (`tft-match-v1`).
	SEA RegionalRoute = "sea"
	// Asia-Pacific, deprecated, for some old matches in `lor-match-v1`.
	//
	// Deprecated
	APAC RegionalRoute = "apac"
	// Special esports platform for `account-v1`. Do not confuse with the `esports` Valorant platform route.
	ESPORTS RegionalRoute = "esports"
)
