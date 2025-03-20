package benchmark_test

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/Kyagara/equinox/v2"
	"github.com/Kyagara/equinox/v2/api"
	"github.com/Kyagara/equinox/v2/cache"
	"github.com/Kyagara/equinox/v2/clients/lol"
	"github.com/Kyagara/equinox/v2/test/util"
	"github.com/jarcoal/httpmock"
)

func BenchmarkCacheDisabledSummonerByPUUID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewJsonResponderOrPanic(200, httpmock.File("../data/summoner.json")))

	client := util.NewBenchmarkEquinoxClient(b)

	ctx := context.Background()

	for b.Loop() {
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
		httpmock.NewJsonResponderOrPanic(200, httpmock.File("../data/summoner.json")))

	config := equinox.DefaultConfig("RGAPI-TEST")
	cache, err := equinox.DefaultCache()
	if err != nil {
		b.Fatal(err)
	}
	client, err := equinox.NewCustomClient(config, nil, cache, nil)
	if err != nil {
		b.Fatal(err)
	}

	ctx := context.Background()

	for b.Loop() {
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
		httpmock.NewJsonResponderOrPanic(200, httpmock.File("../data/summoner.json")))

	client := util.NewBenchmarkRedisCacheEquinoxClient(b)

	ctx := context.Background()

	for b.Loop() {
		data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.KR, "puuid")
		if err != nil {
			b.Fatal(err)
		}

		if data.ProfileIconID != 4933 {
			b.Fatalf("ProfileIconID != 4933, got %d", data.ProfileIconID)
		}
	}
}

func BenchmarkCacheGetKey(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewStringResponder(200, `"response"`))

	req := &http.Request{
		URL: &url.URL{
			Scheme: "http",
			Host:   "example.com",
			Path:   "/path",
		},
		Header: http.Header{},
	}

	equinoxReq := api.EquinoxRequest{Request: req}
	equinoxReq.URL = req.URL.String()

	// Random JWT I asked ChatGPT, its invalid, also around 300 characters longer than the access token I used for testing.
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghij")

	for b.Loop() {
		key, isRSO := cache.GetCacheKey(equinoxReq.URL, equinoxReq.Request.Header.Get("Authorization"))
		if key != "http://example.com/path-ec2cc2a7cbc79c8d8def89cb9b9a1bccf4c2efc56a9c8063f9f4ae806f08c4d7" {
			b.Fatalf("key != http://example.com/path-ec2cc2a7cbc79c8d8def89cb9b9a1bccf4c2efc56a9c8063f9f4ae806f08c4d7, got: %s", key)
		}

		if !isRSO {
			b.Fatal("isRSO != true")
		}
	}
}
