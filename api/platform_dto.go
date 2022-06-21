package api

import "time"

// The PlatformDataDTO is used in a lot of clients.

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
	ArchiveAt         time.Time        `json:"archive_at"`
	Titles            []ContentDTO     `json:"titles"`
	UpdatedAt         time.Time        `json:"updated_at"`
	IncidentSeverity  IncidentSeverity `json:"incident_severity"`
	Platforms         []Platform       `json:"platforms"`
	Updates           []UpdateDTO      `json:"updates"`
	CreatedAt         time.Time        `json:"created_at"`
	ID                int              `json:"id"`
	MaintenanceStatus string           `json:"maintenance_status"`
}

type UpdateDTO struct {
	UpdatedAt        time.Time         `json:"updated_at"`
	Translations     []ContentDTO      `json:"translations"`
	Author           string            `json:"author"`
	Publish          bool              `json:"publish"`
	CreatedAt        time.Time         `json:"created_at"`
	ID               int               `json:"id"`
	PublishLocations []PublishLocation `json:"publish_locations"`
}
