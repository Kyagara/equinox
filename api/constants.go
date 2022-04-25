package api

type Region string

const (
	// Base API URL format
	BaseURLFormat = "https://%s.api.riotgames.com/"
)

// Riot API routes
const (
	RiotRouteAmericas Region = "americas"
	RiotRouteEurope   Region = "europe"
	RiotRouteAsia     Region = "asia"
)

// League of Legends and Teamfight Tactics regions
const (
	LOLRegionBR1  Region = "br1"
	LOLRegionEUN1 Region = "eun1"
	LOLRegionEUW1 Region = "euw1"
	LOLRegionJP1  Region = "jp1"
	LOLRegionKR   Region = "kr"
	LOLRegionLA1  Region = "la1"
	LOLRegionLA2  Region = "la2"
	LOLRegionNA1  Region = "na1"
	LOLRegionOC1  Region = "oc1"
	LOLRegionRU   Region = "ru"
	LOLRegionTR1  Region = "tr1"
)

// Valorant regions
const (
	VALRegionAP    Region = "ap"
	VALRegionBR    Region = "br"
	VALRegionEU    Region = "eu"
	VALRegionKR    Region = "kr"
	VALRegionLATAM Region = "latam"
	VALRegionNA    Region = "na"
)
