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
	"github.com/stretchr/testify/require"
)

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkMatchByID-16 1812 622702 ns/op 70289 B/op 144 allocs/op
BenchmarkMatchByID-16 1819 624619 ns/op 70246 B/op 144 allocs/op
BenchmarkMatchByID-16 1840 630456 ns/op 70261 B/op 144 allocs/op
BenchmarkMatchByID-16 1888 632890 ns/op 70267 B/op 144 allocs/op
BenchmarkMatchByID-16 1779 644411 ns/op 70282 B/op 144 allocs/op
*/
func BenchmarkMatchByID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://americas.api.riotgames.com/lol/match/v5/matches/BR1_2744215970",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/match.json")))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
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
BenchmarkMatchTimeline-16 270 4426118 ns/op 1108450 B/op 828 allocs/op
BenchmarkMatchTimeline-16 270 4383629 ns/op 1108503 B/op 828 allocs/op
BenchmarkMatchTimeline-16 272 4398360 ns/op 1108429 B/op 828 allocs/op
BenchmarkMatchTimeline-16 259 4448896 ns/op 1108504 B/op 828 allocs/op
BenchmarkMatchTimeline-16 270 4403381 ns/op 1108503 B/op 828 allocs/op
*/
func BenchmarkMatchTimeline(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://americas.api.riotgames.com/lol/match/v5/matches/BR1_2744215970/timeline",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/match.timeline.json")))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
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
BenchmarkDDragonAllChampions-16 681 1742204 ns/op 143269 B/op 1490 allocs/op
BenchmarkDDragonAllChampions-16 670 1772683 ns/op 143280 B/op 1490 allocs/op
BenchmarkDDragonAllChampions-16 648 1763162 ns/op 143278 B/op 1490 allocs/op
BenchmarkDDragonAllChampions-16 675 1801580 ns/op 143271 B/op 1490 allocs/op
BenchmarkDDragonAllChampions-16 676 1775978 ns/op 143265 B/op 1490 allocs/op
*/
func BenchmarkDDragonAllChampions(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://ddragon.leagueoflegends.com/cdn/13.22.1/data/en_US/champion.json",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/champions.json")))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
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
BenchmarkVALContentAllLocales-16 24 49882837 ns/op 11794468 B/op 131497 allocs/op
BenchmarkVALContentAllLocales-16 21 50197158 ns/op 11825450 B/op 131498 allocs/op
BenchmarkVALContentAllLocales-16 21 50686294 ns/op 11825441 B/op 131498 allocs/op
BenchmarkVALContentAllLocales-16 21 49946730 ns/op 11824723 B/op 131497 allocs/op
BenchmarkVALContentAllLocales-16 21 50377181 ns/op 11824726 B/op 131497 allocs/op
*/
// Probably the largest response you can get with the Riot API.
func BenchmarkVALContentAllLocales(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br.api.riotgames.com/val/content/v1/contents",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/val.content.all_locales.json")))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
	client := equinox.NewClientWithConfig(config)

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.VAL.ContentV1.Content(ctx, val.BR, "")
		require.NoError(b, err)
		require.Equal(b, "release-07.10", data.Version)
	}
}
