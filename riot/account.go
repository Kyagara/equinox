package riot

import (
	"fmt"
	"net/http"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type AccountEndpoint struct {
	internalClient *internal.InternalClient
}

type AccountDTO struct {
	PUUID string `json:"puuid"`
	// This field may be excluded from the response if the account doesn't have a gameName.
	GameName string `json:"gameName"`
	// This field may be excluded from the response if the account doesn't have a tagLine.
	TagLine string `json:"tagLine"`
}

type ActiveShardDTO struct {
	PUUID       string   `json:"puuid"`
	Game        api.Game `json:"game"`
	ActiveShard string   `json:"activeShard"`
}

// Get active shard for a player.
func (a *AccountEndpoint) PlayerActiveShard(PUUID string, game api.Game) (*ActiveShardDTO, error) {
	logger := a.internalClient.Logger("riot").With("endpoint", "account", "method", "PlayerActiveShard")

	url := fmt.Sprintf(AccountActiveShardURL, game, PUUID)

	var shard *ActiveShardDTO

	err := a.internalClient.Do(http.MethodGet, a.internalClient.Cluster, url, nil, &shard, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return shard, nil
}

// Get account by PUUID.
func (a *AccountEndpoint) ByPUUID(PUUID string) (*AccountDTO, error) {
	return a.getAccount(fmt.Sprintf(AccountByPUUIDURL, PUUID), "", "ByPUUID")
}

// Get account by riot ID.
func (a *AccountEndpoint) ByID(gameName string, tagLine string) (*AccountDTO, error) {
	return a.getAccount(fmt.Sprintf(AccountByRiotIDURL, gameName, tagLine), "", "ByID")
}

// Get account by access token.
func (a *AccountEndpoint) ByAccessToken(accessToken string) (*AccountDTO, error) {
	return a.getAccount(AccountByAccessTokenURL, accessToken, "ByAccessToken")
}

func (a *AccountEndpoint) getAccount(url string, accessToken string, methodName string) (*AccountDTO, error) {
	logger := a.internalClient.Logger("riot").With("endpoint", "account", "method", methodName)

	var account *AccountDTO

	err := a.internalClient.Do(http.MethodGet, a.internalClient.Cluster, url, nil, &account, accessToken)

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return account, nil
}
