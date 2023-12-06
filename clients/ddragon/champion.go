package ddragon

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type ChampionEndpoint struct {
	internal *internal.Client
}

type ChampionsData struct {
	Data    map[string]Champion `json:"data"`
	Type    string              `json:"type"`
	Format  string              `json:"format"`
	Version string              `json:"version"`
}

type Champion struct {
	Version string        `json:"version"`
	ID      string        `json:"id"`
	Key     string        `json:"key"`
	Name    string        `json:"name"`
	Title   string        `json:"title"`
	Blurb   string        `json:"blurb"`
	Partype string        `json:"partype"`
	Tags    []string      `json:"tags"`
	Image   Image         `json:"image"`
	Stats   ChampionStats `json:"stats"`
	Info    struct {
		Attack     int `json:"attack"`
		Defense    int `json:"defense"`
		Magic      int `json:"magic"`
		Difficulty int `json:"difficulty"`
	} `json:"info"`
}

type FullChampionData struct {
	Data    map[string]FullChampion `json:"data"`
	Type    string                  `json:"type"`
	Format  string                  `json:"format"`
	Version string                  `json:"version"`
}

type FullChampion struct {
	ID      string `json:"id"`
	Key     string `json:"key"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	Lore    string `json:"lore"`
	Blurb   string `json:"blurb"`
	Partype string `json:"partype"`
	Skins   []struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Num     int    `json:"num"`
		Chromas bool   `json:"chromas"`
	} `json:"skins"`
	AllyTips  []string `json:"allytips"`
	EnemyTips []string `json:"enemytips"`
	Tags      []string `json:"tags"`
	Spells    []struct {
		DataValues struct {
		} `json:"datavalues"`
		ID           string `json:"id"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		Tooltip      string `json:"tooltip"`
		CooldownBurn string `json:"cooldownBurn"`
		CostBurn     string `json:"costBurn"`
		CostType     string `json:"costType"`
		MaxAmmo      string `json:"maxammo"`
		RangeBurn    string `json:"rangeBurn"`
		Resource     string `json:"resource"`
		Leveltip     struct {
			Label  []string `json:"label"`
			Effect []string `json:"effect"`
		} `json:"leveltip"`
		Cooldown []float64 `json:"cooldown"`
		Cost     []float64 `json:"cost"`
		// Not modeled
		Effect []any `json:"effect,omitempty"`
		// Not modeled
		EffectBurn []any `json:"effectBurn,omitempty"`
		// Not modeled
		Vars    []any `json:"vars,omitempty"`
		Range   []int `json:"range"`
		Image   Image `json:"image"`
		MaxRank int   `json:"maxrank"`
	} `json:"spells"`
	Recommended []any `json:"recommended,omitempty"`
	Passive     struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Image       Image  `json:"image"`
	} `json:"passive"`
	Image Image         `json:"image"`
	Stats ChampionStats `json:"stats"`
	Info  struct {
		Attack     int `json:"attack"`
		Defense    int `json:"defense"`
		Magic      int `json:"magic"`
		Difficulty int `json:"difficulty"`
	} `json:"info"`
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
func (e *ChampionEndpoint) AllChampions(ctx context.Context, version string, language Language) (map[string]Champion, error) {
	logger := e.internal.Logger("DDragon_Champion_AllChampions")
	logger.Debug().Msg("Method started execution")
	equinoxReq, err := e.internal.Request(ctx, logger, api.D_DRAGON_BASE_URL_FORMAT, http.MethodGet, "", fmt.Sprintf(ChampionsURL, version, language), "", nil)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating request")
		return nil, err
	}
	var data ChampionsData
	err = e.internal.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error().Err(err).Msg("Error executing request")
		return nil, err
	}
	return data.Data, nil
}

// Retrieves more information about a champion, includes skins, spells and tips.
func (e *ChampionEndpoint) ByName(ctx context.Context, version string, language Language, champion string) (*FullChampion, error) {
	logger := e.internal.Logger("DDragon_Champion_ByName")
	logger.Debug().Msg("Method started execution")
	equinoxReq, err := e.internal.Request(ctx, logger, api.D_DRAGON_BASE_URL_FORMAT, http.MethodGet, "", fmt.Sprintf(ChampionURL, version, language, champion), "", nil)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating request")
		return nil, err
	}
	var data FullChampionData
	err = e.internal.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error().Err(err).Msg("Error executing request")
		return nil, err
	}
	c := data.Data[champion]
	return &c, nil
}
