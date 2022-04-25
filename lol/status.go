package lol

import (
	"fmt"
	"time"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type StatusEndpoint struct {
	internalClient *internal.InternalClient
}

type PlatformDataDto struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Locales      []string    `json:"locales"`
	Maintenances []StatusDto `json:"maintenances"`
	Incidents    []StatusDto `json:"incidents"`
}

type ContentDto struct {
	Content string `json:"content"`
	Locale  string `json:"locale"`
}

type UpdateDto struct {
	UpdatedAt        string       `json:"updated_at"`
	Translations     []ContentDto `json:"translations"`
	Author           string       `json:"author"`
	Publish          bool         `json:"publish"`
	CreatedAt        time.Time    `json:"created_at"`
	ID               int          `json:"id"`
	PublishLocations []string     `json:"publish_locations"`
}

type StatusDto struct {
	ArchiveAt         time.Time    `json:"archive_at"`
	Titles            []ContentDto `json:"titles"`
	UpdatedAt         time.Time    `json:"updated_at"`
	IncidentSeverity  string       `json:"incident_severity"`
	Platforms         []string     `json:"platforms"`
	Updates           []UpdateDto  `json:"updates"`
	CreatedAt         time.Time    `json:"created_at"`
	ID                int          `json:"id"`
	MaintenanceStatus string       `json:"maintenance_status"`
}

// Get Status
func (c *StatusEndpoint) Status(region api.Region) (*PlatformDataDto, error) {
	res := PlatformDataDto{}

	if err := c.internalClient.SendRequest("GET", fmt.Sprintf(api.BaseURLFormat, region), StatusEndpointURL, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
