package lor

import (
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
	logger := e.internalClient.Logger("LOR", "inventory", "Cards")

	var cards *[]CardDTO

	err := e.internalClient.Get(region, InventoryURL, &cards, accessToken)

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return cards, nil
}
