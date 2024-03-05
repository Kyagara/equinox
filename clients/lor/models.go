package lor

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 031d3e7fc343bd86d82c45559fc79d3a87fa1b82

// lor-inventory-v1.CardDto
type CardV1DTO struct {
	Code  string `json:"code,omitempty"`
	Count string `json:"count,omitempty"`
}

// lor-status-v1.ContentDto
type ContentV1DTO struct {
	Content string `json:"content,omitempty"`
	Locale  string `json:"locale,omitempty"`
}

// lor-deck-v1.DeckDto
type DeckV1DTO struct {
	Code string `json:"code,omitempty"`
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// lor-match-v1.InfoDto
type InfoV1DTO struct {
	// (Legal values:  standard,  eternal)
	GameFormat string `json:"game_format,omitempty"`
	// (Legal values:  Constructed,  Expeditions,  Tutorial)
	GameMode         string `json:"game_mode,omitempty"`
	GameStartTimeUtc string `json:"game_start_time_utc,omitempty"`
	// (Legal values:  Ranked,  Normal,  AI,  Tutorial,  VanillaTrial,  Singleton,  StandardGauntlet)
	GameType    string        `json:"game_type,omitempty"`
	GameVersion string        `json:"game_version,omitempty"`
	Players     []PlayerV1DTO `json:"players,omitempty"`
	// Total turns taken by both players.
	TotalTurnCount int32 `json:"total_turn_count,omitempty"`
}

// lor-ranked-v1.PlayerDto
type LeaderboardPlayerV1DTO struct {
	Name string `json:"name,omitempty"`
	// League points.
	LP   int32 `json:"lp,omitempty"`
	Rank int32 `json:"rank,omitempty"`
}

// lor-ranked-v1.LeaderboardDto
type LeaderboardV1DTO struct {
	// A list of players in Master tier.
	Players []LeaderboardPlayerV1DTO `json:"players,omitempty"`
}

// lor-match-v1.MatchDto
type MatchV1DTO struct {
	// Match metadata.
	Metadata MetadataV1DTO `json:"metadata,omitempty"`
	// Match info.
	Info InfoV1DTO `json:"info,omitempty"`
}

// lor-match-v1.MetadataDto
type MetadataV1DTO struct {
	// Match data version.
	DataVersion string `json:"data_version,omitempty"`
	// Match id.
	MatchID string `json:"match_id,omitempty"`
	// A list of participant PUUIDs.
	Participants []string `json:"participants,omitempty"`
}

// lor-deck-v1.NewDeckDto
type NewDeckV1DTO struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

// lor-status-v1.PlatformDataDto
type PlatformDataV1DTO struct {
	ID           string        `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Incidents    []StatusV1DTO `json:"incidents,omitempty"`
	Locales      []string      `json:"locales,omitempty"`
	Maintenances []StatusV1DTO `json:"maintenances,omitempty"`
}

// lor-match-v1.PlayerDto
type PlayerV1DTO struct {
	// Code for the deck played. Refer to LOR documentation for details on deck codes.
	DeckCode    string   `json:"deck_code,omitempty"`
	DeckID      string   `json:"deck_id,omitempty"`
	GameOutcome string   `json:"game_outcome,omitempty"`
	PUUID       string   `json:"puuid,omitempty"`
	Factions    []string `json:"factions,omitempty"`
	// The order in which the players took turns.
	OrderOfPlay int32 `json:"order_of_play,omitempty"`
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
	Platforms []string       `json:"platforms,omitempty"`
	Titles    []ContentV1DTO `json:"titles,omitempty"`
	Updates   []UpdateV1DTO  `json:"updates,omitempty"`
	ID        int32          `json:"id,omitempty"`
}

// lor-status-v1.UpdateDto
type UpdateV1DTO struct {
	Author    string `json:"author,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	// (Legal values: riotclient, riotstatus, game)
	PublishLocations []string       `json:"publish_locations,omitempty"`
	Translations     []ContentV1DTO `json:"translations,omitempty"`
	ID               int32          `json:"id,omitempty"`
	Publish          bool           `json:"publish,omitempty"`
}
