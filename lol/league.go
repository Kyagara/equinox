package lol

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type LeagueEndpoint struct {
	internalClient *internal.InternalClient
}

type LeagueListDTO struct {
	Tier     string           `json:"tier"`
	LeagueID string           `json:"leagueId"`
	Queue    string           `json:"queue"`
	Name     string           `json:"name"`
	Entries  []LeagueEntryDTO `json:"entries"`
}

type LeagueEntryDTO struct {
	LeagueID  string `json:"leagueId"`
	QueueType string `json:"queueType"`
	Tier      string `json:"tier"`
	// The player's division within a tier.
	Rank string `json:"rank"`
	// Player's encrypted summonerId.
	SummonerID   string `json:"summonerId"`
	SummonerName string `json:"summonerName"`
	LeaguePoints int    `json:"leaguePoints"`
	// Winning team on Summoners Rift.
	Wins int `json:"wins"`
	// Losing team on Summoners Rift.
	Losses     int           `json:"losses"`
	Veteran    bool          `json:"veteran"`
	Inactive   bool          `json:"inactive"`
	FreshBlood bool          `json:"freshBlood"`
	HotStreak  bool          `json:"hotStreak"`
	MiniSeries MiniSeriesDTO `json:"miniSeries,omitempty"`
}

type MiniSeriesDTO struct {
	Progress string `json:"progress"`
	Losses   int    `json:"losses"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}

// Get all the league entries. Page defaults to 1.
func (c *LeagueEndpoint) Entries(region Region, division api.Division, tier Tier, queue QueueType, page int) (*[]LeagueEntryDTO, error) {
	query := url.Values{}

	if page < 1 {
		page = 1
	}

	query.Set("page", strconv.Itoa(page))

	method := fmt.Sprintf(LeagueEntriesURL, division, tier, queue)

	url := fmt.Sprintf("%s?%s", method, query.Encode())

	res := []LeagueEntryDTO{}

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get league with given ID, including inactive entries.
func (c *LeagueEndpoint) ByID(region Region, leagueID string) (*LeagueListDTO, error) {
	url := fmt.Sprintf(LeagueByID, leagueID)

	res := LeagueListDTO{}

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get league entries in all queues for a given summoner ID.
func (c *LeagueEndpoint) SummonerEntries(region Region, summonerID string) (*[]LeagueEntryDTO, error) {
	url := fmt.Sprintf(LeagueEntriesBySummonerURL, summonerID)

	res := []LeagueEntryDTO{}

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get the challenger league for given queue.
func (c *LeagueEndpoint) ChallengerByQueue(region Region, queueType QueueType) (*LeagueListDTO, error) {
	return c.getLeague(LeagueChallengerURL, region, queueType)
}

// Get the grandmaster league for given queue.
func (c *LeagueEndpoint) GrandmasterByQueue(region Region, queueType QueueType) (*LeagueListDTO, error) {
	return c.getLeague(LeagueGrandmasterURL, region, queueType)
}

// Get the master league for given queue.
func (c *LeagueEndpoint) MasterByQueue(region Region, queueType QueueType) (*LeagueListDTO, error) {
	return c.getLeague(LeagueMasterURL, region, queueType)
}

func (c *LeagueEndpoint) getLeague(method string, region Region, queueType QueueType) (*LeagueListDTO, error) {
	url := fmt.Sprintf(method, queueType)

	res := LeagueListDTO{}

	err := c.internalClient.Do(http.MethodGet, region, url, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
