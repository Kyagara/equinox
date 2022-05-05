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
func (i *InventoryEndpoint) Cards(region Region, accessToken string) (*[]CardDTO, error) {
	logger := i.internalClient.Logger("lor").With("endpoint", "inventory", "method", "Cards")

	var cards *[]CardDTO

	err := i.internalClient.Do(http.MethodGet, region, InventoryURL, nil, &cards, accessToken)

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return cards, nil
}