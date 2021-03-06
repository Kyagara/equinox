package riot

import (
	"fmt"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"go.uber.org/zap"
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
func (e *AccountEndpoint) PlayerActiveShard(puuid string, game api.Game) (*ActiveShardDTO, error) {
	logger := e.internalClient.Logger("Riot", "account-v1", "PlayerActiveShard")

	logger.Debug("Method executed")

	url := fmt.Sprintf(AccountActiveShardURL, game, puuid)

	var shard *ActiveShardDTO

	err := e.internalClient.Get(e.internalClient.Cluster, url, &shard, "account-v1", "PlayerActiveShard", "")

	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return shard, nil
}

// Get account by PUUID.
func (e *AccountEndpoint) ByPUUID(puuid string) (*AccountDTO, error) {
	return e.getAccount(fmt.Sprintf(AccountByPUUIDURL, puuid), "", "ByPUUID")
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
	logger := e.internalClient.Logger("Riot", "account-v1", methodName)

	logger.Debug("Method executed")

	var account *AccountDTO

	err := e.internalClient.Get(e.internalClient.Cluster, url, &account, "account-v1", methodName, accessToken)

	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return account, nil
}
