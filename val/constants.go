package val

type Shard string

// Valorant shards.
const (
	AP      Shard = "ap"
	NA      Shard = "na"
	LATAM   Shard = "latam"
	ESPORTS Shard = "esports"
	BR      Shard = "br"
	EU      Shard = "eu"
	KR      Shard = "kr"
)

type Queue string

const (
	CompetitiveQueue Queue = "competitive"
	UnratedQueue     Queue = "unrated"
	SpikeRushQueue   Queue = "spikerush"
	TournamentQueue  Queue = "tournamentmode"
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
