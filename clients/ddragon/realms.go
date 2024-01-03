package ddragon

import (
	"context"
	"net/http"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type RealmEndpoint struct {
	internal *internal.Client
}

func (e *RealmEndpoint) ByName(ctx context.Context, realm Realm) (*RealmData, error) {
	logger := e.internal.Logger("DDragon_Realm_ByName")
	logger.Trace().Msg("Method started execution")
	urlComponents := []string{"https://", "", api.D_DRAGON_BASE_URL_FORMAT, "/realms/", realm.String(), ".json"}
	equinoxReq, err := e.internal.Request(ctx, logger, http.MethodGet, urlComponents, "", nil)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating request")
		return nil, err
	}
	var data RealmData
	err = e.internal.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error().Err(err).Msg("Error executing request")
		return nil, err
	}
	return &data, nil
}
