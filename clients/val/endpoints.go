package val

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 8096d0e7127558ddf4df50a0227b4100b5d54a2f

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Kyagara/equinox/v2/api"
	"github.com/Kyagara/equinox/v2/internal"
)

// # Riot API Reference
//
// [val-console-match-v1]
//
// [val-console-match-v1]: https://developer.riotgames.com/apis#val-console-match-v1
type ConsoleMatchV1 struct {
	internal *internal.Client
}

// Get match by id
//
// # Parameters
//   - route: Route to query.
//   - matchId
//
// # Riot API Reference
//
// [val-console-match-v1.getMatch]
//
// [val-console-match-v1.getMatch]: https://developer.riotgames.com/api-methods/#val-console-match-v1/GET_getMatch
func (endpoint *ConsoleMatchV1) ByID(ctx context.Context, route PlatformRoute, matchId string) (*ConsoleMatchMatchV1DTO, error) {
	logger := endpoint.internal.Logger("VAL_ConsoleMatchV1_ByID")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/val/match/console/v1/matches/", matchId}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "val-console-match-v1.getMatch", nil)
	if err != nil {
		return nil, err
	}
	var data ConsoleMatchMatchV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Get matchlist for games played by puuid and platform type
//
// # Parameters
//   - route: Route to query.
//   - puuid
//   - platformType
//
// # Riot API Reference
//
// [val-console-match-v1.getMatchlist]
//
// [val-console-match-v1.getMatchlist]: https://developer.riotgames.com/api-methods/#val-console-match-v1/GET_getMatchlist
func (endpoint *ConsoleMatchV1) ListByPUUID(ctx context.Context, route PlatformRoute, puuid string, platformType string) (*ConsoleMatchMatchlistV1DTO, error) {
	logger := endpoint.internal.Logger("VAL_ConsoleMatchV1_ListByPUUID")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/val/match/console/v1/matchlists/by-puuid/", puuid}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "val-console-match-v1.getMatchlist", nil)
	if err != nil {
		return nil, err
	}
	values := url.Values{}
	if platformType != "" {
		values.Set("platformType", platformType)
	}
	request.Request.URL.RawQuery = values.Encode()
	var data ConsoleMatchMatchlistV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Get recent matches
//
// # Implementation Notes
//
// Returns a list of match ids that have completed in the last 10 minutes for live regions and 12 hours for the esports routing value. NA/LATAM/BR share a match history deployment. As such, recent matches will return a combined list of matches from those three regions. Requests are load balanced so you may see some inconsistencies as matches are added/removed from the list.
//
// # Parameters
//   - route: Route to query.
//   - queue
//
// # Riot API Reference
//
// [val-console-match-v1.getRecent]
//
// [val-console-match-v1.getRecent]: https://developer.riotgames.com/api-methods/#val-console-match-v1/GET_getRecent
func (endpoint *ConsoleMatchV1) Recent(ctx context.Context, route PlatformRoute, queue string) (*ConsoleMatchRecentMatchesV1DTO, error) {
	logger := endpoint.internal.Logger("VAL_ConsoleMatchV1_Recent")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/val/match/console/v1/recent-matches/by-queue/", queue}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "val-console-match-v1.getRecent", nil)
	if err != nil {
		return nil, err
	}
	var data ConsoleMatchRecentMatchesV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// # Riot API Reference
//
// [val-content-v1]
//
// [val-content-v1]: https://developer.riotgames.com/apis#val-content-v1
type ContentV1 struct {
	internal *internal.Client
}

// Get content optionally filtered by locale
//
// # Parameters
//   - route: Route to query.
//   - locale (optional)
//
// # Riot API Reference
//
// [val-content-v1.getContent]
//
// [val-content-v1.getContent]: https://developer.riotgames.com/api-methods/#val-content-v1/GET_getContent
func (endpoint *ContentV1) Content(ctx context.Context, route PlatformRoute, locale string) (*ContentV1DTO, error) {
	logger := endpoint.internal.Logger("VAL_ContentV1_Content")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/val/content/v1/contents"}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "val-content-v1.getContent", nil)
	if err != nil {
		return nil, err
	}
	values := url.Values{}
	if locale != "" {
		values.Set("locale", locale)
	}
	request.Request.URL.RawQuery = values.Encode()
	var data ContentV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// # Riot API Reference
//
// [val-match-v1]
//
// [val-match-v1]: https://developer.riotgames.com/apis#val-match-v1
type MatchV1 struct {
	internal *internal.Client
}

// Get match by id
//
// # Parameters
//   - route: Route to query.
//   - matchId
//
// # Riot API Reference
//
// [val-match-v1.getMatch]
//
// [val-match-v1.getMatch]: https://developer.riotgames.com/api-methods/#val-match-v1/GET_getMatch
func (endpoint *MatchV1) ByID(ctx context.Context, route PlatformRoute, matchId string) (*MatchV1DTO, error) {
	logger := endpoint.internal.Logger("VAL_MatchV1_ByID")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/val/match/v1/matches/", matchId}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "val-match-v1.getMatch", nil)
	if err != nil {
		return nil, err
	}
	var data MatchV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Get matchlist for games played by puuid
//
// # Parameters
//   - route: Route to query.
//   - puuid
//
// # Riot API Reference
//
// [val-match-v1.getMatchlist]
//
// [val-match-v1.getMatchlist]: https://developer.riotgames.com/api-methods/#val-match-v1/GET_getMatchlist
func (endpoint *MatchV1) ListByPUUID(ctx context.Context, route PlatformRoute, puuid string) (*MatchlistV1DTO, error) {
	logger := endpoint.internal.Logger("VAL_MatchV1_ListByPUUID")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/val/match/v1/matchlists/by-puuid/", puuid}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "val-match-v1.getMatchlist", nil)
	if err != nil {
		return nil, err
	}
	var data MatchlistV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Get recent matches
//
// # Implementation Notes
//
// Returns a list of match ids that have completed in the last 10 minutes for live regions and 12 hours for the esports routing value. NA/LATAM/BR share a match history deployment. As such, recent matches will return a combined list of matches from those three regions. Requests are load balanced so you may see some inconsistencies as matches are added/removed from the list.
//
// # Parameters
//   - route: Route to query.
//   - queue
//
// # Riot API Reference
//
// [val-match-v1.getRecent]
//
// [val-match-v1.getRecent]: https://developer.riotgames.com/api-methods/#val-match-v1/GET_getRecent
func (endpoint *MatchV1) Recent(ctx context.Context, route PlatformRoute, queue string) (*MatchRecentMatchesV1DTO, error) {
	logger := endpoint.internal.Logger("VAL_MatchV1_Recent")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/val/match/v1/recent-matches/by-queue/", queue}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "val-match-v1.getRecent", nil)
	if err != nil {
		return nil, err
	}
	var data MatchRecentMatchesV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// # Riot API Reference
//
// [val-ranked-v1]
//
// [val-ranked-v1]: https://developer.riotgames.com/apis#val-ranked-v1
type RankedV1 struct {
	internal *internal.Client
}

// Get leaderboard for the competitive queue
//
// # Parameters
//   - route: Route to query.
//   - actId: Act ids can be found using the val-content API.
//   - size (optional): Defaults to 200. Valid values: 1 to 200.
//   - startIndex (optional): Defaults to 0.
//
// # Riot API Reference
//
// [val-ranked-v1.getLeaderboard]
//
// [val-ranked-v1.getLeaderboard]: https://developer.riotgames.com/api-methods/#val-ranked-v1/GET_getLeaderboard
func (endpoint *RankedV1) Leaderboard(ctx context.Context, route PlatformRoute, actId string, size int32, startIndex int32) (*RankedLeaderboardV1DTO, error) {
	logger := endpoint.internal.Logger("VAL_RankedV1_Leaderboard")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/val/ranked/v1/leaderboards/by-act/", actId}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "val-ranked-v1.getLeaderboard", nil)
	if err != nil {
		return nil, err
	}
	values := url.Values{}
	if size != -1 {
		values.Set("size", strconv.FormatInt(int64(size), 10))
	}
	if startIndex != -1 {
		values.Set("startIndex", strconv.FormatInt(int64(startIndex), 10))
	}
	request.Request.URL.RawQuery = values.Encode()
	var data RankedLeaderboardV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// # Riot API Reference
//
// [val-status-v1]
//
// [val-status-v1]: https://developer.riotgames.com/apis#val-status-v1
type StatusV1 struct {
	internal *internal.Client
}

// Get VALORANT status for the given platform.
//
// # Parameters
//   - route: Route to query.
//
// # Riot API Reference
//
// [val-status-v1.getPlatformData]
//
// [val-status-v1.getPlatformData]: https://developer.riotgames.com/api-methods/#val-status-v1/GET_getPlatformData
func (endpoint *StatusV1) Platform(ctx context.Context, route PlatformRoute) (*StatusPlatformDataV1DTO, error) {
	logger := endpoint.internal.Logger("VAL_StatusV1_Platform")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/val/status/v1/platform-data"}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "val-status-v1.getPlatformData", nil)
	if err != nil {
		return nil, err
	}
	var data StatusPlatformDataV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
