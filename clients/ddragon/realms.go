package ddragon

import (
	"fmt"

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
	logger.Debug("Method executed")

	url := fmt.Sprintf(RealmURL, realm)

	var data RealmData

	err := e.internalClient.DDragonGet(url, &data, "realm", "ByName")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return &data, nil
}
