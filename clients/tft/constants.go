package tft

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 339cc5986ca34480f2ecf815246cade7105a897a

// LoL and TFT ranked tiers, such as gold, diamond, challenger, etc.
type Tier string

const (
	IRON        Tier = "IRON"
	BRONZE      Tier = "BRONZE"
	SILVER      Tier = "SILVER"
	GOLD        Tier = "GOLD"
	PLATINUM    Tier = "PLATINUM"
	EMERALD     Tier = "EMERALD"
	DIAMOND     Tier = "DIAMOND"
	MASTER      Tier = "MASTER"
	GRANDMASTER Tier = "GRANDMASTER"
	CHALLENGER  Tier = "CHALLENGER"
)

func (tier Tier) String() string {
	switch tier {
	case IRON:
		return "IRON"
	case BRONZE:
		return "BRONZE"
	case SILVER:
		return "SILVER"
	case GOLD:
		return "GOLD"
	case PLATINUM:
		return "PLATINUM"
	case EMERALD:
		return "EMERALD"
	case DIAMOND:
		return "DIAMOND"
	case MASTER:
		return "MASTER"
	case GRANDMASTER:
		return "GRANDMASTER"
	case CHALLENGER:
		return "CHALLENGER"
	default:
		return string(tier)
	}
}

// LoL and TFT rank divisions, I, II, III, IV, and (deprecated) V.
type Division string

const (
	I   Division = "I"
	II  Division = "II"
	III Division = "III"
	IV  Division = "IV"
	// Deprecated
	V Division = "V"
)

func (division Division) String() string {
	switch division {
	case I:
		return "I"
	case II:
		return "II"
	case III:
		return "III"
	case IV:
		return "IV"
	case V:
		return "V"
	default:
		return string(division)
	}
}

// Platform routes for League of Legends (LoL), Teamfight Tactics (TFT).
type PlatformRoute string

const (
	// Brazil.
	BR1 PlatformRoute = "br1"
	// Europe, Northeast.
	EUN1 PlatformRoute = "eun1"
	// Europe, West.
	EUW1 PlatformRoute = "euw1"
	// Japan.
	JP1 PlatformRoute = "jp1"
	// Korea.
	KR PlatformRoute = "kr"
	// Latin America, North.
	LA1 PlatformRoute = "la1"
	// Latin America, South.
	LA2 PlatformRoute = "la2"
	// North America.
	NA1 PlatformRoute = "na1"
	// Oceania.
	OC1 PlatformRoute = "oc1"
	// Public Beta Environment, special beta testing platform. Located in North America.
	PBE1 PlatformRoute = "pbe1"
	// Philippines
	PH2 PlatformRoute = "ph2"
	// Russia
	RU PlatformRoute = "ru"
	// Singapore
	SG2 PlatformRoute = "sg2"
	// Thailand
	TH2 PlatformRoute = "th2"
	// Turkey
	TR1 PlatformRoute = "tr1"
	// Taiwan
	TW2 PlatformRoute = "tw2"
	// Vietnam
	VN2 PlatformRoute = "vn2"
)

func (route PlatformRoute) String() string {
	switch route {
	case BR1:
		return "br1"
	case EUN1:
		return "eun1"
	case EUW1:
		return "euw1"
	case JP1:
		return "jp1"
	case KR:
		return "kr"
	case LA1:
		return "la1"
	case LA2:
		return "la2"
	case NA1:
		return "na1"
	case OC1:
		return "oc1"
	case PBE1:
		return "pbe1"
	case PH2:
		return "ph2"
	case RU:
		return "ru"
	case SG2:
		return "sg2"
	case TH2:
		return "th2"
	case TR1:
		return "tr1"
	case TW2:
		return "tw2"
	case VN2:
		return "vn2"
	default:
		return string(route)
	}
}

// TFT game type: matched game, custom game, or tutorial game.
type GameType string

const (
	// Custom games
	CUSTOM_GAME GameType = "CUSTOM_GAME"
	// all other games
	MATCHED_GAME GameType = "MATCHED_GAME"
	// Tutorial games
	TUTORIAL_GAME GameType = "TUTORIAL_GAME"
)

func (gameType GameType) String() string {
	switch gameType {
	case CUSTOM_GAME:
		return "CUSTOM_GAME"
	case MATCHED_GAME:
		return "MATCHED_GAME"
	case TUTORIAL_GAME:
		return "TUTORIAL_GAME"
	default:
		return string(gameType)
	}
}

// League of Legends game mode, such as Classic,
// ARAM, URF, One For All, Ascension, etc.
type GameMode string

const (
	// ARAM games
	ARAM GameMode = "ARAM"
	// All Random Summoner's Rift games
	ARSR GameMode = "ARSR"
	// Ascension games
	ASCENSION GameMode = "ASCENSION"
	// Blood Hunt Assassin games
	ASSASSINATE GameMode = "ASSASSINATE"
	// 2v2v2v2
	CHERRY GameMode = "CHERRY"
	// Classic Summoner's Rift and Twisted Treeline games
	CLASSIC GameMode = "CLASSIC"
	// Dark Star: Singularity games
	DARKSTAR GameMode = "DARKSTAR"
	// Doom Bot games
	DOOMBOTSTEEMO GameMode = "DOOMBOTSTEEMO"
	// Snowdown Showdown games
	FIRSTBLOOD GameMode = "FIRSTBLOOD"
	// Nexus Blitz games
	GAMEMODEX GameMode = "GAMEMODEX"
	// Legend of the Poro King games
	KINGPORO GameMode = "KINGPORO"
	// Nexus Blitz games
	NEXUSBLITZ GameMode = "NEXUSBLITZ"
	// Dominion/Crystal Scar games
	ODIN GameMode = "ODIN"
	// Odyssey: Extraction games
	ODYSSEY GameMode = "ODYSSEY"
	// One for All games
	ONEFORALL GameMode = "ONEFORALL"
	// Practice tool training games.
	PRACTICETOOL GameMode = "PRACTICETOOL"
	// PROJECT: Hunters games
	PROJECT GameMode = "PROJECT"
	// Nexus Siege games
	SIEGE GameMode = "SIEGE"
	// Star Guardian Invasion games
	STARGUARDIAN GameMode = "STARGUARDIAN"
	// Teamfight Tactics, used in `spectator-v4` endpoints.
	TFT GameMode = "TFT"
	// Tutorial games
	TUTORIAL GameMode = "TUTORIAL"
	// Tutorial: Welcome to League.
	TUTORIAL_MODULE_1 GameMode = "TUTORIAL_MODULE_1"
	// Tutorial: Power Up.
	TUTORIAL_MODULE_2 GameMode = "TUTORIAL_MODULE_2"
	// Tutorial: Shop for Gear.
	TUTORIAL_MODULE_3 GameMode = "TUTORIAL_MODULE_3"
	// Ultimate Spellbook games
	ULTBOOK GameMode = "ULTBOOK"
	// URF games
	URF GameMode = "URF"
)

func (gameMode GameMode) String() string {
	switch gameMode {
	case ARAM:
		return "ARAM"
	case ARSR:
		return "ARSR"
	case ASCENSION:
		return "ASCENSION"
	case ASSASSINATE:
		return "ASSASSINATE"
	case CHERRY:
		return "CHERRY"
	case CLASSIC:
		return "CLASSIC"
	case DARKSTAR:
		return "DARKSTAR"
	case DOOMBOTSTEEMO:
		return "DOOMBOTSTEEMO"
	case FIRSTBLOOD:
		return "FIRSTBLOOD"
	case GAMEMODEX:
		return "GAMEMODEX"
	case KINGPORO:
		return "KINGPORO"
	case NEXUSBLITZ:
		return "NEXUSBLITZ"
	case ODIN:
		return "ODIN"
	case ODYSSEY:
		return "ODYSSEY"
	case ONEFORALL:
		return "ONEFORALL"
	case PRACTICETOOL:
		return "PRACTICETOOL"
	case PROJECT:
		return "PROJECT"
	case SIEGE:
		return "SIEGE"
	case STARGUARDIAN:
		return "STARGUARDIAN"
	case TFT:
		return "TFT"
	case TUTORIAL:
		return "TUTORIAL"
	case TUTORIAL_MODULE_1:
		return "TUTORIAL_MODULE_1"
	case TUTORIAL_MODULE_2:
		return "TUTORIAL_MODULE_2"
	case TUTORIAL_MODULE_3:
		return "TUTORIAL_MODULE_3"
	case ULTBOOK:
		return "ULTBOOK"
	case URF:
		return "URF"
	default:
		return string(gameMode)
	}
}

// TFT ranked queue types.
type QueueType string

const (
	// Ranked Teamfight Tactics games
	RANKED_TFT QueueType = "RANKED_TFT"
	// Ranked Teamfight Tactics (Double Up Workshop) games
	RANKED_TFT_DOUBLE_UP QueueType = "RANKED_TFT_DOUBLE_UP"
	// Ranked Teamfight Tactics (Double Up Workshop) games
	//
	// Deprecated
	RANKED_TFT_PAIRS QueueType = "RANKED_TFT_PAIRS"
	// Ranked Teamfight Tactics (Hyper Roll) games
	RANKED_TFT_TURBO QueueType = "RANKED_TFT_TURBO"
)

func (queue QueueType) String() string {
	switch queue {
	case RANKED_TFT:
		return "RANKED_TFT"
	case RANKED_TFT_DOUBLE_UP:
		return "RANKED_TFT_DOUBLE_UP"
	case RANKED_TFT_PAIRS:
		return "RANKED_TFT_PAIRS"
	case RANKED_TFT_TURBO:
		return "RANKED_TFT_TURBO"
	default:
		return string(queue)
	}
}
