package api

// Base API URLs formats.
const (
	BaseURLFormat = "https://%s.api.riotgames.com"

	DataDragonURLFormat      = "https://ddragon.leagueoflegends.com%s"
	CommunityDragonURLFormat = "https://cdn.communitydragon.org%s"

	DataDragonLOLVersionURL = "/api/versions.json"
)

type Cluster string

// Riot API clusters, used in RiotClient and LOLClient.Tournament/TournamentStub.
const (
	AmericasCluster Cluster = "americas"
	EuropeCluster   Cluster = "europe"
	AsiaCluster     Cluster = "asia"
	EsportsCluster  Cluster = "esports"
)

// Rate limit headers
const (
	RateLimitTypeHeader = "X-Rate-Limit-Type"
	RetryAfterHeader    = "Retry-After"
)

type PublishLocation string

const (
	RiotClientLocation PublishLocation = "riotclient"
	RiotStatusLocation PublishLocation = "riotstatus"
	GameLocation       PublishLocation = "game"
)

type Platform string

const (
	WindowsPlatform Platform = "windows"
	MacOSPlatform   Platform = "macos"
	AndroidPlatform Platform = "android"
	IOSPlatform     Platform = "ios"
	PS4Platform     Platform = "ps4"
	XboxOnePlatform Platform = "xbone"
	SwitchPlatform  Platform = "switch"
)

type IncidentSeverity string

const (
	InfoSeverity     IncidentSeverity = "info"
	WarningSeverity  IncidentSeverity = "warning"
	CriticalSeverity IncidentSeverity = "critical"
)

type Game string

const (
	LOR Game = "lor"
	VAL Game = "val"
)

type LogLevel int8

const (
	// NopLevel won't log anything, this is the default behaviour for the Default client.
	NopLevel LogLevel = iota - 2
	// DebugLevel will log everything.
	DebugLevel
	// InfoLevel will log the requests being made and if they were successful.
	InfoLevel
	// WarnLevel will log when a request was rate limited.
	WarnLevel
	// ErrorLevel will log every error.
	ErrorLevel
)

type Division string

const (
	I   Division = "I"
	II  Division = "II"
	III Division = "III"
	IV  Division = "IV"
)
