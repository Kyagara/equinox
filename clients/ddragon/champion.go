package ddragon

import (
	"context"
	"net/http"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type ChampionEndpoint struct {
	internal *internal.Client
}

// Get all champions basic information, includes stats, tags, title and blurb.
func (e *ChampionEndpoint) AllChampions(ctx context.Context, version string, language Language) (map[string]AllChampionsDataDTO, error) {
	logger := e.internal.Logger("DDragon_Champion_AllChampions")
	logger.Trace().Msg("Method started execution")
	urlComponents := []string{"https://", "", api.D_DRAGON_BASE_URL_FORMAT, "/cdn/", version, "/data/", language.String(), "/champion.json"}
	equinoxReq, err := e.internal.Request(ctx, logger, http.MethodGet, urlComponents, "", nil)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating request")
		return nil, err
	}
	var data AllChampionsDTO
	err = e.internal.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error().Err(err).Msg("Error executing request")
		return nil, err
	}
	return data.Data, nil
}

// Retrieves more information about a champion, includes skins, spells and tips.
func (e *ChampionEndpoint) ByName(ctx context.Context, version string, language Language, champion string) (*FullChampion, error) {
	logger := e.internal.Logger("DDragon_Champion_ByName")
	logger.Trace().Msg("Method started execution")
	urlComponents := []string{"https://", "", api.D_DRAGON_BASE_URL_FORMAT, "/cdn/", version, "/data/", language.String(), "/champion.json"}
	equinoxReq, err := e.internal.Request(ctx, logger, http.MethodGet, urlComponents, "", nil)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating request")
		return nil, err
	}
	var data FullChampionData
	err = e.internal.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error().Err(err).Msg("Error executing request")
		return nil, err
	}
	c := data.Data[champion]
	return &c, nil
}
