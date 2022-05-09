package riot

import (
	"github.com/Kyagara/equinox/internal"
)

// Riot endpoint URLs.
const (
	AccountActiveShardURL   = "/riot/account/v1/active-shards/by-game/%s/by-puuid/%s"
	AccountByPUUIDURL       = "/riot/account/v1/accounts/by-puuid/%s"
	AccountByRiotIDURL      = "/riot/account/v1/accounts/by-riot-id/%s/%s"
	AccountByAccessTokenURL = "/riot/account/v1/accounts/me"
)

type RiotClient struct {
	internalClient *internal.InternalClient
	Account        *AccountEndpoint
}

// Returns a new RiotClient using the InternalClient provided.
func NewRiotClient(client *internal.InternalClient) *RiotClient {
	return &RiotClient{
		internalClient: client,

		Account: &AccountEndpoint{internalClient: client},
	}
}
