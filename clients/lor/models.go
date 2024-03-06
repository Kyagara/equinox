package lor

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 031d3e7fc343bd86d82c45559fc79d3a87fa1b82

// lor-deck-v1.NewDeckDto
type DeckNewDeckV1DTO struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

// lor-deck-v1.DeckDto
type DeckV1DTO struct {
	Code string `json:"code,omitempty"`
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// lor-inventory-v1.CardDto
type InventoryCardV1DTO struct {
	Code  string `json:"code,omitempty"`
	Count string `json:"count,omitempty"`
}

// lor-match-v1.InfoDto
type MatchInfoV1DTO struct {
	// (Legal values:  standard,  eternal)
	GameFormat string `json:"game_format,omitempty"`
	// (Legal values:  Constructed,  Expeditions,  Tutorial)
	GameMode         string `json:"game_mode,omitempty"`
	GameStartTimeUtc string `json:"game_start_time_utc,omitempty"`
	// (Legal values:  Ranked,  Normal,  AI,  Tutorial,  VanillaTrial,  Singleton,  StandardGauntlet)
	GameType    string             `json:"game_type,omitempty"`
	GameVersion string             `json:"game_version,omitempty"`
	Players     []MatchPlayerV1DTO `json:"players,omitempty"`
	// Total turns taken by both players.
	TotalTurnCount int32 `json:"total_turn_count,omitempty"`
}

// lor-match-v1.MetadataDto
type MatchMetadataV1DTO struct {
	// Match data version.
	DataVersion string `json:"data_version,omitempty"`
	// Match id.
	MatchID string `json:"match_id,omitempty"`
	// A list of participant PUUIDs.
	Participants []string `json:"participants,omitempty"`
}

// lor-match-v1.PlayerDto
type MatchPlayerV1DTO struct {
	// Code for the deck played. Refer to LOR documentation for details on deck codes.
	DeckCode    string   `json:"deck_code,omitempty"`
	DeckID      string   `json:"deck_id,omitempty"`
	GameOutcome string   `json:"game_outcome,omitempty"`
	PUUID       string   `json:"puuid,omitempty"`
	Factions    []string `json:"factions,omitempty"`
	// The order in which the players took turns.
	OrderOfPlay int32 `json:"order_of_play,omitempty"`
}

// lor-match-v1.MatchDto
type MatchV1DTO struct {
	// Match metadata.
	Metadata MatchMetadataV1DTO `json:"metadata,omitempty"`
	// Match info.
	Info MatchInfoV1DTO `json:"info,omitempty"`
}

// lor-ranked-v1.LeaderboardDto
type RankedLeaderboardV1DTO struct {
	// A list of players in Master tier.
	Players []RankedPlayerV1DTO `json:"players,omitempty"`
}

// lor-ranked-v1.PlayerDto
type RankedPlayerV1DTO struct {
	Name string `json:"name,omitempty"`
	// League points.
	LP   int32 `json:"lp,omitempty"`
	Rank int32 `json:"rank,omitempty"`
}

// lor-status-v1.ContentDto
type StatusContentV1DTO struct {
	Content string `json:"content,omitempty"`
	Locale  string `json:"locale,omitempty"`
}

// lor-status-v1.PlatformDataDto
type StatusPlatformDataV1DTO struct {
	ID           string        `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Incidents    []StatusV1DTO `json:"incidents,omitempty"`
	Locales      []string      `json:"locales,omitempty"`
	Maintenances []StatusV1DTO `json:"maintenances,omitempty"`
}

// lor-status-v1.UpdateDto
type StatusUpdateV1DTO struct {
	Author    string `json:"author,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	// (Legal values: riotclient, riotstatus, game)
	PublishLocations []string             `json:"publish_locations,omitempty"`
	Translations     []StatusContentV1DTO `json:"translations,omitempty"`
	ID               int32                `json:"id,omitempty"`
	Publish          bool                 `json:"publish,omitempty"`
}

// lor-status-v1.StatusDto
type StatusV1DTO struct {
	ArchiveAt string `json:"archive_at,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	// (Legal values:  info,  warning,  critical)
	IncidentSeverity string `json:"incident_severity,omitempty"`
	// (Legal values:  scheduled,  in_progress,  complete)
	MaintenanceStatus string `json:"maintenance_status,omitempty"`
	UpdatedAt         string `json:"updated_at,omitempty"`
	// (Legal values: windows, macos, android, ios, ps4, xbone, switch)
	Platforms []string             `json:"platforms,omitempty"`
	Titles    []StatusContentV1DTO `json:"titles,omitempty"`
	Updates   []StatusUpdateV1DTO  `json:"updates,omitempty"`
	ID        int32                `json:"id,omitempty"`
}
