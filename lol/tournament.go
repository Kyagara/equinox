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

type TournamentEndpoint struct {
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

type TournamentCodeDTO struct {
	// The tournament code.
	Code string
	// The spectator mode for the tournament code game.
	Spectators SpectatorType
	// The lobby name for the tournament code game.
	LobbyName string
	// The metadata for tournament code.
	MetaData string
	// The password for the tournament code game.
	Password string
	// The team size for the tournament code game.
	TeamSize int
	// The provider's ID.
	ProviderId int
	// 	The pick mode for tournament code game.
	PickType PickType
	// The tournament's ID.
	TournamentID int
	// The tournament code's ID.
	ID int
	// The tournament code's region.
	Region TournamentRegion
	// The game map for the tournament code game
	Map MapType
	// The summonerIds of the participants (Encrypted)
	Participants []string
}

type TournamentCodeParametersDTO struct {
	// Optional list of encrypted summonerIds in order to validate the players eligible to join the lobby. NOTE: We currently do not enforce participants at the team level, but rather the aggregate of teamOne and teamTwo. We may add the ability to enforce at the team level in the future.
	AllowedSummonerIds []string `json:"allowedSummonerIds,omitempty"`
	// The map type of the game.
	MapType MapType `json:"mapType"`
	// Optional string that may contain any data in any format, if specified at all. Used to denote any custom information about the game.
	Metadata string `json:"metadata,omitempty"`
	// The pick type of the game.
	PickType PickType `json:"pickType"`
	// The spectator type of the game.
	SpectatorType SpectatorType `json:"spectatorType"`
	// The team size of the game. Valid values are 1-5.
	TeamSize int `json:"teamSize"`
}

type TournamentCodeUpdateParametersDTO struct {
	// Optional list of encrypted summonerIds in order to validate the players eligible to join the lobby. NOTE: We currently do not enforce participants at the team level, but rather the aggregate of teamOne and teamTwo. We may add the ability to enforce at the team level in the future.
	AllowedSummonerIds []string `json:"allowedSummonerIds,omitempty"`
	// The map type of the game.
	MapType MapType `json:"mapType,omitempty"`
	// The pick type of the game.
	PickType PickType `json:"pickType,omitempty"`
	// The spectator type of the game.
	SpectatorType SpectatorType `json:"spectatorType,omitempty"`
}

// Create a tournament code for the given tournament.
func (t *TournamentEndpoint) CreateCodes(tournamentID int64, count int, parameters *TournamentCodeParametersDTO) ([]string, error) {
	logger := t.internalClient.Logger("lol").With("endpoint", "tournament", "method", "CreateCodes")

	if count < 1 || count > 1000 {
		return nil, fmt.Errorf("count can't be less than 1 or more than 1000")
	}

	if parameters == nil {
		return nil, fmt.Errorf("parameters are required")
	}

	if parameters.MapType == "" && parameters.SpectatorType == "" && parameters.PickType == "" {
		return nil, fmt.Errorf("required values are empty")
	}

	if parameters.MapType == "" || parameters.SpectatorType == "" || parameters.PickType == "" {
		return nil, fmt.Errorf("not all required values are set")
	}

	if parameters.TeamSize < 1 || parameters.TeamSize > 5 {
		return nil, fmt.Errorf("invalid team size: %d, valid values are 1-5", parameters.TeamSize)
	}

	query := url.Values{}

	query.Set("count", strconv.Itoa(count))

	query.Set("tournamentId", strconv.FormatInt(tournamentID, 10))

	url := fmt.Sprintf("%s?%s", TournamentCodesURL, query.Encode())

	body, err := json.Marshal(parameters)

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	var codes []string

	err = t.internalClient.Do(http.MethodPost, api.Americas, url, bytes.NewBuffer(body), &codes, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return codes, nil
}

// Returns the tournament code DTO associated with a tournament code string.
func (t *TournamentEndpoint) ByCode(tournamentCode string) (*TournamentCodeDTO, error) {
	logger := t.internalClient.Logger("lol").With("endpoint", "tournament", "method", "ByCode")

	url := fmt.Sprintf(TournamentByCodeURL, tournamentCode)

	var tournament *TournamentCodeDTO

	err := t.internalClient.Do(http.MethodGet, api.Americas, url, nil, &tournament, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return tournament, nil
}

// Update the pick type, map, spectator type, or allowed summoners for a code.
func (t *TournamentEndpoint) Update(tournamentCode string, parameters *TournamentCodeUpdateParametersDTO) error {
	logger := t.internalClient.Logger("lol").With("endpoint", "tournament", "method", "Update")

	if parameters == nil {
		return fmt.Errorf("parameters are required")
	}

	body, err := json.Marshal(parameters)

	if err != nil {
		logger.Error(err)
		return err
	}

	url := fmt.Sprintf(TournamentByCodeURL, tournamentCode)

	err = t.internalClient.Do(http.MethodPut, api.Americas, url, bytes.NewBuffer(body), nil, "")

	if err != nil {
		logger.Warn(err)
		return err
	}

	return nil
}

// Gets a list of lobby events by tournament code.
func (t *TournamentEndpoint) LobbyEvents(tournamentCode string) (*LobbyEventDTOWrapper, error) {
	logger := t.internalClient.Logger("lol").With("endpoint", "tournament", "method", "LobbyEvents")

	url := fmt.Sprintf(TournamentLobbyEventsURL, tournamentCode)

	var lobbyEvents *LobbyEventDTOWrapper

	err := t.internalClient.Do(http.MethodGet, api.Americas, url, nil, &lobbyEvents, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return lobbyEvents, nil
}

// Creates a tournament provider and returns its ID.
//
// Providers will need to call this endpoint first to register their callback URL and their API key with the tournament system before any other tournament provider endpoints will work.
//
// The region in which the provider will be running tournaments.
//
// The provider's callback URL to which tournament game results in this region should be posted. The URL must be well-formed, use the http or https protocol, and use the default port for the protocol (http URLs must use port 80, https URLs must use port 443).
func (t *TournamentEndpoint) CreateProvider(region TournamentRegion, callbackURL string) (int, error) {
	logger := t.internalClient.Logger("lol").With("endpoint", "tournament", "method", "CreateProvider")

	_, err := url.ParseRequestURI(callbackURL)

	if err != nil {
		logger.Error(err)
		return -1, err
	}

	options := struct {
		Region TournamentRegion `json:"region"`
		URL    string           `json:"url"`
	}{Region: region, URL: callbackURL}

	body, err := json.Marshal(options)

	if err != nil {
		logger.Error(err)
		return -1, err
	}

	var provider int

	err = t.internalClient.Do(http.MethodPost, api.Americas, TournamentProvidersURL, bytes.NewBuffer(body), &provider, "")

	if err != nil {
		logger.Warn(err)
		return -1, err
	}

	return provider, nil
}

// Creates a tournament and returns its ID.
//
// The provider ID to specify the regional registered provider data to associate this tournament.
//
// The optional name of the tournament.
func (t *TournamentEndpoint) Create(providerID int, name string) (int, error) {
	logger := t.internalClient.Logger("lol").With("endpoint", "tournament", "method", "Create")

	options := struct {
		Name       string `json:"name"`
		ProviderId int    `json:"providerId"`
	}{Name: name, ProviderId: providerID}

	body, err := json.Marshal(options)

	if err != nil {
		logger.Error(err)
		return -1, err
	}

	var tournament int

	err = t.internalClient.Do(http.MethodPost, api.Americas, TournamentURL, bytes.NewBuffer(body), &tournament, "")

	if err != nil {
		logger.Warn(err)
		return -1, err
	}

	return tournament, nil
}
