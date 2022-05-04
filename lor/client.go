package lor

import "github.com/Kyagara/equinox/internal"

// Legends of Runeterra endpoints
const (
	MatchByIDURL = "/lor/match/v1/matches/%s"
	MatchListURL = "/lor/match/v1/matches/by-puuid/%s/ids"

	RankedURL = "/lor/ranked/v1/leaderboards"

	StatusURL = "/lor/status/v1/platform-data"
)

type LORClient struct {
	internalClient *internal.InternalClient
	Match          *MatchEndpoint
	Ranked         *RankedEndpoint
	Status         *StatusEndpoint
}

// Returns a new client using the API key provided.
func NewLORClient(client *internal.InternalClient) *LORClient {
	return &LORClient{
		internalClient: client,
		Match:          &MatchEndpoint{internalClient: client},
		Ranked:         &RankedEndpoint{internalClient: client},
		Status:         &StatusEndpoint{internalClient: client},
	}
}
