package ddragon

type Language string

const (
	EnUS Language = "en_US"
	CsCZ Language = "cs_CZ"
	DeDE Language = "de_DE"
	ElGR Language = "el_GR"
	EnAU Language = "en_AU"
	EnGB Language = "en_GB"
	EnPH Language = "en_PH"
	EnSG Language = "en_SG"
	EsAR Language = "es_AR"
	EsES Language = "es_ES"
	EsMX Language = "es_MX"
	FrFR Language = "fr_FR"
	HuHU Language = "hu_HU"
	IdID Language = "id_ID"
	ItIT Language = "it_IT"
	JaJP Language = "ja_JP"
	KoKR Language = "ko_KR"
	PlPL Language = "pl_PL"
	PtBR Language = "pt_BR"
	RoRO Language = "ro_RO"
	RuRU Language = "ru_RU"
	ThTH Language = "th_TH"
	TrTR Language = "tr_TR"
	ViVN Language = "vi_VN"
	ZhCN Language = "zh_CN"
	ZhMY Language = "zh_MY"
	ZhTW Language = "zh_TW"
)

func (l Language) String() string {
	switch l {
	case EnUS:
		return "en_US"
	case CsCZ:
		return "cs_CZ"
	case DeDE:
		return "de_DE"
	case ElGR:
		return "el_GR"
	case EnAU:
		return "en_AU"
	case EnGB:
		return "en_GB"
	case EnPH:
		return "en_PH"
	case EnSG:
		return "en_SG"
	case EsAR:
		return "es_AR"
	case EsES:
		return "es_ES"
	case EsMX:
		return "es_MX"
	case FrFR:
		return "fr_FR"
	case HuHU:
		return "hu_HU"
	case IdID:
		return "id_ID"
	case ItIT:
		return "it_IT"
	case JaJP:
		return "ja_JP"
	case KoKR:
		return "ko_KR"
	case PlPL:
		return "pl_PL"
	case PtBR:
		return "pt_BR"
	case RoRO:
		return "ro_RO"
	case RuRU:
		return "ru_RU"
	case ThTH:
		return "th_TH"
	case TrTR:
		return "tr_TR"
	case ViVN:
		return "vi_VN"
	case ZhCN:
		return "zh_CN"
	case ZhMY:
		return "zh_MY"
	case ZhTW:
		return "zh_TW"
	default:
		return "en_US"
	}
}

type Realm string

const (
	BR   Realm = "br"
	NA   Realm = "na"
	EUNE Realm = "eune"
	EUW  Realm = "euw"
	LAN  Realm = "lan"
	LAS  Realm = "las"
	OCE  Realm = "oce"
	RU   Realm = "ru"
	TR   Realm = "tr"
	JP   Realm = "jp"
	KR   Realm = "kr"
	PBE  Realm = "pbe"
)

func (r Realm) String() string {
	switch r {
	case BR:
		return "br"
	case NA:
		return "na"
	case EUNE:
		return "eune"
	case EUW:
		return "euw"
	case LAN:
		return "lan"
	case LAS:
		return "las"
	case OCE:
		return "oce"
	case RU:
		return "ru"
	case TR:
		return "tr"
	case JP:
		return "jp"
	case KR:
		return "kr"
	case PBE:
		return "pbe"
	default:
		return ""
	}
}
