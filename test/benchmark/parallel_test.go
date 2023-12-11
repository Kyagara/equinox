package benchmark_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/ddragon"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/Kyagara/equinox/test/util"
	"github.com/jarcoal/httpmock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

// This revealed problems with multiple limits in a bucket, only the first bucket was being respected.
func BenchmarkParallelRateLimit(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/summoner.json")).HeaderSet(http.Header{
			ratelimit.APP_RATE_LIMIT_HEADER:       {"20:1,50:3,80:5"},
			ratelimit.APP_RATE_LIMIT_COUNT_HEADER: {"1:1,1:3,1:5"},
		}))

	config := util.NewTestEquinoxConfig()
	config.LogLevel = zerolog.WarnLevel
	config.Retries = 3
	config.RateLimit = ratelimit.NewInternalRateLimit()

	client := equinox.NewClientWithConfig(config)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
			require.NoError(b, err)
			require.Equal(b, "Phanes", data.Name)
		}
	})
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkParallelCachedSummonerByPUUID-16 261210 4524 ns/op 2568 B/op 17 allocs/op
BenchmarkParallelCachedSummonerByPUUID-16 260030 4563 ns/op 2574 B/op 17 allocs/op
BenchmarkParallelCachedSummonerByPUUID-16 263236 4537 ns/op 2558 B/op 17 allocs/op
BenchmarkParallelCachedSummonerByPUUID-16 264544 4840 ns/op 2552 B/op 17 allocs/op
BenchmarkParallelCachedSummonerByPUUID-16 256568 4907 ns/op 2591 B/op 17 allocs/op
*/
func BenchmarkParallelCachedSummonerByPUUID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/summoner.json")))

	client, err := equinox.NewClient("RGAPI-TEST")
	require.NoError(b, err)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
			require.NoError(b, err)
			require.Equal(b, "Phanes", data.Name)
		}
	})
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkParallelSummonerByPUUID-16 208351 5293 ns/op 1770 B/op 27 allocs/op
BenchmarkParallelSummonerByPUUID-16 195762 5755 ns/op 1772 B/op 27 allocs/op
BenchmarkParallelSummonerByPUUID-16 206599 5323 ns/op 1772 B/op 27 allocs/op
BenchmarkParallelSummonerByPUUID-16 195020 5432 ns/op 1773 B/op 27 allocs/op
BenchmarkParallelSummonerByPUUID-16 208340 5557 ns/op 1771 B/op 27 allocs/op
*/
func BenchmarkParallelSummonerByPUUID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/summoner.json")))

	config := util.NewTestEquinoxConfig()
	config.LogLevel = zerolog.WarnLevel
	config.Retries = 3

	client := equinox.NewClientWithConfig(config)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
			require.NoError(b, err)
			require.Equal(b, "Phanes", data.Name)
		}
	})
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkParallelDDragonRealms-16 216250 5642 ns/op 1935 B/op 26 allocs/op
BenchmarkParallelDDragonRealms-16 215046 5630 ns/op 1935 B/op 26 allocs/op
BenchmarkParallelDDragonRealms-16 212302 5614 ns/op 1935 B/op 26 allocs/op
BenchmarkParallelDDragonRealms-16 193950 6363 ns/op 1936 B/op 26 allocs/op
BenchmarkParallelDDragonRealms-16 198546 5908 ns/op 1935 B/op 26 allocs/op
*/
func BenchmarkParallelDDragonRealms(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://ddragon.leagueoflegends.com/realms/na.json",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/realm.json")))

	config := util.NewTestEquinoxConfig()
	config.LogLevel = zerolog.WarnLevel
	config.Retries = 3

	client := equinox.NewClientWithConfig(config)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			data, err := client.DDragon.Realm.ByName(ctx, ddragon.NA)
			require.NoError(b, err)
			require.Equal(b, "13.24.1", data.V)
		}
	})
}

/*
goos: linux
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkParallelMatchListByPUUID-16 187948 6462 ns/op 3235 B/op 44 allocs/op
BenchmarkParallelMatchListByPUUID-16 170966 6547 ns/op 3240 B/op 44 allocs/op
BenchmarkParallelMatchListByPUUID-16 186074 6534 ns/op 3238 B/op 44 allocs/op
BenchmarkParallelMatchListByPUUID-16 176089 6524 ns/op 3239 B/op 44 allocs/op
BenchmarkParallelMatchListByPUUID-16 160318 6569 ns/op 3243 B/op 44 allocs/op
*/
func BenchmarkParallelMatchListByPUUID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://asia.api.riotgames.com/lol/match/v5/matches/by-puuid/puuid/ids?count=20&queue=420&type=ranked",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/match.list.json")))

	config := util.NewTestEquinoxConfig()
	config.LogLevel = zerolog.WarnLevel
	config.Retries = 3

	client := equinox.NewClientWithConfig(config)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			data, err := client.LOL.MatchV5.ListByPUUID(ctx, api.ASIA, "puuid", -1, -1, 420, "ranked", -1, 20)
			require.NoError(b, err)
			require.Equal(b, "KR_6841523755", data[0])
		}
	})
}
