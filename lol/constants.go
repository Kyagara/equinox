package lol

type Tier string

const (
	IronTier        Tier = "IRON"
	BronzeTier      Tier = "BRONZE"
	SilverTier      Tier = "SILVER"
	GoldTier        Tier = "GOLD"
	PlatinumTier    Tier = "PLATINUM"
	DiamondTier     Tier = "DIAMOND"
	MasterTier      Tier = "MASTER"
	GrandmasterTier Tier = "GRANDMASTER"
	ChallengerTier  Tier = "CHALLENGER"
)

type Region string

const (
	BR1  Region = "br1"
	EUN1 Region = "eun1"
	EUW1 Region = "euw1"
	JP1  Region = "jp1"
	KR   Region = "kr"
	LA1  Region = "la1"
	LA2  Region = "la2"
	NA1  Region = "na1"
	OC1  Region = "oc1"
	PBE1 Region = "PBE1"
	RU   Region = "ru"
	TR1  Region = "tr1"
)

type TournamentRegion string

const (
	BR     TournamentRegion = "BR"
	EUN    TournamentRegion = "EUNE"
	EUW    TournamentRegion = "EUW"
	JP     TournamentRegion = "JP"
	LAN    TournamentRegion = "LAN"
	LAS    TournamentRegion = "LAS"
	NA     TournamentRegion = "NA"
	OC     TournamentRegion = "OCE"
	PBE    TournamentRegion = "PBE"
	RUSSIA TournamentRegion = "RU"
	TR     TournamentRegion = "TR"
)

type QueueType string

const (
	Solo5x5Queue QueueType = "RANKED_SOLO_5x5"
	FlexSRQueue  QueueType = "RANKED_FLEX_SR"
	FlexTTQueue  QueueType = "RANKED_FLEX_TT"
)

type MatchType string

const (
	RankedMatch   MatchType = "ranked"
	NormalMatch   MatchType = "normal"
	TourneyMatch  MatchType = "tourney"
	TutorialMatch MatchType = "tutorial"
)

type MapType string

const (
	SummonersRiftMap   MapType = "SUMMONERS_RIFT"
	TwistedTreelineMap MapType = "TWISTED_TREELINE"
	HowlingAbyssMap    MapType = "HOWLING_ABYSS"
)

type PickType string

const (
	BlindPick           PickType = "BLIND_PICK"
	DraftModePick       PickType = "DRAFT_MODE"
	AllRandomPick       PickType = "ALL_RANDOM"
	TournamentDraftPick PickType = "TOURNAMENT_DRAFT"
)

type SpectatorType string

const (
	NoneSpectator      SpectatorType = "NONE"
	LobbyOnlySpectator SpectatorType = "LOBBYONLY"
	AllSpectator       SpectatorType = "ALL"
)

type GameType string

const (
	CustomGame   GameType = "CUSTOM_GAME"
	MatchGame    GameType = "MATCHED_GAME"
	TutorialGame GameType = "TUTORIAL_GAME"
)

type GameMode string

const (
	ClassicMode    GameMode = "CLASSIC"
	OdinMode       GameMode = "ODIN"
	AramMode       GameMode = "ARAM"
	TutorialMode   GameMode = "TUTORIAL"
	OneForAllMode  GameMode = "ONEFORALL"
	AscensionMode  GameMode = "ASCENSION"
	FirstBloodMode GameMode = "FIRSTBLOOD"
	KingPoroMode   GameMode = "KINGPORO"
)

type ChampionTransformation int8

const (
	NoTransformation       ChampionTransformation = 0
	SlayerTransformation   ChampionTransformation = 1
	AssassinTransformation ChampionTransformation = 2
)
