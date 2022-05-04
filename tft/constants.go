package tft

type QueueType string

const (
	RankedTFTTurbo QueueType = "RANKED_TFT_TURBO"
	RankedTFT      QueueType = "RANKED_TFT"
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
	NoStyle   Style = 0
	Bronze    Style = 1
	Silver    Style = 2
	Gold      Style = 3
	Chromatic Style = 4
)
