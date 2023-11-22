package cdragon

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

type ChampionData struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Alias        string `json:"alias"`
	Title        string `json:"title"`
	ShortBio     string `json:"shortBio"`
	TacticalInfo struct {
		Style      int    `json:"style"`
		Difficulty int    `json:"difficulty"`
		DamageType string `json:"damageType"`
	} `json:"tacticalInfo"`
	PlaystyleInfo struct {
		Damage       int `json:"damage"`
		Durability   int `json:"durability"`
		CrowdControl int `json:"crowdControl"`
		Mobility     int `json:"mobility"`
		Utility      int `json:"utility"`
	} `json:"playstyleInfo"`
	SquarePortraitPath      string   `json:"squarePortraitPath"`
	StingerSFXPath          string   `json:"stingerSfxPath"`
	ChooseVoPath            string   `json:"chooseVoPath"`
	BanVoPath               string   `json:"banVoPath"`
	Roles                   []string `json:"roles"`
	RecommendedItemDefaults []any    `json:"recommendedItemDefaults,omitempty"`
	Skins                   []struct {
		ID                        int    `json:"id"`
		IsBase                    bool   `json:"isBase"`
		Name                      string `json:"name"`
		SplashPath                string `json:"splashPath"`
		UncenteredSplashPath      string `json:"uncenteredSplashPath"`
		TilePath                  string `json:"tilePath"`
		LoadScreenPath            string `json:"loadScreenPath"`
		SkinType                  string `json:"skinType"`
		Rarity                    string `json:"rarity"`
		IsLegacy                  bool   `json:"isLegacy"`
		SplashVideoPath           string `json:"splashVideoPath"`
		CollectionSplashVideoPath string `json:"collectionSplashVideoPath"`
		FeaturesText              string `json:"featuresText"`
		ChromaPath                string `json:"chromaPath"`
		Emblems                   any    `json:"emblems,omitempty"`
		RegionRarityID            int    `json:"regionRarityId"`
		RarityGemPath             string `json:"rarityGemPath"`
		SkinLines                 []struct {
			ID int `json:"id"`
		} `json:"skinLines"`
		Description string `json:"description"`
		Chromas     []struct {
			ID           int      `json:"id"`
			Name         string   `json:"name"`
			ChromaPath   string   `json:"chromaPath"`
			Colors       []string `json:"colors"`
			Descriptions []struct {
				Region      string `json:"region"`
				Description string `json:"description"`
			} `json:"descriptions"`
			Rarities []struct {
				Region string `json:"region"`
				Rarity int    `json:"rarity"`
			} `json:"rarities"`
		} `json:"chromas,omitempty"`
	} `json:"skins"`
	Passive struct {
		Name                  string `json:"name"`
		AbilityIconPath       string `json:"abilityIconPath"`
		AbilityVideoPath      string `json:"abilityVideoPath"`
		AbilityVideoImagePath string `json:"abilityVideoImagePath"`
		Description           string `json:"description"`
	} `json:"passive"`
	Spells []struct {
		SpellKey              string    `json:"spellKey"`
		Name                  string    `json:"name"`
		AbilityIconPath       string    `json:"abilityIconPath"`
		AbilityVideoPath      string    `json:"abilityVideoPath"`
		AbilityVideoImagePath string    `json:"abilityVideoImagePath"`
		Cost                  string    `json:"cost"`
		Cooldown              string    `json:"cooldown"`
		Description           string    `json:"description"`
		DynamicDescription    string    `json:"dynamicDescription"`
		Range                 []float64 `json:"range"`
		CostCoefficients      []float64 `json:"costCoefficients"`
		CooldownCoefficients  []float64 `json:"cooldownCoefficients"`
		Coefficients          struct {
			Coefficient1 float64 `json:"coefficient1"`
			Coefficient2 float64 `json:"coefficient2"`
		} `json:"coefficients"`
		EffectAmounts struct {
			Effect1Amount  []float64 `json:"Effect1Amount"`
			Effect2Amount  []float64 `json:"Effect2Amount"`
			Effect3Amount  []float64 `json:"Effect3Amount"`
			Effect4Amount  []float64 `json:"Effect4Amount"`
			Effect5Amount  []float64 `json:"Effect5Amount"`
			Effect6Amount  []float64 `json:"Effect6Amount"`
			Effect7Amount  []float64 `json:"Effect7Amount"`
			Effect8Amount  []float64 `json:"Effect8Amount"`
			Effect9Amount  []float64 `json:"Effect9Amount"`
			Effect10Amount []float64 `json:"Effect10Amount"`
		} `json:"effectAmounts"`
		Ammo struct {
			AmmoRechargeTime []float64 `json:"ammoRechargeTime"`
			MaxAmmo          []int     `json:"maxAmmo"`
		} `json:"ammo"`
		MaxLevel int `json:"maxLevel"`
	} `json:"spells"`
}

// Retrieves more information about a champion, includes skins, spells and tips.
func (e *ChampionEndpoint) ByName(version string, champion string) (*ChampionData, error) {
	logger := e.internalClient.Logger("CDragon_Champion_ByName")
	logger.Debug("Method started execution")
	equinoxReq, err := e.internalClient.Request(logger, api.C_DRAGON_BASE_URL_FORMAT, http.MethodGet, "", fmt.Sprintf(ChampionURL, version, champion), "", nil)
	if err != nil {
		logger.Error("Error creating request", zap.Error(err))
		return nil, err
	}
	var data ChampionData
	err = e.internalClient.Execute(equinoxReq, &data)
	if err != nil {
		logger.Error("Error executing request", zap.Error(err))
		return nil, err
	}
	return &data, nil
}

// Retrieves more information about a champion, includes skins, spells and tips.
func (e *ChampionEndpoint) ByID(version string, id int) (*ChampionData, error) {
	logger := e.internalClient.Logger("CDragon_Champion_ByID")
	logger.Debug("Method started execution")
	equinoxReq, err := e.internalClient.Request(logger, api.C_DRAGON_BASE_URL_FORMAT, http.MethodGet, "", fmt.Sprintf(ChampionURL, version, id), "", nil)
	if err != nil {
		logger.Error("Error creating request", zap.Error(err))
		return nil, err
	}
	var data ChampionData
	err = e.internalClient.Execute(equinoxReq, &data)
	if err != nil {
		logger.Error("Error executing request", zap.Error(err))
		return nil, err
	}
	return &data, nil
}
