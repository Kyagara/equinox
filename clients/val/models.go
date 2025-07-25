package val

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = c5f59a3e27f5101b78b8c7eb9b3fb88318b4225d

// val-console-match-v1.AbilityCastsDto
type ConsoleMatchAbilityCastsV1DTO struct {
	Ability1Casts int `json:"ability1Casts,omitempty"`
	Ability2Casts int `json:"ability2Casts,omitempty"`
	GrenadeCasts  int `json:"grenadeCasts,omitempty"`
	UltimateCasts int `json:"ultimateCasts,omitempty"`
}

// val-console-match-v1.AbilityDto
type ConsoleMatchAbilityV1DTO struct {
	Ability1Effects string `json:"ability1Effects,omitempty"`
	Ability2Effects string `json:"ability2Effects,omitempty"`
	GrenadeEffects  string `json:"grenadeEffects,omitempty"`
	UltimateEffects string `json:"ultimateEffects,omitempty"`
}

// val-console-match-v1.CoachDto
type ConsoleMatchCoachV1DTO struct {
	PUUID  string `json:"puuid,omitempty"`
	TeamID string `json:"teamId,omitempty"`
}

// val-console-match-v1.DamageDto
type ConsoleMatchDamageV1DTO struct {
	// PUUID
	Receiver  string `json:"receiver,omitempty"`
	Bodyshots int    `json:"bodyshots,omitempty"`
	Damage    int    `json:"damage,omitempty"`
	Headshots int    `json:"headshots,omitempty"`
	Legshots  int    `json:"legshots,omitempty"`
}

// val-console-match-v1.EconomyDto
type ConsoleMatchEconomyV1DTO struct {
	Armor        string `json:"armor,omitempty"`
	Weapon       string `json:"weapon,omitempty"`
	LoadoutValue int    `json:"loadoutValue,omitempty"`
	Remaining    int    `json:"remaining,omitempty"`
	Spent        int    `json:"spent,omitempty"`
}

// val-console-match-v1.FinishingDamageDto
type ConsoleMatchFinishingDamageV1DTO struct {
	DamageItem          string `json:"damageItem,omitempty"`
	DamageType          string `json:"damageType,omitempty"`
	IsSecondaryFireMode bool   `json:"isSecondaryFireMode,omitempty"`
}

// val-console-match-v1.KillDto
type ConsoleMatchKillV1DTO struct {
	// PUUID
	Killer string `json:"killer,omitempty"`
	// PUUID
	Victim          string                           `json:"victim,omitempty"`
	FinishingDamage ConsoleMatchFinishingDamageV1DTO `json:"finishingDamage"`
	// List of PUUIDs
	Assistants                []string                           `json:"assistants,omitempty"`
	PlayerLocations           []ConsoleMatchPlayerLocationsV1DTO `json:"playerLocations,omitempty"`
	VictimLocation            ConsoleMatchLocationV1DTO          `json:"victimLocation"`
	TimeSinceGameStartMillis  int                                `json:"timeSinceGameStartMillis,omitempty"`
	TimeSinceRoundStartMillis int                                `json:"timeSinceRoundStartMillis,omitempty"`
}

// val-console-match-v1.LocationDto
type ConsoleMatchLocationV1DTO struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
}

// val-console-match-v1.MatchInfoDto
type ConsoleMatchMatchInfoV1DTO struct {
	CustomGameName     string `json:"customGameName,omitempty"`
	GameMode           string `json:"gameMode,omitempty"`
	MapID              string `json:"mapId,omitempty"`
	MatchID            string `json:"matchId,omitempty"`
	ProvisioningFlowID string `json:"provisioningFlowId,omitempty"`
	QueueID            string `json:"queueId,omitempty"`
	SeasonID           string `json:"seasonId,omitempty"`
	GameLengthMillis   int    `json:"gameLengthMillis,omitempty"`
	GameStartMillis    int    `json:"gameStartMillis,omitempty"`
	IsCompleted        bool   `json:"isCompleted,omitempty"`
	IsRanked           bool   `json:"isRanked,omitempty"`
}

// val-console-match-v1.MatchDto
type ConsoleMatchMatchV1DTO struct {
	Coaches      []ConsoleMatchCoachV1DTO       `json:"coaches,omitempty"`
	Players      []ConsoleMatchPlayerV1DTO      `json:"players,omitempty"`
	RoundResults []ConsoleMatchRoundResultV1DTO `json:"roundResults,omitempty"`
	Teams        []ConsoleMatchTeamV1DTO        `json:"teams,omitempty"`
	MatchInfo    ConsoleMatchMatchInfoV1DTO     `json:"matchInfo"`
}

// val-console-match-v1.MatchlistEntryDto
type ConsoleMatchMatchlistEntryV1DTO struct {
	MatchID             string `json:"matchId,omitempty"`
	QueueID             string `json:"queueId,omitempty"`
	GameStartTimeMillis int    `json:"gameStartTimeMillis,omitempty"`
}

// val-console-match-v1.MatchlistDto
type ConsoleMatchMatchlistV1DTO struct {
	PUUID   string                            `json:"puuid,omitempty"`
	History []ConsoleMatchMatchlistEntryV1DTO `json:"history,omitempty"`
}

// val-console-match-v1.PlayerLocationsDto
type ConsoleMatchPlayerLocationsV1DTO struct {
	PUUID       string                    `json:"puuid,omitempty"`
	Location    ConsoleMatchLocationV1DTO `json:"location"`
	ViewRadians float64                   `json:"viewRadians,omitempty"`
}

// val-console-match-v1.PlayerRoundStatsDto
type ConsoleMatchPlayerRoundStatsV1DTO struct {
	Ability ConsoleMatchAbilityV1DTO  `json:"ability"`
	PUUID   string                    `json:"puuid,omitempty"`
	Damage  []ConsoleMatchDamageV1DTO `json:"damage,omitempty"`
	Kills   []ConsoleMatchKillV1DTO   `json:"kills,omitempty"`
	Economy ConsoleMatchEconomyV1DTO  `json:"economy"`
	Score   int                       `json:"score,omitempty"`
}

// val-console-match-v1.PlayerStatsDto
type ConsoleMatchPlayerStatsV1DTO struct {
	AbilityCasts   ConsoleMatchAbilityCastsV1DTO `json:"abilityCasts"`
	Assists        int                           `json:"assists,omitempty"`
	Deaths         int                           `json:"deaths,omitempty"`
	Kills          int                           `json:"kills,omitempty"`
	PlaytimeMillis int                           `json:"playtimeMillis,omitempty"`
	RoundsPlayed   int                           `json:"roundsPlayed,omitempty"`
	Score          int                           `json:"score,omitempty"`
}

// val-console-match-v1.PlayerDto
type ConsoleMatchPlayerV1DTO struct {
	CharacterID     string                       `json:"characterId,omitempty"`
	GameName        string                       `json:"gameName,omitempty"`
	PUUID           string                       `json:"puuid,omitempty"`
	PartyID         string                       `json:"partyId,omitempty"`
	PlayerCard      string                       `json:"playerCard,omitempty"`
	PlayerTitle     string                       `json:"playerTitle,omitempty"`
	TagLine         string                       `json:"tagLine,omitempty"`
	TeamID          string                       `json:"teamId,omitempty"`
	Stats           ConsoleMatchPlayerStatsV1DTO `json:"stats"`
	CompetitiveTier int                          `json:"competitiveTier,omitempty"`
}

// val-console-match-v1.RecentMatchesDto
type ConsoleMatchRecentMatchesV1DTO struct {
	// A list of recent match ids.
	MatchIDs    []string `json:"matchIds,omitempty"`
	CurrentTime int      `json:"currentTime,omitempty"`
}

// val-console-match-v1.RoundResultDto
type ConsoleMatchRoundResultV1DTO struct {
	// PUUID of player
	BombDefuser string `json:"bombDefuser,omitempty"`
	// PUUID of player
	BombPlanter           string                              `json:"bombPlanter,omitempty"`
	PlantSite             string                              `json:"plantSite,omitempty"`
	RoundCeremony         string                              `json:"roundCeremony,omitempty"`
	RoundResult           string                              `json:"roundResult,omitempty"`
	RoundResultCode       string                              `json:"roundResultCode,omitempty"`
	WinningTeam           string                              `json:"winningTeam,omitempty"`
	DefusePlayerLocations []ConsoleMatchPlayerLocationsV1DTO  `json:"defusePlayerLocations,omitempty"`
	PlantPlayerLocations  []ConsoleMatchPlayerLocationsV1DTO  `json:"plantPlayerLocations,omitempty"`
	PlayerStats           []ConsoleMatchPlayerRoundStatsV1DTO `json:"playerStats,omitempty"`
	DefuseLocation        ConsoleMatchLocationV1DTO           `json:"defuseLocation"`
	PlantLocation         ConsoleMatchLocationV1DTO           `json:"plantLocation"`
	DefuseRoundTime       int                                 `json:"defuseRoundTime,omitempty"`
	PlantRoundTime        int                                 `json:"plantRoundTime,omitempty"`
	RoundNum              int                                 `json:"roundNum,omitempty"`
}

// val-console-match-v1.TeamDto
type ConsoleMatchTeamV1DTO struct {
	// This is an arbitrary string. Red and Blue in bomb modes. The puuid of the player in deathmatch.
	TeamID string `json:"teamId,omitempty"`
	// Team points scored. Number of kills in deathmatch.
	NumPoints    int  `json:"numPoints,omitempty"`
	RoundsPlayed int  `json:"roundsPlayed,omitempty"`
	RoundsWon    int  `json:"roundsWon,omitempty"`
	Won          bool `json:"won,omitempty"`
}

// val-console-ranked-v1.LeaderboardDto
type ConsoleRankedLeaderboardV1DTO struct {
	// The act id for the given leaderboard. Act ids can be found using the val-content API.
	ActID string `json:"actId,omitempty"`
	// The shard for the given leaderboard.
	Shard   string                     `json:"shard,omitempty"`
	Players []ConsoleRankedPlayerV1DTO `json:"players,omitempty"`
	// The total number of players in the leaderboard.
	TotalPlayers int `json:"totalPlayers,omitempty"`
}

// val-console-ranked-v1.PlayerDto
type ConsoleRankedPlayerV1DTO struct {
	// This field may be omitted if the player has been anonymized.
	GameName string `json:"gameName,omitempty"`
	// This field may be omitted if the player has been anonymized.
	PUUID string `json:"puuid,omitempty"`
	// This field may be omitted if the player has been anonymized.
	TagLine         string `json:"tagLine,omitempty"`
	LeaderboardRank int    `json:"leaderboardRank,omitempty"`
	NumberOfWins    int    `json:"numberOfWins,omitempty"`
	RankedRating    int    `json:"rankedRating,omitempty"`
}

// val-content-v1.ActDto
type ContentActV1DTO struct {
	// This field is excluded from the response when a locale is set
	LocalizedNames ContentLocalizedNamesV1DTO `json:"localizedNames"`
	ID             string                     `json:"id,omitempty"`
	Name           string                     `json:"name,omitempty"`
	ParentID       string                     `json:"parentId,omitempty"`
	Type           string                     `json:"type,omitempty"`
	IsActive       bool                       `json:"isActive,omitempty"`
}

// val-content-v1.ContentItemDto
type ContentItemV1DTO struct {
	AssetName string `json:"assetName,omitempty"`
	// This field is only included for maps and game modes. These values are used in the match response.
	AssetPath string `json:"assetPath,omitempty"`
	ID        string `json:"id,omitempty"`
	// This field is excluded from the response when a locale is set
	LocalizedNames ContentLocalizedNamesV1DTO `json:"localizedNames"`
	Name           string                     `json:"name,omitempty"`
}

// val-content-v1.LocalizedNamesDto
type ContentLocalizedNamesV1DTO struct {
	ArAe string `json:"ar-AE,omitempty"`
	DeDe string `json:"de-DE,omitempty"`
	EnGb string `json:"en-GB,omitempty"`
	EnUs string `json:"en-US,omitempty"`
	EsEs string `json:"es-ES,omitempty"`
	EsMx string `json:"es-MX,omitempty"`
	FrFr string `json:"fr-FR,omitempty"`
	IdID string `json:"id-ID,omitempty"`
	ItIt string `json:"it-IT,omitempty"`
	JaJp string `json:"ja-JP,omitempty"`
	KoKr string `json:"ko-KR,omitempty"`
	PlPl string `json:"pl-PL,omitempty"`
	PtBr string `json:"pt-BR,omitempty"`
	RuRu string `json:"ru-RU,omitempty"`
	ThTh string `json:"th-TH,omitempty"`
	TrTr string `json:"tr-TR,omitempty"`
	ViVn string `json:"vi-VN,omitempty"`
	ZhCn string `json:"zh-CN,omitempty"`
	ZhTw string `json:"zh-TW,omitempty"`
}

// val-content-v1.ContentDto
type ContentV1DTO struct {
	Version      string             `json:"version,omitempty"`
	Acts         []ContentActV1DTO  `json:"acts,omitempty"`
	Ceremonies   []ContentItemV1DTO `json:"ceremonies,omitempty"`
	Characters   []ContentItemV1DTO `json:"characters,omitempty"`
	CharmLevels  []ContentItemV1DTO `json:"charmLevels,omitempty"`
	Charms       []ContentItemV1DTO `json:"charms,omitempty"`
	Chromas      []ContentItemV1DTO `json:"chromas,omitempty"`
	Equips       []ContentItemV1DTO `json:"equips,omitempty"`
	GameModes    []ContentItemV1DTO `json:"gameModes,omitempty"`
	Maps         []ContentItemV1DTO `json:"maps,omitempty"`
	PlayerCards  []ContentItemV1DTO `json:"playerCards,omitempty"`
	PlayerTitles []ContentItemV1DTO `json:"playerTitles,omitempty"`
	SkinLevels   []ContentItemV1DTO `json:"skinLevels,omitempty"`
	Skins        []ContentItemV1DTO `json:"skins,omitempty"`
	SprayLevels  []ContentItemV1DTO `json:"sprayLevels,omitempty"`
	Sprays       []ContentItemV1DTO `json:"sprays,omitempty"`
	Totems       []ContentItemV1DTO `json:"totems,omitempty"`
}

// val-match-v1.AbilityCastsDto
type MatchAbilityCastsV1DTO struct {
	Ability1Casts int `json:"ability1Casts,omitempty"`
	Ability2Casts int `json:"ability2Casts,omitempty"`
	GrenadeCasts  int `json:"grenadeCasts,omitempty"`
	UltimateCasts int `json:"ultimateCasts,omitempty"`
}

// val-match-v1.AbilityDto
type MatchAbilityV1DTO struct {
	Ability1Effects string `json:"ability1Effects,omitempty"`
	Ability2Effects string `json:"ability2Effects,omitempty"`
	GrenadeEffects  string `json:"grenadeEffects,omitempty"`
	UltimateEffects string `json:"ultimateEffects,omitempty"`
}

// val-match-v1.CoachDto
type MatchCoachV1DTO struct {
	PUUID  string `json:"puuid,omitempty"`
	TeamID string `json:"teamId,omitempty"`
}

// val-match-v1.DamageDto
type MatchDamageV1DTO struct {
	// PUUID
	Receiver  string `json:"receiver,omitempty"`
	Bodyshots int    `json:"bodyshots,omitempty"`
	Damage    int    `json:"damage,omitempty"`
	Headshots int    `json:"headshots,omitempty"`
	Legshots  int    `json:"legshots,omitempty"`
}

// val-match-v1.EconomyDto
type MatchEconomyV1DTO struct {
	Armor        string `json:"armor,omitempty"`
	Weapon       string `json:"weapon,omitempty"`
	LoadoutValue int    `json:"loadoutValue,omitempty"`
	Remaining    int    `json:"remaining,omitempty"`
	Spent        int    `json:"spent,omitempty"`
}

// val-match-v1.FinishingDamageDto
type MatchFinishingDamageV1DTO struct {
	DamageItem          string `json:"damageItem,omitempty"`
	DamageType          string `json:"damageType,omitempty"`
	IsSecondaryFireMode bool   `json:"isSecondaryFireMode,omitempty"`
}

// val-match-v1.MatchInfoDto
type MatchInfoV1DTO struct {
	PremierMatchInfo   map[string]any `json:"premierMatchInfo,omitempty"`
	CustomGameName     string         `json:"customGameName,omitempty"`
	GameMode           string         `json:"gameMode,omitempty"`
	GameVersion        string         `json:"gameVersion,omitempty"`
	MapID              string         `json:"mapId,omitempty"`
	MatchID            string         `json:"matchId,omitempty"`
	ProvisioningFlowID string         `json:"provisioningFlowId,omitempty"`
	QueueID            string         `json:"queueId,omitempty"`
	Region             string         `json:"region,omitempty"`
	SeasonID           string         `json:"seasonId,omitempty"`
	GameLengthMillis   int            `json:"gameLengthMillis,omitempty"`
	GameStartMillis    int            `json:"gameStartMillis,omitempty"`
	IsCompleted        bool           `json:"isCompleted,omitempty"`
	IsRanked           bool           `json:"isRanked,omitempty"`
}

// val-match-v1.KillDto
type MatchKillV1DTO struct {
	// PUUID
	Killer string `json:"killer,omitempty"`
	// PUUID
	Victim          string                    `json:"victim,omitempty"`
	FinishingDamage MatchFinishingDamageV1DTO `json:"finishingDamage"`
	// List of PUUIDs
	Assistants                []string                    `json:"assistants,omitempty"`
	PlayerLocations           []MatchPlayerLocationsV1DTO `json:"playerLocations,omitempty"`
	VictimLocation            MatchLocationV1DTO          `json:"victimLocation"`
	TimeSinceGameStartMillis  int                         `json:"timeSinceGameStartMillis,omitempty"`
	TimeSinceRoundStartMillis int                         `json:"timeSinceRoundStartMillis,omitempty"`
}

// val-match-v1.LocationDto
type MatchLocationV1DTO struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
}

// val-match-v1.PlayerLocationsDto
type MatchPlayerLocationsV1DTO struct {
	PUUID       string             `json:"puuid,omitempty"`
	Location    MatchLocationV1DTO `json:"location"`
	ViewRadians float64            `json:"viewRadians,omitempty"`
}

// val-match-v1.PlayerRoundStatsDto
type MatchPlayerRoundStatsV1DTO struct {
	Ability MatchAbilityV1DTO  `json:"ability"`
	PUUID   string             `json:"puuid,omitempty"`
	Damage  []MatchDamageV1DTO `json:"damage,omitempty"`
	Kills   []MatchKillV1DTO   `json:"kills,omitempty"`
	Economy MatchEconomyV1DTO  `json:"economy"`
	Score   int                `json:"score,omitempty"`
}

// val-match-v1.PlayerStatsDto
type MatchPlayerStatsV1DTO struct {
	AbilityCasts   MatchAbilityCastsV1DTO `json:"abilityCasts"`
	Assists        int                    `json:"assists,omitempty"`
	Deaths         int                    `json:"deaths,omitempty"`
	Kills          int                    `json:"kills,omitempty"`
	PlaytimeMillis int                    `json:"playtimeMillis,omitempty"`
	RoundsPlayed   int                    `json:"roundsPlayed,omitempty"`
	Score          int                    `json:"score,omitempty"`
}

// val-match-v1.PlayerDto
type MatchPlayerV1DTO struct {
	CharacterID     string                `json:"characterId,omitempty"`
	GameName        string                `json:"gameName,omitempty"`
	PUUID           string                `json:"puuid,omitempty"`
	PartyID         string                `json:"partyId,omitempty"`
	PlayerCard      string                `json:"playerCard,omitempty"`
	PlayerTitle     string                `json:"playerTitle,omitempty"`
	TagLine         string                `json:"tagLine,omitempty"`
	TeamID          string                `json:"teamId,omitempty"`
	Stats           MatchPlayerStatsV1DTO `json:"stats"`
	AccountLevel    int                   `json:"accountLevel,omitempty"`
	CompetitiveTier int                   `json:"competitiveTier,omitempty"`
	IsObserver      bool                  `json:"isObserver,omitempty"`
}

// val-match-v1.RecentMatchesDto
type MatchRecentMatchesV1DTO struct {
	// A list of recent match ids.
	MatchIDs    []string `json:"matchIds,omitempty"`
	CurrentTime int      `json:"currentTime,omitempty"`
}

// val-match-v1.RoundResultDto
type MatchRoundResultV1DTO struct {
	// PUUID of player
	BombDefuser string `json:"bombDefuser,omitempty"`
	// PUUID of player
	BombPlanter           string                       `json:"bombPlanter,omitempty"`
	PlantSite             string                       `json:"plantSite,omitempty"`
	RoundCeremony         string                       `json:"roundCeremony,omitempty"`
	RoundResult           string                       `json:"roundResult,omitempty"`
	RoundResultCode       string                       `json:"roundResultCode,omitempty"`
	WinningTeam           string                       `json:"winningTeam,omitempty"`
	WinningTeamRole       string                       `json:"winningTeamRole,omitempty"`
	DefusePlayerLocations []MatchPlayerLocationsV1DTO  `json:"defusePlayerLocations,omitempty"`
	PlantPlayerLocations  []MatchPlayerLocationsV1DTO  `json:"plantPlayerLocations,omitempty"`
	PlayerStats           []MatchPlayerRoundStatsV1DTO `json:"playerStats,omitempty"`
	DefuseLocation        MatchLocationV1DTO           `json:"defuseLocation"`
	PlantLocation         MatchLocationV1DTO           `json:"plantLocation"`
	DefuseRoundTime       int                          `json:"defuseRoundTime,omitempty"`
	PlantRoundTime        int                          `json:"plantRoundTime,omitempty"`
	RoundNum              int                          `json:"roundNum,omitempty"`
}

// val-match-v1.TeamDto
type MatchTeamV1DTO struct {
	// This is an arbitrary string. Red and Blue in bomb modes. The puuid of the player in deathmatch.
	TeamID string `json:"teamId,omitempty"`
	// Team points scored. Number of kills in deathmatch.
	NumPoints    int  `json:"numPoints,omitempty"`
	RoundsPlayed int  `json:"roundsPlayed,omitempty"`
	RoundsWon    int  `json:"roundsWon,omitempty"`
	Won          bool `json:"won,omitempty"`
}

// val-match-v1.MatchDto
type MatchV1DTO struct {
	Coaches      []MatchCoachV1DTO       `json:"coaches,omitempty"`
	Players      []MatchPlayerV1DTO      `json:"players,omitempty"`
	RoundResults []MatchRoundResultV1DTO `json:"roundResults,omitempty"`
	Teams        []MatchTeamV1DTO        `json:"teams,omitempty"`
	MatchInfo    MatchInfoV1DTO          `json:"matchInfo"`
}

// val-match-v1.MatchlistEntryDto
type MatchlistEntryV1DTO struct {
	MatchID             string `json:"matchId,omitempty"`
	QueueID             string `json:"queueId,omitempty"`
	GameStartTimeMillis int    `json:"gameStartTimeMillis,omitempty"`
}

// val-match-v1.MatchlistDto
type MatchlistV1DTO struct {
	PUUID   string                `json:"puuid,omitempty"`
	History []MatchlistEntryV1DTO `json:"history,omitempty"`
}

// val-ranked-v1.LeaderboardDto
type RankedLeaderboardV1DTO struct {
	TierDetails map[int]RankedTierDetailV1DTO `json:"tierDetails,omitempty"`
	// The act id for the given leaderboard. Act ids can be found using the val-content API.
	ActID string `json:"actId,omitempty"`
	Query string `json:"query,omitempty"`
	// The shard for the given leaderboard.
	Shard                 string              `json:"shard,omitempty"`
	Players               []RankedPlayerV1DTO `json:"players,omitempty"`
	ImmortalStartingIndex int                 `json:"immortalStartingIndex,omitempty"`
	ImmortalStartingPage  int                 `json:"immortalStartingPage,omitempty"`
	StartIndex            int                 `json:"startIndex,omitempty"`
	TopTierRrthreshold    int                 `json:"topTierRRThreshold,omitempty"`
	// The total number of players in the leaderboard.
	TotalPlayers int `json:"totalPlayers,omitempty"`
}

// val-ranked-v1.PlayerDto
type RankedPlayerV1DTO struct {
	// This field may be omitted if the player has been anonymized.
	GameName string `json:"gameName,omitempty"`
	// This field may be omitted if the player has been anonymized.
	PUUID  string `json:"puuid,omitempty"`
	Prefix string `json:"prefix,omitempty"`
	// This field may be omitted if the player has been anonymized.
	TagLine         string `json:"tagLine,omitempty"`
	CompetitiveTier int    `json:"competitiveTier,omitempty"`
	LeaderboardRank int    `json:"leaderboardRank,omitempty"`
	NumberOfWins    int    `json:"numberOfWins,omitempty"`
	RankedRating    int    `json:"rankedRating,omitempty"`
}

// val-ranked-v1.TierDetailDto
type RankedTierDetailV1DTO struct {
	RankedRatingThreshold int `json:"rankedRatingThreshold,omitempty"`
	StartingIndex         int `json:"startingIndex,omitempty"`
	StartingPage          int `json:"startingPage,omitempty"`
}

// val-status-v1.ContentDto
type StatusContentV1DTO struct {
	Content string `json:"content,omitempty"`
	Locale  string `json:"locale,omitempty"`
}

// val-status-v1.PlatformDataDto
type StatusPlatformDataV1DTO struct {
	ID           string        `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Incidents    []StatusV1DTO `json:"incidents,omitempty"`
	Locales      []string      `json:"locales,omitempty"`
	Maintenances []StatusV1DTO `json:"maintenances,omitempty"`
}

// val-status-v1.UpdateDto
type StatusUpdateV1DTO struct {
	Author    string `json:"author,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	// (Legal values: riotclient, riotstatus, game)
	PublishLocations []string             `json:"publish_locations,omitempty"`
	Translations     []StatusContentV1DTO `json:"translations,omitempty"`
	ID               int                  `json:"id,omitempty"`
	Publish          bool                 `json:"publish,omitempty"`
}

// val-status-v1.StatusDto
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
	ID        int                  `json:"id,omitempty"`
}
