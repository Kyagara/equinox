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
	Tier     Tier             `json:"tier"`
	LeagueID string           `json:"leagueId"`
	Queue    QueueType        `json:"queue"`
	Name     string           `json:"name"`
	Entries  []LeagueEntryDTO `json:"entries"`
}

type LeagueEntryDTO struct {
	LeagueID  string    `json:"leagueId"`
	QueueType QueueType `json:"queueType"`
	Tier      Tier      `json:"tier"`
	// The player's division within a tier.
	Rank api.Division `json:"rank"`
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

// Get all the league entries.
//
// Page defaults to 1.
func (e *LeagueEndpoint) Entries(region Region, queue QueueType, tier Tier, division api.Division, page int) (*[]LeagueEntryDTO, error) {
	logger := e.internalClient.Logger("LOL", "league", "Entries")

	if region == PBE1 {
		return nil, fmt.Errorf("the region PBE1 is not available for this method")
	}

	if tier == MasterTier || tier == GrandmasterTier || tier == ChallengerTier {
		return nil, fmt.Errorf("the tier specified is an apex tier, please use the corresponded method instead")
	}

	query := url.Values{}

	if page < 1 {
		page = 1
	}

	query.Set("page", strconv.Itoa(page))

	method := fmt.Sprintf(LeagueEntriesURL, queue, tier, division)

	url := fmt.Sprintf("%s?%s", method, query.Encode())

	var entries *[]LeagueEntryDTO

	err := e.internalClient.Do(http.MethodGet, region, url, nil, &entries, "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return entries, nil
}

// Get league entries in all queues for a given summoner ID.
func (e *LeagueEndpoint) SummonerEntries(region Region, summonerID string) (*[]LeagueEntryDTO, error) {
	logger := e.internalClient.Logger("LOL", "league", "SummonerEntries")

	if region == PBE1 {
		return nil, fmt.Errorf("the region PBE1 is not available for this method")
	}

	url := fmt.Sprintf(LeagueEntriesBySummonerURL, summonerID)

	var entries *[]LeagueEntryDTO

	err := e.internalClient.Do(http.MethodGet, region, url, nil, &entries, "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return entries, nil
}

// Get the challenger league for given queue.
func (e *LeagueEndpoint) ChallengerByQueue(region Region, queueType QueueType) (*LeagueListDTO, error) {
	return e.getLeague(LeagueChallengerURL, region, queueType, "ChallengerByQueue")
}

// Get the grandmaster league for given queue.
func (e *LeagueEndpoint) GrandmasterByQueue(region Region, queueType QueueType) (*LeagueListDTO, error) {
	return e.getLeague(LeagueGrandmasterURL, region, queueType, "GrandmasterByQueue")
}

// Get the master league for given queue.
func (e *LeagueEndpoint) MasterByQueue(region Region, queueType QueueType) (*LeagueListDTO, error) {
	return e.getLeague(LeagueMasterURL, region, queueType, "MasterByQueue")
}

// Get league with given ID, including inactive entries.
func (e *LeagueEndpoint) ByID(region Region, leagueID string) (*LeagueListDTO, error) {
	return e.getLeague(LeagueByIDURL, region, leagueID, "ByID")
}

func (e *LeagueEndpoint) getLeague(endpointMethod string, region Region, queueType interface{}, methodName string) (*LeagueListDTO, error) {
	logger := e.internalClient.Logger("LOL", "league", methodName)

	if region == PBE1 {
		return nil, fmt.Errorf("the region PBE1 is not available for this method")
	}

	url := fmt.Sprintf(endpointMethod, queueType)

	var league *LeagueListDTO

	err := e.internalClient.Do(http.MethodGet, region, url, nil, &league, "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return league, nil
}
