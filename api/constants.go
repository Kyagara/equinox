package api

// Base API URL format.
const (
	BaseURLFormat = "https://%s.api.riotgames.com"
)

type Cluster string

// Riot API clusters.
const (
	Americas Cluster = "americas"
	Europe   Cluster = "europe"
	Asia     Cluster = "asia"
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
	DebugLevel LogLevel = -1
	InfoLevel  LogLevel = 0
	WarnLevel  LogLevel = 1
	ErrorLevel LogLevel = 2
	FatalLevel LogLevel = 5
)

type Division string

const (
	I   Division = "I"
	II  Division = "II"
	III Division = "III"
	IV  Division = "IV"
)
