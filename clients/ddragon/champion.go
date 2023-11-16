package ddragon

import (
	"fmt"
	"net/http"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"go.uber.org/zap"
)

type ChampionEndpoint struct {
	internalClient *internal.InternalClient
}

type ChampionsData struct {
	Type    string              `json:"type"`
	Format  string              `json:"format"`
	Version string              `json:"version"`
	Data    map[string]Champion `json:"data"`
}

type Champion struct {
	Version string `json:"version"`
	ID      string `json:"id"`
	Key     string `json:"key"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	Blurb   string `json:"blurb"`
	Info    struct {
		Attack     int `json:"attack"`
		Defense    int `json:"defense"`
		Magic      int `json:"magic"`
		Difficulty int `json:"difficulty"`
	} `json:"info"`
	Image   Image         `json:"image"`
	Tags    []string      `json:"tags"`
	Partype string        `json:"partype"`
	Stats   ChampionStats `json:"stats"`
}

type FullChampionData struct {
	Type    string                  `json:"type"`
	Format  string                  `json:"format"`
	Version string                  `json:"version"`
	Data    map[string]FullChampion `json:"data"`
}

type FullChampion struct {
	ID    string `json:"id"`
	Key   string `json:"key"`
	Name  string `json:"name"`
	Title string `json:"title"`
	Image Image  `json:"image"`
	Skins []struct {
		ID      string `json:"id"`
		Num     int    `json:"num"`
		Name    string `json:"name"`
		Chromas bool   `json:"chromas"`
	} `json:"skins"`
	Lore      string   `json:"lore"`
	Blurb     string   `json:"blurb"`
	AllyTips  []string `json:"allytips"`
	EnemyTips []string `json:"enemytips"`
	Tags      []string `json:"tags"`
	Partype   string   `json:"partype"`
	Info      struct {
		Attack     int `json:"attack"`
		Defense    int `json:"defense"`
		Magic      int `json:"magic"`
		Difficulty int `json:"difficulty"`
	} `json:"info"`
	Stats  ChampionStats `json:"stats"`
	Spells []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Tooltip     string `json:"tooltip"`
		Leveltip    struct {
			Label  []string `json:"label"`
			Effect []string `json:"effect"`
		} `json:"leveltip"`
		MaxRank      int       `json:"maxrank"`
		Cooldown     []float64 `json:"cooldown"`
		CooldownBurn string    `json:"cooldownBurn"`
		Cost         []float64 `json:"cost"`
		CostBurn     string    `json:"costBurn"`
		DataValues   struct {
		} `json:"datavalues"`
		// Not modeled
		Effect []any `json:"effect,omitempty"`
		// Not modeled
		EffectBurn []any `json:"effectBurn,omitempty"`
		// Not modeled
		Vars      []any  `json:"vars,omitempty"`
		CostType  string `json:"costType"`
		MaxAmmo   string `json:"maxammo"`
		Range     []int  `json:"range"`
		RangeBurn string `json:"rangeBurn"`
		Image     Image  `json:"image"`
		Resource  string `json:"resource"`
	} `json:"spells"`
	Passive struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Image       Image  `json:"image"`
	} `json:"passive"`
	Recommended []any `json:"recommended,omitempty"`
}

type ChampionStats struct {
	HP                   float64 `json:"hp"`
	HPPerLevel           float64 `json:"hpperlevel"`
	MP                   float64 `json:"mp"`
	MPPerLevel           float64 `json:"mpperlevel"`
	MovementSpeed        float64 `json:"movespeed"`
	Armor                float64 `json:"armor"`
	ArmorPerLevel        float64 `json:"armorperlevel"`
	SpellBlock           float64 `json:"spellblock"`
	SpellBlockPerLevel   float64 `json:"spellblockperlevel"`
	AttackRange          float64 `json:"attackrange"`
	HPRegen              float64 `json:"hpregen"`
	HPRegenPerLevel      float64 `json:"hpregenperlevel"`
	MPRegen              float64 `json:"mpregen"`
	MPRegenPerLevel      float64 `json:"mpregenperlevel"`
	Crit                 float64 `json:"crit"`
	CritPerLevel         float64 `json:"critperlevel"`
	AttackDamage         float64 `json:"attackdamage"`
	AttackDamagePerLevel float64 `json:"attackdamageperlevel"`
	AttackSpeedPerLevel  float64 `json:"attackspeedperlevel"`
	AttackSpeed          float64 `json:"attackspeed"`
}

type Image struct {
	Full   string `json:"full"`
	Sprite string `json:"sprite"`
	Group  string `json:"group"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	W      int    `json:"w"`
	H      int    `json:"h"`
}

// Get all champions basic information, includes stats, tags, title and blurb.
func (e *ChampionEndpoint) AllChampions(version string, language Language) (map[string]Champion, error) {
	logger := e.internalClient.Logger("DDragon_Champion_AllChampions")
	logger.Debug("Method started execution")
	request, err := e.internalClient.Request(api.D_DRAGON_BASE_URL_FORMAT, http.MethodGet, "", fmt.Sprintf(ChampionsURL, version, language), nil)
	if err != nil {
		logger.Error("Error creating request", zap.Error(err))
		return nil, err
	}
	var data ChampionsData
	err = e.internalClient.Execute(logger, request, &data)
	if err != nil {
		logger.Error("Error executing request", zap.Error(err))
		return nil, err
	}
	return data.Data, nil
}

// Retrieves more information about a champion, includes skins, spells and tips.
func (e *ChampionEndpoint) ByName(version string, language Language, champion string) (*FullChampion, error) {
	logger := e.internalClient.Logger("DDragon_Champion_ByName")
	logger.Debug("Method started execution")
	request, err := e.internalClient.Request(api.D_DRAGON_BASE_URL_FORMAT, http.MethodGet, "", fmt.Sprintf(ChampionURL, version, language, champion), nil)
	if err != nil {
		logger.Error("Error creating request", zap.Error(err))
		return nil, err
	}
	var data FullChampionData
	err = e.internalClient.Execute(logger, request, &data)
	if err != nil {
		logger.Error("Error executing request", zap.Error(err))
		return nil, err
	}
	c := data.Data[champion]
	return &c, nil
}
