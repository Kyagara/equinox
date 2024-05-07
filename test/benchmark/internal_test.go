package benchmark

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/Kyagara/equinox/v2"
	"github.com/Kyagara/equinox/v2/api"
	"github.com/Kyagara/equinox/v2/internal"
	"github.com/jarcoal/httpmock"
)

// Simulating an endpoint method.
func BenchmarkInternals(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewStringResponder(200, `"response"`))

	config := equinox.DefaultConfig("RGAPI-TEST")
	client, err := internal.NewInternalClient(config, nil, nil, nil)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger := client.Logger("LOL_SummonerV4_ByPUUID")
		ctx := context.Background()
		urlComponents := []string{"https://", "kr", api.RIOT_API_BASE_URL_FORMAT, "/lol/summoner/v4/summoners/by-puuid/puuid"}
		equinoxReq, err := client.Request(ctx, logger, http.MethodGet, urlComponents, "summoner-v4.getByPUUID", nil)
		if err != nil {
			b.Fatal(err)
		}
		var data string
		err = client.Execute(ctx, equinoxReq, &data)
		if err != nil {
			b.Fatal(err)
		}
		if data != "response" {
			b.Fatalf("data != response, got: %s", data)
		}
	}
}

func BenchmarkInternalRequest(b *testing.B) {
	b.ReportAllocs()

	config := equinox.DefaultConfig("RGAPI-TEST")
	client, err := internal.NewInternalClient(config, nil, nil, nil)
	if err != nil {
		b.Fatal(err)
	}

	logger := client.Logger("LOL_SummonerV4_ByPUUID")
	ctx := context.Background()
	urlComponents := []string{"https://", "kr", api.RIOT_API_BASE_URL_FORMAT, "/lol/summoner/v4/summoners/by-puuid/puuid"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		equinoxReq, err := client.Request(ctx, logger, http.MethodGet, urlComponents, "summoner-v4.getByPUUID", nil)
		if err != nil {
			b.Fatal(err)
		}
		if equinoxReq.URL != "https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid" {
			b.Fatalf("URL != https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid, got: %s", equinoxReq.Request.URL.String())
		}
	}
}

func BenchmarkInternalExecute(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewStringResponder(200, `"response"`))

	config := equinox.DefaultConfig("RGAPI-TEST")
	client, err := internal.NewInternalClient(config, nil, nil, nil)
	if err != nil {
		b.Fatal(err)
	}

	logger := client.Logger("LOL_SummonerV4_ByPUUID")
	ctx := context.Background()

	urlComponents := []string{"https://", "kr", api.RIOT_API_BASE_URL_FORMAT, "/lol/summoner/v4/summoners/by-puuid/puuid"}
	equinoxReq, err := client.Request(ctx, logger, http.MethodGet, urlComponents, "summoner-v4.getByPUUID", nil)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data string
		err = client.Execute(ctx, equinoxReq, &data)
		if err != nil {
			b.Fatal(err)
		}
		if data != "response" {
			b.Fatalf("data != response, got: %s", data)
		}
	}
}

func BenchmarkInternalExecuteBytes(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewStringResponder(200, `"response"`))

	config := equinox.DefaultConfig("RGAPI-TEST")
	client, err := internal.NewInternalClient(config, nil, nil, nil)
	if err != nil {
		b.Fatal(err)
	}

	logger := client.Logger("LOL_SummonerV4_ByPUUID")
	ctx := context.Background()

	urlComponents := []string{"https://", "kr", api.RIOT_API_BASE_URL_FORMAT, "/lol/summoner/v4/summoners/by-puuid/", "puuid"}
	equinoxReq, err := client.Request(ctx, logger, http.MethodGet, urlComponents, "summoner-v4.getByPUUID", nil)
	if err != nil {
		b.Fatal(err)
	}

	res := []byte(`"response"`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := client.ExecuteBytes(ctx, equinoxReq)
		if err != nil {
			b.Fatal(err)
		}
		if !bytes.Equal(data, res) {
			b.Fatal("data != response")
		}
	}
}

func BenchmarkInternalGetCacheKey(b *testing.B) {
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		url, isRSO := internal.GetCacheKey(equinoxReq)
		if url != "http://example.com/path-ec2cc2a7cbc79c8d8def89cb9b9a1bccf4c2efc56a9c8063f9f4ae806f08c4d7" {
			b.Fatalf("URL != http://example.com/path-ec2cc2a7cbc79c8d8def89cb9b9a1bccf4c2efc56a9c8063f9f4ae806f08c4d7, got: %s", url)
		}
		if !isRSO {
			b.Fatal("isRSO != true")
		}
	}
}
