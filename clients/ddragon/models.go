package ddragon

type AllChampionsDTO struct {
	Data    map[string]AllChampionsDataDTO `json:"data,omitempty"`
	Type    string                         `json:"type,omitempty"`
	Format  string                         `json:"format,omitempty"`
	Version string                         `json:"version,omitempty"`
}

type AllChampionsDataDTO struct {
	Version string        `json:"version,omitempty"`
	ID      string        `json:"id,omitempty"`
	Key     string        `json:"key,omitempty"`
	Name    string        `json:"name,omitempty"`
	Title   string        `json:"title,omitempty"`
	Blurb   string        `json:"blurb,omitempty"`
	Partype string        `json:"partype,omitempty"`
	Tags    []string      `json:"tags,omitempty"`
	Image   Image         `json:"image,omitempty"`
	Stats   ChampionStats `json:"stats,omitempty"`
	Info    ChampionInfo  `json:"info,omitempty"`
}

type ChampionInfo struct {
	Attack     int `json:"attack,omitempty"`
	Defense    int `json:"defense,omitempty"`
	Magic      int `json:"magic,omitempty"`
	Difficulty int `json:"difficulty,omitempty"`
}

type FullChampionData struct {
	Data    map[string]FullChampion `json:"data,omitempty"`
	Type    string                  `json:"type,omitempty"`
	Format  string                  `json:"format,omitempty"`
	Version string                  `json:"version,omitempty"`
}

type FullChampion struct {
	ID          string           `json:"id,omitempty"`
	Key         string           `json:"key,omitempty"`
	Name        string           `json:"name,omitempty"`
	Title       string           `json:"title,omitempty"`
	Lore        string           `json:"lore,omitempty"`
	Blurb       string           `json:"blurb,omitempty"`
	Partype     string           `json:"partype,omitempty"`
	Skins       []ChampionSkins  `json:"skins,omitempty"`
	AllyTips    []string         `json:"allytips,omitempty"`
	EnemyTips   []string         `json:"enemytips,omitempty"`
	Tags        []string         `json:"tags,omitempty"`
	Spells      []ChampionSpells `json:"spells,omitempty"`
	Recommended []any            `json:"recommended,omitempty"`
	Passive     ChampionPassive  `json:"passive,omitempty"`
	Image       Image            `json:"image,omitempty"`
	Stats       ChampionStats    `json:"stats,omitempty"`
	Info        ChampionInfo     `json:"info,omitempty"`
}

type ChampionSkins struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Num     int    `json:"num,omitempty"`
	Chromas bool   `json:"chromas,omitempty"`
}

type ChampionSpells struct {
	DataValues struct {
	} `json:"datavalues,omitempty"`
	ID           string           `json:"id,omitempty"`
	Name         string           `json:"name,omitempty"`
	Description  string           `json:"description,omitempty"`
	Tooltip      string           `json:"tooltip,omitempty"`
	CooldownBurn string           `json:"cooldownBurn,omitempty"`
	CostBurn     string           `json:"costBurn,omitempty"`
	CostType     string           `json:"costType,omitempty"`
	MaxAmmo      string           `json:"maxammo,omitempty"`
	RangeBurn    string           `json:"rangeBurn,omitempty"`
	Resource     string           `json:"resource,omitempty"`
	LevelTip     ChampionLevelTip `json:"leveltip,omitempty"`
	Cooldown     []float64        `json:"cooldown,omitempty"`
	Cost         []float64        `json:"cost,omitempty"`
	// Not modeled
	Effect []any `json:"effect,omitempty"`
	// Not modeled
	EffectBurn []any `json:"effectBurn,omitempty"`
	// Not modeled
	Vars    []any `json:"vars,omitempty"`
	Range   []int `json:"range,omitempty"`
	Image   Image `json:"image,omitempty"`
	MaxRank int   `json:"maxrank,omitempty"`
}

type ChampionPassive struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Image       Image  `json:"image,omitempty"`
}

type ChampionStats struct {
	HP                   float64 `json:"hp,omitempty"`
	HPPerLevel           float64 `json:"hpperlevel,omitempty"`
	MP                   float64 `json:"mp,omitempty"`
	MPPerLevel           float64 `json:"mpperlevel,omitempty"`
	MovementSpeed        float64 `json:"movespeed,omitempty"`
	Armor                float64 `json:"armor,omitempty"`
	ArmorPerLevel        float64 `json:"armorperlevel,omitempty"`
	SpellBlock           float64 `json:"spellblock,omitempty"`
	SpellBlockPerLevel   float64 `json:"spellblockperlevel,omitempty"`
	AttackRange          float64 `json:"attackrange,omitempty"`
	HPRegen              float64 `json:"hpregen,omitempty"`
	HPRegenPerLevel      float64 `json:"hpregenperlevel,omitempty"`
	MPRegen              float64 `json:"mpregen,omitempty"`
	MPRegenPerLevel      float64 `json:"mpregenperlevel,omitempty"`
	Crit                 float64 `json:"crit,omitempty"`
	CritPerLevel         float64 `json:"critperlevel,omitempty"`
	AttackDamage         float64 `json:"attackdamage,omitempty"`
	AttackDamagePerLevel float64 `json:"attackdamageperlevel,omitempty"`
	AttackSpeedPerLevel  float64 `json:"attackspeedperlevel,omitempty"`
	AttackSpeed          float64 `json:"attackspeed,omitempty"`
}

type ChampionLevelTip struct {
	Label  []string `json:"label,omitempty"`
	Effect []string `json:"effect,omitempty"`
}

type Image struct {
	Full   string `json:"full,omitempty"`
	Sprite string `json:"sprite,omitempty"`
	Group  string `json:"group,omitempty"`
	X      int    `json:"x,omitempty"`
	Y      int    `json:"y,omitempty"`
	W      int    `json:"w,omitempty"`
	H      int    `json:"h,omitempty"`
}

type RealmData struct {
	Store          any    `json:"store,omitempty"`
	N              RealmN `json:"n,omitempty"`
	V              string `json:"v,omitempty"`
	L              string `json:"l,omitempty"`
	CDN            string `json:"cdn,omitempty"`
	DD             string `json:"dd,omitempty"`
	LG             string `json:"lg,omitempty"`
	CSS            string `json:"css,omitempty"`
	ProfileIconMax int    `json:"profileiconmax,omitempty"`
}

type RealmN struct {
	Item        string `json:"item,omitempty"`
	Rune        string `json:"rune,omitempty"`
	Mastery     string `json:"mastery,omitempty"`
	Summoner    string `json:"summoner,omitempty"`
	Champion    string `json:"champion,omitempty"`
	ProfileIcon string `json:"profileicon,omitempty"`
	Map         string `json:"map,omitempty"`
	Language    string `json:"language,omitempty"`
	Sticker     string `json:"sticker,omitempty"`
}
