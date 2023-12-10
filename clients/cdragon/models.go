package cdragon

type ChampionData struct {
	Passive                 ChampionPassive       `json:"passive,omitempty"`
	Name                    string                `json:"name,omitempty"`
	Alias                   string                `json:"alias,omitempty"`
	Title                   string                `json:"title,omitempty"`
	ShortBio                string                `json:"shortBio,omitempty"`
	SquarePortraitPath      string                `json:"squarePortraitPath,omitempty"`
	StingerSFXPath          string                `json:"stingerSfxPath,omitempty"`
	ChooseVoPath            string                `json:"chooseVoPath,omitempty"`
	BanVoPath               string                `json:"banVoPath,omitempty"`
	Roles                   []string              `json:"roles,omitempty"`
	RecommendedItemDefaults []any                 `json:"recommendedItemDefaults,omitempty"`
	Skins                   []ChampionSkins       `json:"skins,omitempty"`
	Spells                  []ChampionSpells      `json:"spells,omitempty"`
	TacticalInfo            ChampionTacticalInfo  `json:"tacticalInfo,omitempty"`
	PlaystyleInfo           ChampionPlaystyleInfo `json:"playstyleInfo,omitempty"`
	ID                      int                   `json:"id,omitempty"`
}

type ChampionPassive struct {
	Name                  string `json:"name,omitempty"`
	AbilityIconPath       string `json:"abilityIconPath,omitempty"`
	AbilityVideoPath      string `json:"abilityVideoPath,omitempty"`
	AbilityVideoImagePath string `json:"abilityVideoImagePath,omitempty"`
	Description           string `json:"description,omitempty"`
}

type ChampionSkins struct {
	Emblems                   any                   `json:"emblems,omitempty"`
	Name                      string                `json:"name,omitempty"`
	SplashPath                string                `json:"splashPath,omitempty"`
	UncenteredSplashPath      string                `json:"uncenteredSplashPath,omitempty"`
	TilePath                  string                `json:"tilePath,omitempty"`
	LoadScreenPath            string                `json:"loadScreenPath,omitempty"`
	SkinType                  string                `json:"skinType,omitempty"`
	Rarity                    string                `json:"rarity,omitempty"`
	SplashVideoPath           string                `json:"splashVideoPath,omitempty"`
	CollectionSplashVideoPath string                `json:"collectionSplashVideoPath,omitempty"`
	FeaturesText              string                `json:"featuresText,omitempty"`
	ChromaPath                string                `json:"chromaPath,omitempty"`
	RarityGemPath             string                `json:"rarityGemPath,omitempty"`
	Description               string                `json:"description,omitempty"`
	SkinLines                 []ChampionSkinLines   `json:"skinLines,omitempty"`
	Chromas                   []ChampionSkinChromas `json:"chromas,omitempty"`
	ID                        int                   `json:"id,omitempty"`
	RegionRarityID            int                   `json:"regionRarityId,omitempty"`
	IsBase                    bool                  `json:"isBase,omitempty"`
	IsLegacy                  bool                  `json:"isLegacy,omitempty"`
}

type ChampionSkinLines struct {
	ID int `json:"id,omitempty"`
}

type ChampionSkinChromas struct {
	Name         string                           `json:"name,omitempty"`
	ChromaPath   string                           `json:"chromaPath,omitempty"`
	Colors       []string                         `json:"colors,omitempty"`
	Descriptions []ChampionSkinChromaDescriptions `json:"descriptions,omitempty"`
	Rarities     []ChampionSkinChromaRarities     `json:"rarities,omitempty"`
	ID           int                              `json:"id,omitempty"`
}

type ChampionSkinChromaDescriptions struct {
	Region      string `json:"region,omitempty"`
	Description string `json:"description,omitempty"`
}

type ChampionSkinChromaRarities struct {
	Region string `json:"region,omitempty"`
	Rarity int    `json:"rarity,omitempty"`
}

type ChampionSpells struct {
	SpellKey              string                     `json:"spellKey,omitempty"`
	Name                  string                     `json:"name,omitempty"`
	AbilityIconPath       string                     `json:"abilityIconPath,omitempty"`
	AbilityVideoPath      string                     `json:"abilityVideoPath,omitempty"`
	AbilityVideoImagePath string                     `json:"abilityVideoImagePath,omitempty"`
	Cost                  string                     `json:"cost,omitempty"`
	Cooldown              string                     `json:"cooldown,omitempty"`
	Description           string                     `json:"description,omitempty"`
	DynamicDescription    string                     `json:"dynamicDescription,omitempty"`
	EffectAmounts         ChampionSpellEffectAmounts `json:"effectAmounts,omitempty"`
	Ammo                  ChampionSpellAmmo          `json:"ammo,omitempty"`
	Range                 []float64                  `json:"range,omitempty"`
	CostCoefficients      []float64                  `json:"costCoefficients,omitempty"`
	CooldownCoefficients  []float64                  `json:"cooldownCoefficients,omitempty"`
	Coefficients          ChampionSpellCoefficients  `json:"coefficients,omitempty"`
	MaxLevel              int                        `json:"maxLevel,omitempty"`
}

type ChampionSpellEffectAmounts struct {
	Effect1Amount  []float64 `json:"Effect1Amount,omitempty"`
	Effect2Amount  []float64 `json:"Effect2Amount,omitempty"`
	Effect3Amount  []float64 `json:"Effect3Amount,omitempty"`
	Effect4Amount  []float64 `json:"Effect4Amount,omitempty"`
	Effect5Amount  []float64 `json:"Effect5Amount,omitempty"`
	Effect6Amount  []float64 `json:"Effect6Amount,omitempty"`
	Effect7Amount  []float64 `json:"Effect7Amount,omitempty"`
	Effect8Amount  []float64 `json:"Effect8Amount,omitempty"`
	Effect9Amount  []float64 `json:"Effect9Amount,omitempty"`
	Effect10Amount []float64 `json:"Effect10Amount,omitempty"`
}

type ChampionSpellAmmo struct {
	AmmoRechargeTime []float64 `json:"ammoRechargeTime,omitempty"`
	MaxAmmo          []int     `json:"maxAmmo,omitempty"`
}

type ChampionSpellCoefficients struct {
	Coefficient1 float64 `json:"coefficient1,omitempty"`
	Coefficient2 float64 `json:"coefficient2,omitempty"`
}

type ChampionPlaystyleInfo struct {
	Damage       int `json:"damage,omitempty"`
	Durability   int `json:"durability,omitempty"`
	CrowdControl int `json:"crowdControl,omitempty"`
	Mobility     int `json:"mobility,omitempty"`
	Utility      int `json:"utility,omitempty"`
}

type ChampionTacticalInfo struct {
	DamageType string `json:"damageType,omitempty"`
	Style      int    `json:"style,omitempty"`
	Difficulty int    `json:"difficulty,omitempty"`
}
