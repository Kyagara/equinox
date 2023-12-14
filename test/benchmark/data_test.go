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
BenchmarkMatchByID-16 1904 621687 ns/op 70432 B/op 168 allocs/op
BenchmarkMatchByID-16 1897 618226 ns/op 70432 B/op 168 allocs/op
BenchmarkMatchByID-16 1909 609363 ns/op 70432 B/op 168 allocs/op
BenchmarkMatchByID-16 1858 618229 ns/op 70425 B/op 168 allocs/op
BenchmarkMatchByID-16 1856 611426 ns/op 70434 B/op 168 allocs/op
*/
func BenchmarkMatchByID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://americas.api.riotgames.com/lol/match/v5/matches/BR1_2744215970",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/match.json")))

	config := util.NewTestEquinoxConfig()
	config.Logger = api.Logger{
		Level:               zerolog.WarnLevel,
		Pretty:              false,
		TimeFieldFormat:     zerolog.TimeFormatUnix,
		EnableConfigLogging: true,
		EnableTimestamp:     true,
	}
	config.Retry = api.Retry{MaxRetries: 3}
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
BenchmarkMatchTimeline-16 270 4339316 ns/op 1107714 B/op 831 allocs/op
BenchmarkMatchTimeline-16 273 4409977 ns/op 1107698 B/op 831 allocs/op
BenchmarkMatchTimeline-16 278 4296051 ns/op 1107672 B/op 831 allocs/op
BenchmarkMatchTimeline-16 272 4302664 ns/op 1107766 B/op 831 allocs/op
BenchmarkMatchTimeline-16 273 4266458 ns/op 1107759 B/op 831 allocs/op
*/
func BenchmarkMatchTimeline(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://americas.api.riotgames.com/lol/match/v5/matches/BR1_2744215970/timeline",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/match.timeline.json")))

	config := util.NewTestEquinoxConfig()
	config.Logger = api.Logger{
		Level:               zerolog.WarnLevel,
		Pretty:              false,
		TimeFieldFormat:     zerolog.TimeFormatUnix,
		EnableConfigLogging: true,
		EnableTimestamp:     true,
	}
	config.Retry = api.Retry{MaxRetries: 3}
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
BenchmarkDDragonAllChampions-16 674 1739610 ns/op 143302 B/op 1494 allocs/op
BenchmarkDDragonAllChampions-16 616 1718938 ns/op 143321 B/op 1494 allocs/op
BenchmarkDDragonAllChampions-16 697 1755047 ns/op 143293 B/op 1494 allocs/op
BenchmarkDDragonAllChampions-16 698 1723216 ns/op 143269 B/op 1493 allocs/op
BenchmarkDDragonAllChampions-16 693 1729216 ns/op 143288 B/op 1494 allocs/op
*/
func BenchmarkDDragonAllChampions(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://ddragon.leagueoflegends.com/cdn/13.22.1/data/en_US/champion.json",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/champions.json")))

	config := util.NewTestEquinoxConfig()
	config.Logger = api.Logger{
		Level:               zerolog.WarnLevel,
		Pretty:              false,
		TimeFieldFormat:     zerolog.TimeFormatUnix,
		EnableConfigLogging: true,
		EnableTimestamp:     true,
	}
	config.Retry = api.Retry{MaxRetries: 3}
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
BenchmarkVALContentAllLocales-16 22 47399728 ns/op 11812547 B/op 131498 allocs/op
BenchmarkVALContentAllLocales-16 24 48412437 ns/op 11793214 B/op 131498 allocs/op
BenchmarkVALContentAllLocales-16 24 47967851 ns/op 11792628 B/op 131496 allocs/op
BenchmarkVALContentAllLocales-16 22 48317562 ns/op 11812547 B/op 131498 allocs/op
BenchmarkVALContentAllLocales-16 22 47355273 ns/op 11812547 B/op 131498 allocs/op
*/
// Probably the largest response you can get with the Riot API.
func BenchmarkVALContentAllLocales(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br.api.riotgames.com/val/content/v1/contents",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/val.content.all_locales.json")))

	config := util.NewTestEquinoxConfig()
	config.Logger = api.Logger{
		Level:               zerolog.WarnLevel,
		Pretty:              false,
		TimeFieldFormat:     zerolog.TimeFormatUnix,
		EnableConfigLogging: true,
		EnableTimestamp:     true,
	}
	config.Retry = api.Retry{MaxRetries: 3}
	client := equinox.NewClientWithConfig(config)

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.VAL.ContentV1.Content(ctx, val.BR, "")
		require.NoError(b, err)
		require.Equal(b, "release-07.10", data.Version)
	}
}
