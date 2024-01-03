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
BenchmarkParallelRateLimit-16 100 100043976 ns/op 3007 B/op 33 allocs/op
BenchmarkParallelRateLimit-16 100 100038252 ns/op 2894 B/op 32 allocs/op
BenchmarkParallelRateLimit-16 100 100059004 ns/op 3062 B/op 33 allocs/op
BenchmarkParallelRateLimit-16 100 100040742 ns/op 2934 B/op 32 allocs/op
BenchmarkParallelRateLimit-16 100 100045917 ns/op 2868 B/op 32 allocs/op
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
	config.RateLimit = ratelimit.NewInternalRateLimit(0.99, 1*time.Second)
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
BenchmarkParallelCachedSummonerByPUUID-16 278990 3961 ns/op 2246 B/op 8 allocs/op
BenchmarkParallelCachedSummonerByPUUID-16 300795 3995 ns/op 2159 B/op 8 allocs/op
BenchmarkParallelCachedSummonerByPUUID-16 284236 4077 ns/op 2224 B/op 8 allocs/op
BenchmarkParallelCachedSummonerByPUUID-16 300584 4184 ns/op 2159 B/op 8 allocs/op
BenchmarkParallelCachedSummonerByPUUID-16 276688 4668 ns/op 2256 B/op 8 allocs/op
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
BenchmarkParallelSummonerByPUUID-16 232520 4927 ns/op 1525 B/op 18 allocs/op
BenchmarkParallelSummonerByPUUID-16 243302 4908 ns/op 1525 B/op 18 allocs/op
BenchmarkParallelSummonerByPUUID-16 224158 4896 ns/op 1525 B/op 18 allocs/op
BenchmarkParallelSummonerByPUUID-16 231002 4952 ns/op 1526 B/op 18 allocs/op
BenchmarkParallelSummonerByPUUID-16 230145 4934 ns/op 1527 B/op 18 allocs/op
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
BenchmarkParallelDDragonRealms-16 227568 5248 ns/op 1705 B/op 18 allocs/op
BenchmarkParallelDDragonRealms-16 229098 5130 ns/op 1705 B/op 18 allocs/op
BenchmarkParallelDDragonRealms-16 200659 5458 ns/op 1706 B/op 18 allocs/op
BenchmarkParallelDDragonRealms-16 221775 5168 ns/op 1705 B/op 18 allocs/op
BenchmarkParallelDDragonRealms-16 234630 5219 ns/op 1704 B/op 18 allocs/op
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
BenchmarkParallelMatchListByPUUID-16 177630 6101 ns/op 2823 B/op 34 allocs/op
BenchmarkParallelMatchListByPUUID-16 193851 6562 ns/op 2823 B/op 34 allocs/op
BenchmarkParallelMatchListByPUUID-16 168358 6158 ns/op 2822 B/op 34 allocs/op
BenchmarkParallelMatchListByPUUID-16 190045 5902 ns/op 2823 B/op 34 allocs/op
BenchmarkParallelMatchListByPUUID-16 182811 6005 ns/op 2825 B/op 34 allocs/op
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
