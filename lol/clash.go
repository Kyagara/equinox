package lol

import (
	"fmt"
	"net/http"

	"github.com/Kyagara/equinox/internal"
)

type ClashEndpoint struct {
	internalClient *internal.InternalClient
}

type TournamentDTO struct {
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
func (c *ClashEndpoint) Tournaments(region Region) (*[]TournamentDTO, error) {
	res := []TournamentDTO{}

	err := c.internalClient.Do(http.MethodGet, region, ClashURL, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get players by summoner ID.
//
// This endpoint returns a list of active Clash players for a given summoner ID. If a summoner registers for multiple tournaments at the same time (e.g., Saturday and Sunday) then both registrations would appear in this list.
func (c *ClashEndpoint) SummonerEntries(region Region, summonerID string) (*[]TournamentPlayerDTO, error) {
	url := fmt.Sprintf(ClashSummonerEntriesURL, summonerID)

	res := []TournamentPlayerDTO{}

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get team by ID.
func (c *ClashEndpoint) TournamentTeamByID(region Region, teamID string) (*TournamentTeamDto, error) {
	url := fmt.Sprintf(ClashTournamentTeamByIDURL, teamID)

	res := TournamentTeamDto{}

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get tournament by ID.
func (c *ClashEndpoint) ByID(region Region, tournamentID string) (*TournamentDTO, error) {
	return c.getClash(ClashByIDURL, region, tournamentID)
}

// Get tournament by team ID.
func (c *ClashEndpoint) ByTeamID(region Region, teamID string) (*TournamentDTO, error) {
	return c.getClash(ClashByTeamIDURL, region, teamID)
}

func (c *ClashEndpoint) getClash(method string, region Region, id string) (*TournamentDTO, error) {
	url := fmt.Sprintf(method, id)

	res := TournamentDTO{}

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
