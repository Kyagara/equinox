package lor

import (
	"bytes"
	"encoding/json"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
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
	logger := e.internalClient.Logger("LOR", "deck", "List")

	var deck *[]DeckDTO

	err := e.internalClient.Get(region, DeckURL, &deck, accessToken)

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return deck, nil
}

// Create a new deck for the calling user.
func (e *DeckEndpoint) Create(region Region, accessToken string, code string, name string) (string, error) {
	logger := e.internalClient.Logger("LOR", "deck", "Create")

	options := struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}{Code: code, Name: name}

	// This shouldn't fail since the values are checked before getting here
	body, _ := json.Marshal(options)

	var deck api.PlainTextResponse

	err := e.internalClient.Post(region, DeckURL, bytes.NewBuffer(body), &deck, accessToken)

	if err != nil {
		logger.Error(err)
		return "", err
	}

	return deck.Response.(string), nil
}
