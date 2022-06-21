package val

import "github.com/Kyagara/equinox/internal"

// Valorant endpoint URLs.
const (
	ContentURL = "/val/content/v1/contents"

	MatchByIDURL   = "/val/match/v1/matches/%s"
	MatchListURL   = "/val/match/v1/matchlists/by-puuid/%s"
	MatchRecentURL = "/val/match/v1/recent-matches/by-queue/%s"

	RankedURL = "/val/ranked/v1/leaderboards/by-act/%s"

	StatusURL = "/val/status/v1/platform-data"
)

type VALClient struct {
	internalClient *internal.InternalClient
	Content        *ContentEndpoint
	Match          *MatchEndpoint
	Ranked         *RankedEndpoint
	Status         *StatusEndpoint
}

// Returns a new VALClient using the InternalClient provided.
func NewVALClient(client *internal.InternalClient) *VALClient {
	return &VALClient{
		internalClient: client,
		Content:        &ContentEndpoint{internalClient: client},
		Match:          &MatchEndpoint{internalClient: client},
		Ranked:         &RankedEndpoint{internalClient: client},
		Status:         &StatusEndpoint{internalClient: client},
	}
}
