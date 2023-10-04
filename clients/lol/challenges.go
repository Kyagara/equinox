package lol

import (
	"fmt"

	"github.com/Kyagara/equinox/internal"
	"go.uber.org/zap"
)

type ChallengesEndpoint struct {
	internalClient *internal.InternalClient
}

type ChallengeConfigInfoDTO struct {
	ID             int           `json:"id"`
	LocalizedNames NamesDTO      `json:"localizedNames"`
	State          State         `json:"state"`
	Tracking       Tracking      `json:"tracking"`
	Thresholds     PercentileDTO `json:"thresholds,omitempty"`
	Leaderboard    bool          `json:"leaderboard"`
	StartTimestamp int64         `json:"startTimestamp,omitempty"`
	EndTimestamp   int64         `json:"endTimestamp,omitempty"`
}

type LocaleDTO struct {
	Description      string `json:"description"`
	Name             string `json:"name"`
	ShortDescription string `json:"shortDescription"`
}

type NamesDTO struct {
	ArAE LocaleDTO `json:"ar_AE"`
	CsCZ LocaleDTO `json:"cs_CZ"`
	DeDE LocaleDTO `json:"de_DE"`
	ElGR LocaleDTO `json:"el_GR"`
	EnAU LocaleDTO `json:"en_AU"`
	EnGB LocaleDTO `json:"en_GB"`
	EnPH LocaleDTO `json:"en_PH"`
	EnSG LocaleDTO `json:"en_SG"`
	EnUS LocaleDTO `json:"en_US"`
	EsAR LocaleDTO `json:"es_AR"`
	EsES LocaleDTO `json:"es_ES"`
	EsMX LocaleDTO `json:"es_MX"`
	FrFR LocaleDTO `json:"fr_FR"`
	HuHU LocaleDTO `json:"hu_HU"`
	ItIT LocaleDTO `json:"it_IT"`
	JaJP LocaleDTO `json:"ja_JP"`
	KoKR LocaleDTO `json:"ko_KR"`
	PlPL LocaleDTO `json:"pl_PL"`
	PtBR LocaleDTO `json:"pt_BR"`
	RoRO LocaleDTO `json:"ro_RO"`
	RuRU LocaleDTO `json:"ru_RU"`
	ThTH LocaleDTO `json:"th_TH"`
	TrTR LocaleDTO `json:"tr_TR"`
	ViVN LocaleDTO `json:"vi_VN"`
	ZhCN LocaleDTO `json:"zh_CN"`
	ZhMY LocaleDTO `json:"zh_MY"`
	ZhTW LocaleDTO `json:"zh_TW"`
}

type PercentileDTO struct {
	Iron        float64 `json:"IRON"`
	Bronze      float64 `json:"BRONZE"`
	Silver      float64 `json:"SILVER"`
	Gold        float64 `json:"GOLD"`
	Platinum    float64 `json:"PLATINUM"`
	Diamond     float64 `json:"DIAMOND"`
	Master      float64 `json:"MASTER"`
	Grandmaster float64 `json:"GRANDMASTER"`
	Challenger  float64 `json:"CHALLENGER"`
}

type ApexPlayerInfoDTO struct {
	PUUID    string  `json:"puuid"`
	Value    float64 `json:"value"`
	Position int     `json:"position"`
}

type PlayerInfoDTO struct {
	TotalPoints    TotalPoints     `json:"totalPoints"`
	CategoryPoints CategoryPoints  `json:"categoryPoints"`
	Challenges     []ChallengeInfo `json:"challenges"`
	Preferences    interface{}     `json:"preferences"`
}

type TotalPoints struct {
	Level      string  `json:"level"`
	Current    int     `json:"current"`
	Max        int     `json:"max"`
	Percentile float64 `json:"percentile"`
}

type CategoryPoints struct {
	Expertise   ChallengePoints `json:"EXPERTISE"`
	Collection  ChallengePoints `json:"COLLECTION"`
	Imagination ChallengePoints `json:"IMAGINATION"`
	Veterancy   ChallengePoints `json:"VETERANCY"`
	Teamwork    ChallengePoints `json:"TEAMWORK"`
}

type ChallengePoints struct {
	Level      string `json:"level"`
	Current    int    `json:"current"`
	Max        int    `json:"max"`
	Percentile int    `json:"percentile"`
}

type ChallengeInfo struct {
	ChallengeID  int     `json:"challengeId"`
	Percentile   float64 `json:"percentile"`
	Level        string  `json:"level"`
	Value        int     `json:"value"`
	AchievedTime int64   `json:"achievedTime"`
}

// List of all basic challenge configuration information (includes all translations for names and descriptions.
func (e *ChallengesEndpoint) List(region Region) (*[]ChallengeConfigInfoDTO, error) {
	logger := e.internalClient.Logger("LOL", "lol-challenges-v1", "List")
	logger.Debug("Method executed")

	var challenges *[]ChallengeConfigInfoDTO

	err := e.internalClient.Get(region, ChallengesConfigurationsURL, &challenges, "lol-challenges-v1", "List", "")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return challenges, nil
}

// Get challenge configuration.
func (e *ChallengesEndpoint) ByID(region Region, challengeID int64) (*ChallengeConfigInfoDTO, error) {
	logger := e.internalClient.Logger("LOL", "lol-challenges-v1", "ByID")
	logger.Debug("Method executed")

	var challenge *ChallengeConfigInfoDTO

	url := fmt.Sprintf(ChallengesConfigurationByIDURL, challengeID)

	err := e.internalClient.Get(region, url, &challenge, "lol-challenges-v1", "ByID", "")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return challenge, nil
}

// Map of level to percentile of players who have achieved it - keys: ChallengeId -> Season -> Level -> percentile of players who achieved it.
func (e *ChallengesEndpoint) Percentiles(region Region) (*map[int64]PercentileDTO, error) {
	logger := e.internalClient.Logger("LOL", "lol-challenges-v1", "Percentiles")
	logger.Debug("Method executed")

	var percentiles *map[int64]PercentileDTO

	err := e.internalClient.Get(region, ChallengesPercentilesURL, &percentiles, "lol-challenges-v1", "Percentiles", "")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return percentiles, nil
}

// Map of level to percentile of players who have achieved it.
func (e *ChallengesEndpoint) PercentilesByID(region Region, challengeID int64) (*PercentileDTO, error) {
	logger := e.internalClient.Logger("LOL", "lol-challenges-v1", "PercentileByID")
	logger.Debug("Method executed")

	var percentile *PercentileDTO

	url := fmt.Sprintf(ChallengesPercentileByIDURL, challengeID)

	err := e.internalClient.Get(region, url, &percentile, "lol-challenges-v1", "PercentileByID", "")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return percentile, nil
}

// Return top players for each level. Level must be MASTER, GRANDMASTER or CHALLENGER.
//
// Limit is optional, if 0 is provided, a limit will not be set.
func (e *ChallengesEndpoint) Leaderboards(region Region, challengeID int64, level Level, limit int) (*[]ApexPlayerInfoDTO, error) {
	logger := e.internalClient.Logger("LOL", "lol-challenges-v1", "Leaderboards")
	logger.Debug("Method executed")

	var leaderboards *[]ApexPlayerInfoDTO

	url := fmt.Sprintf(ChallengesLeaderboardsByLevelURL, challengeID, level)

	if limit > 0 {
		url = fmt.Sprintf("%s?limit=%d", url, limit)
	}

	err := e.internalClient.Get(region, url, &leaderboards, "lol-challenges-v1", "Leaderboards", "")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return leaderboards, nil
}

// Returns player information with list of all progressed challenges.
func (e *ChallengesEndpoint) ByPUUID(region Region, puuid string) (*PlayerInfoDTO, error) {
	logger := e.internalClient.Logger("LOL", "lol-challenges-v1", "ByPUUID")
	logger.Debug("Method executed")

	var challenges *PlayerInfoDTO

	url := fmt.Sprintf(ChallengesByPUUIDURL, puuid)

	err := e.internalClient.Get(region, url, &challenges, "lol-challenges-v1", "ByPUUID", "")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return challenges, nil
}
