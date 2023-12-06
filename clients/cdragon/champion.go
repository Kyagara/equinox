package cdragon

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

type ChampionData struct {
	Passive struct {
		Name                  string `json:"name"`
		AbilityIconPath       string `json:"abilityIconPath"`
		AbilityVideoPath      string `json:"abilityVideoPath"`
		AbilityVideoImagePath string `json:"abilityVideoImagePath"`
		Description           string `json:"description"`
	} `json:"passive"`
	Name                    string   `json:"name"`
	Alias                   string   `json:"alias"`
	Title                   string   `json:"title"`
	ShortBio                string   `json:"shortBio"`
	SquarePortraitPath      string   `json:"squarePortraitPath"`
	StingerSFXPath          string   `json:"stingerSfxPath"`
	ChooseVoPath            string   `json:"chooseVoPath"`
	BanVoPath               string   `json:"banVoPath"`
	Roles                   []string `json:"roles"`
	RecommendedItemDefaults []any    `json:"recommendedItemDefaults,omitempty"`
	Skins                   []struct {
		Emblems                   any    `json:"emblems,omitempty"`
		Name                      string `json:"name"`
		SplashPath                string `json:"splashPath"`
		UncenteredSplashPath      string `json:"uncenteredSplashPath"`
		TilePath                  string `json:"tilePath"`
		LoadScreenPath            string `json:"loadScreenPath"`
		SkinType                  string `json:"skinType"`
		Rarity                    string `json:"rarity"`
		SplashVideoPath           string `json:"splashVideoPath"`
		CollectionSplashVideoPath string `json:"collectionSplashVideoPath"`
		FeaturesText              string `json:"featuresText"`
		ChromaPath                string `json:"chromaPath"`
		RarityGemPath             string `json:"rarityGemPath"`
		Description               string `json:"description"`
		SkinLines                 []struct {
			ID int `json:"id"`
		} `json:"skinLines"`
		Chromas []struct {
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
			ID int `json:"id"`
		} `json:"chromas,omitempty"`
		ID             int  `json:"id"`
		RegionRarityID int  `json:"regionRarityId"`
		IsBase         bool `json:"isBase"`
		IsLegacy       bool `json:"isLegacy"`
	} `json:"skins"`
	Spells []struct {
		SpellKey              string `json:"spellKey"`
		Name                  string `json:"name"`
		AbilityIconPath       string `json:"abilityIconPath"`
		AbilityVideoPath      string `json:"abilityVideoPath"`
		AbilityVideoImagePath string `json:"abilityVideoImagePath"`
		Cost                  string `json:"cost"`
		Cooldown              string `json:"cooldown"`
		Description           string `json:"description"`
		DynamicDescription    string `json:"dynamicDescription"`
		EffectAmounts         struct {
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
		Range                []float64 `json:"range"`
		CostCoefficients     []float64 `json:"costCoefficients"`
		CooldownCoefficients []float64 `json:"cooldownCoefficients"`
		Coefficients         struct {
			Coefficient1 float64 `json:"coefficient1"`
			Coefficient2 float64 `json:"coefficient2"`
		} `json:"coefficients"`
		MaxLevel int `json:"maxLevel"`
	} `json:"spells"`
	TacticalInfo struct {
		DamageType string `json:"damageType"`
		Style      int    `json:"style"`
		Difficulty int    `json:"difficulty"`
	} `json:"tacticalInfo"`
	PlaystyleInfo struct {
		Damage       int `json:"damage"`
		Durability   int `json:"durability"`
		CrowdControl int `json:"crowdControl"`
		Mobility     int `json:"mobility"`
		Utility      int `json:"utility"`
	} `json:"playstyleInfo"`
	ID int `json:"id"`
}

// Retrieves more information about a champion, includes skins, spells and tips.
func (e *ChampionEndpoint) ByName(ctx context.Context, version string, champion string) (*ChampionData, error) {
	logger := e.internal.Logger("CDragon_Champion_ByName")
	logger.Debug().Msg("Method started execution")
	equinoxReq, err := e.internal.Request(ctx, logger, api.C_DRAGON_BASE_URL_FORMAT, http.MethodGet, "", fmt.Sprintf(ChampionURL, version, champion), "", nil)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating request")
		return nil, err
	}
	var data ChampionData
	err = e.internal.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error().Err(err).Msg("Error executing request")
		return nil, err
	}
	return &data, nil
}

// Retrieves more information about a champion, includes skins, spells and tips.
func (e *ChampionEndpoint) ByID(ctx context.Context, version string, id int) (*ChampionData, error) {
	logger := e.internal.Logger("CDragon_Champion_ByID")
	logger.Debug().Msg("Method started execution")
	equinoxReq, err := e.internal.Request(ctx, logger, api.C_DRAGON_BASE_URL_FORMAT, http.MethodGet, "", fmt.Sprintf(ChampionURL, version, id), "", nil)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating request")
		return nil, err
	}
	var data ChampionData
	err = e.internal.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error().Err(err).Msg("Error executing request")
		return nil, err
	}
	return &data, nil
}
