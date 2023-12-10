package benchmark_test

import (
	"context"
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/ddragon"
	"github.com/Kyagara/equinox/clients/val"
	"github.com/Kyagara/equinox/test/util"
	"github.com/jarcoal/httpmock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkMatchByID-16 2778 396701 ns/op 225703 B/op 198 allocs/op
BenchmarkMatchByID-16 2748 399298 ns/op 225839 B/op 199 allocs/op
BenchmarkMatchByID-16 2838 393265 ns/op 226074 B/op 200 allocs/op
BenchmarkMatchByID-16 2846 397531 ns/op 226099 B/op 200 allocs/op
BenchmarkMatchByID-16 2767 394891 ns/op 225970 B/op 199 allocs/op
*/
func BenchmarkMatchByID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://americas.api.riotgames.com/lol/match/v5/matches/BR1_2744215970",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/match.json")))

	config := util.NewTestEquinoxConfig()
	config.LogLevel = zerolog.WarnLevel
	config.Retries = 3

	client, err := equinox.NewClientWithConfig(config)
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.LOL.MatchV5.ByID(ctx, api.AMERICAS, "BR1_2744215970")
		require.NoError(b, err)
		require.Equal(b, int64(1686266124922), data.Info.GameCreation)
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkMatchTimeline-16 307 3813587 ns/op 2666872 B/op 911 allocs/op
BenchmarkMatchTimeline-16 309 3709283 ns/op 2663917 B/op 887 allocs/op
BenchmarkMatchTimeline-16 314 3676451 ns/op 2663491 B/op 892 allocs/op
BenchmarkMatchTimeline-16 321 3715784 ns/op 2661674 B/op 877 allocs/op
BenchmarkMatchTimeline-16 324 3627076 ns/op 2661302 B/op 877 allocs/op
*/
func BenchmarkMatchTimeline(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://americas.api.riotgames.com/lol/match/v5/matches/BR1_2744215970/timeline",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/match.timeline.json")))

	config := util.NewTestEquinoxConfig()
	config.LogLevel = zerolog.WarnLevel
	config.Retries = 3

	client, err := equinox.NewClientWithConfig(config)
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.LOL.MatchV5.Timeline(ctx, api.AMERICAS, "BR1_2744215970")
		require.NoError(b, err)
		require.Equal(b, int64(2744215970), data.Info.GameID)
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkDDragonAllChampions-16 622 1848744 ns/op 835533 B/op 1536 allocs/op
BenchmarkDDragonAllChampions-16 622 1865572 ns/op 835403 B/op 1533 allocs/op
BenchmarkDDragonAllChampions-16 634 1874150 ns/op 836084 B/op 1540 allocs/op
BenchmarkDDragonAllChampions-16 638 1839908 ns/op 834754 B/op 1531 allocs/op
BenchmarkDDragonAllChampions-16 637 1906226 ns/op 834147 B/op 1528 allocs/op
*/
func BenchmarkDDragonAllChampions(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://ddragon.leagueoflegends.com/cdn/13.22.1/data/en_US/champion.json",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/champions.json")))

	config := util.NewTestEquinoxConfig()
	config.LogLevel = zerolog.WarnLevel
	config.Retries = 3

	client, err := equinox.NewClientWithConfig(config)
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.DDragon.Champion.AllChampions(ctx, "13.22.1", ddragon.EnUS)
		require.NoError(b, err)
		require.Equal(b, "Ahri", data["Ahri"].Name)
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkVALContentAllLocales-16 18 57826153 ns/op 42062045 B/op 138842 allocs/op
BenchmarkVALContentAllLocales-16 19 56859852 ns/op 41867027 B/op 138453 allocs/op
BenchmarkVALContentAllLocales-16 18 55964393 ns/op 42059940 B/op 138841 allocs/op
BenchmarkVALContentAllLocales-16 19 56303605 ns/op 41865869 B/op 138453 allocs/op
BenchmarkVALContentAllLocales-16 18 55905806 ns/op 42058712 B/op 138839 allocs/op
*/
// Probably the largest response you can get with the Riot API.
func BenchmarkVALContentAllLocales(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br.api.riotgames.com/val/content/v1/contents",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/val.content.all_locales.json")))

	config := util.NewTestEquinoxConfig()
	config.LogLevel = zerolog.WarnLevel
	config.Retries = 3

	client, err := equinox.NewClientWithConfig(config)
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.VAL.ContentV1.Content(ctx, val.BR, "")
		require.NoError(b, err)
		require.Equal(b, "release-07.10", data.Version)
	}
}
