package val

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = a5a3a5f5d5f2a617a56302a0afac77c745e4fd56

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"go.uber.org/zap"
)

// # Riot API Reference
//
// [val-content-v1]
//
// Note: this struct is automatically generated.
//
// [val-content-v1]: https://developer.riotgames.com/apis#val-content-v1
type ContentV1 struct {
	internalClient *internal.InternalClient
}

// Get content optionally filtered by locale
//
// # Parameters
//   - `route` - Route to query.
//   - `locale` (optional, in query)
//
// # Riot API Reference
//
// [val-content-v1.getContent]
//
// Note: this method is automatically generated.
//
// [val-content-v1.getContent]: https://developer.riotgames.com/api-methods/#val-content-v1/GET_getContent
func (e *ContentV1) Content(ctx context.Context, route PlatformRoute, locale string) (*ContentV1DTO, error) {
	logger := e.internalClient.Logger("VAL_ContentV1_Content")
	logger.Debug("Method started execution")
	equinoxReq, err := e.internalClient.Request(ctx, logger, api.RIOT_API_BASE_URL_FORMAT, http.MethodGet, route, "/val/content/v1/contents", "val-content-v1.getContent", nil)
	if err != nil {
		logger.Error("Error creating request", zap.Error(err))
		return nil, err
	}
	values := equinoxReq.Request.URL.Query()
	if locale != "" {
		values.Set("locale", locale)
	}
	equinoxReq.Request.URL.RawQuery = values.Encode()
	var data ContentV1DTO
	err = e.internalClient.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error("Error executing request", zap.Error(err))
		return nil, err
	}
	logger.Debug("Method executed successfully")
	return &data, nil
}

// # Riot API Reference
//
// [val-match-v1]
//
// Note: this struct is automatically generated.
//
// [val-match-v1]: https://developer.riotgames.com/apis#val-match-v1
type MatchV1 struct {
	internalClient *internal.InternalClient
}

// Get match by id
//
// # Parameters
//   - `route` - Route to query.
//   - `matchId` (required, in path)
//
// # Riot API Reference
//
// [val-match-v1.getMatch]
//
// Note: this method is automatically generated.
//
// [val-match-v1.getMatch]: https://developer.riotgames.com/api-methods/#val-match-v1/GET_getMatch
func (e *MatchV1) ByID(ctx context.Context, route PlatformRoute, matchId string) (*MatchV1DTO, error) {
	logger := e.internalClient.Logger("VAL_MatchV1_ByID")
	logger.Debug("Method started execution")
	equinoxReq, err := e.internalClient.Request(ctx, logger, api.RIOT_API_BASE_URL_FORMAT, http.MethodGet, route, fmt.Sprintf("/val/match/v1/matches/%v", matchId), "val-match-v1.getMatch", nil)
	if err != nil {
		logger.Error("Error creating request", zap.Error(err))
		return nil, err
	}
	var data MatchV1DTO
	err = e.internalClient.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error("Error executing request", zap.Error(err))
		return nil, err
	}
	logger.Debug("Method executed successfully")
	return &data, nil
}

// Get matchlist for games played by puuid
//
// # Parameters
//   - `route` - Route to query.
//   - `puuid` (required, in path)
//
// # Riot API Reference
//
// [val-match-v1.getMatchlist]
//
// Note: this method is automatically generated.
//
// [val-match-v1.getMatchlist]: https://developer.riotgames.com/api-methods/#val-match-v1/GET_getMatchlist
func (e *MatchV1) ListByPUUID(ctx context.Context, route PlatformRoute, puuid string) (*MatchlistV1DTO, error) {
	logger := e.internalClient.Logger("VAL_MatchV1_ListByPUUID")
	logger.Debug("Method started execution")
	equinoxReq, err := e.internalClient.Request(ctx, logger, api.RIOT_API_BASE_URL_FORMAT, http.MethodGet, route, fmt.Sprintf("/val/match/v1/matchlists/by-puuid/%v", puuid), "val-match-v1.getMatchlist", nil)
	if err != nil {
		logger.Error("Error creating request", zap.Error(err))
		return nil, err
	}
	var data MatchlistV1DTO
	err = e.internalClient.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error("Error executing request", zap.Error(err))
		return nil, err
	}
	logger.Debug("Method executed successfully")
	return &data, nil
}

// Get recent matches
//
// # Implementation Notes
//
// Returns a list of match ids that have completed in the last 10 minutes for live regions and 12 hours for the esports routing value. NA/LATAM/BR share a match history deployment. As such, recent matches will return a combined list of matches from those three regions. Requests are load balanced so you may see some inconsistencies as matches are added/removed from the list.
//
// # Parameters
//   - `route` - Route to query.
//   - `queue` (required, in path)
//
// # Riot API Reference
//
// [val-match-v1.getRecent]
//
// Note: this method is automatically generated.
//
// [val-match-v1.getRecent]: https://developer.riotgames.com/api-methods/#val-match-v1/GET_getRecent
func (e *MatchV1) Recent(ctx context.Context, route PlatformRoute, queue string) (*RecentMatchesV1DTO, error) {
	logger := e.internalClient.Logger("VAL_MatchV1_Recent")
	logger.Debug("Method started execution")
	equinoxReq, err := e.internalClient.Request(ctx, logger, api.RIOT_API_BASE_URL_FORMAT, http.MethodGet, route, fmt.Sprintf("/val/match/v1/recent-matches/by-queue/%v", queue), "val-match-v1.getRecent", nil)
	if err != nil {
		logger.Error("Error creating request", zap.Error(err))
		return nil, err
	}
	var data RecentMatchesV1DTO
	err = e.internalClient.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error("Error executing request", zap.Error(err))
		return nil, err
	}
	logger.Debug("Method executed successfully")
	return &data, nil
}

// # Riot API Reference
//
// [val-ranked-v1]
//
// Note: this struct is automatically generated.
//
// [val-ranked-v1]: https://developer.riotgames.com/apis#val-ranked-v1
type RankedV1 struct {
	internalClient *internal.InternalClient
}

// Get leaderboard for the competitive queue
//
// # Parameters
//   - `route` - Route to query.
//   - `actId` (required, in path) - Act ids can be found using the val-content API.
//   - `size` (optional, in query) - Defaults to 200. Valid values: 1 to 200.
//   - `startIndex` (optional, in query) - Defaults to 0.
//
// # Riot API Reference
//
// [val-ranked-v1.getLeaderboard]
//
// Note: this method is automatically generated.
//
// [val-ranked-v1.getLeaderboard]: https://developer.riotgames.com/api-methods/#val-ranked-v1/GET_getLeaderboard
func (e *RankedV1) Leaderboard(ctx context.Context, route PlatformRoute, actId string, size int32, startIndex int32) (*LeaderboardV1DTO, error) {
	logger := e.internalClient.Logger("VAL_RankedV1_Leaderboard")
	logger.Debug("Method started execution")
	equinoxReq, err := e.internalClient.Request(ctx, logger, api.RIOT_API_BASE_URL_FORMAT, http.MethodGet, route, fmt.Sprintf("/val/ranked/v1/leaderboards/by-act/%v", actId), "val-ranked-v1.getLeaderboard", nil)
	if err != nil {
		logger.Error("Error creating request", zap.Error(err))
		return nil, err
	}
	values := equinoxReq.Request.URL.Query()
	if size != -1 {
		values.Set("size", fmt.Sprint(size))
	}
	if startIndex != -1 {
		values.Set("startIndex", fmt.Sprint(startIndex))
	}
	equinoxReq.Request.URL.RawQuery = values.Encode()
	var data LeaderboardV1DTO
	err = e.internalClient.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error("Error executing request", zap.Error(err))
		return nil, err
	}
	logger.Debug("Method executed successfully")
	return &data, nil
}

// # Riot API Reference
//
// [val-status-v1]
//
// Note: this struct is automatically generated.
//
// [val-status-v1]: https://developer.riotgames.com/apis#val-status-v1
type StatusV1 struct {
	internalClient *internal.InternalClient
}

// Get VALORANT status for the given platform.
//
// # Parameters
//   - `route` - Route to query.
//
// # Riot API Reference
//
// [val-status-v1.getPlatformData]
//
// Note: this method is automatically generated.
//
// [val-status-v1.getPlatformData]: https://developer.riotgames.com/api-methods/#val-status-v1/GET_getPlatformData
func (e *StatusV1) Platform(ctx context.Context, route PlatformRoute) (*PlatformDataV1DTO, error) {
	logger := e.internalClient.Logger("VAL_StatusV1_Platform")
	logger.Debug("Method started execution")
	equinoxReq, err := e.internalClient.Request(ctx, logger, api.RIOT_API_BASE_URL_FORMAT, http.MethodGet, route, "/val/status/v1/platform-data", "val-status-v1.getPlatformData", nil)
	if err != nil {
		logger.Error("Error creating request", zap.Error(err))
		return nil, err
	}
	var data PlatformDataV1DTO
	err = e.internalClient.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error("Error executing request", zap.Error(err))
		return nil, err
	}
	logger.Debug("Method executed successfully")
	return &data, nil
}
