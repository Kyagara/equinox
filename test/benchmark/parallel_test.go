package benchmark_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/ddragon"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/Kyagara/equinox/test/util"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

// This revealed problems with multiple limits in a bucket, only the first bucket was being respected.
//
// The only thing that really matters here for benchmark purposes is B/op and allocs/op.
/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkParallelRateLimit-16 100 90047134 ns/op 3209 B/op 38 allocs/op
BenchmarkParallelRateLimit-16 100 90036735 ns/op 2983 B/op 37 allocs/op
BenchmarkParallelRateLimit-16 100 90033642 ns/op 3161 B/op 37 allocs/op
BenchmarkParallelRateLimit-16 100 90035373 ns/op 3055 B/op 37 allocs/op
BenchmarkParallelRateLimit-16 100 90035765 ns/op 3156 B/op 37 allocs/op
*/
func BenchmarkParallelRateLimit(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/summoner.json")).HeaderSet(http.Header{
			ratelimit.APP_RATE_LIMIT_HEADER:          {"20:1,40:4"},
			ratelimit.APP_RATE_LIMIT_COUNT_HEADER:    {"1:1,1:4"},
			ratelimit.METHOD_RATE_LIMIT_HEADER:       {"1300:60"},
			ratelimit.METHOD_RATE_LIMIT_COUNT_HEADER: {"1:60"},
		}))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
	config.RateLimit = ratelimit.NewInternalRateLimit(1.0, 1*time.Second)
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
BenchmarkParallelCachedSummonerByPUUID-16 264927 4124 ns/op 2406 B/op 12 allocs/op
BenchmarkParallelCachedSummonerByPUUID-16 248793 4197 ns/op 2489 B/op 12 allocs/op
BenchmarkParallelCachedSummonerByPUUID-16 232384 4440 ns/op 2584 B/op 12 allocs/op
BenchmarkParallelCachedSummonerByPUUID-16 291516 4244 ns/op 2290 B/op 12 allocs/op
BenchmarkParallelCachedSummonerByPUUID-16 290792 4308 ns/op 2293 B/op 12 allocs/op
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
BenchmarkParallelSummonerByPUUID-16 276248 4852 ns/op 1608 B/op 22 allocs/op
BenchmarkParallelSummonerByPUUID-16 266298 4773 ns/op 1608 B/op 22 allocs/op
BenchmarkParallelSummonerByPUUID-16 278318 4361 ns/op 1608 B/op 22 allocs/op
BenchmarkParallelSummonerByPUUID-16 267100 4356 ns/op 1608 B/op 22 allocs/op
BenchmarkParallelSummonerByPUUID-16 286087 4565 ns/op 1608 B/op 22 allocs/op
*/
func BenchmarkParallelSummonerByPUUID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/summoner.json")))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
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
BenchmarkParallelDDragonRealms-16 256857 4534 ns/op 1736 B/op 21 allocs/op
BenchmarkParallelDDragonRealms-16 260287 4796 ns/op 1736 B/op 21 allocs/op
BenchmarkParallelDDragonRealms-16 250515 4585 ns/op 1736 B/op 21 allocs/op
BenchmarkParallelDDragonRealms-16 247935 4640 ns/op 1736 B/op 21 allocs/op
BenchmarkParallelDDragonRealms-16 255441 4639 ns/op 1736 B/op 21 allocs/op
*/
func BenchmarkParallelDDragonRealms(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://ddragon.leagueoflegends.com/realms/na.json",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/realm.json")))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
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
BenchmarkParallelMatchListByPUUID-16 240942 5311 ns/op 2888 B/op 38 allocs/op
BenchmarkParallelMatchListByPUUID-16 243670 4943 ns/op 2888 B/op 38 allocs/op
BenchmarkParallelMatchListByPUUID-16 237004 4918 ns/op 2888 B/op 38 allocs/op
BenchmarkParallelMatchListByPUUID-16 249390 4986 ns/op 2888 B/op 38 allocs/op
BenchmarkParallelMatchListByPUUID-16 224121 4900 ns/op 2888 B/op 38 allocs/op
*/
func BenchmarkParallelMatchListByPUUID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://asia.api.riotgames.com/lol/match/v5/matches/by-puuid/puuid/ids?count=20&queue=420&type=ranked",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/match.list.json")))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
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
