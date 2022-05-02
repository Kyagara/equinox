package lol

import (
	"fmt"
	"net/http"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type SpectatorEndpoint struct {
	internalClient *internal.InternalClient
}

type FeaturedGamesDTO struct {
	// The list of featured games.
	GameList []FeaturedGameInfoDTO `json:"gameList"`
	// The suggested interval to wait before requesting FeaturedGames again.
	ClientRefreshInterval int `json:"clientRefreshInterval"`
}

type FeaturedGameInfoDTO struct {
	// The ID of the game.
	GameID int `json:"gameId"`
	// The ID of the map.
	MapID int `json:"mapId"`
	// The game mode (Legal values: CLASSIC, ODIN, ARAM, TUTORIAL, ONEFORALL, ASCENSION, FIRSTBLOOD, KINGPORO).
	GameMode string `json:"gameMode"`
	// The game type (Legal values: CUSTOM_GAME, MATCHED_GAME, TUTORIAL_GAME).
	GameType string `json:"gameType"`
	// The queue type (queue types are documented on the Game Constants page).
	GameQueueConfigID int `json:"gameQueueConfigId"`
	// The participant information.
	Participants []ParticipantDTO `json:"participants"`
	// The observer information.
	Observers ObserverDTO `json:"observers"`
	// The ID of the platform on which the game is being played.
	PlatformID string `json:"platformId"`
	// Banned champion information.
	BannedChampions []BannedChampionDTO `json:"bannedChampions"`
	// The game start time represented in epoch milliseconds.
	GameStartTime int `json:"gameStartTime"`
	// The amount of time in seconds that has passed since the game started.
	GameLength int `json:"gameLength"`
}

type ParticipantDTO struct {
	// The team ID of this participant, indicating the participant's team.
	TeamID int `json:"teamId"`
	// The ID of the first summoner spell used by this participant.
	Spell1ID int `json:"spell1Id"`
	// The ID of the second summoner spell used by this participant.
	Spell2ID int `json:"spell2Id"`
	// The ID of the champion played by this participant.
	ChampionID int `json:"championId"`
	// The ID of the profile icon used by this participant.
	ProfileIconID int `json:"profileIconId"`
	// The summoner name of this participant.
	SummonerName string `json:"summonerName"`
	// Flag indicating whether or not this participant is a bot.
	Bot bool `json:"bot"`
}

type ObserverDTO struct {
	// Key used to decrypt the spectator grid game data for playback.
	EncryptionKey string `json:"encryptionKey"`
}

type BannedChampionDTO struct {
	// The ID of the banned champion.
	ChampionID int `json:"championId"`
	// The ID of the team that banned the champion.
	TeamID int `json:"teamId"`
	// The turn during which the champion was banned.
	PickTurn int `json:"pickTurn"`
}

type CurrentGameInfoDTO struct {
	// The ID of the game.
	GameID int `json:"gameId"`
	// The ID of the map.
	MapID int `json:"mapId"`
	// The game mode (Legal values: CLASSIC, ODIN, ARAM, TUTORIAL, ONEFORALL, ASCENSION, FIRSTBLOOD, KINGPORO).
	GameMode string `json:"gameMode"`
	// The game type (Legal values: CUSTOM_GAME, MATCHED_GAME, TUTORIAL_GAME).
	GameType string `json:"gameType"`
	// The queue type (queue types are documented on the Game Constants page).
	GameQueueConfigID int `json:"gameQueueConfigId"`
	// The participant information.
	Participants []CurrentGameParticipantDTO `json:"participants"`
	// The observer information.
	Observers ObserverDTO `json:"observers"`
	// The ID of the platform on which the game is being played.
	PlatformID string `json:"platformId"`
	// Banned champion information.
	BannedChampions []BannedChampionDTO `json:"bannedChampions"`
	// The game start time represented in epoch milliseconds.
	GameStartTime int `json:"gameStartTime"`
	// The amount of time in seconds that has passed since the game started.
	GameLength int `json:"gameLength"`
}

type CurrentGameParticipantDTO struct {
	// The team ID of this participant, indicating the participant's team.
	TeamID int `json:"teamId"`
	// The ID of the first summoner spell used by this participant.
	Spell1ID int `json:"spell1Id"`
	// The ID of the second summoner spell used by this participant.
	Spell2ID int `json:"spell2Id"`
	// The ID of the champion played by this participant.
	ChampionID int `json:"championId"`
	// The ID of the profile icon used by this participant.
	ProfileIconID int `json:"profileIconId"`
	// The encrypted summoner ID of this participant.
	SummonerID string `json:"summonerId"`
	// The summoner name of this participant.
	SummonerName string `json:"summonerName"`
	// Flag indicating whether or not this participant is a bot.
	Bot bool `json:"bot"`
	// List of Game Customizations.
	GameCustomizationObjects []GameCustomizationObject `json:"gameCustomizationObjects"`
	// Perks/Runes Reforged Information.
	Perks `json:"perks"`
}

type GameCustomizationObject struct {
	// Category identifier for Game Customization.
	Category string `json:"category"`
	// Game Customization content.
	Content string `json:"content"`
}

type Perks struct {
	// IDs of the perks/runes assigned.
	PerkIDs []int `json:"perkIds"`
	// Primary runes path.
	PerkStyle int `json:"perkStyle"`
	// Secondary runes path.
	PerkSubStyle int `json:"perkSubStyle"`
}

// Get featured games in a region.
func (c *SpectatorEndpoint) FeaturedGames(region api.LOLRegion) (*FeaturedGamesDTO, error) {
	res := FeaturedGamesDTO{}

	err := c.internalClient.Do(http.MethodGet, region, SpectatorURL, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get the current game information for the given summoner ID.
func (c *SpectatorEndpoint) CurrentGame(region api.LOLRegion, summonerID string) (*CurrentGameInfoDTO, error) {
	res := CurrentGameInfoDTO{}

	err := c.internalClient.Do(http.MethodGet, region, fmt.Sprintf(SpectatorCurrentGameURL, summonerID), nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
