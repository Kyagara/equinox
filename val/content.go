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
	Version      string            `json:"version"`
	Characters   []CharactersDTO   `json:"characters"`
	Maps         []MapsDTO         `json:"maps"`
	Chromas      []ChromasDTO      `json:"chromas"`
	Skins        []SkinsDTO        `json:"skins"`
	SkinLevels   []SkinLevelsDTO   `json:"skinLevels"`
	Equips       []EquipsDTO       `json:"equips"`
	GameModes    []GameModesDTO    `json:"gameModes"`
	Sprays       []SpraysDTO       `json:"sprays"`
	SprayLevels  []SprayLevelsDTO  `json:"sprayLevels"`
	Charms       []CharmsDTO       `json:"charms"`
	CharmLevels  []CharmLevelsDTO  `json:"charmLevels"`
	PlayerCards  []PlayerCardsDTO  `json:"playerCards"`
	PlayerTitles []PlayerTitlesDTO `json:"playerTitles"`
	Acts         []ActsDTO         `json:"acts"`
	Ceremonies   []CeremoniesDTO   `json:"ceremonies"`
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

type CharactersDTO struct {
	Name           string            `json:"name"`
	LocalizedNames LocalizedNamesDTO `json:"localizedNames,omitempty"`
	ID             string            `json:"id"`
	AssetName      string            `json:"assetName"`
}

type MapsDTO struct {
	Name              string            `json:"name"`
	LocalizedNamesDTO LocalizedNamesDTO `json:"localizedNames,omitempty"`
	ID                string            `json:"id"`
	AssetName         string            `json:"assetName"`
	AssetPath         string            `json:"assetPath,omitempty"`
}

type ChromasDTO struct {
	Name              string            `json:"name"`
	LocalizedNamesDTO LocalizedNamesDTO `json:"localizedNames,omitempty"`
	ID                string            `json:"id"`
	AssetName         string            `json:"assetName"`
}

type SkinsDTO struct {
	Name           string            `json:"name"`
	LocalizedNames LocalizedNamesDTO `json:"localizedNames,omitempty"`
	ID             string            `json:"id"`
	AssetName      string            `json:"assetName"`
}

type SkinLevelsDTO struct {
	Name              string            `json:"name"`
	LocalizedNamesDTO LocalizedNamesDTO `json:"localizedNames,omitempty"`
	ID                string            `json:"id"`
	AssetName         string            `json:"assetName"`
}

type EquipsDTO struct {
	Name           string            `json:"name"`
	LocalizedNames LocalizedNamesDTO `json:"localizedNames,omitempty"`
	ID             string            `json:"id"`
	AssetName      string            `json:"assetName"`
}

type GameModesDTO struct {
	Name           string            `json:"name"`
	LocalizedNames LocalizedNamesDTO `json:"localizedNames,omitempty"`
	ID             string            `json:"id"`
	AssetName      string            `json:"assetName"`
	AssetPath      string            `json:"assetPath"`
}

type SpraysDTO struct {
	Name           string            `json:"name"`
	LocalizedNames LocalizedNamesDTO `json:"localizedNames,omitempty"`
	ID             string            `json:"id"`
	AssetName      string            `json:"assetName"`
}

type SprayLevelsDTO struct {
	Name           string            `json:"name"`
	LocalizedNames LocalizedNamesDTO `json:"localizedNames,omitempty"`
	ID             string            `json:"id"`
	AssetName      string            `json:"assetName"`
}

type CharmsDTO struct {
	Name           string            `json:"name"`
	LocalizedNames LocalizedNamesDTO `json:"localizedNames,omitempty"`
	ID             string            `json:"id"`
	AssetName      string            `json:"assetName"`
}

type CharmLevelsDTO struct {
	Name           string            `json:"name"`
	LocalizedNames LocalizedNamesDTO `json:"localizedNames,omitempty"`
	ID             string            `json:"id"`
	AssetName      string            `json:"assetName"`
}

type PlayerCardsDTO struct {
	Name           string            `json:"name"`
	LocalizedNames LocalizedNamesDTO `json:"localizedNames,omitempty"`
	ID             string            `json:"id"`
	AssetName      string            `json:"assetName"`
}

type PlayerTitlesDTO struct {
	Name           string            `json:"name"`
	LocalizedNames LocalizedNamesDTO `json:"localizedNames,omitempty"`
	ID             string            `json:"id"`
	AssetName      string            `json:"assetName"`
}

type ActsDTO struct {
	ID             string            `json:"id"`
	ParentID       string            `json:"parentId"`
	Type           string            `json:"type"`
	Name           string            `json:"name"`
	LocalizedNames LocalizedNamesDTO `json:"localizedNames,omitempty"`
	IsActive       bool              `json:"isActive"`
}

type CeremoniesDTO struct {
	Name           string            `json:"name"`
	LocalizedNames LocalizedNamesDTO `json:"localizedNames,omitempty"`
	ID             string            `json:"id"`
	AssetName      string            `json:"assetName"`
}

type LocalizedContentDTO struct {
	Version    string `json:"version"`
	Characters []struct {
		Name      string `json:"name"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
	} `json:"characters"`
	Maps []struct {
		Name      string `json:"name"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
		AssetPath string `json:"assetPath,omitempty"`
	} `json:"maps"`
	Chromas []struct {
		Name      string `json:"name"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
	} `json:"chromas"`
	Skins []struct {
		Name      string `json:"name"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
	} `json:"skins"`
	SkinLevels []struct {
		Name      string `json:"name"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
	} `json:"skinLevels"`
	Equips []struct {
		Name      string `json:"name"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
	} `json:"equips"`
	GameModes []struct {
		Name      string `json:"name"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
		AssetPath string `json:"assetPath"`
	} `json:"gameModes"`
	Sprays []struct {
		Name      string `json:"name"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
	} `json:"sprays"`
	SprayLevels []struct {
		Name      string `json:"name"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
	} `json:"sprayLevels"`
	Charms []struct {
		Name      string `json:"name"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
	} `json:"charms"`
	CharmLevels []struct {
		Name      string `json:"name"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
	} `json:"charmLevels"`
	PlayerCards []struct {
		Name      string `json:"name"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
	} `json:"playerCards"`
	PlayerTitles []struct {
		Name      string `json:"name"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
	} `json:"playerTitles"`
	Acts []struct {
		ID       string `json:"id"`
		ParentID string `json:"parentId"`
		Type     string `json:"type"`
		Name     string `json:"name"`
		IsActive bool   `json:"isActive"`
	} `json:"acts"`
	Ceremonies []struct {
		Name      string `json:"name"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
	} `json:"ceremonies"`
}

// Get content filtered by locale.
//
// Locale defaults to en-US.
func (s *ContentEndpoint) ByLocale(region Region, locale Locale) (*LocalizedContentDTO, error) {
	logger := s.internalClient.Logger("val").With("endpoint", "content", "method", "ByLocale")

	if locale != "" {
		locale = EnglishUS
	}

	url := fmt.Sprintf("%s?locale=%s", ContentURL, locale)

	var content *LocalizedContentDTO

	err := s.internalClient.Do(http.MethodGet, region, url, nil, &content, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return content, nil
}

// Get content with all available locales.
func (s *ContentEndpoint) AllLocales(region Region) (*ContentDTO, error) {
	logger := s.internalClient.Logger("val").With("endpoint", "content", "method", "AllLocales")

	var content *ContentDTO

	err := s.internalClient.Do(http.MethodGet, region, ContentURL, nil, &content, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return content, nil
}
