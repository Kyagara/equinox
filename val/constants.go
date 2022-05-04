package val

type Queue string

const (
	Competitive Queue = "competitive"
	Unrated     Queue = "unrated"
	SpikeRush   Queue = "spikerush"
	Tournament  Queue = "tournamentmode"
)

type Region string

// Valorant regions
const (
	AP      Region = "ap"
	NA      Region = "na"
	LATAM   Region = "latam"
	ESPORTS Region = "esports"
	BR      Region = "br"
	EU      Region = "eu"
	KR      Region = "kr"
)

type Locale string

const (
	ArabicAE     Locale = "ar-AE"
	GermanDE     Locale = "de-DE"
	EnglishUS    Locale = "en-US"
	SpanishES    Locale = "es-ES"
	SpanishMX    Locale = "es-MX"
	FrenchFR     Locale = "fr-FR"
	IndonesianID Locale = "id-ID"
	ItalianIT    Locale = "it-IT"
	JapaneseJP   Locale = "ja-JP"
	KoreanKR     Locale = "ko-KR"
	PortuguesePL Locale = "pl-PL"
	PortugueseBR Locale = "pt-BR"
	RussianRU    Locale = "RuRU"
	ThaiTH       Locale = "th-TH"
	TurkishTR    Locale = "tr-TR"
	VietnameseVN Locale = "vi-VN"
	ChineseCN    Locale = "zh-CN"
	ChineseTW    Locale = "zh-TW"
)
