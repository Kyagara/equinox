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

type Division string

const (
	I   Division = "I"
	II  Division = "II"
	III Division = "III"
	IV  Division = "IV"
)
