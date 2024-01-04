package tft

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = d712d94a43004a22ad9f31b9ebfbcaa9e0820305

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
