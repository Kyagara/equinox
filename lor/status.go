package lor

import (
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type StatusEndpoint struct {
	internalClient *internal.InternalClient
}

// Get Legends of Runeterra status for the given platform.
func (e *StatusEndpoint) PlatformStatus(region Region) (*api.PlatformDataDTO, error) {
	logger := e.internalClient.Logger("LOR", "status", "PlatformStatus")

	var status *api.PlatformDataDTO

	err := e.internalClient.Get(region, StatusURL, &status, "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return status, nil
}
