package api

type RiotRoute string

// Riot API routes
const (
	RouteAmericas RiotRoute = "americas"
	RouteEurope   RiotRoute = "europe"
	RouteAsia     RiotRoute = "asia"
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
