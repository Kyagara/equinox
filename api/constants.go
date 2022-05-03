package api

const (
	// Base API URL format
	BaseURLFormat = "https://%s.api.riotgames.com"
)

type RiotRoute string

// Riot API routes
const (
	RouteAmericas RiotRoute = "americas"
	RouteEurope   RiotRoute = "europe"
	RouteAsia     RiotRoute = "asia"
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
