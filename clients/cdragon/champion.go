package cdragon

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type ChampionEndpoint struct {
	internal *internal.Client
}

// Retrieves more information about a champion, includes skins, spells and tips.
func (e *ChampionEndpoint) ByName(ctx context.Context, version string, champion string) (*ChampionData, error) {
	logger := e.internal.Logger("CDragon_Champion_ByName")
	logger.Trace().Msg("Method started execution")
	urlComponents := []string{"https://", "", api.C_DRAGON_BASE_URL_FORMAT, "/", version, "/champion/", champion, "/data"}
	equinoxReq, err := e.internal.Request(ctx, logger, http.MethodGet, urlComponents, "", nil)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating request")
		return nil, err
	}
	var data ChampionData
	err = e.internal.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error().Err(err).Msg("Error executing request")
		return nil, err
	}
	return &data, nil
}

// Retrieves more information about a champion, includes skins, spells and tips.
func (e *ChampionEndpoint) ByID(ctx context.Context, version string, id int64) (*ChampionData, error) {
	logger := e.internal.Logger("CDragon_Champion_ByID")
	logger.Trace().Msg("Method started execution")
	urlComponents := []string{"https://", "", api.C_DRAGON_BASE_URL_FORMAT, "/", version, "/champion/", strconv.FormatInt(id, 10), "/data"}
	equinoxReq, err := e.internal.Request(ctx, logger, http.MethodGet, urlComponents, "", nil)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating request")
		return nil, err
	}
	var data ChampionData
	err = e.internal.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error().Err(err).Msg("Error executing request")
		return nil, err
	}
	return &data, nil
}
