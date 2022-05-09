package val

import (
	"fmt"
	"net/http"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type StatusEndpoint struct {
	internalClient *internal.InternalClient
}

// Get VALORANT status for the given platform.
func (e *StatusEndpoint) PlatformStatus(region Region) (*api.PlatformDataDTO, error) {
	logger := e.internalClient.Logger("VAL", "status", "PlatformStatus")

	if region == ESPORTS {
		return nil, fmt.Errorf("the region ESPORTS is not available for this method")
	}

	var status *api.PlatformDataDTO

	err := e.internalClient.Do(http.MethodGet, region, StatusURL, nil, &status, "")

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return status, nil
}
