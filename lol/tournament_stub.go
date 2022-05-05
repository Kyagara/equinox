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

// Create a mock tournament code for the given tournament.
func (t *TournamentStubEndpoint) CreateCodes(tournamentID int64, count int, parameters *TournamentCodeParametersDTO) ([]string, error) {
	logger := t.internalClient.Logger("lol").With("endpoint", "tournament-stub", "method", "CreateCodes")

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

	url := fmt.Sprintf("%s?%s", TournamentStubCodesURL, query.Encode())

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

// Gets a mock list of lobby events by tournament code.
func (t *TournamentStubEndpoint) LobbyEvents(tournamentCode string) (*LobbyEventDTOWrapper, error) {
	logger := t.internalClient.Logger("lol").With("endpoint", "tournament-stub", "method", "LobbyEvents")

	url := fmt.Sprintf(TournamentStubLobbyEventsURL, tournamentCode)

	var lobbyEvents *LobbyEventDTOWrapper

	err := t.internalClient.Do(http.MethodGet, api.Americas, url, nil, &lobbyEvents, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return lobbyEvents, nil
}

// Creates a mock tournament provider and returns its ID.
//
// Providers will need to call this endpoint first to register their callback URL and their API key with the tournament system before any other tournament provider endpoints will work.
//
// The region in which the provider will be running tournaments.
//
// The provider's callback URL to which tournament game results in this region should be posted. The URL must be well-formed, use the http or https protocol, and use the default port for the protocol (http URLs must use port 80, https URLs must use port 443).
func (t *TournamentStubEndpoint) CreateProvider(region TournamentRegion, callbackURL string) (int, error) {
	logger := t.internalClient.Logger("lol").With("endpoint", "tournament-stub", "method", "CreateProvider")

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

	err = t.internalClient.Do(http.MethodPost, api.Americas, TournamentStubProvidersURL, bytes.NewBuffer(body), &provider, "")

	if err != nil {
		logger.Warn(err)
		return -1, err
	}

	return provider, nil
}

// Creates a mock tournament and returns its ID.
//
// The provider ID to specify the regional registered provider data to associate this tournament.
//
// The optional name of the tournament.
func (t *TournamentStubEndpoint) Create(providerID int, name string) (int, error) {
	logger := t.internalClient.Logger("lol").With("endpoint", "tournament-stub", "method", "Create")

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

	err = t.internalClient.Do(http.MethodPost, api.Americas, TournamentStubURL, bytes.NewBuffer(body), &tournament, "")

	if err != nil {
		logger.Warn(err)
		return -1, err
	}

	return tournament, nil
}
