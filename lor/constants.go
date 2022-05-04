package lor

type Region string

// Legends of Runeterra regions
const (
	Americas Region = "americas"
	Europe   Region = "europe"
	SEA      Region = "sea"
)

type GameMode string

const (
	ConstructedMode GameMode = "Constructed"
	ExpeditionsMode GameMode = "Expeditions"
	TutorialMode    GameMode = "Tutorial"
)

type GameType string

const (
	Ranked           GameType = "Ranked"
	Normal           GameType = "Normal"
	AI               GameType = "AI"
	Tutorial         GameType = "Tutorial"
	VanillaTrial     GameType = "VanillaTrial"
	Singleton        GameType = "Singleton"
	StandardGauntlet GameType = "StandardGauntlet"
)
