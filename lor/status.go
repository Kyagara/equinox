package lor

import (
	"net/http"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type StatusEndpoint struct {
	internalClient *internal.InternalClient
}

// Get Legends of Runeterra status for the given platform.
func (e *StatusEndpoint) PlatformStatus(region Region) (*api.PlatformDataDTO, error) {
	logger := e.internalClient.Logger("lor").With("endpoint", "status", "method", "PlatformStatus")

	var status *api.PlatformDataDTO

	err := e.internalClient.Do(http.MethodGet, region, StatusURL, nil, &status, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return status, nil
}
