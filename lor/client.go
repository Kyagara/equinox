package lor

import "github.com/Kyagara/equinox/internal"

// Legends of Runeterra endpoints
const (
	StatusURL = "/lor/status/v1/platform-data"
)

type LORClient struct {
	internalClient *internal.InternalClient
	Status         *StatusEndpoint
}

// Returns a new client using the API key provided.
func NewLORClient(client *internal.InternalClient) *LORClient {
	return &LORClient{
		internalClient: client,
		Status:         &StatusEndpoint{internalClient: client},
	}
}
