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
	GameName string `json:"gameName,omitempty"`
	// This field may be excluded from the response if the account doesn't have a tagLine.
	TagLine string `json:"tagLine,omitempty"`
}

type ActiveShardDTO struct {
	PUUID       string   `json:"puuid"`
	Game        api.Game `json:"game"`
	ActiveShard string   `json:"activeShard"`
}

// Get active shard for a player.
func (e *AccountEndpoint) PlayerActiveShard(PUUID string, game api.Game) (*ActiveShardDTO, error) {
	logger := e.internalClient.Logger("Riot", "account", "PlayerActiveShard")

	url := fmt.Sprintf(AccountActiveShardURL, game, PUUID)

	var shard *ActiveShardDTO

	err := e.internalClient.Do(http.MethodGet, e.internalClient.Cluster, url, nil, &shard, "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return shard, nil
}

// Get account by PUUID.
func (e *AccountEndpoint) ByPUUID(PUUID string) (*AccountDTO, error) {
	return e.getAccount(fmt.Sprintf(AccountByPUUIDURL, PUUID), "", "ByPUUID")
}

// Get account by riot ID.
func (e *AccountEndpoint) ByID(gameName string, tagLine string) (*AccountDTO, error) {
	return e.getAccount(fmt.Sprintf(AccountByRiotIDURL, gameName, tagLine), "", "ByID")
}

// Get account by access token.
func (e *AccountEndpoint) ByAccessToken(accessToken string) (*AccountDTO, error) {
	return e.getAccount(AccountByAccessTokenURL, accessToken, "ByAccessToken")
}

func (e *AccountEndpoint) getAccount(url string, accessToken string, methodName string) (*AccountDTO, error) {
	logger := e.internalClient.Logger("Riot", "account", methodName)

	var account *AccountDTO

	err := e.internalClient.Do(http.MethodGet, e.internalClient.Cluster, url, nil, &account, accessToken)

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return account, nil
}
