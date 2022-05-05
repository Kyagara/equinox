package lor

import (
	"net/http"

	"github.com/Kyagara/equinox/internal"
)

type InventoryEndpoint struct {
	internalClient *internal.InternalClient
}

type CardDTO struct {
	Code  string `json:"code"`
	Count string `json:"count"`
}

// Return a list of cards owned by the calling user.
func (e *InventoryEndpoint) Cards(region Region, accessToken string) (*[]CardDTO, error) {
	logger := e.internalClient.Logger("lor").With("endpoint", "inventory", "method", "Cards")

	var cards *[]CardDTO

	err := e.internalClient.Do(http.MethodGet, region, InventoryURL, nil, &cards, accessToken)

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return cards, nil
}
