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
BenchmarkMatchByID-16 1884 605817 ns/op 70609 B/op 173 allocs/op
BenchmarkMatchByID-16 1950 602600 ns/op 70607 B/op 173 allocs/op
BenchmarkMatchByID-16 2008 608089 ns/op 70606 B/op 173 allocs/op
BenchmarkMatchByID-16 1962 613049 ns/op 70607 B/op 173 allocs/op
BenchmarkMatchByID-16 1998 615987 ns/op 70606 B/op 173 allocs/op
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

	client := equinox.NewClientWithConfig(config)

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
BenchmarkMatchTimeline-16 272 4362339 ns/op 1107932 B/op 836 allocs/op
BenchmarkMatchTimeline-16 274 4316877 ns/op 1107921 B/op 836 allocs/op
BenchmarkMatchTimeline-16 277 4232793 ns/op 1107906 B/op 836 allocs/op
BenchmarkMatchTimeline-16 282 4209472 ns/op 1107879 B/op 836 allocs/op
BenchmarkMatchTimeline-16 276 4208447 ns/op 1107912 B/op 836 allocs/op
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

	client := equinox.NewClientWithConfig(config)

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
BenchmarkDDragonAllChampions-16 697 1706249 ns/op 143455 B/op 1499 allocs/op
BenchmarkDDragonAllChampions-16 688 1696351 ns/op 143476 B/op 1499 allocs/op
BenchmarkDDragonAllChampions-16 694 1682909 ns/op 143479 B/op 1499 allocs/op
BenchmarkDDragonAllChampions-16 687 1683969 ns/op 143465 B/op 1499 allocs/op
BenchmarkDDragonAllChampions-16 667 1723390 ns/op 143470 B/op 1499 allocs/op
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

	client := equinox.NewClientWithConfig(config)

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
BenchmarkVALContentAllLocales-16 21 50809133 ns/op 11827537 B/op 131509 allocs/op
BenchmarkVALContentAllLocales-16 22 53467450 ns/op 11816480 B/op 131509 allocs/op
BenchmarkVALContentAllLocales-16 20 51978820 ns/op 11839869 B/op 131510 allocs/op
BenchmarkVALContentAllLocales-16 20 52869176 ns/op 11839236 B/op 131509 allocs/op
BenchmarkVALContentAllLocales-16 21 50267745 ns/op 11827097 B/op 131508 allocs/op
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

	client := equinox.NewClientWithConfig(config)

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.VAL.ContentV1.Content(ctx, val.BR, "")
		require.NoError(b, err)
		require.Equal(b, "release-07.10", data.Version)
	}
}
