package api

const (
	// Base API URL format
	BaseURLFormat = "https://%s.api.riotgames.com"
)

type Division string

const (
	I   Division = "I"
	II  Division = "II"
	III Division = "III"
	IV  Division = "IV"
)

type LOLTier string

const (
	LOLTierIron     LOLTier = "IRON"
	LOLTierBronze   LOLTier = "BRONZE"
	LOLTierSilver   LOLTier = "SILVER"
	LOLTierGold     LOLTier = "GOLD"
	LOLTierPlatinum LOLTier = "PLATINUM"
	LOLTierDiamond  LOLTier = "DIAMOND"
)

type LOLQueueType string

const (
	RankedSoloQueueType   LOLQueueType = "RANKED_SOLO_5x5"
	RankedFlexSRQueueType LOLQueueType = "RANKED_FLEX_SR"
	RankedFlexTTQueueType LOLQueueType = "RANKED_FLEX_TT"
)

type LOLMatchType string

const (
	RankedMatchType   LOLMatchType = "ranked"
	NormalMatchType   LOLMatchType = "normal"
	TourneyMatchType  LOLMatchType = "tourney"
	TutorialMatchType LOLMatchType = "tutorial"
)
