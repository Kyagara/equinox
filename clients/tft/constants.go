package tft

type Route string

// Teamfight Tactics routes, used on the match endpoint.
const (
	Americas Route = "americas"
	Europe   Route = "europe"
	Asia     Route = "asia"
)

type QueueType string

const (
	RankedTFTTurboQueue QueueType = "RANKED_TFT_TURBO"
	RankedTFTQueue      QueueType = "RANKED_TFT"
)

type RatedTier string

const (
	OrangeTier RatedTier = "ORANGE"
	PurpleTier RatedTier = "PURPLE"
	BlueTier   RatedTier = "BLUE"
	GreenTier  RatedTier = "GREEN"
	GrayTier   RatedTier = "GRAY"
)

type Style int8

const (
	NoStyle        Style = 0
	BronzeStyle    Style = 1
	SilverStyle    Style = 2
	GoldStyle      Style = 3
	ChromaticStyle Style = 4
)
