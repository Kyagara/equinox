package benchmark_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/Kyagara/equinox/v2"
	"github.com/Kyagara/equinox/v2/api"
	"github.com/Kyagara/equinox/v2/clients/lol"
	"github.com/Kyagara/equinox/v2/ratelimit"
	"github.com/Kyagara/equinox/v2/test/util"
	"github.com/jarcoal/httpmock"
)

// The only thing that really matters here for benchmark purposes is bytes/op and allocs/op.
func BenchmarkParallelTestRateLimit(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewJsonResponderOrPanic(200, httpmock.File("../data/summoner.json")).
			HeaderSet(http.Header{
				ratelimit.APP_RATE_LIMIT_HEADER:          {"20:1,40:4"},
				ratelimit.APP_RATE_LIMIT_COUNT_HEADER:    {"1:1,1:4"},
				ratelimit.METHOD_RATE_LIMIT_HEADER:       {"1300:60"},
				ratelimit.METHOD_RATE_LIMIT_COUNT_HEADER: {"1:60"},
			}))

	config := equinox.DefaultConfig("RGAPI-TEST")
	rateLimit := ratelimit.NewInternalRateLimit(0.99, time.Second)
	client, err := equinox.NewCustomClient(config, nil, nil, rateLimit)
	if err != nil {
		b.Fatal(err)
	}

	ctx := context.Background()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.KR, "puuid")
			if err != nil {
				b.Fatal(err)
			}
			if data.ProfileIconID != 4933 {
				b.Fatalf("ProfileIconID != 4933, got %d", data.ProfileIconID)
			}
		}
	})
}

func BenchmarkParallelSummonerByPUUID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewJsonResponderOrPanic(200, httpmock.File("../data/summoner.json")))

	client := util.NewBenchmarkEquinoxClient(b)

	ctx := context.Background()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.KR, "puuid")
			if err != nil {
				b.Fatal(err)
			}
			if data.ProfileIconID != 4933 {
				b.Fatalf("ProfileIconID != 4933, got %d", data.ProfileIconID)
			}
		}
	})
}

func BenchmarkParallelRedisCachedSummonerByPUUID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewJsonResponderOrPanic(200, httpmock.File("../data/summoner.json")))

	client := util.NewBenchmarkRedisCacheEquinoxClient(b)

	ctx := context.Background()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.KR, "puuid")
			if err != nil {
				b.Fatal(err)
			}
			if data.ProfileIconID != 4933 {
				b.Fatalf("ProfileIconID != 4933, got %d", data.ProfileIconID)
			}
		}
	})
}

// This endpoint method clones apiHeaders and adds a new Authorization header.
func BenchmarkParallelSummonerByAccessToken(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://kr.api.riotgames.com/lol/summoner/v4/summoners/me",
		httpmock.NewJsonResponderOrPanic(200, httpmock.File("../data/summoner.json")))

	client := util.NewBenchmarkEquinoxClient(b)

	ctx := context.Background()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data, err := client.LOL.SummonerV4.ByAccessToken(ctx, lol.KR, "accesstoken")
			if err != nil {
				b.Fatal(err)
			}
			if data.ProfileIconID != 4933 {
				b.Fatalf("ProfileIconID != 4933, got %d", data.ProfileIconID)
			}
		}
	})
}

// This endpoint method has multiple http queries.
func BenchmarkParallelMatchListByPUUID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://asia.api.riotgames.com/lol/match/v5/matches/by-puuid/puuid/ids?count=20&queue=420&type=ranked",
		httpmock.NewJsonResponderOrPanic(200, httpmock.File("../data/match.list.json")))

	client := util.NewBenchmarkEquinoxClient(b)

	ctx := context.Background()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data, err := client.LOL.MatchV5.ListByPUUID(ctx, api.ASIA, "puuid", -1, -1, lol.SUMMONERS_RIFT_5V5_RANKED_SOLO_QUEUE, "ranked", -1, 20)
			if err != nil {
				b.Fatal(err)
			}
			if data[0] != "KR_7050905124" {
				b.Fatalf("data[0] != KR_7050905124, got %s", data[0])
			}
		}
	})
}
