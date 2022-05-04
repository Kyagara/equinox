package lor

import (
	"bytes"
	"encoding/json"
	"net/http"

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
func (d *DeckEndpoint) List(region Region, accessToken string) (*[]DeckDTO, error) {
	logger := d.internalClient.Logger("lor").With("endpoint", "deck", "method", "List")

	var deck *[]DeckDTO

	err := d.internalClient.Do(http.MethodGet, region, DeckURL, nil, &deck, accessToken)

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return deck, nil
}

// Create a new deck for the calling user.
func (d *DeckEndpoint) Create(region Region, accessToken string, code string, name string) (string, error) {
	logger := d.internalClient.Logger("lor").With("endpoint", "deck", "method", "List")

	options := struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}{Code: code, Name: name}

	body, err := json.Marshal(options)

	if err != nil {
		logger.Error(err)
		return "", err
	}

	var deck api.PlainTextResponse

	err = d.internalClient.Do(http.MethodPost, region, DeckURL, bytes.NewBuffer(body), &deck, accessToken)

	if err != nil {
		logger.Warn(err)
		return "", err
	}

	return deck.Response.(string), nil
}
