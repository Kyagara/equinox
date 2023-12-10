package ddragon

import (
	"context"
	"fmt"
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
	logger.Debug().Msg("Method started execution")
	equinoxReq, err := e.internal.Request(ctx, logger, api.D_DRAGON_BASE_URL_FORMAT, http.MethodGet, "", fmt.Sprintf(ChampionsURL, version, language), "", nil)
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
	logger.Debug().Msg("Method started execution")
	equinoxReq, err := e.internal.Request(ctx, logger, api.D_DRAGON_BASE_URL_FORMAT, http.MethodGet, "", fmt.Sprintf(ChampionURL, version, language, champion), "", nil)
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
