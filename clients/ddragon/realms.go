package ddragon

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"go.uber.org/zap"
)

type RealmEndpoint struct {
	internal *internal.InternalClient
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
	V              string `json:"v"`
	L              string `json:"l"`
	CDN            string `json:"cdn"`
	DD             string `json:"dd"`
	LG             string `json:"lg"`
	CSS            string `json:"css"`
	ProfileIconMax int    `json:"profileiconmax"`
	Store          any    `json:"store,omitempty"`
}

func (e *RealmEndpoint) ByName(ctx context.Context, realm Realm) (*RealmData, error) {
	logger := e.internal.Logger("DDragon_Realm_ByName")
	logger.Debug("Method started execution")
	equinoxReq, err := e.internal.Request(ctx, logger, api.D_DRAGON_BASE_URL_FORMAT, http.MethodGet, "", fmt.Sprintf(RealmURL, realm), "", nil)
	if err != nil {
		logger.Error("Error creating request", zap.Error(err))
		return nil, err
	}
	var data RealmData
	err = e.internal.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error("Error executing request", zap.Error(err))
		return nil, err
	}
	return &data, nil
}
