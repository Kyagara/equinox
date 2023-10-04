package lol

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"go.uber.org/zap"
)

type TournamentEndpoint struct {
	internalClient *internal.InternalClient
}

type LobbyEventDTOWrapper struct {
	EventList []LobbyEventDTO `json:"eventList"`
}

type LobbyEventDTO struct {
	// The PUUID that triggered the event (Encrypted).
	PUUID string `json:"PUUID"`
	// The type of event that was triggered.
	EventType string `json:"eventType"`
	// Timestamp from the event.
	Timestamp string `json:"timestamp"`
}

type TournamentCodeDTO struct {
	// The tournament code.
	Code string `json:"code"`
	// The spectator mode for the tournament code game.
	Spectators SpectatorType `json:"spectators"`
	// The lobby name for the tournament code game.
	LobbyName string `json:"lobbyName"`
	// The metadata for tournament code.
	MetaData string `json:"metaData"`
	// The password for the tournament code game.
	Password string `json:"password"`
	// The team size for the tournament code game.
	TeamSize int `json:"teamSize"`
	// The provider's ID.
	ProviderId int `json:"providerId"`
	// The pick mode for tournament code game.
	PickType PickType `json:"pickType"`
	// The tournament's ID.
	TournamentID int `json:"tournamentId"`
	// The tournament code's ID.
	ID int `json:"id"`
	// The tournament code's region.
	Region TournamentRegion `json:"region"`
	// The game map for the tournament code game.
	Map MapType `json:"map"`
	// The PUUIDs of the participants (Encrypted).
	Participants []string `json:"participants"`
}

type TournamentCodeParametersDTO struct {
	// Optional list of encrypted PUUIDs in order to validate the players eligible to join the lobby. NOTE: We currently do not enforce participants at the team level, but rather the aggregate of teamOne and teamTwo. We may add the ability to enforce at the team level in the future.
	AllowedParticipants []string `json:"allowedParticipants,omitempty"`
	// The map type of the game. (Legal values: SUMMONERS_RIFT, HOWLING_ABYSS).
	MapType MapType `json:"mapType"`
	// Optional string that may contain any data in any format, if specified at all. Used to denote any custom information about the game.
	Metadata string `json:"metadata,omitempty"`
	// The pick type of the game. (Legal values: BLIND_PICK, DRAFT_MODE, ALL_RANDOM, TOURNAMENT_DRAFT).
	PickType PickType `json:"pickType"`
	// The spectator type of the game. (Legal values: NONE, LOBBYONLY, ALL).
	SpectatorType SpectatorType `json:"spectatorType"`
	// The team size of the game. Valid values are 1-5.
	TeamSize int `json:"teamSize"`
	// Checks if allowed participants are enough to make full teams.
	EnoughPlayers int `json:"enoughPlayers"`
}

type TournamentCodeUpdateParametersDTO struct {
	// Optional list of encrypted PUUIDs in order to validate the players eligible to join the lobby. NOTE: We currently do not enforce participants at the team level, but rather the aggregate of teamOne and teamTwo. We may add the ability to enforce at the team level in the future.
	AllowedParticipants []string `json:"allowedParticipants,omitempty"`
	// The pick type. (Legal values: BLIND_PICK, DRAFT_MODE, ALL_RANDOM, TOURNAMENT_DRAFT).
	MapType MapType `json:"mapType,omitempty"`
	// The map type. (Legal values: SUMMONERS_RIFT, HOWLING_ABYSS).
	PickType PickType `json:"pickType,omitempty"`
	// The spectator type. (Legal values: NONE, LOBBYONLY, ALL).
	SpectatorType SpectatorType `json:"spectatorType,omitempty"`
}

// Create a tournament code for the given tournament.
func (e *TournamentEndpoint) CreateCodes(tournamentID int64, count int, parameters *TournamentCodeParametersDTO) (*[]string, error) {
	logger := e.internalClient.Logger("LOL", "tournament-v5", "CreateCodes")
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

	url := fmt.Sprintf("%s?%s", TournamentCodesURL, query.Encode())

	var codes []string

	err := e.internalClient.Post(api.AmericasCluster, url, parameters, &codes, "tournament-v5", "CreateCodes", "")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return &codes, nil
}

// Returns the tournament code DTO associated with a tournament code string.
func (e *TournamentEndpoint) ByCode(tournamentCode string) (*TournamentCodeDTO, error) {
	logger := e.internalClient.Logger("LOL", "tournament-v5", "ByCode")
	logger.Debug("Method executed")

	url := fmt.Sprintf(TournamentByCodeURL, tournamentCode)

	var tournament TournamentCodeDTO

	err := e.internalClient.Get(api.AmericasCluster, url, &tournament, "tournament-v5", "ByCode", "")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return &tournament, nil
}

// Update the pick type, map, spectator type, or allowed summoners for a code.
func (e *TournamentEndpoint) Update(tournamentCode string, parameters *TournamentCodeUpdateParametersDTO) error {
	logger := e.internalClient.Logger("LOL", "tournament-v5", "Update")
	logger.Debug("Method executed")

	if parameters == nil {
		return fmt.Errorf("parameters are required")
	}

	url := fmt.Sprintf(TournamentByCodeURL, tournamentCode)

	err := e.internalClient.Put(api.AmericasCluster, url, parameters, "tournament-v5", "Update")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return err
	}

	return nil
}

// Gets a list of lobby events by tournament code.
func (e *TournamentEndpoint) LobbyEvents(tournamentCode string) (*LobbyEventDTOWrapper, error) {
	logger := e.internalClient.Logger("LOL", "tournament-v5", "LobbyEvents")
	logger.Debug("Method executed")

	url := fmt.Sprintf(TournamentLobbyEventsURL, tournamentCode)

	var lobbyEvents LobbyEventDTOWrapper

	err := e.internalClient.Get(api.AmericasCluster, url, &lobbyEvents, "tournament-v5", "LobbyEvents", "")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return &lobbyEvents, nil
}

// Creates a tournament provider and returns its ID.
//
// Providers will need to call this endpoint first to register their callback URL and their API key with the tournament system before any other tournament provider endpoints will work.
//
// The region in which the provider will be running tournaments.
//
// The provider's callback URL to which tournament game results in this region should be posted. The URL must be well-formed, use the http or https protocol, and use the default port for the protocol (http URLs must use port 80, https URLs must use port 443).
func (e *TournamentEndpoint) CreateProvider(region TournamentRegion, callbackURL string) (int, error) {
	logger := e.internalClient.Logger("LOL", "tournament-v5", "CreateProvider")
	logger.Debug("Method executed")

	_, err := url.ParseRequestURI(callbackURL)
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return -1, err
	}

	options := struct {
		Region TournamentRegion `json:"region"`
		URL    string           `json:"url"`
	}{Region: region, URL: callbackURL}

	var provider int

	err = e.internalClient.Post(api.AmericasCluster, TournamentProvidersURL, options, &provider, "tournament-v5", "CreateProvider", "")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return -1, err
	}

	return provider, nil
}

// Creates a tournament and returns its ID.
//
// The provider ID to specify the regional registered provider data to associate this tournament.
//
// The name of the tournament is optional.
func (e *TournamentEndpoint) Create(providerID int, name string) (int, error) {
	logger := e.internalClient.Logger("LOL", "tournament-v5", "Create")
	logger.Debug("Method executed")

	options := struct {
		Name       string `json:"name"`
		ProviderId int    `json:"providerId"`
	}{Name: name, ProviderId: providerID}

	var tournament int

	err := e.internalClient.Post(api.AmericasCluster, TournamentURL, options, &tournament, "tournament-v5", "Create", "")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return -1, err
	}

	return tournament, nil
}
