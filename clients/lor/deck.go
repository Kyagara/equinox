package lor

import (
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"go.uber.org/zap"
)

type DeckEndpoint struct {
	internalClient *internal.InternalClient
}

type DeckDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// Get a list of the calling user's decks.
func (e *DeckEndpoint) List(region Region, accessToken string) (*[]DeckDTO, error) {
	logger := e.internalClient.Logger("LOR", "lor-deck-v1", "List")

	logger.Debug("Method executed")

	var deck *[]DeckDTO

	err := e.internalClient.Get(region, DeckURL, &deck, "lor-deck-v1", "List", accessToken)

	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return deck, nil
}

// Create a new deck for the calling user.
func (e *DeckEndpoint) Create(region Region, accessToken string, code string, name string) (string, error) {
	logger := e.internalClient.Logger("LOR", "lor-deck-v1", "Create")

	logger.Debug("Method executed")

	options := struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}{Code: code, Name: name}

	var deck *api.PlainTextResponse

	err := e.internalClient.Post(region, DeckURL, options, &deck, "lor-deck-v1", "Create", accessToken)

	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return "", err
	}

	return deck.Response.(string), nil
}
