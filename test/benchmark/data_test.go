package benchmark_test

import (
	"context"
	"testing"

	"github.com/Kyagara/equinox/v2/api"
	"github.com/Kyagara/equinox/v2/clients/val"
	"github.com/Kyagara/equinox/v2/test/util"
	"github.com/jarcoal/httpmock"
)

func BenchmarkDataMatchByID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://asia.api.riotgames.com/lol/match/v5/matches/KR_7014499581",
		httpmock.NewJsonResponderOrPanic(200, httpmock.File("../data/match.json")))

	client := util.NewBenchmarkEquinoxClient(b)

	ctx := context.Background()

	for b.Loop() {
		data, err := client.LOL.MatchV5.ByID(ctx, api.ASIA, "KR_7014499581")
		if err != nil {
			b.Fatal(err)
		}

		if data.Info.GameCreation != 1712161609888 {
			b.Fatalf("GameCreation != 1712161609888, got: %d", data.Info.GameCreation)
		}
	}
}

func BenchmarkDataMatchTimeline(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://asia.api.riotgames.com/lol/match/v5/matches/KR_7014499581/timeline",
		httpmock.NewJsonResponderOrPanic(200, httpmock.File("../data/match.timeline.json")))

	client := util.NewBenchmarkEquinoxClient(b)

	ctx := context.Background()

	for b.Loop() {
		data, err := client.LOL.MatchV5.Timeline(ctx, api.ASIA, "KR_7014499581")
		if err != nil {
			b.Fatal()
		}

		if data.Info.GameID != 7014499581 {
			b.Fatalf("GameID != 7014499581, got %d", data.Info.GameID)
		}
	}
}

// Probably the largest response you can get with the Riot API, very good for benchmarking json unmarshaling.
func BenchmarkDataVALContentAllLocales(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://na.api.riotgames.com/val/content/v1/contents",
		httpmock.NewJsonResponderOrPanic(200, httpmock.File("../data/val.content.all_locales.json")))

	client := util.NewBenchmarkEquinoxClient(b)

	ctx := context.Background()

	for b.Loop() {
		data, err := client.VAL.ContentV1.Content(ctx, val.NA, "")
		if err != nil {
			b.Fatal()
		}

		if data.Version != "release-08.08" {
			b.Fatalf("Version != release-08.08, got %s", data.Version)
		}
	}
}
