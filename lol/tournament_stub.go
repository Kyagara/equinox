package lol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type TournamentStubEndpoint struct {
	internalClient *internal.InternalClient
}

type LobbyEventDTOWrapper struct {
	EventList []LobbyEventDTO `json:"eventList"`
}

type LobbyEventDTO struct {
	// The summonerId that triggered the event (Encrypted)
	SummonerID string `json:"summonerId"`
	// The type of event that was triggered
	EventType string `json:"eventType"`
	// Timestamp from the event
	Timestamp string `json:"timestamp"`
}

type TournamentCodeParameters struct {
	// Optional list of encrypted summonerIds in order to validate the players eligible to join the lobby. NOTE: We currently do not enforce participants at the team level, but rather the aggregate of teamOne and teamTwo. We may add the ability to enforce at the team level in the future.
	AllowedSummonerIds []string `json:"allowedSummonerIds,omitempty"`
	// The map type of the game. (Legal values: SUMMONERS_RIFT, TWISTED_TREELINE, HOWLING_ABYSS)
	MapType MapType `json:"mapType"`
	// Optional string that may contain any data in any format, if specified at all. Used to denote any custom information about the game.
	Metadata string `json:"metadata,omitempty"`
	// The pick type of the game. (Legal values: BLIND_PICK, DRAFT_MODE, ALL_RANDOM, TOURNAMENT_DRAFT)
	PickType PickType `json:"pickType"`
	// The spectator type of the game. (Legal values: NONE, LOBBYONLY, ALL)
	SpectatorType SpectatorType `json:"spectatorType"`
	// The team size of the game. Valid values are 1-5.
	TeamSize int `json:"teamSize"`
}

// Create a mock tournament code for the given tournament. Count defaults to 20 (max 1000).
func (c *TournamentStubEndpoint) CreateCodes(tournamentID int64, count int, options TournamentCodeParameters) ([]string, error) {
	if count < 0 {
		count = 0
	}

	if options.TeamSize < 1 || options.TeamSize < 5 {
		return nil, fmt.Errorf("invalid team size: %d, valid values are 1-5", options.TeamSize)
	}

	if options.MapType == "" || options.SpectatorType == "" || options.PickType == "" {
		return nil, fmt.Errorf("required values are empty")
	}

	query := url.Values{}

	query.Set("count", strconv.Itoa(count))

	query.Set("tournamentId", strconv.FormatInt(tournamentID, 10))

	url := fmt.Sprintf("%s?%s", TournamentStubCodesURL, query.Encode())

	body, err := json.Marshal(options)

	if err != nil {
		return nil, err
	}

	res := []string{}

	err = c.internalClient.Do(http.MethodPost, api.RouteAmericas, url, bytes.NewBuffer(body), &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// Gets a mock list of lobby events by tournament code.
func (c *TournamentStubEndpoint) LobbyEvents(tournamentCode string) (*LobbyEventDTOWrapper, error) {
	url := fmt.Sprintf(TournamentStubLobbyEventsURL, tournamentCode)

	res := LobbyEventDTOWrapper{}

	err := c.internalClient.Do(http.MethodGet, api.RouteAmericas, url, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Creates a mock tournament provider and returns its ID.
//
// Providers will need to call this endpoint first to register their callback URL and their API key with the tournament system before any other tournament provider endpoints will work.
//
// The region in which the provider will be running tournaments. (Legal values: BR, EUNE, EUW, JP, LAN, LAS, NA, OCE, PBE, RU, TR)
//
// The provider's callback URL to which tournament game results in this region should be posted. The URL must be well-formed, use the http or https protocol, and use the default port for the protocol (http URLs must use port 80, https URLs must use port 443).
func (c *TournamentStubEndpoint) CreateTournamentProvider(region TournamentRegion, callbackURL string) (int, error) {
	_, err := url.ParseRequestURI(callbackURL)

	if err != nil {
		return -1, err
	}

	options := struct {
		Region TournamentRegion `json:"region"`
		URL    string           `json:"url"`
	}{Region: region, URL: callbackURL}

	body, err := json.Marshal(options)

	if err != nil {
		return -1, err
	}

	res := 0

	err = c.internalClient.Do(http.MethodPost, api.RouteAmericas, TournamentStubProvidersURL, bytes.NewBuffer(body), &res)

	if err != nil {
		return -1, err
	}

	return res, nil
}

// Creates a mock tournament and returns its ID.
//
// The provider ID to specify the regional registered provider data to associate this tournament.
//
// The optional name of the tournament.
func (c *TournamentStubEndpoint) CreateTournament(providerID int, name string) (int, error) {
	options := struct {
		Name       string `json:"name"`
		ProviderId int    `json:"providerId"`
	}{Name: name, ProviderId: providerID}

	body, err := json.Marshal(options)

	if err != nil {
		return -1, err
	}

	res := 0

	err = c.internalClient.Do(http.MethodPost, api.RouteAmericas, TournamentStubURL, bytes.NewBuffer(body), &res)

	if err != nil {
		return -1, err
	}

	return res, nil
}
