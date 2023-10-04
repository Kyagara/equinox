package lol

import (
	"fmt"

	"github.com/Kyagara/equinox/internal"
	"go.uber.org/zap"
)

type ClashEndpoint struct {
	internalClient *internal.InternalClient
}

type ClashTournamentDTO struct {
	ID               int    `json:"id"`
	ThemeID          int    `json:"themeId"`
	NameKey          string `json:"nameKey"`
	NameKeySecondary string `json:"nameKeySecondary"`
	// Tournament phase.
	Schedule []TournamentPhaseDto `json:"schedule"`
}

type TournamentPhaseDto struct {
	ID               int   `json:"id"`
	RegistrationTime int64 `json:"registrationTime"`
	StartTime        int64 `json:"startTime"`
	Cancelled        bool  `json:"cancelled"`
}

type TournamentTeamDto struct {
	ID           string `json:"id"`
	TournamentID int    `json:"tournamentId"`
	Name         string `json:"name"`
	IconID       int    `json:"iconId"`
	Tier         int    `json:"tier"`
	// Summoner ID of the team captain.
	Captain      string `json:"captain"`
	Abbreviation string `json:"abbreviation"`
	// Team members.
	Players []TournamentPlayerDTO `json:"players"`
}

type TournamentPlayerDTO struct {
	SummonerID string `json:"summonerId"`
	TeamID     string `json:"teamId,omitempty"`
	// (Legal values: UNSELECTED, FILL, TOP, JUNGLE, MIDDLE, BOTTOM, UTILITY).
	Position string `json:"position"`
	// (Legal values: CAPTAIN, MEMBER).
	Role string `json:"role"`
}

// Get all active or upcoming tournaments.
func (e *ClashEndpoint) Tournaments(region Region) (*[]ClashTournamentDTO, error) {
	logger := e.internalClient.Logger("LOL", "clash-v1", "Tournaments")
	logger.Debug("Method executed")

	if region == PBE1 {
		return nil, fmt.Errorf("the region PBE1 is not available for this method")
	}

	var tournaments []ClashTournamentDTO

	err := e.internalClient.Get(region, ClashURL, &tournaments, "clash-v1", "Tournaments", "")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return &tournaments, nil
}

// Get players by summoner ID.
//
// This endpoint returns a list of active Clash players for a given summoner ID. If a summoner registers for multiple tournaments at the same time (e.g., Saturday and Sunday) then both registrations would appear in this list.
func (e *ClashEndpoint) SummonerEntries(region Region, summonerID string) (*[]TournamentPlayerDTO, error) {
	logger := e.internalClient.Logger("LOL", "clash-v1", "SummonerEntries")
	logger.Debug("Method executed")

	if region == PBE1 {
		return nil, fmt.Errorf("the region PBE1 is not available for this method")
	}

	url := fmt.Sprintf(ClashSummonerEntriesURL, summonerID)

	var players []TournamentPlayerDTO

	err := e.internalClient.Get(region, url, &players, "clash-v1", "SummonerEntries", "")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return &players, nil
}

// Get team by ID.
func (e *ClashEndpoint) TournamentTeamByID(region Region, teamID string) (*TournamentTeamDto, error) {
	logger := e.internalClient.Logger("LOL", "clash-v1", "TournamentTeamByID")
	logger.Debug("Method executed")

	if region == PBE1 {
		return nil, fmt.Errorf("the region PBE1 is not available for this method")
	}

	url := fmt.Sprintf(ClashTournamentTeamByIDURL, teamID)

	var team TournamentTeamDto

	err := e.internalClient.Get(region, url, &team, "clash-v1", "TournamentTeamByID", "")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return &team, nil
}

// Get tournament by ID.
func (e *ClashEndpoint) ByID(region Region, tournamentID string) (*ClashTournamentDTO, error) {
	return e.getClash(ClashByIDURL, region, tournamentID, "ByID")
}

// Get tournament by team ID.
func (e *ClashEndpoint) ByTeamID(region Region, teamID string) (*ClashTournamentDTO, error) {
	return e.getClash(ClashByTeamIDURL, region, teamID, "ByTeamID")
}

func (e *ClashEndpoint) getClash(endpointMethod string, region Region, id string, methodName string) (*ClashTournamentDTO, error) {
	logger := e.internalClient.Logger("LOL", "clash-v1", methodName)
	logger.Debug("Method executed")

	if region == PBE1 {
		return nil, fmt.Errorf("the region PBE1 is not available for this method")
	}

	url := fmt.Sprintf(endpointMethod, id)

	var tournament ClashTournamentDTO

	err := e.internalClient.Get(region, url, &tournament, "clash-v1", methodName, "")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return &tournament, nil
}
