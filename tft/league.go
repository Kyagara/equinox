package tft

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
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
	LeagueID  string    `json:"leagueId"`
	QueueType QueueType `json:"queueType"`
	Tier      lol.Tier  `json:"tier"`
	// Only included for the RANKED_TFT_TURBO queueType.
	RatedTier RatedTier `json:"ratedTier"`
	// Only included for the RANKED_TFT_TURBO queueType.
	RatedRating string       `json:"ratedRating"`
	Rank        api.Division `json:"rank"`
	// Player's encrypted summonerId.
	SummonerID   string `json:"summonerId"`
	SummonerName string `json:"summonerName"`
	LeaguePoints int    `json:"leaguePoints"`
	// First placement.
	Wins int `json:"wins"`
	// Second through eighth placement.
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

type TopRatedLadderEntryDTO struct {
	SummonerID   string    `json:"summonerId"`
	SummonerName string    `json:"summonerName"`
	RatedTier    RatedTier `json:"ratedTier"`
	RatedRating  int       `json:"ratedRating"`
	// First placement.
	Wins                         int `json:"wins"`
	PreviousUpdateLadderPosition int `json:"previousUpdateLadderPosition"`
}

// Get all the league entries.
//
// Page defaults to 1.
func (l *LeagueEndpoint) Entries(region lol.Region, tier lol.Tier, division api.Division, page int) (*[]LeagueEntryDTO, error) {
	logger := l.internalClient.Logger("tft").With("endpoint", "league", "method", "Entries")

	if tier == lol.MasterTier || tier == lol.GrandmasterTier || tier == lol.ChallengerTier {
		return nil, fmt.Errorf("the tier specified is an apex tier, please use the corresponded method instead")
	}

	query := url.Values{}

	if page < 1 {
		page = 1
	}

	query.Set("page", strconv.Itoa(page))

	method := fmt.Sprintf(LeagueEntriesURL, tier, division)

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
func (l *LeagueEndpoint) ByID(region lol.Region, leagueID string) (*LeagueListDTO, error) {
	return l.getLeague(fmt.Sprintf(LeagueByIDURL, leagueID), region, "ByID")
}

// Get league entries in all queues for a given summoner ID.
func (l *LeagueEndpoint) SummonerEntries(region lol.Region, summonerID string) (*[]LeagueEntryDTO, error) {
	logger := l.internalClient.Logger("tft").With("endpoint", "league", "method", "SummonerEntries")

	url := fmt.Sprintf(LeagueEntriesBySummonerURL, summonerID)

	var entries *[]LeagueEntryDTO

	err := l.internalClient.Do(http.MethodGet, region, url, nil, &entries, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return entries, nil
}

// Get the top rated ladder for given queue.
func (l *LeagueEndpoint) TopRatedLadder(region lol.Region, queue QueueType) (*[]TopRatedLadderEntryDTO, error) {
	logger := l.internalClient.Logger("tft").With("endpoint", "league", "method", "TopRatedLadder")

	if queue == RankedTFT {
		return nil, fmt.Errorf("the queue specified is not available for the top rated ladder endpoint, please use the RankedTFTTurbo queue")
	}

	url := fmt.Sprintf(LeagueRatedLaddersURL, queue)

	var entries *[]TopRatedLadderEntryDTO

	err := l.internalClient.Do(http.MethodGet, region, url, nil, &entries, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return entries, nil
}

// Get the challenger league.
func (l *LeagueEndpoint) Challenger(region lol.Region) (*LeagueListDTO, error) {
	return l.getLeague(LeagueChallengerURL, region, "Challenger")
}

// Get the grandmaster league.
func (l *LeagueEndpoint) Grandmaster(region lol.Region) (*LeagueListDTO, error) {
	return l.getLeague(LeagueGrandmasterURL, region, "Grandmaster")
}

// Get the master league.
func (l *LeagueEndpoint) Master(region lol.Region) (*LeagueListDTO, error) {
	return l.getLeague(LeagueMasterURL, region, "Master")
}

func (l *LeagueEndpoint) getLeague(url string, region lol.Region, methodName string) (*LeagueListDTO, error) {
	logger := l.internalClient.Logger("tft").With("endpoint", "league", "method", methodName)

	var league *LeagueListDTO

	err := l.internalClient.Do(http.MethodGet, region, url, nil, &league, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return league, nil
}
