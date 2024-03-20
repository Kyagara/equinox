package lor

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 339cc5986ca34480f2ecf815246cade7105a897a

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

// # Riot API Reference
//
// [lor-deck-v1]
//
// [lor-deck-v1]: https://developer.riotgames.com/apis#lor-deck-v1
type DeckV1 struct {
	internal *internal.Client
}

// Create a new deck for the calling user.
//
// # Parameters
//   - route : Route to query.
//   - Authorization
//
// # Riot API Reference
//
// [lor-deck-v1.createDeck]
//
// [lor-deck-v1.createDeck]: https://developer.riotgames.com/api-methods/#lor-deck-v1/POST_createDeck
func (e *DeckV1) CreateDeck(ctx context.Context, route api.RegionalRoute, body *DeckNewDeckV1DTO, authorization string) (string, error) {
	logger := e.internal.Logger("LOR_DeckV1_CreateDeck")
	logger.Trace().Msg("Method started execution")
	if authorization == "" {
		return "", fmt.Errorf("'authorization' header is required")
	}
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/lor/deck/v1/decks/me"}
	equinoxReq, err := e.internal.Request(ctx, logger, http.MethodPost, urlComponents, "lor-deck-v1.createDeck", body)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating request")
		return "", err
	}
	equinoxReq.Request.Header = equinoxReq.Request.Header.Clone()
	equinoxReq.Request.Header.Add("Authorization", authorization)
	var data string
	err = e.internal.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error().Err(err).Msg("Error executing request")
		return "", err
	}
	return data, nil
}

// Get a list of the calling user's decks.
//
// # Parameters
//   - route : Route to query.
//   - Authorization
//
// # Riot API Reference
//
// [lor-deck-v1.getDecks]
//
// [lor-deck-v1.getDecks]: https://developer.riotgames.com/api-methods/#lor-deck-v1/GET_getDecks
func (e *DeckV1) Decks(ctx context.Context, route api.RegionalRoute, authorization string) ([]DeckV1DTO, error) {
	logger := e.internal.Logger("LOR_DeckV1_Decks")
	logger.Trace().Msg("Method started execution")
	if authorization == "" {
		return nil, fmt.Errorf("'authorization' header is required")
	}
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/lor/deck/v1/decks/me"}
	equinoxReq, err := e.internal.Request(ctx, logger, http.MethodGet, urlComponents, "lor-deck-v1.getDecks", nil)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating request")
		return nil, err
	}
	equinoxReq.Request.Header = equinoxReq.Request.Header.Clone()
	equinoxReq.Request.Header.Add("Authorization", authorization)
	var data []DeckV1DTO
	err = e.internal.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error().Err(err).Msg("Error executing request")
		return nil, err
	}
	return data, nil
}

// # Riot API Reference
//
// [lor-inventory-v1]
//
// [lor-inventory-v1]: https://developer.riotgames.com/apis#lor-inventory-v1
type InventoryV1 struct {
	internal *internal.Client
}

// Return a list of cards owned by the calling user.
//
// # Parameters
//   - route : Route to query.
//   - Authorization
//
// # Riot API Reference
//
// [lor-inventory-v1.getCards]
//
// [lor-inventory-v1.getCards]: https://developer.riotgames.com/api-methods/#lor-inventory-v1/GET_getCards
func (e *InventoryV1) Cards(ctx context.Context, route api.RegionalRoute, authorization string) ([]InventoryCardV1DTO, error) {
	logger := e.internal.Logger("LOR_InventoryV1_Cards")
	logger.Trace().Msg("Method started execution")
	if authorization == "" {
		return nil, fmt.Errorf("'authorization' header is required")
	}
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/lor/inventory/v1/cards/me"}
	equinoxReq, err := e.internal.Request(ctx, logger, http.MethodGet, urlComponents, "lor-inventory-v1.getCards", nil)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating request")
		return nil, err
	}
	equinoxReq.Request.Header = equinoxReq.Request.Header.Clone()
	equinoxReq.Request.Header.Add("Authorization", authorization)
	var data []InventoryCardV1DTO
	err = e.internal.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error().Err(err).Msg("Error executing request")
		return nil, err
	}
	return data, nil
}

// # Riot API Reference
//
// [lor-match-v1]
//
// [lor-match-v1]: https://developer.riotgames.com/apis#lor-match-v1
type MatchV1 struct {
	internal *internal.Client
}

// Get match by id
//
// # Parameters
//   - route : Route to query.
//   - matchId
//
// # Riot API Reference
//
// [lor-match-v1.getMatch]
//
// [lor-match-v1.getMatch]: https://developer.riotgames.com/api-methods/#lor-match-v1/GET_getMatch
func (e *MatchV1) ByID(ctx context.Context, route api.RegionalRoute, matchId string) (*MatchV1DTO, error) {
	logger := e.internal.Logger("LOR_MatchV1_ByID")
	logger.Trace().Msg("Method started execution")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/lor/match/v1/matches/", matchId}
	equinoxReq, err := e.internal.Request(ctx, logger, http.MethodGet, urlComponents, "lor-match-v1.getMatch", nil)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating request")
		return nil, err
	}
	var data MatchV1DTO
	err = e.internal.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error().Err(err).Msg("Error executing request")
		return nil, err
	}
	return &data, nil
}

// Get a list of match ids by PUUID
//
// # Parameters
//   - route : Route to query.
//   - puuid
//
// # Riot API Reference
//
// [lor-match-v1.getMatchIdsByPUUID]
//
// [lor-match-v1.getMatchIdsByPUUID]: https://developer.riotgames.com/api-methods/#lor-match-v1/GET_getMatchIdsByPUUID
func (e *MatchV1) ListByPUUID(ctx context.Context, route api.RegionalRoute, puuid string) ([]string, error) {
	logger := e.internal.Logger("LOR_MatchV1_ListByPUUID")
	logger.Trace().Msg("Method started execution")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/lor/match/v1/matches/by-puuid/", puuid, "/ids"}
	equinoxReq, err := e.internal.Request(ctx, logger, http.MethodGet, urlComponents, "lor-match-v1.getMatchIdsByPUUID", nil)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating request")
		return nil, err
	}
	var data []string
	err = e.internal.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error().Err(err).Msg("Error executing request")
		return nil, err
	}
	return data, nil
}

// # Riot API Reference
//
// [lor-ranked-v1]
//
// [lor-ranked-v1]: https://developer.riotgames.com/apis#lor-ranked-v1
type RankedV1 struct {
	internal *internal.Client
}

// Get the players in Master tier.
//
// # Parameters
//   - route : Route to query.
//
// # Riot API Reference
//
// [lor-ranked-v1.getLeaderboards]
//
// [lor-ranked-v1.getLeaderboards]: https://developer.riotgames.com/api-methods/#lor-ranked-v1/GET_getLeaderboards
func (e *RankedV1) Leaderboards(ctx context.Context, route api.RegionalRoute) (*RankedLeaderboardV1DTO, error) {
	logger := e.internal.Logger("LOR_RankedV1_Leaderboards")
	logger.Trace().Msg("Method started execution")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/lor/ranked/v1/leaderboards"}
	equinoxReq, err := e.internal.Request(ctx, logger, http.MethodGet, urlComponents, "lor-ranked-v1.getLeaderboards", nil)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating request")
		return nil, err
	}
	var data RankedLeaderboardV1DTO
	err = e.internal.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error().Err(err).Msg("Error executing request")
		return nil, err
	}
	return &data, nil
}

// # Riot API Reference
//
// [lor-status-v1]
//
// [lor-status-v1]: https://developer.riotgames.com/apis#lor-status-v1
type StatusV1 struct {
	internal *internal.Client
}

// Get Legends of Runeterra status for the given platform.
//
// # Parameters
//   - route : Route to query.
//
// # Riot API Reference
//
// [lor-status-v1.getPlatformData]
//
// [lor-status-v1.getPlatformData]: https://developer.riotgames.com/api-methods/#lor-status-v1/GET_getPlatformData
func (e *StatusV1) Platform(ctx context.Context, route api.RegionalRoute) (*StatusPlatformDataV1DTO, error) {
	logger := e.internal.Logger("LOR_StatusV1_Platform")
	logger.Trace().Msg("Method started execution")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/lor/status/v1/platform-data"}
	equinoxReq, err := e.internal.Request(ctx, logger, http.MethodGet, urlComponents, "lor-status-v1.getPlatformData", nil)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating request")
		return nil, err
	}
	var data StatusPlatformDataV1DTO
	err = e.internal.Execute(ctx, equinoxReq, &data)
	if err != nil {
		logger.Error().Err(err).Msg("Error executing request")
		return nil, err
	}
	return &data, nil
}
