package lol

import "strconv"

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 996d171a2b79e9bb85c549f47b07c6ef2721fc8a

// Platform routes for League of Legends.
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
	// Middle East and North Africa.
	ME1 PlatformRoute = "me1"
	// North America.
	NA1 PlatformRoute = "na1"
	// Oceania.
	OC1 PlatformRoute = "oc1"
	// Public Beta Environment, special beta testing platform. Located in North America.
	PBE1 PlatformRoute = "pbe1"
	// # Deprecated
	//
	// Philippines, moved into `sg2` on 2025-01-08.
	PH2 PlatformRoute = "ph2"
	// Russia
	RU PlatformRoute = "ru"
	// Singapore, Thailand, Philippines
	SG2 PlatformRoute = "sg2"
	// # Deprecated
	//
	// Thailand, moved into `sg2` on 2025-01-08.
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
	case ME1:
		return "me1"
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

// Tournament regions for League of Legends.
type TournamentRegion string

const (
	// Brazil.
	BR TournamentRegion = "br"
	// Europe, Northeast.
	EUNE TournamentRegion = "eune"
	// Europe, West.
	EUW TournamentRegion = "euw"
	// Japan.
	JP TournamentRegion = "jp"
	// Latin America, North.
	LAN TournamentRegion = "lan"
	// Latin America, South.
	LAS TournamentRegion = "las"
	// North America.
	NA TournamentRegion = "na"
	// Oceania.
	OCE TournamentRegion = "oce"
	// Public Beta Environment, special beta testing platform. Located in North America.
	PBE TournamentRegion = "pbe"
	// Turkey
	TR TournamentRegion = "tr"
)

func (route TournamentRegion) String() string {
	switch route {
	case BR:
		return "br"
	case EUNE:
		return "eune"
	case EUW:
		return "euw"
	case JP:
		return "jp"
	case LAN:
		return "lan"
	case LAS:
		return "las"
	case NA:
		return "na"
	case OCE:
		return "oce"
	case PBE:
		return "pbe"
	case TR:
		return "tr"
	default:
		return string(route)
	}
}

// League of Legends ranked tiers, such as gold, diamond, challenger, etc.
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

// League of Legends rank divisions, I, II, III, IV, and (deprecated) V.
type Division string

const (
	I   Division = "I"
	II  Division = "II"
	III Division = "III"
	IV  Division = "IV"
	// # Deprecated
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

// Team IDs for League of Legends.
type Team int

// https://github.com/MingweiSamuel/Riven/blob/v/2.x.x/riven/src/consts/team.rs

const (
	// Team ID zero for 2v2v2v2 Arena `CHERRY` game mode. (TODO: SUBJECT TO CHANGE?)
	ZERO Team = 0
	// Blue team (bottom left on Summoner's Rift).
	BLUE Team = 100
	// Red team (top right on Summoner's Rift).
	RED Team = 200
	// "killerTeamId" when Baron Nashor spawns and kills Rift Herald.
	OTHER Team = 300
)

func (team Team) String() string {
	switch team {
	case ZERO:
		return "0"
	case BLUE:
		return "100"
	case RED:
		return "200"
	case OTHER:
		return "300"
	default:
		return strconv.FormatInt(int64(team), 10)
	}
}

// League of Legends game types: matched game, custom game, or tutorial game.
type GameType string

const (
	// Custom games
	CUSTOM_GAME_GAMETYPE GameType = "CUSTOM"
	// all other games
	MATCHED_GAME_GAMETYPE GameType = "MATCHED"
	// Tutorial games
	TUTORIAL_GAME_GAMETYPE GameType = "TUTORIAL"
)

func (gameType GameType) String() string {
	switch gameType {
	case CUSTOM_GAME_GAMETYPE:
		return "CUSTOM"
	case MATCHED_GAME_GAMETYPE:
		return "MATCHED"
	case TUTORIAL_GAME_GAMETYPE:
		return "TUTORIAL"
	default:
		return string(gameType)
	}
}

// League of Legends ranked queue types.
type QueueType string

const (
	// "Arena" games
	CHERRY_QUEUETYPE QueueType = "CHERRY"
	// 5v5 Ranked Flex games
	RANKED_FLEX_SR_QUEUETYPE QueueType = "RANKED_FLEX_SR"
	// # Deprecated
	//
	// 3v3 Ranked Flex games
	RANKED_FLEX_TT_QUEUETYPE QueueType = "RANKED_FLEX_TT"
	// 5v5 Ranked Solo games
	RANKED_SOLO_5X5_QUEUETYPE QueueType = "RANKED_SOLO_5X5"
)

func (queueType QueueType) String() string {
	switch queueType {
	case CHERRY_QUEUETYPE:
		return "CHERRY"
	case RANKED_FLEX_SR_QUEUETYPE:
		return "RANKED_FLEX_SR"
	case RANKED_FLEX_TT_QUEUETYPE:
		return "RANKED_FLEX_TT"
	case RANKED_SOLO_5X5_QUEUETYPE:
		return "RANKED_SOLO_5X5"
	default:
		return string(queueType)
	}
}

// League of Legends game modes, such as Classic,
// ARAM, URF, One For All, Ascension, etc.
type GameMode string

const (
	// ARAM games
	ARAM_GAMEMODE GameMode = "ARAM"
	// All Random Summoner's Rift games
	ARSR_GAMEMODE GameMode = "ARSR"
	// Ascension games
	ASCENSION_GAMEMODE GameMode = "ASCENSION"
	// Blood Hunt Assassin games
	ASSASSINATE_GAMEMODE GameMode = "ASSASSINATE"
	// 2v2v2v2 Arena
	CHERRY_GAMEMODE GameMode = "CHERRY"
	// Classic Summoner's Rift and Twisted Treeline games
	CLASSIC_GAMEMODE GameMode = "CLASSIC"
	// Dark Star: Singularity games
	DARKSTAR_GAMEMODE GameMode = "DARKSTAR"
	// Doom Bot games
	DOOMBOTSTEEMO_GAMEMODE GameMode = "DOOMBOTSTEEMO"
	// Snowdown Showdown games
	FIRSTBLOOD_GAMEMODE GameMode = "FIRSTBLOOD"
	// Nexus Blitz games
	GAMEMODEX_GAMEMODE GameMode = "GAMEMODEX"
	// Legend of the Poro King games
	KINGPORO_GAMEMODE GameMode = "KINGPORO"
	// Nexus Blitz games
	NEXUSBLITZ_GAMEMODE GameMode = "NEXUSBLITZ"
	// Dominion/Crystal Scar games
	ODIN_GAMEMODE GameMode = "ODIN"
	// Odyssey: Extraction games
	ODYSSEY_GAMEMODE GameMode = "ODYSSEY"
	// One for All games
	ONEFORALL_GAMEMODE GameMode = "ONEFORALL"
	// Practice tool training games.
	PRACTICETOOL_GAMEMODE GameMode = "PRACTICETOOL"
	// PROJECT: Hunters games
	PROJECT_GAMEMODE GameMode = "PROJECT"
	// Nexus Siege games
	SIEGE_GAMEMODE GameMode = "SIEGE"
	// Star Guardian Invasion games
	STARGUARDIAN_GAMEMODE GameMode = "STARGUARDIAN"
	// Swarm
	STRAWBERRY_GAMEMODE GameMode = "STRAWBERRY"
	// Swiftplay Summoner's Rift
	SWIFTPLAY_GAMEMODE GameMode = "SWIFTPLAY"
	// Tutorial games
	TUTORIAL_GAMEMODE GameMode = "TUTORIAL"
	// Tutorial: Welcome to League.
	TUTORIAL_MODULE_1_GAMEMODE GameMode = "TUTORIAL_MODULE_1"
	// Tutorial: Power Up.
	TUTORIAL_MODULE_2_GAMEMODE GameMode = "TUTORIAL_MODULE_2"
	// Tutorial: Shop for Gear.
	TUTORIAL_MODULE_3_GAMEMODE GameMode = "TUTORIAL_MODULE_3"
	// Ultimate Spellbook games
	ULTBOOK_GAMEMODE GameMode = "ULTBOOK"
	// URF games
	URF_GAMEMODE GameMode = "URF"
)

func (gameMode GameMode) String() string {
	switch gameMode {
	case ARAM_GAMEMODE:
		return "ARAM"
	case ARSR_GAMEMODE:
		return "ARSR"
	case ASCENSION_GAMEMODE:
		return "ASCENSION"
	case ASSASSINATE_GAMEMODE:
		return "ASSASSINATE"
	case CHERRY_GAMEMODE:
		return "CHERRY"
	case CLASSIC_GAMEMODE:
		return "CLASSIC"
	case DARKSTAR_GAMEMODE:
		return "DARKSTAR"
	case DOOMBOTSTEEMO_GAMEMODE:
		return "DOOMBOTSTEEMO"
	case FIRSTBLOOD_GAMEMODE:
		return "FIRSTBLOOD"
	case GAMEMODEX_GAMEMODE:
		return "GAMEMODEX"
	case KINGPORO_GAMEMODE:
		return "KINGPORO"
	case NEXUSBLITZ_GAMEMODE:
		return "NEXUSBLITZ"
	case ODIN_GAMEMODE:
		return "ODIN"
	case ODYSSEY_GAMEMODE:
		return "ODYSSEY"
	case ONEFORALL_GAMEMODE:
		return "ONEFORALL"
	case PRACTICETOOL_GAMEMODE:
		return "PRACTICETOOL"
	case PROJECT_GAMEMODE:
		return "PROJECT"
	case SIEGE_GAMEMODE:
		return "SIEGE"
	case STARGUARDIAN_GAMEMODE:
		return "STARGUARDIAN"
	case STRAWBERRY_GAMEMODE:
		return "STRAWBERRY"
	case SWIFTPLAY_GAMEMODE:
		return "SWIFTPLAY"
	case TUTORIAL_GAMEMODE:
		return "TUTORIAL"
	case TUTORIAL_MODULE_1_GAMEMODE:
		return "TUTORIAL_MODULE_1"
	case TUTORIAL_MODULE_2_GAMEMODE:
		return "TUTORIAL_MODULE_2"
	case TUTORIAL_MODULE_3_GAMEMODE:
		return "TUTORIAL_MODULE_3"
	case ULTBOOK_GAMEMODE:
		return "ULTBOOK"
	case URF_GAMEMODE:
		return "URF"
	default:
		return string(gameMode)
	}
}

// League of Legends maps.
type Map int

const (
	// Arena
	//
	// Map for 2v2v2v2 (`CHERRY`). Team up with a friend or venture solo in this new game mode. Face against multiple teams in chaotic battles across diverse arenas
	ARENA_MAP Map = 30
	// Butcher's Bridge
	//
	// Alternate ARAM map
	BUTCHERS_BRIDGE_MAP Map = 14
	// Cosmic Ruins
	//
	// Dark Star: Singularity map
	COSMIC_RUINS_MAP Map = 16
	// Crash Site
	//
	// Odyssey: Extraction map
	CRASH_SITE_MAP Map = 20
	// Howling Abyss
	//
	// ARAM map
	HOWLING_ABYSS_MAP Map = 12
	// Nexus Blitz
	//
	// Nexus Blitz map
	NEXUS_BLITZ_MAP Map = 21
	// Substructure 43
	//
	// PROJECT: Hunters map
	SUBSTRUCTURE_43_MAP Map = 19
	// Summoner's Rift
	//
	// Current Version
	SUMMONERS_RIFT_MAP Map = 11
	// # Deprecated
	//
	// Summoner's Rift
	//
	// Original Autumn variant
	SUMMONERS_RIFT_ORIGINAL_AUTUMN_VARIANT_MAP Map = 2
	// # Deprecated
	//
	// Summoner's Rift
	//
	// Original Summer variant
	SUMMONERS_RIFT_ORIGINAL_SUMMER_VARIANT_MAP Map = 1
	// Swarm
	//
	// Map for Swarm (`STRAWBERRY`). Team up with a friend or venture solo in this horde survival mode.
	SWARM_MAP Map = 33
	// The Crystal Scar
	//
	// Dominion map
	THE_CRYSTAL_SCAR_MAP Map = 8
	// The Proving Grounds
	//
	// Tutorial Map
	THE_PROVING_GROUNDS_MAP Map = 3
	// Twisted Treeline
	//
	// Last TT map
	TWISTED_TREELINE_MAP Map = 10
	// # Deprecated
	//
	// Twisted Treeline
	//
	// Original Version
	TWISTED_TREELINE_ORIGINAL_VERSION_MAP Map = 4
	// Valoran City Park
	//
	// Star Guardian Invasion map
	VALORAN_CITY_PARK_MAP Map = 18
)

func (gameMap Map) String() string {
	switch gameMap {
	case ARENA_MAP:
		return "30"
	case BUTCHERS_BRIDGE_MAP:
		return "14"
	case COSMIC_RUINS_MAP:
		return "16"
	case CRASH_SITE_MAP:
		return "20"
	case HOWLING_ABYSS_MAP:
		return "12"
	case NEXUS_BLITZ_MAP:
		return "21"
	case SUBSTRUCTURE_43_MAP:
		return "19"
	case SUMMONERS_RIFT_MAP:
		return "11"
	case SUMMONERS_RIFT_ORIGINAL_AUTUMN_VARIANT_MAP:
		return "2"
	case SUMMONERS_RIFT_ORIGINAL_SUMMER_VARIANT_MAP:
		return "1"
	case SWARM_MAP:
		return "33"
	case THE_CRYSTAL_SCAR_MAP:
		return "8"
	case THE_PROVING_GROUNDS_MAP:
		return "3"
	case TWISTED_TREELINE_MAP:
		return "10"
	case TWISTED_TREELINE_ORIGINAL_VERSION_MAP:
		return "4"
	case VALORAN_CITY_PARK_MAP:
		return "18"
	default:
		return strconv.FormatInt(int64(gameMap), 10)
	}
}

// League of Legends queues.
type Queue int

const (
	// 2v2v2v2 `CHERRY` games on Arena
	ARENA_2V2V2V2_CHERRY_QUEUE Queue = 1700
	// 5v5 ARAM games on Butcher's Bridge
	BUTCHERS_BRIDGE_5V5_ARAM_QUEUE Queue = 100
	// Dark Star: Singularity games on Cosmic Ruins
	COSMIC_RUINS_DARK_STAR_SINGULARITY_QUEUE Queue = 610
	// Odyssey Extraction: Cadet games on Crash Site
	CRASH_SITE_ODYSSEY_EXTRACTION_CADET_QUEUE Queue = 1040
	// Odyssey Extraction: Captain games on Crash Site
	CRASH_SITE_ODYSSEY_EXTRACTION_CAPTAIN_QUEUE Queue = 1060
	// Odyssey Extraction: Crewmember games on Crash Site
	CRASH_SITE_ODYSSEY_EXTRACTION_CREWMEMBER_QUEUE Queue = 1050
	// Odyssey Extraction: Intro games on Crash Site
	CRASH_SITE_ODYSSEY_EXTRACTION_INTRO_QUEUE Queue = 1030
	// Odyssey Extraction: Onslaught games on Crash Site
	CRASH_SITE_ODYSSEY_EXTRACTION_ONSLAUGHT_QUEUE Queue = 1070
	// # Deprecated
	//
	// 5v5 Dominion Blind Pick games on Crystal Scar
	CRYSTAL_SCAR_5V5_DOMINION_BLIND_PICK_QUEUE Queue = 16
	// # Deprecated
	//
	// 5v5 Dominion Draft Pick games on Crystal Scar
	CRYSTAL_SCAR_5V5_DOMINION_DRAFT_PICK_QUEUE Queue = 17
	// # Deprecated
	//
	// Ascension games on Crystal Scar
	CRYSTAL_SCAR_ASCENSION_96_QUEUE Queue = 96
	// Ascension games on Crystal Scar
	CRYSTAL_SCAR_ASCENSION_QUEUE Queue = 910
	// Definitely Not Dominion games on Crystal Scar
	CRYSTAL_SCAR_DEFINITELY_NOT_DOMINION_QUEUE Queue = 317
	// # Deprecated
	//
	// Dominion Co-op vs AI games on Crystal Scar
	CRYSTAL_SCAR_DOMINION_CO_OP_VS_AI_QUEUE Queue = 25
	// Games on Custom games
	CUSTOM_QUEUE Queue = 0
	// 1v1 Snowdown Showdown games on Howling Abyss
	HOWLING_ABYSS_1V1_SNOWDOWN_SHOWDOWN_QUEUE Queue = 72
	// 2v2 Snowdown Showdown games on Howling Abyss
	HOWLING_ABYSS_2V2_SNOWDOWN_SHOWDOWN_QUEUE Queue = 73
	// # Deprecated
	//
	// 5v5 ARAM games on Howling Abyss
	HOWLING_ABYSS_5V5_ARAM_65_QUEUE Queue = 65
	// 5v5 ARAM games on Howling Abyss
	HOWLING_ABYSS_5V5_ARAM_QUEUE Queue = 450
	// ARAM Clash games on Howling Abyss
	HOWLING_ABYSS_ARAM_CLASH_QUEUE Queue = 720
	// # Deprecated
	//
	// ARAM Co-op vs AI games on Howling Abyss
	HOWLING_ABYSS_ARAM_CO_OP_VS_AI_QUEUE Queue = 67
	// # Deprecated
	//
	// Legend of the Poro King games on Howling Abyss
	HOWLING_ABYSS_LEGEND_OF_THE_PORO_KING_300_QUEUE Queue = 300
	// Legend of the Poro King games on Howling Abyss
	HOWLING_ABYSS_LEGEND_OF_THE_PORO_KING_QUEUE Queue = 920
	// One For All: Mirror Mode games on Howling Abyss
	HOWLING_ABYSS_ONE_FOR_ALL_MIRROR_MODE_QUEUE Queue = 78
	// # Deprecated
	//
	// Nexus Blitz games on Nexus Blitz
	NEXUS_BLITZ_1200_QUEUE Queue = 1200
	// Nexus Blitz games on Nexus Blitz
	NEXUS_BLITZ_QUEUE Queue = 1300
	// PROJECT: Hunters games on Overcharge
	OVERCHARGE_PROJECT_HUNTERS_QUEUE Queue = 1000
	// Arena (`CHERRY` games) games on Rings of Wrath
	RINGS_OF_WRATH_ARENA_CHERRY_GAMES_QUEUE Queue = 1710
	// # Deprecated
	//
	// 5v5 Blind Pick games on Summoner's Rift
	SUMMONERS_RIFT_5V5_BLIND_PICK_2_QUEUE Queue = 2
	// 5v5 Blind Pick games on Summoner's Rift
	SUMMONERS_RIFT_5V5_BLIND_PICK_QUEUE Queue = 430
	// # Deprecated
	//
	// 5v5 Draft Pick games on Summoner's Rift
	SUMMONERS_RIFT_5V5_DRAFT_PICK_14_QUEUE Queue = 14
	// 5v5 Draft Pick games on Summoner's Rift
	SUMMONERS_RIFT_5V5_DRAFT_PICK_QUEUE Queue = 400
	// # Deprecated
	//
	// 5v5 Ranked Dynamic games on Summoner's Rift
	SUMMONERS_RIFT_5V5_RANKED_DYNAMIC_QUEUE Queue = 410
	// 5v5 Ranked Flex games on Summoner's Rift
	SUMMONERS_RIFT_5V5_RANKED_FLEX_QUEUE Queue = 440
	// # Deprecated
	//
	// 5v5 Ranked Premade games on Summoner's Rift
	SUMMONERS_RIFT_5V5_RANKED_PREMADE_QUEUE Queue = 6
	// # Deprecated
	//
	// 5v5 Ranked Solo games on Summoner's Rift
	SUMMONERS_RIFT_5V5_RANKED_SOLO_4_QUEUE Queue = 4
	// 5v5 Ranked Solo games on Summoner's Rift
	SUMMONERS_RIFT_5V5_RANKED_SOLO_QUEUE Queue = 420
	// # Deprecated
	//
	// 5v5 Ranked Team games on Summoner's Rift
	SUMMONERS_RIFT_5V5_RANKED_TEAM_QUEUE Queue = 42
	// # Deprecated
	//
	// 5v5 Team Builder games on Summoner's Rift
	SUMMONERS_RIFT_5V5_TEAM_BUILDER_QUEUE Queue = 61
	// 6v6 Hexakill games on Summoner's Rift
	SUMMONERS_RIFT_6V6_HEXAKILL_QUEUE Queue = 75
	// All Random games on Summoner's Rift
	SUMMONERS_RIFT_ALL_RANDOM_QUEUE Queue = 325
	// # Deprecated
	//
	// ARURF games on Summoner's Rift
	SUMMONERS_RIFT_ARURF_318_QUEUE Queue = 318
	// ARURF games on Summoner's Rift
	SUMMONERS_RIFT_ARURF_QUEUE Queue = 900
	// Black Market Brawlers games on Summoner's Rift
	SUMMONERS_RIFT_BLACK_MARKET_BRAWLERS_QUEUE Queue = 313
	// Blood Hunt Assassin games on Summoner's Rift
	SUMMONERS_RIFT_BLOOD_HUNT_ASSASSIN_QUEUE Queue = 600
	// Summoner's Rift Clash games on Summoner's Rift
	SUMMONERS_RIFT_CLASH_QUEUE Queue = 700
	// # Deprecated
	//
	// Co-op vs AI Beginner Bot games on Summoner's Rift
	SUMMONERS_RIFT_CO_OP_VS_AI_BEGINNER_BOT_32_QUEUE Queue = 32
	// # Deprecated
	//
	// Co-op vs. AI Beginner Bot games on Summoner's Rift
	SUMMONERS_RIFT_CO_OP_VS_AI_BEGINNER_BOT_840_QUEUE Queue = 840
	// Co-op vs. AI Beginner Bot games on Summoner's Rift
	SUMMONERS_RIFT_CO_OP_VS_AI_BEGINNER_BOT_QUEUE Queue = 880
	// # Deprecated
	//
	// Co-op vs AI Intermediate Bot games on Summoner's Rift
	SUMMONERS_RIFT_CO_OP_VS_AI_INTERMEDIATE_BOT_33_QUEUE Queue = 33
	// # Deprecated
	//
	// Co-op vs. AI Intermediate Bot games on Summoner's Rift
	SUMMONERS_RIFT_CO_OP_VS_AI_INTERMEDIATE_BOT_850_QUEUE Queue = 850
	// Co-op vs. AI Intermediate Bot games on Summoner's Rift
	SUMMONERS_RIFT_CO_OP_VS_AI_INTERMEDIATE_BOT_QUEUE Queue = 890
	// # Deprecated
	//
	// Co-op vs AI Intro Bot games on Summoner's Rift
	SUMMONERS_RIFT_CO_OP_VS_AI_INTRO_BOT_31_QUEUE Queue = 31
	// # Deprecated
	//
	// Co-op vs. AI Intro Bot games on Summoner's Rift
	SUMMONERS_RIFT_CO_OP_VS_AI_INTRO_BOT_830_QUEUE Queue = 830
	// Co-op vs. AI Intro Bot games on Summoner's Rift
	SUMMONERS_RIFT_CO_OP_VS_AI_INTRO_BOT_QUEUE Queue = 870
	// # Deprecated
	//
	// Co-op vs AI games on Summoner's Rift
	SUMMONERS_RIFT_CO_OP_VS_AI_QUEUE Queue = 7
	// Co-op vs AI Ultra Rapid Fire games on Summoner's Rift
	SUMMONERS_RIFT_CO_OP_VS_AI_ULTRA_RAPID_FIRE_QUEUE Queue = 83
	// # Deprecated
	//
	// Doom Bots Rank 1 games on Summoner's Rift
	SUMMONERS_RIFT_DOOM_BOTS_RANK_1_QUEUE Queue = 91
	// # Deprecated
	//
	// Doom Bots Rank 2 games on Summoner's Rift
	SUMMONERS_RIFT_DOOM_BOTS_RANK_2_QUEUE Queue = 92
	// # Deprecated
	//
	// Doom Bots Rank 5 games on Summoner's Rift
	SUMMONERS_RIFT_DOOM_BOTS_RANK_5_QUEUE Queue = 93
	// Doom Bots Standard games on Summoner's Rift
	SUMMONERS_RIFT_DOOM_BOTS_STANDARD_QUEUE Queue = 960
	// Doom Bots Voting games on Summoner's Rift
	SUMMONERS_RIFT_DOOM_BOTS_VOTING_QUEUE Queue = 950
	// Nemesis games on Summoner's Rift
	SUMMONERS_RIFT_NEMESIS_QUEUE Queue = 310
	// # Deprecated
	//
	// Nexus Siege games on Summoner's Rift
	SUMMONERS_RIFT_NEXUS_SIEGE_315_QUEUE Queue = 315
	// Nexus Siege games on Summoner's Rift
	SUMMONERS_RIFT_NEXUS_SIEGE_QUEUE Queue = 940
	// Normal (Quickplay) games on Summoner's Rift
	SUMMONERS_RIFT_NORMAL_QUICKPLAY_QUEUE Queue = 490
	// Normal (Swiftplay) games on Summoner's Rift
	SUMMONERS_RIFT_NORMAL_SWIFTPLAY_QUEUE Queue = 480
	// # Deprecated
	//
	// One for All games on Summoner's Rift
	SUMMONERS_RIFT_ONE_FOR_ALL_70_QUEUE Queue = 70
	// One for All games on Summoner's Rift
	SUMMONERS_RIFT_ONE_FOR_ALL_QUEUE Queue = 1020
	// Pick URF games on Summoner's Rift
	SUMMONERS_RIFT_PICK_URF_QUEUE Queue = 1900
	// Snow ARURF games on Summoner's Rift
	SUMMONERS_RIFT_SNOW_ARURF_QUEUE Queue = 1010
	// Tutorial 1 games on Summoner's Rift
	SUMMONERS_RIFT_TUTORIAL_1_QUEUE Queue = 2000
	// Tutorial 2 games on Summoner's Rift
	SUMMONERS_RIFT_TUTORIAL_2_QUEUE Queue = 2010
	// Tutorial 3 games on Summoner's Rift
	SUMMONERS_RIFT_TUTORIAL_3_QUEUE Queue = 2020
	// Ultimate Spellbook games on Summoner's Rift
	SUMMONERS_RIFT_ULTIMATE_SPELLBOOK_QUEUE Queue = 1400
	// Ultra Rapid Fire games on Summoner's Rift
	SUMMONERS_RIFT_ULTRA_RAPID_FIRE_QUEUE Queue = 76
	// Swarm duo (`STRAWBERRY` games) games on Swarm
	SWARM_DUO_STRAWBERRY_GAMES_QUEUE Queue = 1820
	// Swarm quad (`STRAWBERRY` games) games on Swarm
	SWARM_QUAD_STRAWBERRY_GAMES_QUEUE Queue = 1840
	// Swarm solo (`STRAWBERRY` games) games on Swarm
	SWARM_SOLO_STRAWBERRY_GAMES_QUEUE Queue = 1810
	// Swarm trio (`STRAWBERRY` games) games on Swarm
	SWARM_TRIO_STRAWBERRY_GAMES_QUEUE Queue = 1830
	// # Deprecated
	//
	// 3v3 Blind Pick games on Twisted Treeline
	TWISTED_TREELINE_3V3_BLIND_PICK_QUEUE Queue = 460
	// # Deprecated
	//
	// 3v3 Normal games on Twisted Treeline
	TWISTED_TREELINE_3V3_NORMAL_QUEUE Queue = 8
	// # Deprecated
	//
	// 3v3 Ranked Flex games on Twisted Treeline
	TWISTED_TREELINE_3V3_RANKED_FLEX_470_QUEUE Queue = 470
	// # Deprecated
	//
	// 3v3 Ranked Flex games on Twisted Treeline
	TWISTED_TREELINE_3V3_RANKED_FLEX_9_QUEUE Queue = 9
	// # Deprecated
	//
	// 3v3 Ranked Team games on Twisted Treeline
	TWISTED_TREELINE_3V3_RANKED_TEAM_QUEUE Queue = 41
	// 6v6 Hexakill games on Twisted Treeline
	TWISTED_TREELINE_6V6_HEXAKILL_QUEUE Queue = 98
	// Co-op vs. AI Beginner Bot games on Twisted Treeline
	TWISTED_TREELINE_CO_OP_VS_AI_BEGINNER_BOT_QUEUE Queue = 820
	// # Deprecated
	//
	// Co-op vs. AI Intermediate Bot games on Twisted Treeline
	TWISTED_TREELINE_CO_OP_VS_AI_INTERMEDIATE_BOT_QUEUE Queue = 800
	// # Deprecated
	//
	// Co-op vs. AI Intro Bot games on Twisted Treeline
	TWISTED_TREELINE_CO_OP_VS_AI_INTRO_BOT_QUEUE Queue = 810
	// # Deprecated
	//
	// Co-op vs AI games on Twisted Treeline
	TWISTED_TREELINE_CO_OP_VS_AI_QUEUE Queue = 52
	// Star Guardian Invasion: Normal games on Valoran City Park
	VALORAN_CITY_PARK_STAR_GUARDIAN_INVASION_NORMAL_QUEUE Queue = 980
	// Star Guardian Invasion: Onslaught games on Valoran City Park
	VALORAN_CITY_PARK_STAR_GUARDIAN_INVASION_ONSLAUGHT_QUEUE Queue = 990
)

func (queue Queue) String() string {
	switch queue {
	case ARENA_2V2V2V2_CHERRY_QUEUE:
		return "1700"
	case BUTCHERS_BRIDGE_5V5_ARAM_QUEUE:
		return "100"
	case COSMIC_RUINS_DARK_STAR_SINGULARITY_QUEUE:
		return "610"
	case CRASH_SITE_ODYSSEY_EXTRACTION_CADET_QUEUE:
		return "1040"
	case CRASH_SITE_ODYSSEY_EXTRACTION_CAPTAIN_QUEUE:
		return "1060"
	case CRASH_SITE_ODYSSEY_EXTRACTION_CREWMEMBER_QUEUE:
		return "1050"
	case CRASH_SITE_ODYSSEY_EXTRACTION_INTRO_QUEUE:
		return "1030"
	case CRASH_SITE_ODYSSEY_EXTRACTION_ONSLAUGHT_QUEUE:
		return "1070"
	case CRYSTAL_SCAR_5V5_DOMINION_BLIND_PICK_QUEUE:
		return "16"
	case CRYSTAL_SCAR_5V5_DOMINION_DRAFT_PICK_QUEUE:
		return "17"
	case CRYSTAL_SCAR_ASCENSION_96_QUEUE:
		return "96"
	case CRYSTAL_SCAR_ASCENSION_QUEUE:
		return "910"
	case CRYSTAL_SCAR_DEFINITELY_NOT_DOMINION_QUEUE:
		return "317"
	case CRYSTAL_SCAR_DOMINION_CO_OP_VS_AI_QUEUE:
		return "25"
	case CUSTOM_QUEUE:
		return "0"
	case HOWLING_ABYSS_1V1_SNOWDOWN_SHOWDOWN_QUEUE:
		return "72"
	case HOWLING_ABYSS_2V2_SNOWDOWN_SHOWDOWN_QUEUE:
		return "73"
	case HOWLING_ABYSS_5V5_ARAM_65_QUEUE:
		return "65"
	case HOWLING_ABYSS_5V5_ARAM_QUEUE:
		return "450"
	case HOWLING_ABYSS_ARAM_CLASH_QUEUE:
		return "720"
	case HOWLING_ABYSS_ARAM_CO_OP_VS_AI_QUEUE:
		return "67"
	case HOWLING_ABYSS_LEGEND_OF_THE_PORO_KING_300_QUEUE:
		return "300"
	case HOWLING_ABYSS_LEGEND_OF_THE_PORO_KING_QUEUE:
		return "920"
	case HOWLING_ABYSS_ONE_FOR_ALL_MIRROR_MODE_QUEUE:
		return "78"
	case NEXUS_BLITZ_1200_QUEUE:
		return "1200"
	case NEXUS_BLITZ_QUEUE:
		return "1300"
	case OVERCHARGE_PROJECT_HUNTERS_QUEUE:
		return "1000"
	case RINGS_OF_WRATH_ARENA_CHERRY_GAMES_QUEUE:
		return "1710"
	case SUMMONERS_RIFT_5V5_BLIND_PICK_2_QUEUE:
		return "2"
	case SUMMONERS_RIFT_5V5_BLIND_PICK_QUEUE:
		return "430"
	case SUMMONERS_RIFT_5V5_DRAFT_PICK_14_QUEUE:
		return "14"
	case SUMMONERS_RIFT_5V5_DRAFT_PICK_QUEUE:
		return "400"
	case SUMMONERS_RIFT_5V5_RANKED_DYNAMIC_QUEUE:
		return "410"
	case SUMMONERS_RIFT_5V5_RANKED_FLEX_QUEUE:
		return "440"
	case SUMMONERS_RIFT_5V5_RANKED_PREMADE_QUEUE:
		return "6"
	case SUMMONERS_RIFT_5V5_RANKED_SOLO_4_QUEUE:
		return "4"
	case SUMMONERS_RIFT_5V5_RANKED_SOLO_QUEUE:
		return "420"
	case SUMMONERS_RIFT_5V5_RANKED_TEAM_QUEUE:
		return "42"
	case SUMMONERS_RIFT_5V5_TEAM_BUILDER_QUEUE:
		return "61"
	case SUMMONERS_RIFT_6V6_HEXAKILL_QUEUE:
		return "75"
	case SUMMONERS_RIFT_ALL_RANDOM_QUEUE:
		return "325"
	case SUMMONERS_RIFT_ARURF_318_QUEUE:
		return "318"
	case SUMMONERS_RIFT_ARURF_QUEUE:
		return "900"
	case SUMMONERS_RIFT_BLACK_MARKET_BRAWLERS_QUEUE:
		return "313"
	case SUMMONERS_RIFT_BLOOD_HUNT_ASSASSIN_QUEUE:
		return "600"
	case SUMMONERS_RIFT_CLASH_QUEUE:
		return "700"
	case SUMMONERS_RIFT_CO_OP_VS_AI_BEGINNER_BOT_32_QUEUE:
		return "32"
	case SUMMONERS_RIFT_CO_OP_VS_AI_BEGINNER_BOT_840_QUEUE:
		return "840"
	case SUMMONERS_RIFT_CO_OP_VS_AI_BEGINNER_BOT_QUEUE:
		return "880"
	case SUMMONERS_RIFT_CO_OP_VS_AI_INTERMEDIATE_BOT_33_QUEUE:
		return "33"
	case SUMMONERS_RIFT_CO_OP_VS_AI_INTERMEDIATE_BOT_850_QUEUE:
		return "850"
	case SUMMONERS_RIFT_CO_OP_VS_AI_INTERMEDIATE_BOT_QUEUE:
		return "890"
	case SUMMONERS_RIFT_CO_OP_VS_AI_INTRO_BOT_31_QUEUE:
		return "31"
	case SUMMONERS_RIFT_CO_OP_VS_AI_INTRO_BOT_830_QUEUE:
		return "830"
	case SUMMONERS_RIFT_CO_OP_VS_AI_INTRO_BOT_QUEUE:
		return "870"
	case SUMMONERS_RIFT_CO_OP_VS_AI_QUEUE:
		return "7"
	case SUMMONERS_RIFT_CO_OP_VS_AI_ULTRA_RAPID_FIRE_QUEUE:
		return "83"
	case SUMMONERS_RIFT_DOOM_BOTS_RANK_1_QUEUE:
		return "91"
	case SUMMONERS_RIFT_DOOM_BOTS_RANK_2_QUEUE:
		return "92"
	case SUMMONERS_RIFT_DOOM_BOTS_RANK_5_QUEUE:
		return "93"
	case SUMMONERS_RIFT_DOOM_BOTS_STANDARD_QUEUE:
		return "960"
	case SUMMONERS_RIFT_DOOM_BOTS_VOTING_QUEUE:
		return "950"
	case SUMMONERS_RIFT_NEMESIS_QUEUE:
		return "310"
	case SUMMONERS_RIFT_NEXUS_SIEGE_315_QUEUE:
		return "315"
	case SUMMONERS_RIFT_NEXUS_SIEGE_QUEUE:
		return "940"
	case SUMMONERS_RIFT_NORMAL_QUICKPLAY_QUEUE:
		return "490"
	case SUMMONERS_RIFT_NORMAL_SWIFTPLAY_QUEUE:
		return "480"
	case SUMMONERS_RIFT_ONE_FOR_ALL_70_QUEUE:
		return "70"
	case SUMMONERS_RIFT_ONE_FOR_ALL_QUEUE:
		return "1020"
	case SUMMONERS_RIFT_PICK_URF_QUEUE:
		return "1900"
	case SUMMONERS_RIFT_SNOW_ARURF_QUEUE:
		return "1010"
	case SUMMONERS_RIFT_TUTORIAL_1_QUEUE:
		return "2000"
	case SUMMONERS_RIFT_TUTORIAL_2_QUEUE:
		return "2010"
	case SUMMONERS_RIFT_TUTORIAL_3_QUEUE:
		return "2020"
	case SUMMONERS_RIFT_ULTIMATE_SPELLBOOK_QUEUE:
		return "1400"
	case SUMMONERS_RIFT_ULTRA_RAPID_FIRE_QUEUE:
		return "76"
	case SWARM_DUO_STRAWBERRY_GAMES_QUEUE:
		return "1820"
	case SWARM_QUAD_STRAWBERRY_GAMES_QUEUE:
		return "1840"
	case SWARM_SOLO_STRAWBERRY_GAMES_QUEUE:
		return "1810"
	case SWARM_TRIO_STRAWBERRY_GAMES_QUEUE:
		return "1830"
	case TWISTED_TREELINE_3V3_BLIND_PICK_QUEUE:
		return "460"
	case TWISTED_TREELINE_3V3_NORMAL_QUEUE:
		return "8"
	case TWISTED_TREELINE_3V3_RANKED_FLEX_470_QUEUE:
		return "470"
	case TWISTED_TREELINE_3V3_RANKED_FLEX_9_QUEUE:
		return "9"
	case TWISTED_TREELINE_3V3_RANKED_TEAM_QUEUE:
		return "41"
	case TWISTED_TREELINE_6V6_HEXAKILL_QUEUE:
		return "98"
	case TWISTED_TREELINE_CO_OP_VS_AI_BEGINNER_BOT_QUEUE:
		return "820"
	case TWISTED_TREELINE_CO_OP_VS_AI_INTERMEDIATE_BOT_QUEUE:
		return "800"
	case TWISTED_TREELINE_CO_OP_VS_AI_INTRO_BOT_QUEUE:
		return "810"
	case TWISTED_TREELINE_CO_OP_VS_AI_QUEUE:
		return "52"
	case VALORAN_CITY_PARK_STAR_GUARDIAN_INVASION_NORMAL_QUEUE:
		return "980"
	case VALORAN_CITY_PARK_STAR_GUARDIAN_INVASION_ONSLAUGHT_QUEUE:
		return "990"
	default:
		return strconv.FormatInt(int64(queue), 10)
	}
}
