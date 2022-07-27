package data_dragon

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
	VnVN Language = "vn_VN"
	ZhCN Language = "zh_CN"
	ZhMY Language = "zh_MY"
	ZhTW Language = "zh_TW"
)

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
