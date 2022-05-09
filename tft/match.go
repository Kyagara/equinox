package tft

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Kyagara/equinox/internal"
)

type MatchEndpoint struct {
	internalClient *internal.InternalClient
}

type MatchDTO struct {
	// Match metadata.
	Metadata Metadata `json:"metadata"`
	// Match info.
	Info InfoDTO `json:"info"`
}

type Metadata struct {
	// Match data version.
	DataVersion string `json:"data_version"`
	// Match ID.
	MatchID string `json:"match_id"`
	// A list of participant PUUIDs.
	Participants []string `json:"participants"`
}

type CompanionDTO struct {
	ContentID string `json:"content_ID"`
	SkinID    int    `json:"skin_ID"`
	Species   string `json:"species"`
}

type TraitsDTO struct {
	// Trait name.
	Name string `json:"name"`
	// Number of units with this trait.
	NumUnits int `json:"num_units"`
	// Current style for this trait.
	Style Style `json:"style"`
	// Current active tier for the trait.
	TierCurrent int `json:"tier_current"`
	// Total tiers for the trait.
	TierTotal int `json:"tier_total"`
}

type UnitsDTO struct {
	// This field was introduced in patch 9.22 with data_version 2.
	CharacterID string   `json:"character_id"`
	ItemNames   []string `json:"itemNames"`
	// A list of the unit's items. Please refer to the Teamfight Tactics documentation for item ids.
	Items []int `json:"items"`
	// Unit name. This field is often left blank.
	Name string `json:"name"`
	// Unit rarity. This doesn't equate to the unit cost.
	Rarity int `json:"rarity"`
	// Unit tier.
	Tier int `json:"tier"`
	// If a unit is chosen as part of the Fates set mechanic, the chosen trait will be indicated by this field. Otherwise this field is excluded from the response.
	Chosen string `json:"chosen,omitempty"`
}

type ParticipantsDTO struct {
	Augments []string `json:"augments"`
	// Participant's companion.
	Companion CompanionDTO `json:"companion"`
	// Gold left after participant was eliminated.
	GoldLeft int `json:"gold_left"`
	// The round the participant was eliminated in. Note: If the player was eliminated in stage 2-1 their last_round would be 5.
	LastRound int `json:"last_round"`
	// Participant Little Legend level. Note: This is not the number of active units.
	Level int `json:"level"`
	// Participant placement upon elimination.
	Placement int `json:"placement"`
	// Number of players the participant eliminated.
	PlayersEliminated int    `json:"players_eliminated"`
	PUUID             string `json:"puuid"`
	// The number of seconds before the participant was eliminated.
	TimeEliminated float64 `json:"time_eliminated"`
	// Damage the participant dealt to other players.
	TotalDamageToPlayers int `json:"total_damage_to_players"`
	// A complete list of traits for the participant's active units.
	Traits []TraitsDTO `json:"traits"`
	// A list of active units for the participant.
	Units []UnitsDTO `json:"units"`
}

type InfoDTO struct {
	// Unix timestamp.
	GameDatetime int64 `json:"game_datetime"`
	// Game length in seconds.
	GameLength float64 `json:"game_length"`
	// Game variation key. Game variations documented in TFT static data.
	GameVariation string `json:"game_variation"`
	// Game client version.
	GameVersion  string            `json:"game_version"`
	Participants []ParticipantsDTO `json:"participants"`
	// Please refer to the League of Legends documentation.
	QueueID     int    `json:"queue_id"`
	TFTGameType string `json:"tft_game_type,omitempty"`
	// Teamfight Tactics set number.
	TFTSetNumber int `json:"tft_set_number"`
}

// Get a list of match IDs by PUUID.
//
// Count defaults to 20.
func (e *MatchEndpoint) List(PUUID string, count int) (*[]string, error) {
	logger := e.internalClient.Logger("TFT", "match", "List")

	if count > 100 || count < 1 {
		count = 20
	}

	query := url.Values{}

	query.Set("count", strconv.Itoa(count))

	method := fmt.Sprintf(MatchListURL, PUUID)

	url := fmt.Sprintf("%s?%s", method, query.Encode())

	var list *[]string

	err := e.internalClient.Do(http.MethodGet, e.internalClient.Cluster, url, nil, &list, "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return list, nil
}

// Get a match by match ID.
func (e *MatchEndpoint) ByID(matchID string) (*MatchDTO, error) {
	logger := e.internalClient.Logger("TFT", "match", "ByID")

	url := fmt.Sprintf(MatchByIDURL, matchID)

	var match *MatchDTO

	err := e.internalClient.Do(http.MethodGet, e.internalClient.Cluster, url, nil, &match, "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return match, nil
}
