package lor

import "github.com/Kyagara/equinox/internal"

// Legends of Runeterra endpoint URLs.
const (
	DeckURL = "/lor/deck/v1/decks/me"

	InventoryURL = "/lor/inventory/v1/cards/me"

	MatchByIDURL = "/lor/match/v1/matches/%s"
	MatchListURL = "/lor/match/v1/matches/by-puuid/%s/ids"

	RankedURL = "/lor/ranked/v1/leaderboards"

	StatusURL = "/lor/status/v1/platform-data"
)

type LORClient struct {
	internalClient *internal.InternalClient
	Deck           *DeckEndpoint
	Inventory      *InventoryEndpoint
	Match          *MatchEndpoint
	Ranked         *RankedEndpoint
	Status         *StatusEndpoint
}

// Returns a new LORClient using the InternalClient provided.
func NewLORClient(client *internal.InternalClient) *LORClient {
	if client.IsDataDragonOnly {
		return nil
	}

	return &LORClient{
		internalClient: client,
		Deck:           &DeckEndpoint{internalClient: client},
		Inventory:      &InventoryEndpoint{internalClient: client},
		Match:          &MatchEndpoint{internalClient: client},
		Ranked:         &RankedEndpoint{internalClient: client},
		Status:         &StatusEndpoint{internalClient: client},
	}
}
