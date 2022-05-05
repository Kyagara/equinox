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
	RankedGame           GameType = "Ranked"
	NormalGame           GameType = "Normal"
	AIGame               GameType = "AI"
	TutorialGame         GameType = "Tutorial"
	VanillaTrialGame     GameType = "VanillaTrial"
	SingletonGame        GameType = "Singleton"
	StandardGauntletGame GameType = "StandardGauntlet"
)
