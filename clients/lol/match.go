package lol

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/Kyagara/equinox/internal"
)

type MatchEndpoint struct {
	internalClient *internal.InternalClient
}

type MatchDTO struct {
	// Match metadata.
	Metadata struct {
		// Match data version.
		DataVersion string `json:"dataVersion"`
		// Match ID.
		MatchID string `json:"matchId"`
		// A list of participant PUUIDs.
		Participants []string `json:"participants"`
	} `json:"metadata"`
	// Match info.
	Info struct {
		// Unix timestamp for when the game is created on the game server (i.e., the loading screen).
		GameCreation int64 `json:"gameCreation"`
		// Prior to patch 11.20, this field returns the game length in milliseconds calculated from gameEndTimestamp - gameStartTimestamp. Post patch 11.20, this field returns the max timePlayed of any participant in the game in seconds, which makes the behavior of this field consistent with that of match-v4. The best way to handling the change in this field is to treat the value as milliseconds if the gameEndTimestamp field isn't in the response and to treat the value as seconds if gameEndTimestamp is in the response.
		GameDuration int `json:"gameDuration"`
		// Unix timestamp for when match ends on the game server. This timestamp can occasionally be significantly longer than when the match "ends". The most reliable way of determining the timestamp for the end of the match would be to add the max time played of any participant to the gameStartTimestamp. This field was added to match-v5 in patch 11.20 on Oct 5th, 2021.
		GameEndTimestamp int64 `json:"gameEndTimestamp"`
		GameID           int   `json:"gameId"`
		// Refer to the Game Constants documentation.
		GameMode string `json:"gameMode"`
		GameName string `json:"gameName"`
		// Unix timestamp for when match starts on the game server.
		GameStartTimestamp int64  `json:"gameStartTimestamp"`
		GameType           string `json:"gameType"`
		// The first two parts can be used to determine the patch a game was played on.
		GameVersion string `json:"gameVersion"`
		// Refer to the Game Constants documentation.
		MapID        int                `json:"mapId"`
		Participants []MatchParticipant `json:"participants"`
		// Platform where the match was played.
		PlatformID Region `json:"platformId"`
		// Refer to the Game Constants documentation.
		QueueID int `json:"queueId"`
		Teams   []struct {
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
		// Tournament code used to generate the match. This field was added to match-v5 in patch 11.13 on June 23rd, 2021.
		TournamentCode string `json:"tournamentCode"`
	} `json:"info"`
}

type MatchParticipant struct {
	Assists     int `json:"assists"`
	BaronKills  int `json:"baronKills"`
	BountyLevel int `json:"bountyLevel"`
	Challenges  struct {
		One2AssistStreakCount                     int `json:"12AssistStreakCount"`
		AbilityUses                               int `json:"abilityUses"`
		AcesBefore15Minutes                       int `json:"acesBefore15Minutes"`
		AlliedJungleMonsterKills                  int `json:"alliedJungleMonsterKills"`
		BaronBuffGoldAdvantageOverThreshold       int `json:"baronBuffGoldAdvantageOverThreshold"`
		BaronTakedowns                            int `json:"baronTakedowns"`
		BlastConeOppositeOpponentCount            int `json:"blastConeOppositeOpponentCount"`
		BountyGold                                int `json:"bountyGold"`
		BuffsStolen                               int `json:"buffsStolen"`
		CompleteSupportQuestInTime                int `json:"completeSupportQuestInTime"`
		ControlWardsPlaced                        int `json:"controlWardsPlaced"`
		ControlWardTimeCoverageInRiverOrEnemyHalf int `json:"controlWardTimeCoverageInRiverOrEnemyHalf"`
		DamagePerMinute                           int `json:"damagePerMinute"`
		DamageTakenOnTeamPercentage               int `json:"damageTakenOnTeamPercentage"`
		DancedWithRiftHerald                      int `json:"dancedWithRiftHerald"`
		DeathsByEnemyChamps                       int `json:"deathsByEnemyChamps"`
		DodgeSkillShotsSmallWindow                int `json:"dodgeSkillShotsSmallWindow"`
		DoubleAces                                int `json:"doubleAces"`
		DragonTakedowns                           int `json:"dragonTakedowns"`
		EarliestBaron                             int `json:"earliestBaron"`
		EarliestDragonTakedown                    int `json:"earliestDragonTakedown"`
		EarliestElderDragon                       int `json:"earliestElderDragon"`
		EarlyLaningPhaseGoldExpAdvantage          int `json:"earlyLaningPhaseGoldExpAdvantage"`
		EffectiveHealAndShielding                 int `json:"effectiveHealAndShielding"`
		ElderDragonKillsWithOpposingSoul          int `json:"elderDragonKillsWithOpposingSoul"`
		ElderDragonMultikills                     int `json:"elderDragonMultikills"`
		EnemyChampionImmobilizations              int `json:"enemyChampionImmobilizations"`
		EnemyJungleMonsterKills                   int `json:"enemyJungleMonsterKills"`
		EpicMonsterKillsNearEnemyJungler          int `json:"epicMonsterKillsNearEnemyJungler"`
		EpicMonsterKillsWithin30SecondsOfSpawn    int `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
		EpicMonsterSteals                         int `json:"epicMonsterSteals"`
		EpicMonsterStolenWithoutSmite             int `json:"epicMonsterStolenWithoutSmite"`
		FasterSupportQuestCompletion              int `json:"fasterSupportQuestCompletion"`
		FastestLegendary                          int `json:"fastestLegendary"`
		FirstTurretKilledTime                     int `json:"firstTurretKilledTime"`
		FlawlessAces                              int `json:"flawlessAces"`
		FullTeamTakedown                          int `json:"fullTeamTakedown"`
		GameLength                                int `json:"gameLength"`
		GetTakedownsInAllLanesEarlyJungleAsLaner  int `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
		GoldPerMinute                             int `json:"goldPerMinute"`
		HadAfkTeammate                            int `json:"hadAfkTeammate"`
		HadOpenNexus                              int `json:"hadOpenNexus"`
		HighestChampionDamage                     int `json:"highestChampionDamage"`
		HighestCrowdControlScore                  int `json:"highestCrowdControlScore"`
		HighestWardKills                          int `json:"highestWardKills"`
		ImmobilizeAndKillWithAlly                 int `json:"immobilizeAndKillWithAlly"`
		InitialBuffCount                          int `json:"initialBuffCount"`
		InitialCrabCount                          int `json:"initialCrabCount"`
		JungleCsBefore10Minutes                   int `json:"jungleCsBefore10Minutes"`
		JunglerKillsEarlyJungle                   int `json:"junglerKillsEarlyJungle"`
		JunglerTakedownsNearDamagedEpicMonster    int `json:"junglerTakedownsNearDamagedEpicMonster"`
		Kda                                       int `json:"kda"`
		KillAfterHiddenWithAlly                   int `json:"killAfterHiddenWithAlly"`
		KilledChampTookFullTeamDamageSurvived     int `json:"killedChampTookFullTeamDamageSurvived"`
		KillParticipation                         int `json:"killParticipation"`
		KillsNearEnemyTurret                      int `json:"killsNearEnemyTurret"`
		KillsOnLanersEarlyJungleAsJungler         int `json:"killsOnLanersEarlyJungleAsJungler"`
		KillsOnOtherLanesEarlyJungleAsLaner       int `json:"killsOnOtherLanesEarlyJungleAsLaner"`
		KillsOnRecentlyHealedByAramPack           int `json:"killsOnRecentlyHealedByAramPack"`
		KillsUnderOwnTurret                       int `json:"killsUnderOwnTurret"`
		KillsWithHelpFromEpicMonster              int `json:"killsWithHelpFromEpicMonster"`
		KnockEnemyIntoTeamAndKill                 int `json:"knockEnemyIntoTeamAndKill"`
		KTurretsDestroyedBeforePlatesFall         int `json:"kTurretsDestroyedBeforePlatesFall"`
		LandSkillShotsEarlyGame                   int `json:"landSkillShotsEarlyGame"`
		LaneMinionsFirst10Minutes                 int `json:"laneMinionsFirst10Minutes"`
		LaningPhaseGoldExpAdvantage               int `json:"laningPhaseGoldExpAdvantage"`
		LegendaryCount                            int `json:"legendaryCount"`
		LostAnInhibitor                           int `json:"lostAnInhibitor"`
		MaxCsAdvantageOnLaneOpponent              int `json:"maxCsAdvantageOnLaneOpponent"`
		MaxKillDeficit                            int `json:"maxKillDeficit"`
		MaxLevelLeadLaneOpponent                  int `json:"maxLevelLeadLaneOpponent"`
		MejaisFullStackInTime                     int `json:"mejaisFullStackInTime"`
		MoreEnemyJungleThanOpponent               int `json:"moreEnemyJungleThanOpponent"`
		MostWardsDestroyedOneSweeper              int `json:"mostWardsDestroyedOneSweeper"`
		MultiKillOneSpell                         int `json:"multiKillOneSpell"`
		Multikills                                int `json:"multikills"`
		MultikillsAfterAggressiveFlash            int `json:"multikillsAfterAggressiveFlash"`
		MultiTurretRiftHeraldCount                int `json:"multiTurretRiftHeraldCount"`
		MythicItemUsed                            int `json:"mythicItemUsed"`
		OuterTurretExecutesBefore10Minutes        int `json:"outerTurretExecutesBefore10Minutes"`
		OutnumberedKills                          int `json:"outnumberedKills"`
		OutnumberedNexusKill                      int `json:"outnumberedNexusKill"`
		PerfectDragonSoulsTaken                   int `json:"perfectDragonSoulsTaken"`
		PerfectGame                               int `json:"perfectGame"`
		PickKillWithAlly                          int `json:"pickKillWithAlly"`
		PoroExplosions                            int `json:"poroExplosions"`
		QuickCleanse                              int `json:"quickCleanse"`
		QuickFirstTurret                          int `json:"quickFirstTurret"`
		QuickSoloKills                            int `json:"quickSoloKills"`
		RiftHeraldTakedowns                       int `json:"riftHeraldTakedowns"`
		SaveAllyFromDeath                         int `json:"saveAllyFromDeath"`
		ScuttleCrabKills                          int `json:"scuttleCrabKills"`
		ShortestTimeToAceFromFirstTakedown        int `json:"shortestTimeToAceFromFirstTakedown"`
		SkillshotsDodged                          int `json:"skillshotsDodged"`
		SkillshotsHit                             int `json:"skillshotsHit"`
		SnowballsHit                              int `json:"snowballsHit"`
		SoloBaronKills                            int `json:"soloBaronKills"`
		SoloKills                                 int `json:"soloKills"`
		SoloTurretsLategame                       int `json:"soloTurretsLategame"`
		StealthWardsPlaced                        int `json:"stealthWardsPlaced"`
		SurvivedSingleDigitHpCount                int `json:"survivedSingleDigitHpCount"`
		SurvivedThreeImmobilizesInFight           int `json:"survivedThreeImmobilizesInFight"`
		TakedownOnFirstTurret                     int `json:"takedownOnFirstTurret"`
		Takedowns                                 int `json:"takedowns"`
		TakedownsAfterGainingLevelAdvantage       int `json:"takedownsAfterGainingLevelAdvantage"`
		TakedownsBeforeJungleMinionSpawn          int `json:"takedownsBeforeJungleMinionSpawn"`
		TakedownsFirst25Minutes                   int `json:"takedownsFirst25Minutes"`
		TakedownsInAlcove                         int `json:"takedownsInAlcove"`
		TakedownsInEnemyFountain                  int `json:"takedownsInEnemyFountain"`
		TeamBaronKills                            int `json:"teamBaronKills"`
		TeamDamagePercentage                      int `json:"teamDamagePercentage"`
		TeamElderDragonKills                      int `json:"teamElderDragonKills"`
		TeamRiftHeraldKills                       int `json:"teamRiftHeraldKills"`
		TeleportTakedowns                         int `json:"teleportTakedowns"`
		ThirdInhibitorDestroyedTime               int `json:"thirdInhibitorDestroyedTime"`
		ThreeWardsOneSweeperCount                 int `json:"threeWardsOneSweeperCount"`
		TookLargeDamageSurvived                   int `json:"tookLargeDamageSurvived"`
		TurretPlatesTaken                         int `json:"turretPlatesTaken"`
		TurretsTakenWithRiftHerald                int `json:"turretsTakenWithRiftHerald"`
		TurretTakedowns                           int `json:"turretTakedowns"`
		TwentyMinionsIn3SecondsCount              int `json:"twentyMinionsIn3SecondsCount"`
		UnseenRecalls                             int `json:"unseenRecalls"`
		VisionScoreAdvantageLaneOpponent          int `json:"visionScoreAdvantageLaneOpponent"`
		VisionScorePerMinute                      int `json:"visionScorePerMinute"`
		WardsGuarded                              int `json:"wardsGuarded"`
		WardTakedowns                             int `json:"wardTakedowns"`
		WardTakedownsBefore20M                    int `json:"wardTakedownsBefore20M"`
	} `json:"challenges,omitempty"`
	ChampExperience int `json:"champExperience"`
	ChampLevel      int `json:"champLevel"`
	// Prior to patch 11.4, on Feb 18th, 2021, this field returned invalid championIds. We recommend determining the champion based on the championName field for matches played prior to patch 11.4.
	ChampionID   int    `json:"championId"`
	ChampionName string `json:"championName"`
	// This field is currently only utilized for Kayn's transformations.
	ChampionTransform         ChampionTransformation `json:"championTransform"`
	ConsumablesPurchased      int                    `json:"consumablesPurchased"`
	DamageDealtToBuildings    int                    `json:"damageDealtToBuildings"`
	DamageDealtToObjectives   int                    `json:"damageDealtToObjectives"`
	DamageDealtToTurrets      int                    `json:"damageDealtToTurrets"`
	DamageSelfMitigated       int                    `json:"damageSelfMitigated"`
	Deaths                    int                    `json:"deaths"`
	DetectorWardsPlaced       int                    `json:"detectorWardsPlaced"`
	DoubleKills               int                    `json:"doubleKills"`
	DragonKills               int                    `json:"dragonKills"`
	EligibleForProgression    bool                   `json:"eligibleForProgression"`
	FirstBloodAssist          bool                   `json:"firstBloodAssist"`
	FirstBloodKill            bool                   `json:"firstBloodKill"`
	FirstTowerAssist          bool                   `json:"firstTowerAssist"`
	FirstTowerKill            bool                   `json:"firstTowerKill"`
	GameEndedInEarlySurrender bool                   `json:"gameEndedInEarlySurrender"`
	GameEndedInSurrender      bool                   `json:"gameEndedInSurrender"`
	GoldEarned                int                    `json:"goldEarned"`
	GoldSpent                 int                    `json:"goldSpent"`
	// Both individualPosition and teamPosition are computed by the game server and are different versions of the most likely position played by a player. The individualPosition is the best guess for which position the player actually played in isolation of anything else. The teamPosition is the best guess for which position the player actually played if we add the constraint that each team must have one top player, one jungle, one middle, etc. Generally the recommendation is to use the teamPosition field over the individualPosition field.
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
	PUUID                          string `json:"puuid"`
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
	// Both individualPosition and teamPosition are computed by the game server and are different versions of the most likely position played by a player. The individualPosition is the best guess for which position the player actually played in isolation of anything else. The teamPosition is the best guess for which position the player actually played if we add the constraint that each team must have one top player, one jungle, one middle, etc. Generally the recommendation is to use the teamPosition field over the individualPosition field.
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
}

type MatchTimelineDTO struct {
	// Match metadata.
	Metadata struct {
		// Match data version.
		DataVersion string `json:"dataVersion"`
		// Match ID.
		MatchID string `json:"matchId"`
		// A list of participant PUUIDs.
		Participants []string `json:"participants"`
	} `json:"metadata"`
	// Match info.
	Info struct {
		FrameInterval int `json:"frameInterval"`
		Frames        []struct {
			Events []struct {
				RealTimestamp           int    `json:"realTimestamp"`
				Timestamp               int    `json:"timestamp"`
				Type                    string `json:"type"`
				ItemID                  int    `json:"itemId"`
				ParticipantID           int    `json:"participantId"`
				LevelUpType             string `json:"levelUpType"`
				SkillSlot               int    `json:"skillSlot"`
				CreatorID               int    `json:"creatorId"`
				WardType                string `json:"wardType"`
				Level                   int    `json:"level"`
				AssistingParticipantIds []int  `json:"assistingParticipantIds"`
				Bounty                  int    `json:"bounty"`
				KillStreakLength        int    `json:"killStreakLength"`
				KillerID                int    `json:"killerId"`
				Position                struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"position"`
				VictimDamageDealt []struct {
					Basic          bool   `json:"basic"`
					MagicDamage    int    `json:"magicDamage"`
					Name           string `json:"name"`
					ParticipantID  int    `json:"participantId"`
					PhysicalDamage int    `json:"physicalDamage"`
					SpellName      string `json:"spellName"`
					SpellSlot      int    `json:"spellSlot"`
					TrueDamage     int    `json:"trueDamage"`
					Type           string `json:"type"`
				} `json:"victimDamageDealt"`
				VictimDamageReceived []struct {
					Basic          bool   `json:"basic"`
					MagicDamage    int    `json:"magicDamage"`
					Name           string `json:"name"`
					ParticipantID  int    `json:"participantId"`
					PhysicalDamage int    `json:"physicalDamage"`
					SpellName      string `json:"spellName"`
					SpellSlot      int    `json:"spellSlot"`
					TrueDamage     int    `json:"trueDamage"`
					Type           string `json:"type"`
				} `json:"victimDamageReceived"`
				VictimID        int    `json:"victimId"`
				KillType        string `json:"killType"`
				LaneType        string `json:"laneType"`
				TeamID          int    `json:"teamId"`
				MultiKillLength int    `json:"multiKillLength"`
				KillerTeamID    int    `json:"killerTeamId"`
				MonsterType     string `json:"monsterType"`
				MonsterSubType  string `json:"monsterSubType"`
				BuildingType    string `json:"buildingType"`
				TowerType       string `json:"towerType"`
				AfterID         int    `json:"afterId"`
				BeforeID        int    `json:"beforeId"`
				GoldGain        int    `json:"goldGain"`
				GameID          int    `json:"gameId"`
				WinningTeam     int    `json:"winningTeam"`
				TransformType   string `json:"transformType"`
				Name            string `json:"name"`
				ShutdownBounty  int    `json:"shutdownBounty"`
				ActualStartTime int    `json:"actualStartTime"`
			} `json:"events"`
			ParticipantFrames struct {
				Participant1  MatchTimelineParticipantFrame `json:"1"`
				Participant2  MatchTimelineParticipantFrame `json:"2,omitempty"`
				Participant3  MatchTimelineParticipantFrame `json:"3,omitempty"`
				Participant4  MatchTimelineParticipantFrame `json:"4,omitempty"`
				Participant5  MatchTimelineParticipantFrame `json:"5,omitempty"`
				Participant6  MatchTimelineParticipantFrame `json:"6,omitempty"`
				Participant7  MatchTimelineParticipantFrame `json:"7,omitempty"`
				Participant8  MatchTimelineParticipantFrame `json:"8,omitempty"`
				Participant9  MatchTimelineParticipantFrame `json:"9,omitempty"`
				Participant10 MatchTimelineParticipantFrame `json:"10,omitempty"`
			} `json:"participantFrames"`
			Timestamp int `json:"timestamp"`
		} `json:"frames"`
		GameID       int64 `json:"gameId"`
		Participants []struct {
			ParticipantID int    `json:"participantId"`
			Puuid         string `json:"puuid"`
		} `json:"participants"`
	} `json:"info"`
}

type MatchTimelineParticipantFrame struct {
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
}

type MatchlistOptions struct {
	// The matchlist started storing timestamps on June 16th, 2021.
	// Any matches played before June 16th, 2021 won't be included in the results if the startTime filter is set.
	StartTime int `json:"startTime"`
	// Epoch timestamp in seconds.
	EndTime int `json:"endTime"`
	// Filter the list of match ids by a specific queue ID.
	// This filter is mutually inclusive of the type filter meaning any match ids returned must match both the queue and type filters.
	Queue int `json:"queue"`
	// Filter the list of match ids by the type of match.
	// This filter is mutually inclusive of the queue filter meaning any match ids returned must match both the queue and type filters.
	Type MatchType `json:"type"`
	// Defaults to 0. Start index.
	Start int `json:"start"`
	// Defaults to 20. Valid values: 0 to 100. Number of match ids to return.
	Count int `json:"count"`
}

// Get a list of match IDs by PUUID.

// Start defaults to 0.
//
// Count defaults to 20.
func (e *MatchEndpoint) List(route Route, puuid string, options *MatchlistOptions) (*[]string, error) {
	logger := e.internalClient.Logger("LOL", "match-v5", "List")

	logger.Debug("Method executed")

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

	method := fmt.Sprintf(MatchListURL, puuid)

	url := fmt.Sprintf("%s?%s", method, query.Encode())

	var list *[]string

	err := e.internalClient.Get(route, url, &list, "match-v5", "List", "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return list, nil
}

// Get a match by match ID.
func (e *MatchEndpoint) ByID(route Route, matchID string) (*MatchDTO, error) {
	logger := e.internalClient.Logger("LOL", "match-v5", "ByID")

	logger.Debug("Method executed")

	url := fmt.Sprintf(MatchByIDURL, matchID)

	var match *MatchDTO

	err := e.internalClient.Get(route, url, &match, "match-v5", "ByID", "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return match, nil
}

// Get a match timeline by match ID.
func (e *MatchEndpoint) Timeline(route Route, matchID string) (*MatchTimelineDTO, error) {
	logger := e.internalClient.Logger("LOL", "match-v5", "Timeline")

	logger.Debug("Method executed")

	url := fmt.Sprintf(MatchTimelineURL, matchID)

	var timeline *MatchTimelineDTO

	err := e.internalClient.Get(route, url, &timeline, "match-v5", "Timeline", "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return timeline, nil
}
