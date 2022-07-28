package lol

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"go.uber.org/zap"
)

type TournamentStubEndpoint struct {
	internalClient *internal.InternalClient
}

// Create a mock tournament code for the given tournament.
func (e *TournamentStubEndpoint) CreateCodes(tournamentID int64, count int, parameters *TournamentCodeParametersDTO) (*[]string, error) {
	logger := e.internalClient.Logger("LOL", "tournament-stub-v4", "CreateCodes")

	logger.Debug("Method executed")

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

	var codes *[]string

	err := e.internalClient.Post(api.AmericasCluster, url, parameters, &codes, "tournament-stub-v4", "CreateCodes", "")

	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return codes, nil
}

// Gets a mock list of lobby events by tournament code.
func (e *TournamentStubEndpoint) LobbyEvents(tournamentCode string) (*LobbyEventDTOWrapper, error) {
	logger := e.internalClient.Logger("LOL", "tournament-stub-v4", "LobbyEvents")

	logger.Debug("Method executed")

	url := fmt.Sprintf(TournamentStubLobbyEventsURL, tournamentCode)

	var lobbyEvents *LobbyEventDTOWrapper

	err := e.internalClient.Get(api.AmericasCluster, url, &lobbyEvents, "tournament-stub-v4", "LobbyEvents", "")

	if err != nil {
		logger.Error("Method failed", zap.Error(err))
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
func (e *TournamentStubEndpoint) CreateProvider(region TournamentRegion, callbackURL string) (*int, error) {
	logger := e.internalClient.Logger("LOL", "tournament-stub", "CreateProvider")

	logger.Debug("Method executed")

	_, err := url.ParseRequestURI(callbackURL)

	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	options := struct {
		Region TournamentRegion `json:"region"`
		URL    string           `json:"url"`
	}{Region: region, URL: callbackURL}

	var provider *int

	err = e.internalClient.Post(api.AmericasCluster, TournamentStubProvidersURL, options, &provider, "tournament-stub", "CreateProvider", "")

	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return provider, nil
}

// Creates a mock tournament and returns its ID.
//
// The provider ID to specify the regional registered provider data to associate this tournament.
//
// The optional name of the tournament.
func (e *TournamentStubEndpoint) Create(providerID int, name string) (*int, error) {
	logger := e.internalClient.Logger("LOL", "tournament-stub-v4", "Create")

	logger.Debug("Method executed")

	options := struct {
		Name       string `json:"name"`
		ProviderId int    `json:"providerId"`
	}{Name: name, ProviderId: providerID}

	var tournament *int

	err := e.internalClient.Post(api.AmericasCluster, TournamentStubURL, options, &tournament, "tournament-stub-v4", "Create", "")

	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return tournament, nil
}
