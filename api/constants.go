package api

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = a5a3a5f5d5f2a617a56302a0afac77c745e4fd56

// Base API URLs formats.
const (
	RIOT_API_BASE_URL_FORMAT = "https://%s.api.riotgames.com"
	D_DRAGON_BASE_URL_FORMAT = "https://ddragon.leagueoflegends.com%s"
	C_DRAGON_BASE_URL_FORMAT = "https://cdn.communitydragon.org%s"
)

// Rate limit headers
const (
	X_RATE_LIMIT_TYPE_HEADER = "X-Rate-Limit-Type"
	RETRY_AFTER_HEADER       = "Retry-After"
)

type LogLevel int8

const (
	// NOP_LOG_LEVEL won't log anything, this is the default behaviour for the Default client.
	NOP_LOG_LEVEL LogLevel = iota - 2
	// DEBUG_LOG_LEVEL will log everything.
	DEBUG_LOG_LEVEL
	// INFO_LOG_LEVEL will log the requests being made and if they were successful.
	INFO_LOG_LEVEL
	// WARN_LOG_LEVEL will log when a request was rate limited.
	WARN_LOG_LEVEL
	// ERROR_LOG_LEVEL will log every error.
	ERROR_LOG_LEVEL
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
	// Deprecated
	APAC RegionalRoute = "apac"
	// Special esports platform for `account-v1`. Do not confuse with the `esports` Valorant platform route.
	ESPORTS RegionalRoute = "esports"
)
