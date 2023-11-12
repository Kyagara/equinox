package ddragon

import (
	"fmt"
	"net/http"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"go.uber.org/zap"
)

type RealmEndpoint struct {
	internalClient *internal.InternalClient
}

type RealmData struct {
	N struct {
		Item        string `json:"item"`
		Rune        string `json:"rune"`
		Mastery     string `json:"mastery"`
		Summoner    string `json:"summoner"`
		Champion    string `json:"champion"`
		ProfileIcon string `json:"profileicon"`
		Map         string `json:"map"`
		Language    string `json:"language"`
		Sticker     string `json:"sticker"`
	} `json:"n"`
	V              string      `json:"v"`
	L              string      `json:"l"`
	CDN            string      `json:"cdn"`
	DD             string      `json:"dd"`
	LG             string      `json:"lg"`
	CSS            string      `json:"css"`
	ProfileIconMax int         `json:"profileiconmax"`
	Store          interface{} `json:"store"`
}

func (e *RealmEndpoint) ByName(realm Realm) (*RealmData, error) {
	logger := e.internalClient.Logger("DDragon", "realm", "ByName")
	logger.Debug("Method started execution")
	request, err := e.internalClient.Request(api.D_DRAGON_BASE_URL_FORMAT, http.MethodGet, "", fmt.Sprintf(RealmURL, realm), nil)
	if err != nil {
		logger.Error("Error creating request", zap.Error(err))
		return nil, err
	}
	var data RealmData
	err = e.internalClient.Execute(request, &data)
	if err != nil {
		logger.Error("Error executing request", zap.Error(err))
		return nil, err
	}
	return &data, nil
}
