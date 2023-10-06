package data_dragon

import (
	"encoding/json"
	"fmt"

	"github.com/Kyagara/equinox/internal"
	"go.uber.org/zap"
)

type ChampionEndpoint struct {
	internalClient *internal.InternalClient
}

type ChampionData struct {
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
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   struct {
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
	} `json:"stats"`
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

func (e *ChampionEndpoint) AllChampions(version string, language Language) (map[string]*ChampionData, error) {
	logger := e.internalClient.Logger("Data Dragon", "champion", "AllChampions")
	logger.Debug("Method executed")

	url := fmt.Sprintf(ChampionsURL, version, language)

	var data DataDragonMetadata

	err := e.internalClient.DataDragonGet(url, &data, "champion", "AllChampions")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	champions, err := json.Marshal(data.Data)
	if err != nil {
		logger.Error("Failed to encode champions data", zap.Error(err))
		return nil, err
	}

	var championsData map[string]*ChampionData

	err = json.Unmarshal(champions, &championsData)
	if err != nil {
		logger.Error("Failed to parse champions data", zap.Error(err))
		return nil, err
	}

	return championsData, nil
}

func (e *ChampionEndpoint) ByName(version string, language Language, champion string) (*ChampionData, error) {
	logger := e.internalClient.Logger("Data Dragon", "champion", "ByName")
	logger.Debug("Method executed")

	url := fmt.Sprintf(ChampionURL, version, language, champion)

	var data DataDragonMetadata

	err := e.internalClient.DataDragonGet(url, &data, "champion", "ByName")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	champions, err := json.Marshal(data.Data)
	if err != nil {
		logger.Error("Failed to encode champion data", zap.Error(err))
		return nil, err
	}

	var championsData map[string]*ChampionData

	err = json.Unmarshal(champions, &championsData)
	if err != nil {
		logger.Error("Failed to parse champion data", zap.Error(err))
		return nil, err
	}

	return championsData[champion], nil
}
