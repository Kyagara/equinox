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
)

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkMatchByID-16 1940 605811 ns/op 70219 B/op 143 allocs/op
BenchmarkMatchByID-16 1929 600429 ns/op 70211 B/op 143 allocs/op
BenchmarkMatchByID-16 1981 597861 ns/op 70225 B/op 143 allocs/op
BenchmarkMatchByID-16 1868 599994 ns/op 70206 B/op 143 allocs/op
BenchmarkMatchByID-16 1902 600510 ns/op 70204 B/op 143 allocs/op
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.LOL.MatchV5.ByID(ctx, api.AMERICAS, "BR1_2744215970")
		if err != nil {
			b.Fail()
		}
		if data.Info.GameCreation != 1686266124922 {
			b.Fail()
		}
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkMatchTimeline-16 276 4267819 ns/op 1106977 B/op 826 allocs/op
BenchmarkMatchTimeline-16 271 4270326 ns/op 1107037 B/op 827 allocs/op
BenchmarkMatchTimeline-16 277 4295245 ns/op 1107033 B/op 827 allocs/op
BenchmarkMatchTimeline-16 278 4271286 ns/op 1106980 B/op 826 allocs/op
BenchmarkMatchTimeline-16 278 4328674 ns/op 1106978 B/op 826 allocs/op
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.LOL.MatchV5.Timeline(ctx, api.AMERICAS, "BR1_2744215970")
		if err != nil {
			b.Fail()
		}
		if data.Info.GameID != 2744215970 {
			b.Fail()
		}
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkDDragonAllChampions-16 709 1716146 ns/op 143037 B/op 1489 allocs/op
BenchmarkDDragonAllChampions-16 704 1688116 ns/op 143021 B/op 1489 allocs/op
BenchmarkDDragonAllChampions-16 687 1679027 ns/op 143038 B/op 1489 allocs/op
BenchmarkDDragonAllChampions-16 703 1695122 ns/op 143033 B/op 1489 allocs/op
BenchmarkDDragonAllChampions-16 692 1690217 ns/op 143030 B/op 1489 allocs/op
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.DDragon.Champion.AllChampions(ctx, "13.22.1", ddragon.EnUS)
		if err != nil {
			b.Fail()
		}
		if data["Ahri"].Name != "Ahri" {
			b.Fail()
		}
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkVALContentAllLocales-16 24 49510713 ns/op 11582650 B/op 131494 allocs/op
BenchmarkVALContentAllLocales-16 22 49665224 ns/op 11582552 B/op 131494 allocs/op
BenchmarkVALContentAllLocales-16 24 49389100 ns/op 11582479 B/op 131494 allocs/op
BenchmarkVALContentAllLocales-16 21 49429627 ns/op 11582548 B/op 131494 allocs/op
BenchmarkVALContentAllLocales-16 22 49535290 ns/op 11582562 B/op 131494 allocs/op
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.VAL.ContentV1.Content(ctx, val.BR, "")
		if err != nil {
			b.Fail()
		}
		if data.Version != "release-07.10" {
			b.Fail()
		}
	}
}
