package lor

import (
	"github.com/Kyagara/equinox/internal"
	"go.uber.org/zap"
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
	logger := e.internalClient.Logger("LOR", "lor-inventory-v1", "Cards")

	logger.Debug("Method executed")

	var cards *[]CardDTO

	err := e.internalClient.Get(region, InventoryURL, &cards, "lor-inventory-v1", "Cards", accessToken)

	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return cards, nil
}
