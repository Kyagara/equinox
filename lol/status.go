package lol

import (
	"net/http"
	"time"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type StatusEndpoint struct {
	internalClient *internal.InternalClient
}

type PlatformDataDTO struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Locales      []string    `json:"locales"`
	Maintenances []StatusDTO `json:"maintenances"`
	Incidents    []StatusDTO `json:"incidents"`
}

type ContentDTO struct {
	Content string `json:"content"`
	Locale  string `json:"locale"`
}

type StatusDTO struct {
	ArchiveAt time.Time    `json:"archive_at"`
	Titles    []ContentDTO `json:"titles"`
	UpdatedAt time.Time    `json:"updated_at"`
	// (Legal values: info, warning, critical)
	IncidentSeverity string `json:"incident_severity"`
	// (Legal values: windows, macos, android, ios, ps4, xbone, switch)
	Platforms         []string    `json:"platforms"`
	Updates           []UpdateDTO `json:"updates"`
	CreatedAt         time.Time   `json:"created_at"`
	ID                int         `json:"id"`
	MaintenanceStatus string      `json:"maintenance_status"`
}

type UpdateDTO struct {
	UpdatedAt    time.Time    `json:"updated_at"`
	Translations []ContentDTO `json:"translations"`
	Author       string       `json:"author"`
	Publish      bool         `json:"publish"`
	CreatedAt    time.Time    `json:"created_at"`
	ID           int          `json:"id"`
	// (Legal values: riotclient, riotstatus, game)
	PublishLocations []string `json:"publish_locations"`
}

// Get League of Legends status for the given platform.
func (c *StatusEndpoint) PlatformStatus(region api.LOLRegion) (*PlatformDataDTO, error) {
	res := PlatformDataDTO{}

	err := c.internalClient.Do(http.MethodGet, region, StatusURL, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
