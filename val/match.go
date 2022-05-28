package val

import (
	"fmt"

	"github.com/Kyagara/equinox/internal"
)

type MatchEndpoint struct {
	internalClient *internal.InternalClient
}

type MatchListDTO struct {
	PUUID   string       `json:"puuid"`
	History []HistoryDTO `json:"history"`
}

type HistoryDTO struct {
	MatchID             string `json:"matchId"`
	GameStartTimeMillis int    `json:"gameStartTimeMillis"`
	TeamID              string `json:"teamId"`
}

type RecentMatchesDTO struct {
	CurrentTime int      `json:"currentTime"`
	MatchIDs    []string `json:"matchIds"`
}

type MatchDTO struct {
	MatchInfo    MatchInfoDTO     `json:"matchInfo"`
	Players      []MatchPlayerDTO `json:"players"`
	Coaches      []CoachDTO       `json:"coaches"`
	Teams        []TeamDTO        `json:"teams"`
	RoundResults []RoundResultDTO `json:"roundResults"`
}

type MatchInfoDTO struct {
	MatchID            string `json:"matchId"`
	MapID              string `json:"mapId"`
	GameLengthMillis   int    `json:"gameLengthMillis"`
	GameStartMillis    int    `json:"gameStartMillis"`
	ProvisioningFlowID string `json:"provisioningFlowId"`
	IsCompleted        bool   `json:"isCompleted"`
	CustomGameName     string `json:"customGameName"`
	QueueID            Queue  `json:"queueId"`
	GameMode           string `json:"gameMode"`
	IsRanked           bool   `json:"isRanked"`
	SeasonID           string `json:"seasonId"`
}

type AbilityCastDTO struct {
	GrenadeCasts  int `json:"grenadeCasts"`
	Ability1Casts int `json:"ability1Casts"`
	Ability2Casts int `json:"ability2Casts"`
	UltimateCasts int `json:"ultimateCasts"`
}

type StatsDTO struct {
	Score          int            `json:"score"`
	RoundsPlayed   int            `json:"roundsPlayed"`
	Kills          int            `json:"kills"`
	Deaths         int            `json:"deaths"`
	Assists        int            `json:"assists"`
	PlaytimeMillis int            `json:"playtimeMillis"`
	AbilityCasts   AbilityCastDTO `json:"abilityCasts"`
}

type MatchPlayerDTO struct {
	PUUID           string   `json:"puuid"`
	GameName        string   `json:"gameName"`
	TagLine         string   `json:"tagLine"`
	TeamID          string   `json:"teamId"`
	PartyID         string   `json:"partyId"`
	CharacterID     string   `json:"characterId"`
	Stats           StatsDTO `json:"stats"`
	CompetitiveTier int      `json:"competitiveTier"`
	PlayerCard      string   `json:"playerCard"`
	PlayerTitle     string   `json:"playerTitle"`
}

type CoachDTO struct {
	PUUID  string `json:"puuid"`
	TeamID string `json:"teamId"`
}

type TeamDTO struct {
	TeamID       string `json:"teamId"`
	Won          bool   `json:"won"`
	RoundsPlayed int    `json:"roundsPlayed"`
	RoundsWon    int    `json:"roundsWon"`
	NumPoints    int    `json:"numPoints"`
}

type LocationDTO struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PlayerLocationDTO struct {
	PUUID       string      `json:"puuid"`
	ViewRadians int         `json:"viewRadians"`
	Location    LocationDTO `json:"location"`
}

type FinishingDamageDTO struct {
	DamageType          string `json:"damageType"`
	DamageItem          string `json:"damageItem"`
	IsSecondaryFireMode bool   `json:"isSecondaryFireMode"`
}

type KillDTO struct {
	TimeSinceGameStartMillis  int                 `json:"timeSinceGameStartMillis"`
	TimeSinceRoundStartMillis int                 `json:"timeSinceRoundStartMillis"`
	Killer                    string              `json:"killer"`
	Victim                    string              `json:"victim"`
	VictimLocation            LocationDTO         `json:"victimLocation"`
	Assistants                []string            `json:"assistants"`
	PlayerLocations           []PlayerLocationDTO `json:"playerLocations"`
	FinishingDamage           FinishingDamageDTO  `json:"finishingDamage"`
}

type DamageDTO struct {
	Receiver  string `json:"receiver"`
	Damage    int    `json:"damage"`
	Legshots  int    `json:"legshots"`
	Bodyshots int    `json:"bodyshots"`
	Headshots int    `json:"headshots"`
}

type EconomyDTO struct {
	LoadoutValue int    `json:"loadoutValue"`
	Weapon       string `json:"weapon"`
	Armor        string `json:"armor"`
	Remaining    int    `json:"remaining"`
	Spent        int    `json:"spent"`
}

type AbilityDTO struct {
	GrenadeEffects  string `json:"grenadeEffects"`
	Ability1Effects string `json:"ability1Effects"`
	Ability2Effects string `json:"ability2Effects"`
	UltimateEffects string `json:"ultimateEffects"`
}

type PlayerStatsDTO struct {
	PUUID   string      `json:"puuid"`
	Kills   []KillDTO   `json:"kills"`
	Damage  []DamageDTO `json:"damage"`
	Score   int         `json:"score"`
	Economy EconomyDTO  `json:"economy"`
	Ability AbilityDTO  `json:"ability"`
}

type RoundResultDTO struct {
	RoundNum              int                 `json:"roundNum"`
	RoundResult           string              `json:"roundResult"`
	RoundCeremony         string              `json:"roundCeremony"`
	WinningTeam           string              `json:"winningTeam"`
	BombPlanter           string              `json:"bombPlanter"`
	BombDefuser           string              `json:"bombDefuser"`
	PlantRoundTime        int                 `json:"plantRoundTime"`
	PlantPlayerLocations  []PlayerLocationDTO `json:"plantPlayerLocations"`
	PlantLocation         LocationDTO         `json:"plantLocation"`
	PlantSite             string              `json:"plantSite"`
	DefuseRoundTime       int                 `json:"defuseRoundTime"`
	DefusePlayerLocations []PlayerLocationDTO `json:"defusePlayerLocations"`
	DefuseLocation        LocationDTO         `json:"defuseLocation"`
	PlayerStats           []PlayerStatsDTO    `json:"playerStats"`
	RoundResultCode       string              `json:"roundResultCode"`
}

// Get matchlist for games played by PUUID.
func (e *MatchEndpoint) List(shard Shard, puuid string) (*MatchListDTO, error) {
	logger := e.internalClient.Logger("VAL", "match", "List")

	url := fmt.Sprintf(MatchListURL, puuid)

	var list *MatchListDTO

	err := e.internalClient.Get(shard, url, &list, "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return list, nil
}

// Get match by ID.
func (e *MatchEndpoint) ByID(shard Shard, matchID string) (*MatchDTO, error) {
	logger := e.internalClient.Logger("VAL", "match", "ByID")

	url := fmt.Sprintf(MatchByIDURL, matchID)

	var match *MatchDTO

	err := e.internalClient.Get(shard, url, &match, "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return match, nil
}

// Get recent matches.
//
// Returns a list of match ids that have completed in the last 10 minutes for live regions and 12 hours for the esports routing value.
//
// NA/LATAM/BR share a match history deployment. As such, recent matches will return a combined list of matches from those three regions.
//
// Requests are load balanced so you may see some inconsistencies as matches are added/removed from the list.
func (e *MatchEndpoint) Recent(shard Shard, queue Queue) (*RecentMatchesDTO, error) {
	logger := e.internalClient.Logger("VAL", "match", "Recent")

	url := fmt.Sprintf(MatchRecentURL, queue)

	var recent *RecentMatchesDTO

	err := e.internalClient.Get(shard, url, &recent, "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return recent, nil
}
