package lor

import (
	"fmt"

	"github.com/Kyagara/equinox/internal"
)

type MatchEndpoint struct {
	internalClient *internal.InternalClient
}

type MatchDTO struct {
	// Match metadata.
	Metadata MetadataDTO `json:"metadata"`
	// Match info.
	Info InfoDTO `json:"info"`
}

type MetadataDTO struct {
	// Match data version.
	DataVersion string `json:"data_version"`
	// Match ID.
	MatchID string `json:"match_id"`
	// A list of participant PUUIDs.
	Participants []string `json:"participants"`
}

type InfoDTO struct {
	GameMode         GameMode     `json:"game_mode"`
	GameType         GameType     `json:"game_type"`
	GameStartTimeUtc string       `json:"game_start_time_utc"`
	GameVersion      string       `json:"game_version"`
	Players          []PlayersDTO `json:"players"`
	// Total turns taken by both players.
	TotalTurnCount int `json:"total_turn_count"`
}

type PlayersDTO struct {
	PUUID  string `json:"puuid"`
	DeckID string `json:"deck_id"`
	// Code for the deck played. Refer to LOR documentation for details on deck codes.
	DeckCode    string   `json:"deck_code"`
	Factions    []string `json:"factions"`
	GameOutcome string   `json:"game_outcome"`
	// The order in which the players took turns.
	OrderOfPlay int `json:"order_of_play"`
}

// Get match by ID.
func (e *MatchEndpoint) ByID(region Region, matchID string) (*MatchDTO, error) {
	logger := e.internalClient.Logger("LOR", "lor-match-v1", "ByID")

	logger.Debug("Method executed")

	url := fmt.Sprintf(MatchByIDURL, matchID)

	var match *MatchDTO

	err := e.internalClient.Get(region, url, &match, "lor-match-v1", "ByID", "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return match, nil
}

// Get a list of match ids by PUUID.
func (e *MatchEndpoint) List(region Region, puuid string) (*[]string, error) {
	logger := e.internalClient.Logger("LOR", "lor-match-v1", "List")

	logger.Debug("Method executed")

	url := fmt.Sprintf(MatchListURL, puuid)

	var list *[]string

	err := e.internalClient.Get(region, url, &list, "lor-match-v1", "List", "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return list, nil
}
