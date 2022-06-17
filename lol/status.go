package lol

import (
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type StatusEndpoint struct {
	internalClient *internal.InternalClient
}

// Get League of Legends status for the given platform.
func (e *StatusEndpoint) PlatformStatus(region Region) (*api.PlatformDataDTO, error) {
	logger := e.internalClient.Logger("LOL", "lol-status-v4", "PlatformStatus")

	logger.Debug("Method executed")

	var status *api.PlatformDataDTO

	err := e.internalClient.Get(region, StatusURL, &status, "lol-status-v4", "PlatformStatus", "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return status, nil
}
