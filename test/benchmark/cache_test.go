package benchmark_test

import (
	"context"
	"testing"
	"time"

	"github.com/Kyagara/equinox/v2"
	"github.com/Kyagara/equinox/v2/cache"
	"github.com/Kyagara/equinox/v2/clients/lol"
	"github.com/Kyagara/equinox/v2/test/util"
	"github.com/jarcoal/httpmock"
	"github.com/redis/go-redis/v9"
)

func BenchmarkCacheDisabledSummonerByPUUID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/summoner.json")))

	client := util.NewBenchmarkEquinoxClient(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.KR, "puuid")
		if err != nil {
			b.Fatal(err)
		}
		if data.ProfileIconID != 4933 {
			b.Fatalf("ProfileIconID != 4933, got %d", data.ProfileIconID)
		}
	}
}

func BenchmarkCacheBigCacheSummonerByPUUID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/summoner.json")))

	client, err := equinox.NewClient("RGAPI-TEST")
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.KR, "puuid")
		if err != nil {
			b.Fatal(err)
		}
		if data.ProfileIconID != 4933 {
			b.Fatalf("ProfileIconID != 4933, got %d", data.ProfileIconID)
		}
	}
}

func BenchmarkCacheRedisSummonerByPUUID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/summoner.json")))

	ctx := context.Background()
	redisConfig := &redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	}

	cache, err := cache.NewRedis(ctx, redisConfig, 4*time.Minute)
	if err != nil {
		b.Fatal(err)
	}

	config := equinox.DefaultConfig("RGAPI-TEST")
	client, err := equinox.NewCustomClient(config, nil, cache, nil)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.KR, "puuid")
		if err != nil {
			b.Fatal(err)
		}
		if data.ProfileIconID != 4933 {
			b.Fatalf("ProfileIconID != 4933, got %d", data.ProfileIconID)
		}
	}
}
