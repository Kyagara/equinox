package val

import "github.com/Kyagara/equinox/internal"

// Riot endpoints
const (
	ContentURL = "/val/content/v1/contents"

	RankedURL = "/val/ranked/v1/leaderboards/by-act/%s"

	StatusURL = "/val/status/v1/platform-data"
)

type VALClient struct {
	internalClient *internal.InternalClient

	Content *ContentEndpoint
	Ranked  *RankedEndpoint
	Status  *StatusEndpoint
}

// Returns a new client using the API key provided.
func NewVALClient(client *internal.InternalClient) *VALClient {
	return &VALClient{
		internalClient: client,
		Content:        &ContentEndpoint{internalClient: client},
		Ranked:         &RankedEndpoint{internalClient: client},
		Status:         &StatusEndpoint{internalClient: client},
	}
}
