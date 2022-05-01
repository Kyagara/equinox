package api

const (
	// Base API URL format
	BaseURLFormat = "https://%s.api.riotgames.com"
)

type Route string

// Riot API routes
const (
	RouteAmericas Route = "americas"
	RouteEurope   Route = "europe"
	RouteAsia     Route = "asia"
)

type LOLRegion string

// League of Legends and Teamfight Tactics regions
const (
	LOLRegionBR1  LOLRegion = "br1"
	LOLRegionEUN1 LOLRegion = "eun1"
	LOLRegionEUW1 LOLRegion = "euw1"
	LOLRegionJP1  LOLRegion = "jp1"
	LOLRegionKR   LOLRegion = "kr"
	LOLRegionLA1  LOLRegion = "la1"
	LOLRegionLA2  LOLRegion = "la2"
	LOLRegionNA1  LOLRegion = "na1"
	LOLRegionOC1  LOLRegion = "oc1"
	LOLRegionRU   LOLRegion = "ru"
	LOLRegionTR1  LOLRegion = "tr1"
)
