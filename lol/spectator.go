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
	GameList []struct {
		GameID            int    `json:"gameId"`
		MapID             int    `json:"mapId"`
		GameMode          string `json:"gameMode"`
		GameType          string `json:"gameType"`
		GameQueueConfigID int    `json:"gameQueueConfigId"`
		Participants      []struct {
			TeamID        int    `json:"teamId"`
			Spell1ID      int    `json:"spell1Id"`
			Spell2ID      int    `json:"spell2Id"`
			ChampionID    int    `json:"championId"`
			ProfileIconID int    `json:"profileIconId"`
			SummonerName  string `json:"summonerName"`
			Bot           bool   `json:"bot"`
		} `json:"participants"`
		Observers struct {
			EncryptionKey string `json:"encryptionKey"`
		} `json:"observers"`
		PlatformID      string `json:"platformId"`
		BannedChampions []struct {
			ChampionID int `json:"championId"`
			TeamID     int `json:"teamId"`
			PickTurn   int `json:"pickTurn"`
		} `json:"bannedChampions"`
		GameStartTime int `json:"gameStartTime"`
		GameLength    int `json:"gameLength"`
	} `json:"gameList"`
	ClientRefreshInterval int `json:"clientRefreshInterval"`
}

type ActiveGameBySummonerID struct {
	GameID            int    `json:"gameId"`
	MapID             int    `json:"mapId"`
	GameMode          string `json:"gameMode"`
	GameType          string `json:"gameType"`
	GameQueueConfigID int    `json:"gameQueueConfigId"`
	Participants      []struct {
		TeamID                   int    `json:"teamId"`
		Spell1ID                 int    `json:"spell1Id"`
		Spell2ID                 int    `json:"spell2Id"`
		ChampionID               int    `json:"championId"`
		ProfileIconID            int    `json:"profileIconId"`
		SummonerName             string `json:"summonerName"`
		Bot                      bool   `json:"bot"`
		SummonerID               string `json:"summonerId"`
		GameCustomizationObjects []struct {
			Category string `json:"category"`
			Content  string `json:"content"`
		} `json:"gameCustomizationObjects"`
		Perks struct {
			PerkIds      []int `json:"perkIds"`
			PerkStyle    int   `json:"perkStyle"`
			PerkSubStyle int   `json:"perkSubStyle"`
		} `json:"perks"`
	} `json:"participants"`
	Observers struct {
		EncryptionKey string `json:"encryptionKey"`
	} `json:"observers"`
	PlatformID      string `json:"platformId"`
	BannedChampions []struct {
		ChampionID int `json:"championId"`
		TeamID     int `json:"teamId"`
		PickTurn   int `json:"pickTurn"`
	} `json:"bannedChampions"`
	GameStartTime int `json:"gameStartTime"`
	GameLength    int `json:"gameLength"`
}

// Get featured games in a region
func (c *SpectatorEndpoint) FeaturedGames(region api.Region) (*FeaturedGamesDTO, error) {
	res := FeaturedGamesDTO{}

	err := c.internalClient.Do(http.MethodGet, region, SpectatorEndpointURL, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get an active game by summoner ID
func (c *SpectatorEndpoint) ActiveGameBySummonerID(region api.Region, summonerId string) (*ActiveGameBySummonerID, error) {
	res := ActiveGameBySummonerID{}

	err := c.internalClient.Do(http.MethodGet, region, fmt.Sprintf(SpectatorBySummonerEndpointURL, summonerId), nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
