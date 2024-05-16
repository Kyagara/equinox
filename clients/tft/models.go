package tft

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 6461993a9c4165ddca053929f19f6d0e3eb1ca14

// tft-league-v1.LeagueEntryDTO
type LeagueEntryV1DTO struct {
	// Not included for the RANKED_TFT_TURBO queueType.
	LeagueID string `json:"leagueId,omitempty"`
	// Player Universal Unique Identifier. Exact length of 78 characters. (Encrypted)
	PUUID string `json:"puuid,omitempty"`
	// The player's division within a tier. Not included for the RANKED_TFT_TURBO queueType.
	Rank Division `json:"rank,omitempty"`
	// Only included for the RANKED_TFT_TURBO queueType.
	//
	// (Legal values:  ORANGE,  PURPLE,  BLUE,  GREEN,  GRAY)
	RatedTier string `json:"ratedTier,omitempty"`
	// Player's encrypted summonerId.
	SummonerID string `json:"summonerId,omitempty"`
	// Not included for the RANKED_TFT_TURBO queueType.
	Tier Tier `json:"tier,omitempty"`
	// Not included for the RANKED_TFT_TURBO queueType.
	MiniSeries LeagueMiniSeriesV1DTO `json:"miniSeries,omitempty"`
	// Not included for the RANKED_TFT_TURBO queueType.
	LeaguePoints int32 `json:"leaguePoints,omitempty"`
	// Second through eighth placement.
	Losses    int32     `json:"losses,omitempty"`
	QueueType QueueType `json:"queueType,omitempty"`
	// Only included for the RANKED_TFT_TURBO queueType.
	RatedRating int32 `json:"ratedRating,omitempty"`
	// First placement.
	Wins int32 `json:"wins,omitempty"`
	// Not included for the RANKED_TFT_TURBO queueType.
	FreshBlood bool `json:"freshBlood,omitempty"`
	// Not included for the RANKED_TFT_TURBO queueType.
	HotStreak bool `json:"hotStreak,omitempty"`
	// Not included for the RANKED_TFT_TURBO queueType.
	Inactive bool `json:"inactive,omitempty"`
	// Not included for the RANKED_TFT_TURBO queueType.
	Veteran bool `json:"veteran,omitempty"`
}

// tft-league-v1.LeagueItemDTO
type LeagueItemV1DTO struct {
	Rank Division `json:"rank,omitempty"`
	// Player's encrypted summonerId.
	SummonerID   string                `json:"summonerId,omitempty"`
	MiniSeries   LeagueMiniSeriesV1DTO `json:"miniSeries,omitempty"`
	LeaguePoints int32                 `json:"leaguePoints,omitempty"`
	// Second through eighth placement.
	Losses int32 `json:"losses,omitempty"`
	// First placement.
	Wins       int32 `json:"wins,omitempty"`
	FreshBlood bool  `json:"freshBlood,omitempty"`
	HotStreak  bool  `json:"hotStreak,omitempty"`
	Inactive   bool  `json:"inactive,omitempty"`
	Veteran    bool  `json:"veteran,omitempty"`
}

// tft-league-v1.LeagueListDTO
type LeagueListV1DTO struct {
	LeagueID string            `json:"leagueId,omitempty"`
	Name     string            `json:"name,omitempty"`
	Tier     Tier              `json:"tier,omitempty"`
	Entries  []LeagueItemV1DTO `json:"entries,omitempty"`
	Queue    QueueType         `json:"queue,omitempty"`
}

// tft-league-v1.MiniSeriesDTO
type LeagueMiniSeriesV1DTO struct {
	Progress string `json:"progress,omitempty"`
	Losses   int32  `json:"losses,omitempty"`
	Target   int32  `json:"target,omitempty"`
	Wins     int32  `json:"wins,omitempty"`
}

// tft-league-v1.TopRatedLadderEntryDto
type LeagueTopRatedLadderEntryV1DTO struct {
	// (Legal values:  ORANGE,  PURPLE,  BLUE,  GREEN,  GRAY)
	RatedTier                    string `json:"ratedTier,omitempty"`
	SummonerID                   string `json:"summonerId,omitempty"`
	PreviousUpdateLadderPosition int32  `json:"previousUpdateLadderPosition,omitempty"`
	RatedRating                  int32  `json:"ratedRating,omitempty"`
	// First placement.
	Wins int32 `json:"wins,omitempty"`
}

// tft-match-v1.CompanionDto
type MatchCompanionV1DTO struct {
	ContentID string `json:"content_ID,omitempty"`
	Species   string `json:"species,omitempty"`
	ItemID    int32  `json:"item_ID,omitempty"`
	SkinID    int32  `json:"skin_ID,omitempty"`
}

// tft-match-v1.InfoDto
type MatchInfoV1DTO struct {
	EndOfGameResult string `json:"endOfGameResult,omitempty"`
	// Game variation key. Game variations documented in TFT static data.
	GameVariation string `json:"game_variation,omitempty"`
	// Game client version.
	GameVersion    string                  `json:"game_version,omitempty"`
	TFTGameType    string                  `json:"tft_game_type,omitempty"`
	TFTSetCoreName string                  `json:"tft_set_core_name,omitempty"`
	Participants   []MatchParticipantV1DTO `json:"participants,omitempty"`
	GameCreation   int64                   `json:"gameCreation,omitempty"`
	// Unix timestamp.
	GameDatetime int64 `json:"game_datetime,omitempty"`
	GameID       int64 `json:"gameId,omitempty"`
	MapID        int64 `json:"mapId,omitempty"`
	// Game length in seconds.
	GameLength float32 `json:"game_length,omitempty"`
	// Please refer to the League of Legends documentation.
	QueueID int32 `json:"queueId,omitempty"`
	// Please refer to the League of Legends documentation.
	QueueID_ int32 `json:"queue_id,omitempty"`
	// Teamfight Tactics set number.
	TFTSetNumber int32 `json:"tft_set_number,omitempty"`
}

// tft-match-v1.MetadataDto
type MatchMetadataV1DTO struct {
	// Match data version.
	DataVersion string `json:"data_version,omitempty"`
	// Match id.
	MatchID string `json:"match_id,omitempty"`
	// A list of participant PUUIDs.
	Participants []string `json:"participants,omitempty"`
}

// tft-match-v1.ParticipantMissionsDto
type MatchParticipantMissionsV1DTO struct {
	Assists                        int32 `json:"Assists,omitempty"`
	DamageDealt                    int32 `json:"DamageDealt,omitempty"`
	DamageDealtToObjectives        int32 `json:"DamageDealtToObjectives,omitempty"`
	DamageDealtToTurrets           int32 `json:"DamageDealtToTurrets,omitempty"`
	DamageTaken                    int32 `json:"DamageTaken,omitempty"`
	Deaths                         int32 `json:"Deaths,omitempty"`
	DoubleKills                    int32 `json:"DoubleKills,omitempty"`
	GoldEarned                     int32 `json:"GoldEarned,omitempty"`
	GoldSpent                      int32 `json:"GoldSpent,omitempty"`
	InhibitorsDestroyed            int32 `json:"InhibitorsDestroyed,omitempty"`
	KillingSprees                  int32 `json:"KillingSprees,omitempty"`
	Kills                          int32 `json:"Kills,omitempty"`
	LargestKillingSpree            int32 `json:"LargestKillingSpree,omitempty"`
	LargestMultiKill               int32 `json:"LargestMultiKill,omitempty"`
	MagicDamageDealt               int32 `json:"MagicDamageDealt,omitempty"`
	MagicDamageDealtToChampions    int32 `json:"MagicDamageDealtToChampions,omitempty"`
	MagicDamageTaken               int32 `json:"MagicDamageTaken,omitempty"`
	NeutralMinionsKilledTeamJungle int32 `json:"NeutralMinionsKilledTeamJungle,omitempty"`
	PentaKills                     int32 `json:"PentaKills,omitempty"`
	PhysicalDamageDealt            int32 `json:"PhysicalDamageDealt,omitempty"`
	PhysicalDamageDealtToChampions int32 `json:"PhysicalDamageDealtToChampions,omitempty"`
	PhysicalDamageTaken            int32 `json:"PhysicalDamageTaken,omitempty"`
	PlayerScore0                   int32 `json:"PlayerScore0,omitempty"`
	PlayerScore1                   int32 `json:"PlayerScore1,omitempty"`
	PlayerScore10                  int32 `json:"PlayerScore10,omitempty"`
	PlayerScore11                  int32 `json:"PlayerScore11,omitempty"`
	PlayerScore2                   int32 `json:"PlayerScore2,omitempty"`
	PlayerScore3                   int32 `json:"PlayerScore3,omitempty"`
	PlayerScore4                   int32 `json:"PlayerScore4,omitempty"`
	PlayerScore5                   int32 `json:"PlayerScore5,omitempty"`
	PlayerScore6                   int32 `json:"PlayerScore6,omitempty"`
	PlayerScore9                   int32 `json:"PlayerScore9,omitempty"`
	QuadraKills                    int32 `json:"QuadraKills,omitempty"`
	Spell1Casts                    int32 `json:"Spell1Casts,omitempty"`
	Spell2Casts                    int32 `json:"Spell2Casts,omitempty"`
	Spell3Casts                    int32 `json:"Spell3Casts,omitempty"`
	Spell4Casts                    int32 `json:"Spell4Casts,omitempty"`
	SummonerSpell1Casts            int32 `json:"SummonerSpell1Casts,omitempty"`
	TimeCcothers                   int32 `json:"TimeCCOthers,omitempty"`
	TotalDamageDealtToChampions    int32 `json:"TotalDamageDealtToChampions,omitempty"`
	TotalMinionsKilled             int32 `json:"TotalMinionsKilled,omitempty"`
	TripleKills                    int32 `json:"TripleKills,omitempty"`
	TrueDamageDealt                int32 `json:"TrueDamageDealt,omitempty"`
	TrueDamageDealtToChampions     int32 `json:"TrueDamageDealtToChampions,omitempty"`
	TrueDamageTaken                int32 `json:"TrueDamageTaken,omitempty"`
	UnrealKills                    int32 `json:"UnrealKills,omitempty"`
	VisionScore                    int32 `json:"VisionScore,omitempty"`
	WardsKilled                    int32 `json:"WardsKilled,omitempty"`
}

// tft-match-v1.ParticipantDto
type MatchParticipantV1DTO struct {
	// Participant's companion.
	Companion MatchCompanionV1DTO `json:"companion,omitempty"`
	PUUID     string              `json:"puuid,omitempty"`
	Augments  []string            `json:"augments,omitempty"`
	// A complete list of traits for the participant's active units.
	Traits []MatchTraitV1DTO `json:"traits,omitempty"`
	// A list of active units for the participant.
	Units    []MatchUnitV1DTO              `json:"units,omitempty"`
	Missions MatchParticipantMissionsV1DTO `json:"missions,omitempty"`
	// Gold left after participant was eliminated.
	GoldLeft int32 `json:"gold_left,omitempty"`
	// The round the participant was eliminated in. Note: If the player was eliminated in stage 2-1 their last_round would be 5.
	LastRound int32 `json:"last_round,omitempty"`
	// Participant Little Legend level. Note: This is not the number of active units.
	Level          int32 `json:"level,omitempty"`
	PartnerGroupID int32 `json:"partner_group_id,omitempty"`
	// Participant placement upon elimination.
	Placement int32 `json:"placement,omitempty"`
	// Number of players the participant eliminated.
	PlayersEliminated int32 `json:"players_eliminated,omitempty"`
	// The number of seconds before the participant was eliminated.
	TimeEliminated float32 `json:"time_eliminated,omitempty"`
	// Damage the participant dealt to other players.
	TotalDamageToPlayers int32 `json:"total_damage_to_players,omitempty"`
}

// tft-match-v1.TraitDto
type MatchTraitV1DTO struct {
	// Trait name.
	Name string `json:"name,omitempty"`
	// Number of units with this trait.
	NumUnits int32 `json:"num_units,omitempty"`
	// Current style for this trait. (0 = No style, 1 = Bronze, 2 = Silver, 3 = Gold, 4 = Chromatic)
	Style int32 `json:"style,omitempty"`
	// Current active tier for the trait.
	TierCurrent int32 `json:"tier_current,omitempty"`
	// Total tiers for the trait.
	TierTotal int32 `json:"tier_total,omitempty"`
}

// tft-match-v1.UnitDto
type MatchUnitV1DTO struct {
	// This field was introduced in patch 9.22 with data_version 2.
	CharacterID string `json:"character_id,omitempty"`
	// If a unit is chosen as part of the Fates set mechanic, the chosen trait will be indicated by this field. Otherwise this field is excluded from the response.
	Chosen string `json:"chosen,omitempty"`
	// Unit name. This field is often left blank.
	Name      string   `json:"name,omitempty"`
	ItemNames []string `json:"itemNames,omitempty"`
	// A list of the unit's items. Please refer to the Teamfight Tactics documentation for item ids.
	Items []int32 `json:"items,omitempty"`
	// Unit rarity. This doesn't equate to the unit cost.
	Rarity int32 `json:"rarity,omitempty"`
	// Unit tier.
	Tier int32 `json:"tier,omitempty"`
}

// tft-match-v1.MatchDto
type MatchV1DTO struct {
	// Match metadata.
	Metadata MatchMetadataV1DTO `json:"metadata,omitempty"`
	// Match info.
	Info MatchInfoV1DTO `json:"info,omitempty"`
}

// spectator-tft-v5.BannedChampion
type SpectatorBannedChampionV5DTO struct {
	// The ID of the banned champion
	ChampionID int64 `json:"championId,omitempty"`
	// The turn during which the champion was banned
	PickTurn int32 `json:"pickTurn,omitempty"`
	// The ID of the team that banned the champion
	TeamID int64 `json:"teamId,omitempty"`
}

// spectator-tft-v5.CurrentGameInfo
type SpectatorCurrentGameInfoV5DTO struct {
	// The game mode
	GameMode GameMode `json:"gameMode,omitempty"`
	// The game type
	GameType GameType `json:"gameType,omitempty"`
	// The observer information
	Observers SpectatorObserverV5DTO `json:"observers,omitempty"`
	// The ID of the platform on which the game is being played
	PlatformID string `json:"platformId,omitempty"`
	// Banned champion information
	BannedChampions []SpectatorBannedChampionV5DTO `json:"bannedChampions,omitempty"`
	// The participant information
	Participants []SpectatorCurrentGameParticipantV5DTO `json:"participants,omitempty"`
	// The ID of the game
	GameID int64 `json:"gameId,omitempty"`
	// The amount of time in seconds that has passed since the game started
	GameLength int64 `json:"gameLength,omitempty"`
	// The queue type (queue types are documented on the Game Constants page)
	GameQueueConfigID int64 `json:"gameQueueConfigId,omitempty"`
	// The game start time represented in epoch milliseconds
	GameStartTime int64 `json:"gameStartTime,omitempty"`
	// The ID of the map
	MapID int64 `json:"mapId,omitempty"`
}

// spectator-tft-v5.CurrentGameParticipant
type SpectatorCurrentGameParticipantV5DTO struct {
	// The encrypted puuid of this participant
	PUUID  string `json:"puuid,omitempty"`
	RiotID string `json:"riotId,omitempty"`
	// The encrypted summoner ID of this participant
	SummonerID string `json:"summonerId,omitempty"`
	// List of Game Customizations
	GameCustomizationObjects []SpectatorGameCustomizationObjectV5DTO `json:"gameCustomizationObjects,omitempty"`
	// Perks/Runes Reforged Information
	Perks SpectatorPerksV5DTO `json:"perks,omitempty"`
	// The ID of the champion played by this participant
	ChampionID int64 `json:"championId,omitempty"`
	// The ID of the profile icon used by this participant
	ProfileIconID int64 `json:"profileIconId,omitempty"`
	// The ID of the first summoner spell used by this participant
	Spell1ID int64 `json:"spell1Id,omitempty"`
	// The ID of the second summoner spell used by this participant
	Spell2ID int64 `json:"spell2Id,omitempty"`
	// The team ID of this participant, indicating the participant's team
	TeamID int64 `json:"teamId,omitempty"`
}

// spectator-tft-v5.FeaturedGameInfo
type SpectatorFeaturedGameInfoV5DTO struct {
	// The game mode
	//
	// (Legal values:  TFT)
	GameMode GameMode `json:"gameMode,omitempty"`
	// The game type
	//
	// (Legal values:  MATCHED)
	GameType GameType `json:"gameType,omitempty"`
	// The observer information
	Observers SpectatorObserverV5DTO `json:"observers,omitempty"`
	// The ID of the platform on which the game is being played
	PlatformID string `json:"platformId,omitempty"`
	// Banned champion information
	BannedChampions []SpectatorBannedChampionV5DTO `json:"bannedChampions,omitempty"`
	// The participant information
	Participants []SpectatorParticipantV5DTO `json:"participants,omitempty"`
	// The ID of the game
	GameID int64 `json:"gameId,omitempty"`
	// The amount of time in seconds that has passed since the game started
	GameLength int64 `json:"gameLength,omitempty"`
	// The queue type (queue types are documented on the Game Constants page)
	GameQueueConfigID int64 `json:"gameQueueConfigId,omitempty"`
	// The ID of the map
	MapID int64 `json:"mapId,omitempty"`
}

// spectator-tft-v5.FeaturedGames
type SpectatorFeaturedGamesV5DTO struct {
	// The list of featured games
	GameList []SpectatorFeaturedGameInfoV5DTO `json:"gameList,omitempty"`
	// The suggested interval to wait before requesting FeaturedGames again
	ClientRefreshInterval int64 `json:"clientRefreshInterval,omitempty"`
}

// spectator-tft-v5.GameCustomizationObject
type SpectatorGameCustomizationObjectV5DTO struct {
	// Category identifier for Game Customization
	Category string `json:"category,omitempty"`
	// Game Customization content
	Content string `json:"content,omitempty"`
}

// spectator-tft-v5.Observer
type SpectatorObserverV5DTO struct {
	// Key used to decrypt the spectator grid game data for playback
	EncryptionKey string `json:"encryptionKey,omitempty"`
}

// spectator-tft-v5.Participant
type SpectatorParticipantV5DTO struct {
	// Encrypted puuid of this participant
	PUUID  string `json:"puuid,omitempty"`
	RiotID string `json:"riotId,omitempty"`
	// Encrypted summoner ID of this participant
	SummonerID string `json:"summonerId,omitempty"`
	// The ID of the champion played by this participant
	ChampionID int64 `json:"championId,omitempty"`
	// The ID of the profile icon used by this participant
	ProfileIconID int64 `json:"profileIconId,omitempty"`
	// The ID of the first summoner spell used by this participant
	Spell1ID int64 `json:"spell1Id,omitempty"`
	// The ID of the second summoner spell used by this participant
	Spell2ID int64 `json:"spell2Id,omitempty"`
	// The team ID of this participant, indicating the participant's team
	TeamID int64 `json:"teamId,omitempty"`
}

// spectator-tft-v5.Perks
type SpectatorPerksV5DTO struct {
	// IDs of the perks/runes assigned.
	PerkIDs []int64 `json:"perkIds,omitempty"`
	// Primary runes path
	PerkStyle int64 `json:"perkStyle,omitempty"`
	// Secondary runes path
	PerkSubStyle int64 `json:"perkSubStyle,omitempty"`
}

// tft-status-v1.ContentDto
type StatusContentV1DTO struct {
	Content string `json:"content,omitempty"`
	Locale  string `json:"locale,omitempty"`
}

// tft-status-v1.PlatformDataDto
type StatusPlatformDataV1DTO struct {
	ID           string        `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Incidents    []StatusV1DTO `json:"incidents,omitempty"`
	Locales      []string      `json:"locales,omitempty"`
	Maintenances []StatusV1DTO `json:"maintenances,omitempty"`
}

// tft-status-v1.UpdateDto
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

// tft-status-v1.StatusDto
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

// tft-summoner-v1.SummonerDTO
type SummonerV1DTO struct {
	// Encrypted account ID. Max length 56 characters.
	AccountID string `json:"accountId,omitempty"`
	// Encrypted summoner ID. Max length 63 characters.
	ID string `json:"id,omitempty"`
	// Encrypted PUUID. Exact length of 78 characters.
	PUUID string `json:"puuid,omitempty"`
	// ID of the summoner icon associated with the summoner.
	ProfileIconID int32 `json:"profileIconId,omitempty"`
	// Date summoner was last modified specified as epoch milliseconds. The following events will update this timestamp: summoner name change, summoner level change, or profile icon change.
	RevisionDate int64 `json:"revisionDate,omitempty"`
	// Summoner level associated with the summoner.
	SummonerLevel int64 `json:"summonerLevel,omitempty"`
}
