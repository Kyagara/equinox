package benchmark_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/ddragon"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/clients/val"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkMatchByID-16 2730 409422 ns/op 226113 B/op 203 allocs/op
BenchmarkMatchByID-16 2724 406036 ns/op 226282 B/op 204 allocs/op
BenchmarkMatchByID-16 2676 408483 ns/op 226565 B/op 205 allocs/op
BenchmarkMatchByID-16 2661 420415 ns/op 226676 B/op 206 allocs/op
BenchmarkMatchByID-16 2768 418691 ns/op 226694 B/op 206 allocs/op
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

	config := equinox.NewTestEquinoxConfig()
	config.LogLevel = api.WARN_LOG_LEVEL
	config.Retry = true

	client, err := equinox.NewClientWithConfig(config)
	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.LOL.MatchV5.ByID(ctx, api.AMERICAS, "BR1_2744215970")
		require.Nil(b, err)
		require.Equal(b, res.Info.GameCreation, data.Info.GameCreation)
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkMatchTimeline-16 303 3988132 ns/op 2666936 B/op 899 allocs/op
BenchmarkMatchTimeline-16 306 3821935 ns/op 2665303 B/op 890 allocs/op
BenchmarkMatchTimeline-16 310 3783340 ns/op 2664774 B/op 893 allocs/op
BenchmarkMatchTimeline-16 313 3726846 ns/op 2663926 B/op 888 allocs/op
BenchmarkMatchTimeline-16 316 3723217 ns/op 2663558 B/op 889 allocs/op
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

	config := equinox.NewTestEquinoxConfig()
	config.LogLevel = api.WARN_LOG_LEVEL
	config.Retry = true

	client, err := equinox.NewClientWithConfig(config)
	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.LOL.MatchV5.Timeline(ctx, api.AMERICAS, "BR1_2744215970")
		require.Nil(b, err)
		require.Equal(b, res.Info.GameID, data.Info.GameID)
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkDDragonAllChampions-16 588 1895630 ns/op 836539 B/op 1540 allocs/op
BenchmarkDDragonAllChampions-16 553 1912140 ns/op 836887 B/op 1542 allocs/op
BenchmarkDDragonAllChampions-16 624 1874974 ns/op 835884 B/op 1538 allocs/op
BenchmarkDDragonAllChampions-16 620 1875622 ns/op 836504 B/op 1545 allocs/op
BenchmarkDDragonAllChampions-16 639 1872955 ns/op 835547 B/op 1539 allocs/op
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

	config := equinox.NewTestEquinoxConfig()
	config.LogLevel = api.WARN_LOG_LEVEL
	config.Retry = true

	client, err := equinox.NewClientWithConfig(config)
	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.DDragon.Champion.AllChampions(ctx, "13.22.1", ddragon.EnUS)
		require.Nil(b, err)
		require.Equal(b, "Ahri", data["Ahri"].Name)
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkVALContentAllLocales-16 19 59340126 ns/op 41893341 B/op 138466 allocs/op
BenchmarkVALContentAllLocales-16 18 59339266 ns/op 42086452 B/op 138851 allocs/op
BenchmarkVALContentAllLocales-16 18 58743392 ns/op 42085314 B/op 138849 allocs/op
BenchmarkVALContentAllLocales-16 19 58060897 ns/op 41890770 B/op 138464 allocs/op
BenchmarkVALContentAllLocales-16 19 58286993 ns/op 41891660 B/op 138467 allocs/op
*/
// Probably the largest response you can get with the Riot API.
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

	config := equinox.NewTestEquinoxConfig()
	config.LogLevel = api.WARN_LOG_LEVEL
	config.Retry = true

	client, err := equinox.NewClientWithConfig(config)
	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.VAL.ContentV1.Content(ctx, val.BR, "")
		require.Nil(b, err)
		require.Equal(b, res.Version, data.Version)
	}
}
