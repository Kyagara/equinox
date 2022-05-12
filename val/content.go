package val

import (
	"fmt"
	"net/http"

	"github.com/Kyagara/equinox/internal"
)

type ContentEndpoint struct {
	internalClient *internal.InternalClient
}

type ContentDTO struct {
	Version      string     `json:"version"`
	Characters   []AssetDTO `json:"characters"`
	Maps         []AssetDTO `json:"maps"`
	Chromas      []AssetDTO `json:"chromas"`
	Skins        []AssetDTO `json:"skins"`
	SkinLevels   []AssetDTO `json:"skinLevels"`
	Equips       []AssetDTO `json:"equips"`
	GameModes    []AssetDTO `json:"gameModes"`
	Sprays       []AssetDTO `json:"sprays"`
	SprayLevels  []AssetDTO `json:"sprayLevels"`
	Charms       []AssetDTO `json:"charms"`
	CharmLevels  []AssetDTO `json:"charmLevels"`
	PlayerCards  []AssetDTO `json:"playerCards"`
	PlayerTitles []AssetDTO `json:"playerTitles"`
	Acts         []ActDTO   `json:"acts"`
	Ceremonies   []AssetDTO `json:"ceremonies"`
}

type LocalizedNamesDTO struct {
	ArAE string `json:"ar-AE"`
	DeDE string `json:"de-DE"`
	EnUS string `json:"en-US"`
	EsES string `json:"es-ES"`
	EsMX string `json:"es-MX"`
	FrFR string `json:"fr-FR"`
	IDID string `json:"id-ID"`
	ItIT string `json:"it-IT"`
	JaJP string `json:"ja-JP"`
	KoKR string `json:"ko-KR"`
	PlPL string `json:"pl-PL"`
	PtBR string `json:"pt-BR"`
	RuRU string `json:"ru-RU"`
	ThTH string `json:"th-TH"`
	TrTR string `json:"tr-TR"`
	ViVN string `json:"vi-VN"`
	ZhCN string `json:"zh-CN"`
	ZhTW string `json:"zh-TW"`
}

type LocalizedContentDTO struct {
	Version      string              `json:"version"`
	Characters   []LocalizedAssetDTO `json:"characters"`
	Maps         []LocalizedAssetDTO `json:"maps"`
	Chromas      []LocalizedAssetDTO `json:"chromas"`
	Skins        []LocalizedAssetDTO `json:"skins"`
	SkinLevels   []LocalizedAssetDTO `json:"skinLevels"`
	Equips       []LocalizedAssetDTO `json:"equips"`
	GameModes    []LocalizedAssetDTO `json:"gameModes"`
	Sprays       []LocalizedAssetDTO `json:"sprays"`
	SprayLevels  []LocalizedAssetDTO `json:"sprayLevels"`
	Charms       []LocalizedAssetDTO `json:"charms"`
	CharmLevels  []LocalizedAssetDTO `json:"charmLevels"`
	PlayerCards  []LocalizedAssetDTO `json:"playerCards"`
	PlayerTitles []LocalizedAssetDTO `json:"playerTitles"`
	Ceremonies   []LocalizedAssetDTO `json:"ceremonies"`
	Acts         []LocalizedActDTO   `json:"acts"`
}

type AssetDTO struct {
	Name           string            `json:"name"`
	ID             string            `json:"id"`
	AssetName      string            `json:"assetName"`
	AssetPath      string            `json:"assetPath,omitempty"`
	LocalizedNames LocalizedNamesDTO `json:"localizedNames"`
}

type LocalizedAssetDTO struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	AssetName string `json:"assetName"`
	AssetPath string `json:"assetPath,omitempty"`
}

type ActDTO struct {
	ID             string            `json:"id"`
	ParentID       string            `json:"parentId"`
	Type           string            `json:"type"`
	Name           string            `json:"name"`
	LocalizedNames LocalizedNamesDTO `json:"localizedNames"`
	IsActive       bool              `json:"isActive"`
}

type LocalizedActDTO struct {
	ID       string `json:"id"`
	ParentID string `json:"parentId"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	IsActive bool   `json:"isActive"`
}

// Get content filtered by locale.
//
// Locale defaults to en-US.
func (e *ContentEndpoint) ByLocale(shard Shard, locale Locale) (*LocalizedContentDTO, error) {
	logger := e.internalClient.Logger("VAL", "content", "ByLocale")

	if locale != "" {
		locale = EnglishUS
	}

	url := fmt.Sprintf("%s?locale=%s", ContentURL, locale)

	var content *LocalizedContentDTO

	err := e.internalClient.Do(http.MethodGet, shard, url, nil, &content, "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return content, nil
}

// Get content with all available locales.
func (e *ContentEndpoint) AllLocales(shard Shard) (*ContentDTO, error) {
	logger := e.internalClient.Logger("VAL", "content", "AllLocales")

	var content *ContentDTO

	err := e.internalClient.Do(http.MethodGet, shard, ContentURL, nil, &content, "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return content, nil
}
