package lol

import (
	"fmt"
	"net/http"

	"github.com/Kyagara/equinox/internal"
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
	// (Legal values: CAPTAIN, MEMBER)
	Role string `json:"role"`
}

// Get all active or upcoming tournaments.
func (c *ClashEndpoint) Tournaments(region Region) (*[]ClashTournamentDTO, error) {
	logger := c.internalClient.Logger().With("endpoint", "clash", "method", "Tournaments")

	var tournaments *[]ClashTournamentDTO

	err := c.internalClient.Do(http.MethodGet, region, ClashURL, nil, &tournaments, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return tournaments, nil
}

// Get players by summoner ID.
//
// This endpoint returns a list of active Clash players for a given summoner ID. If a summoner registers for multiple tournaments at the same time (e.g., Saturday and Sunday) then both registrations would appear in this list.
func (c *ClashEndpoint) SummonerEntries(region Region, summonerID string) (*[]TournamentPlayerDTO, error) {
	logger := c.internalClient.Logger().With("endpoint", "clash", "method", "SummonerEntries")

	url := fmt.Sprintf(ClashSummonerEntriesURL, summonerID)

	var players *[]TournamentPlayerDTO

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &players, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return players, nil
}

// Get team by ID.
func (c *ClashEndpoint) TournamentTeamByID(region Region, teamID string) (*TournamentTeamDto, error) {
	logger := c.internalClient.Logger().With("endpoint", "clash", "method", "TournamentTeamByID")

	url := fmt.Sprintf(ClashTournamentTeamByIDURL, teamID)

	var team *TournamentTeamDto

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &team, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return team, nil
}

// Get tournament by ID.
func (c *ClashEndpoint) ByID(region Region, tournamentID string) (*ClashTournamentDTO, error) {
	return c.getClash(ClashByIDURL, region, tournamentID, "ByID")
}

// Get tournament by team ID.
func (c *ClashEndpoint) ByTeamID(region Region, teamID string) (*ClashTournamentDTO, error) {
	return c.getClash(ClashByTeamIDURL, region, teamID, "ByTeamID")
}

func (c *ClashEndpoint) getClash(endpointMethod string, region Region, id string, methodName string) (*ClashTournamentDTO, error) {
	logger := c.internalClient.Logger().With("endpoint", "clash", "method", methodName)

	url := fmt.Sprintf(endpointMethod, id)

	var tournament *ClashTournamentDTO

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &tournament, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return tournament, nil
}
