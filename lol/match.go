package lol

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type MatchEndpoint struct {
	internalClient *internal.InternalClient
}

type MatchDTO struct {
	Metadata struct {
		DataVersion  string   `json:"dataVersion"`
		MatchID      string   `json:"matchId"`
		Participants []string `json:"participants"`
	} `json:"metadata"`
	Info struct {
		GameCreation       time.Time `json:"gameCreation"`
		GameDuration       int       `json:"gameDuration"`
		GameEndTimestamp   time.Time `json:"gameEndTimestamp"`
		GameID             int       `json:"gameId"`
		GameMode           string    `json:"gameMode"`
		GameName           string    `json:"gameName"`
		GameStartTimestamp time.Time `json:"gameStartTimestamp"`
		GameType           string    `json:"gameType"`
		GameVersion        string    `json:"gameVersion"`
		MapID              int       `json:"mapId"`
		Participants       []struct {
			Assists                     int    `json:"assists"`
			BaronKills                  int    `json:"baronKills"`
			BountyLevel                 int    `json:"bountyLevel"`
			ChampExperience             int    `json:"champExperience"`
			ChampLevel                  int    `json:"champLevel"`
			ChampionID                  int    `json:"championId"`
			ChampionName                string `json:"championName"`
			ChampionTransform           int    `json:"championTransform"`
			ConsumablesPurchased        int    `json:"consumablesPurchased"`
			DamageDealtToBuildings      int    `json:"damageDealtToBuildings"`
			DamageDealtToObjectives     int    `json:"damageDealtToObjectives"`
			DamageDealtToTurrets        int    `json:"damageDealtToTurrets"`
			DamageSelfMitigated         int    `json:"damageSelfMitigated"`
			Deaths                      int    `json:"deaths"`
			DetectorWardsPlaced         int    `json:"detectorWardsPlaced"`
			DoubleKills                 int    `json:"doubleKills"`
			DragonKills                 int    `json:"dragonKills"`
			EligibleForProgression      bool   `json:"eligibleForProgression"`
			FirstBloodAssist            bool   `json:"firstBloodAssist"`
			FirstBloodKill              bool   `json:"firstBloodKill"`
			FirstTowerAssist            bool   `json:"firstTowerAssist"`
			FirstTowerKill              bool   `json:"firstTowerKill"`
			GameEndedInEarlySurrender   bool   `json:"gameEndedInEarlySurrender"`
			GameEndedInSurrender        bool   `json:"gameEndedInSurrender"`
			GoldEarned                  int    `json:"goldEarned"`
			GoldSpent                   int    `json:"goldSpent"`
			IndividualPosition          string `json:"individualPosition"`
			InhibitorKills              int    `json:"inhibitorKills"`
			InhibitorTakedowns          int    `json:"inhibitorTakedowns"`
			InhibitorsLost              int    `json:"inhibitorsLost"`
			Item0                       int    `json:"item0"`
			Item1                       int    `json:"item1"`
			Item2                       int    `json:"item2"`
			Item3                       int    `json:"item3"`
			Item4                       int    `json:"item4"`
			Item5                       int    `json:"item5"`
			Item6                       int    `json:"item6"`
			ItemsPurchased              int    `json:"itemsPurchased"`
			KillingSprees               int    `json:"killingSprees"`
			Kills                       int    `json:"kills"`
			Lane                        string `json:"lane"`
			LargestCriticalStrike       int    `json:"largestCriticalStrike"`
			LargestKillingSpree         int    `json:"largestKillingSpree"`
			LargestMultiKill            int    `json:"largestMultiKill"`
			LongestTimeSpentLiving      int    `json:"longestTimeSpentLiving"`
			MagicDamageDealt            int    `json:"magicDamageDealt"`
			MagicDamageDealtToChampions int    `json:"magicDamageDealtToChampions"`
			MagicDamageTaken            int    `json:"magicDamageTaken"`
			NeutralMinionsKilled        int    `json:"neutralMinionsKilled"`
			NexusKills                  int    `json:"nexusKills"`
			NexusLost                   int    `json:"nexusLost"`
			NexusTakedowns              int    `json:"nexusTakedowns"`
			ObjectivesStolen            int    `json:"objectivesStolen"`
			ObjectivesStolenAssists     int    `json:"objectivesStolenAssists"`
			ParticipantID               int    `json:"participantId"`
			PentaKills                  int    `json:"pentaKills"`
			Perks                       struct {
				StatPerks struct {
					Defense int `json:"defense"`
					Flex    int `json:"flex"`
					Offense int `json:"offense"`
				} `json:"statPerks"`
				Styles []struct {
					Description string `json:"description"`
					Selections  []struct {
						Perk int `json:"perk"`
						Var1 int `json:"var1"`
						Var2 int `json:"var2"`
						Var3 int `json:"var3"`
					} `json:"selections"`
					Style int `json:"style"`
				} `json:"styles"`
			} `json:"perks"`
			PhysicalDamageDealt            int    `json:"physicalDamageDealt"`
			PhysicalDamageDealtToChampions int    `json:"physicalDamageDealtToChampions"`
			PhysicalDamageTaken            int    `json:"physicalDamageTaken"`
			ProfileIcon                    int    `json:"profileIcon"`
			Puuid                          string `json:"puuid"`
			QuadraKills                    int    `json:"quadraKills"`
			RiotIDName                     string `json:"riotIdName"`
			RiotIDTagline                  string `json:"riotIdTagline"`
			Role                           string `json:"role"`
			SightWardsBoughtInGame         int    `json:"sightWardsBoughtInGame"`
			Spell1Casts                    int    `json:"spell1Casts"`
			Spell2Casts                    int    `json:"spell2Casts"`
			Spell3Casts                    int    `json:"spell3Casts"`
			Spell4Casts                    int    `json:"spell4Casts"`
			Summoner1Casts                 int    `json:"summoner1Casts"`
			Summoner1ID                    int    `json:"summoner1Id"`
			Summoner2Casts                 int    `json:"summoner2Casts"`
			Summoner2ID                    int    `json:"summoner2Id"`
			SummonerID                     string `json:"summonerId"`
			SummonerLevel                  int    `json:"summonerLevel"`
			SummonerName                   string `json:"summonerName"`
			TeamEarlySurrendered           bool   `json:"teamEarlySurrendered"`
			TeamID                         int    `json:"teamId"`
			TeamPosition                   string `json:"teamPosition"`
			TimeCCingOthers                int    `json:"timeCCingOthers"`
			TimePlayed                     int    `json:"timePlayed"`
			TotalDamageDealt               int    `json:"totalDamageDealt"`
			TotalDamageDealtToChampions    int    `json:"totalDamageDealtToChampions"`
			TotalDamageShieldedOnTeammates int    `json:"totalDamageShieldedOnTeammates"`
			TotalDamageTaken               int    `json:"totalDamageTaken"`
			TotalHeal                      int    `json:"totalHeal"`
			TotalHealsOnTeammates          int    `json:"totalHealsOnTeammates"`
			TotalMinionsKilled             int    `json:"totalMinionsKilled"`
			TotalTimeCCDealt               int    `json:"totalTimeCCDealt"`
			TotalTimeSpentDead             int    `json:"totalTimeSpentDead"`
			TotalUnitsHealed               int    `json:"totalUnitsHealed"`
			TripleKills                    int    `json:"tripleKills"`
			TrueDamageDealt                int    `json:"trueDamageDealt"`
			TrueDamageDealtToChampions     int    `json:"trueDamageDealtToChampions"`
			TrueDamageTaken                int    `json:"trueDamageTaken"`
			TurretKills                    int    `json:"turretKills"`
			TurretTakedowns                int    `json:"turretTakedowns"`
			TurretsLost                    int    `json:"turretsLost"`
			UnrealKills                    int    `json:"unrealKills"`
			VisionScore                    int    `json:"visionScore"`
			VisionWardsBoughtInGame        int    `json:"visionWardsBoughtInGame"`
			WardsKilled                    int    `json:"wardsKilled"`
			WardsPlaced                    int    `json:"wardsPlaced"`
			Win                            bool   `json:"win"`
		} `json:"participants"`
		PlatformID string `json:"platformId"`
		QueueID    int    `json:"queueId"`
		Teams      []struct {
			Bans       []interface{} `json:"bans"`
			Objectives struct {
				Baron struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"baron"`
				Champion struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"champion"`
				Dragon struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"dragon"`
				Inhibitor struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"inhibitor"`
				RiftHerald struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"riftHerald"`
				Tower struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"tower"`
			} `json:"objectives"`
			TeamID int  `json:"teamId"`
			Win    bool `json:"win"`
		} `json:"teams"`
		TournamentCode string `json:"tournamentCode"`
	} `json:"info"`
}

type MatchTimelineDTO struct {
	Metadata struct {
		DataVersion  string   `json:"dataVersion"`
		MatchID      string   `json:"matchId"`
		Participants []string `json:"participants"`
	} `json:"metadata"`
	Info struct {
		FrameInterval int `json:"frameInterval"`
		Frames        []struct {
			Events []struct {
				RealTimestamp time.Time `json:"realTimestamp"`
				Timestamp     int       `json:"timestamp"`
				Type          string    `json:"type"`
			} `json:"events"`
			ParticipantFrames struct {
				Num1 struct {
					ChampionStats struct {
						AbilityHaste         int `json:"abilityHaste"`
						AbilityPower         int `json:"abilityPower"`
						Armor                int `json:"armor"`
						ArmorPen             int `json:"armorPen"`
						ArmorPenPercent      int `json:"armorPenPercent"`
						AttackDamage         int `json:"attackDamage"`
						AttackSpeed          int `json:"attackSpeed"`
						BonusArmorPenPercent int `json:"bonusArmorPenPercent"`
						BonusMagicPenPercent int `json:"bonusMagicPenPercent"`
						CcReduction          int `json:"ccReduction"`
						CooldownReduction    int `json:"cooldownReduction"`
						Health               int `json:"health"`
						HealthMax            int `json:"healthMax"`
						HealthRegen          int `json:"healthRegen"`
						Lifesteal            int `json:"lifesteal"`
						MagicPen             int `json:"magicPen"`
						MagicPenPercent      int `json:"magicPenPercent"`
						MagicResist          int `json:"magicResist"`
						MovementSpeed        int `json:"movementSpeed"`
						Omnivamp             int `json:"omnivamp"`
						PhysicalVamp         int `json:"physicalVamp"`
						Power                int `json:"power"`
						PowerMax             int `json:"powerMax"`
						PowerRegen           int `json:"powerRegen"`
						SpellVamp            int `json:"spellVamp"`
					} `json:"championStats"`
					CurrentGold int `json:"currentGold"`
					DamageStats struct {
						MagicDamageDone               int `json:"magicDamageDone"`
						MagicDamageDoneToChampions    int `json:"magicDamageDoneToChampions"`
						MagicDamageTaken              int `json:"magicDamageTaken"`
						PhysicalDamageDone            int `json:"physicalDamageDone"`
						PhysicalDamageDoneToChampions int `json:"physicalDamageDoneToChampions"`
						PhysicalDamageTaken           int `json:"physicalDamageTaken"`
						TotalDamageDone               int `json:"totalDamageDone"`
						TotalDamageDoneToChampions    int `json:"totalDamageDoneToChampions"`
						TotalDamageTaken              int `json:"totalDamageTaken"`
						TrueDamageDone                int `json:"trueDamageDone"`
						TrueDamageDoneToChampions     int `json:"trueDamageDoneToChampions"`
						TrueDamageTaken               int `json:"trueDamageTaken"`
					} `json:"damageStats"`
					GoldPerSecond       int `json:"goldPerSecond"`
					JungleMinionsKilled int `json:"jungleMinionsKilled"`
					Level               int `json:"level"`
					MinionsKilled       int `json:"minionsKilled"`
					ParticipantID       int `json:"participantId"`
					Position            struct {
						X int `json:"x"`
						Y int `json:"y"`
					} `json:"position"`
					TimeEnemySpentControlled int `json:"timeEnemySpentControlled"`
					TotalGold                int `json:"totalGold"`
					Xp                       int `json:"xp"`
				} `json:"1"`
			} `json:"participantFrames"`
			Timestamp int `json:"timestamp"`
		} `json:"frames"`
		GameID       int `json:"gameId"`
		Participants []struct {
			ParticipantID int    `json:"participantId"`
			PUUID         string `json:"puuid"`
		} `json:"participants"`
	} `json:"info"`
}

type LOLMatchType string

const (
	RankedMatchType   LOLMatchType = "ranked"
	NormalMatchType   LOLMatchType = "normal"
	TourneyMatchType  LOLMatchType = "tourney"
	TutorialMatchType LOLMatchType = "tutorial"
)

type MatchlistOptions struct {
	// The matchlist started storing timestamps on June 16th, 2021.
	// Any matches played before June 16th, 2021 won't be included in the results if the startTime filter is set.
	StartTime int `json:"startTime"`
	// Epoch timestamp in seconds.
	EndTime int `json:"endTime"`
	// Filter the list of match ids by a specific queue id.
	// This filter is mutually inclusive of the type filter meaning any match ids returned must match both the queue and type filters.
	Queue int `json:"queue"`
	// Filter the list of match ids by the type of match.
	// This filter is mutually inclusive of the queue filter meaning any match ids returned must match both the queue and type filters.
	Type LOLMatchType `json:"type"`
	// Defaults to 0. Start index.
	Start int `json:"start"`
	// Defaults to 20. Valid values: 0 to 100. Number of match ids to return.
	Count int `json:"count"`
}

// Get a list of match IDs by PUUID.
func (c *MatchEndpoint) ListByPUUID(region api.Route, PUUID string, options *MatchlistOptions) ([]string, error) {
	if options == nil {
		options = &MatchlistOptions{Start: 0, Count: 20}
	}

	if options.Count > 100 || options.Count < 1 {
		options.Count = 20
	}

	query := url.Values{}

	if options.StartTime != 0 {
		query.Set("startTime", strconv.Itoa(options.StartTime))
	}

	if options.EndTime != 0 {
		query.Set("endTime", strconv.Itoa(options.EndTime))
	}

	if options.Queue != 0 {
		query.Set("queue", strconv.Itoa(options.Queue))
	}

	if options.Type != "" {
		query.Set("queue", string(options.Type))
	}

	query.Set("start", strconv.Itoa(options.Start))

	query.Set("count", strconv.Itoa(options.Count))

	endpoint := fmt.Sprintf(MatchlistURL, PUUID)

	url := fmt.Sprintf("%s?%s", endpoint, query.Encode())

	res := []string{}

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get a match by match ID.
func (c *MatchEndpoint) ByID(region api.Route, matchID string) (*MatchDTO, error) {
	url := fmt.Sprintf(MatchURL, matchID)

	res := MatchDTO{}

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get a match timeline by match ID.
func (c *MatchEndpoint) Timeline(region api.Route, matchID string) (*MatchTimelineDTO, error) {
	url := fmt.Sprintf(MatchTimelineURL, matchID)

	res := MatchTimelineDTO{}

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
