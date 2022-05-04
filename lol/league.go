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
func (l *LeagueEndpoint) Entries(region Region, division api.Division, tier Tier, queue QueueType, page int) (*[]LeagueEntryDTO, error) {
	logger := l.internalClient.Logger().With("endpoint", "league", "method", "Entries")

	query := url.Values{}

	if page < 1 {
		page = 1
	}

	query.Set("page", strconv.Itoa(page))

	method := fmt.Sprintf(LeagueEntriesURL, division, tier, queue)

	url := fmt.Sprintf("%s?%s", method, query.Encode())

	var entries *[]LeagueEntryDTO

	err := l.internalClient.Do(http.MethodGet, region, url, nil, &entries, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return entries, nil
}

// Get league with given ID, including inactive entries.
func (l *LeagueEndpoint) ByID(region Region, leagueID string) (*LeagueListDTO, error) {
	return l.getLeague(LeagueByIDURL, region, QueueType(leagueID), "ByID")
}

// Get league entries in all queues for a given summoner ID.
func (l *LeagueEndpoint) SummonerEntries(region Region, summonerID string) (*[]LeagueEntryDTO, error) {
	logger := l.internalClient.Logger().With("endpoint", "league", "method", "SummonerEntries")

	url := fmt.Sprintf(LeagueEntriesBySummonerURL, summonerID)

	var entries *[]LeagueEntryDTO

	err := l.internalClient.Do(http.MethodGet, region, url, nil, &entries, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return entries, nil
}

// Get the challenger league for given queue.
func (l *LeagueEndpoint) ChallengerByQueue(region Region, queueType QueueType) (*LeagueListDTO, error) {
	return l.getLeague(LeagueChallengerURL, region, queueType, "ChallengerByQueue")
}

// Get the grandmaster league for given queue.
func (l *LeagueEndpoint) GrandmasterByQueue(region Region, queueType QueueType) (*LeagueListDTO, error) {
	return l.getLeague(LeagueGrandmasterURL, region, queueType, "GrandmasterByQueue")
}

// Get the master league for given queue.
func (l *LeagueEndpoint) MasterByQueue(region Region, queueType QueueType) (*LeagueListDTO, error) {
	return l.getLeague(LeagueMasterURL, region, queueType, "MasterByQueue")
}

func (l *LeagueEndpoint) getLeague(endpointMethod string, region Region, queueType QueueType, methodName string) (*LeagueListDTO, error) {
	logger := l.internalClient.Logger().With("endpoint", "league", "method", methodName)

	url := fmt.Sprintf(endpointMethod, queueType)

	var league *LeagueListDTO

	err := l.internalClient.Do(http.MethodGet, region, url, nil, &league, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return league, nil
}
