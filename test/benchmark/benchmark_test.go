package benchmark_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/ddragon"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/clients/val"
	"github.com/Kyagara/equinox/internal"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

/*
goos: windows
goarch: amd64
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkCachedSummonerByName-16 119008 10058 ns/op 4976 B/op 33 allocs/op
BenchmarkCachedSummonerByName-16 118884  9912 ns/op 4979 B/op 33 allocs/op
BenchmarkCachedSummonerByName-16 118033  9800 ns/op 4999 B/op 33 allocs/op
BenchmarkCachedSummonerByName-16 116996  9861 ns/op 5025 B/op 33 allocs/op
BenchmarkCachedSummonerByName-16 111648  9993 ns/op 5163 B/op 33 allocs/op
*/
// This version and the non cached version are used to estimate how the cache impacts performance.
func BenchmarkCachedSummonerByName(b *testing.B) {
	b.ReportAllocs()

	summoner := &lol.SummonerV4DTO{
		ID:            "5kIdR5x9LO0pVU_v01FtNVlb-dOws-D04GZCbNOmxCrB7A",
		AccountID:     "NkJ3FK5BQcrpKtF6Rj4PrAe9Nqodd2rwa5qJL8kJIPN_BkM",
		PUUID:         "6WQtgEvp61ZJ6f48qDZVQea1RYL9akRy7lsYOIHH8QDPnXr4E02E-JRwtNVE6n6GoGSU1wdXdCs5EQ",
		Name:          "Phanes",
		ProfileIconID: 1386,
		RevisionDate:  1657211888000,
		SummonerLevel: 68,
	}

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, lol.BR1)).
		Get("/lol/summoner/v4/summoners/by-name/Phanes").
		Persist().
		Reply(200).
		JSON(summoner)

	client, err := equinox.NewClient("RGAPI-TEST")
	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		data, err := client.LOL.SummonerV4.ByName(lol.BR1, "Phanes")
		require.Nil(b, err)
		require.Equal(b, "Phanes", data.Name)
	}
}

/*
goos: windows
goarch: amd64
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkSummonerByName-16 61848 19319 ns/op 4727 B/op 69 allocs/op
BenchmarkSummonerByName-16 61212 20062 ns/op 4858 B/op 70 allocs/op
BenchmarkSummonerByName-16 58365 20639 ns/op 5111 B/op 71 allocs/op
BenchmarkSummonerByName-16 55460 20802 ns/op 5112 B/op 71 allocs/op
BenchmarkSummonerByName-16 55996 21634 ns/op 5625 B/op 72 allocs/op
*/
func BenchmarkSummonerByName(b *testing.B) {
	b.ReportAllocs()

	summoner := &lol.SummonerV4DTO{
		ID:            "5kIdR5x9LO0pVU_v01FtNVlb-dOws-D04GZCbNOmxCrB7A",
		AccountID:     "NkJ3FK5BQcrpKtF6Rj4PrAe9Nqodd2rwa5qJL8kJIPN_BkM",
		PUUID:         "6WQtgEvp61ZJ6f48qDZVQea1RYL9akRy7lsYOIHH8QDPnXr4E02E-JRwtNVE6n6GoGSU1wdXdCs5EQ",
		Name:          "Phanes",
		ProfileIconID: 1386,
		RevisionDate:  1657211888000,
		SummonerLevel: 68,
	}

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, lol.BR1)).
		Get("/lol/summoner/v4/summoners/by-name/Phanes").
		Persist().
		Reply(200).
		JSON(summoner)

	config := internal.NewTestEquinoxConfig()
	config.LogLevel = api.WARN_LOG_LEVEL
	config.Retry = 1

	client, err := equinox.NewClientWithConfig(config)
	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		data, err := client.LOL.SummonerV4.ByName(lol.BR1, "Phanes")
		require.Nil(b, err)
		require.Equal(b, "Phanes", data.Name)
	}
}

/*
goos: windows
goarch: amd64
pkg: github.com/Kyagara/equinox
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkMatchByID-16 1399 856941 ns/op 235085 B/op 1093 allocs/op
BenchmarkMatchByID-16 1332 870089 ns/op 235232 B/op 1094 allocs/op
BenchmarkMatchByID-16 1352 911695 ns/op 235468 B/op 1095 allocs/op
BenchmarkMatchByID-16 1351 880334 ns/op 235495 B/op 1095 allocs/op
BenchmarkMatchByID-16 1374 892069 ns/op 235439 B/op 1095 allocs/op
*/
func BenchmarkMatchByID(b *testing.B) {
	b.ReportAllocs()

	var res lol.MatchV5DTO
	err := ReadFile("../data/match.json", &res)
	require.Nil(b, err)

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, api.AMERICAS)).
		Get(fmt.Sprintf("/lol/match/v5/matches/%v", "BR1_2744215970")).
		Persist().
		Reply(200).
		JSON(res)

	config := internal.NewTestEquinoxConfig()
	config.LogLevel = api.WARN_LOG_LEVEL
	config.Retry = 1

	client, err := equinox.NewClientWithConfig(config)
	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		data, err := client.LOL.MatchV5.ByID(api.AMERICAS, "BR1_2744215970")
		require.Nil(b, err)
		require.Equal(b, res.Info.GameCreation, data.Info.GameCreation)
	}
}

/*
goos: windows
goarch: amd64
pkg: github.com/Kyagara/equinox
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkMatchTimeline-16 138 9019941 ns/op 2779315 B/op 11979 allocs/op
BenchmarkMatchTimeline-16 132 8438055 ns/op 2780317 B/op 11983 allocs/op
BenchmarkMatchTimeline-16 140 8080165 ns/op 2778002 B/op 11979 allocs/op
BenchmarkMatchTimeline-16 147 8166768 ns/op 2775993 B/op 11974 allocs/op
BenchmarkMatchTimeline-16 150 7756563 ns/op 2775373 B/op 11973 allocs/op
*/
func BenchmarkMatchTimeline(b *testing.B) {
	b.ReportAllocs()

	var res lol.MatchTimelineV5DTO
	err := ReadFile("../data/match.timeline.json", &res)
	require.Nil(b, err)

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, api.AMERICAS)).
		Get(fmt.Sprintf("/lol/match/v5/matches/%v/timeline", "BR1_2744215970")).
		Persist().
		Reply(200).
		JSON(res)

	config := internal.NewTestEquinoxConfig()
	config.LogLevel = api.WARN_LOG_LEVEL
	config.Retry = 1

	client, err := equinox.NewClientWithConfig(config)
	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		data, err := client.LOL.MatchV5.Timeline(api.AMERICAS, "BR1_2744215970")
		require.Nil(b, err)
		require.Equal(b, res.Info.GameID, data.Info.GameID)
	}
}

/*
goos: windows
goarch: amd64
pkg: github.com/Kyagara/equinox
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkVALContentAllLocales-16 10 100190160 ns/op 45592683 B/op 166085 allocs/op
BenchmarkVALContentAllLocales-16 12  97401892 ns/op 44478412 B/op 163570 allocs/op
BenchmarkVALContentAllLocales-16 10 100424060 ns/op 45591764 B/op 166083 allocs/op
BenchmarkVALContentAllLocales-16 10 100606000 ns/op 45591508 B/op 166082 allocs/op
BenchmarkVALContentAllLocales-16 12  96558058 ns/op 44478152 B/op 163568 allocs/op
*/
// Probably the biggest api call you can make with the Riot API.
func BenchmarkVALContentAllLocales(b *testing.B) {
	b.ReportAllocs()

	var res val.ContentV1DTO
	err := ReadFile("../data/val.content.all_locales.json", &res)
	require.Nil(b, err)

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, val.BR)).
		Get("/val/content/v1/contents").
		Persist().
		Reply(200).
		JSON(res)

	config := internal.NewTestEquinoxConfig()
	config.LogLevel = api.WARN_LOG_LEVEL
	config.Retry = 1

	client, err := equinox.NewClientWithConfig(config)
	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		data, err := client.VAL.ContentV1.Content(val.BR, "")
		require.Nil(b, err)
		require.Equal(b, res.Version, data.Version)
	}
}

/*
goos: windows
goarch: amd64
pkg: github.com/Kyagara/equinox
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkDDragonAllChampions-16 312 3879322 ns/op 868727 B/op 6213 allocs/op
BenchmarkDDragonAllChampions-16 302 3891187 ns/op 868926 B/op 6214 allocs/op
BenchmarkDDragonAllChampions-16 310 3833106 ns/op 868675 B/op 6215 allocs/op
BenchmarkDDragonAllChampions-16 301 3832329 ns/op 868830 B/op 6215 allocs/op
BenchmarkDDragonAllChampions-16 310 3703218 ns/op 868455 B/op 6214 allocs/op
*/
func BenchmarkDDragonAllChampions(b *testing.B) {
	b.ReportAllocs()

	var data ddragon.ChampionsData
	err := ReadFile("../data/champions.json", &data)
	require.Nil(b, err)

	gock.New(fmt.Sprintf(api.D_DRAGON_BASE_URL_FORMAT, "")).
		Get(fmt.Sprintf(ddragon.ChampionsURL, "13.22.1", ddragon.EnUS)).
		Persist().
		Reply(200).
		JSON(data)

	config := internal.NewTestEquinoxConfig()
	config.LogLevel = api.WARN_LOG_LEVEL
	config.Retry = 1

	client, err := equinox.NewClientWithConfig(config)
	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		data, err := client.DDragon.Champion.AllChampions("13.22.1", ddragon.EnUS)
		require.Nil(b, err)
		require.Equal(b, "Ahri", data["Ahri"].Name)
	}
}

func ReadFile(filename string, target any) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, target)
	if err != nil {
		return err
	}
	return nil
}
