package data_dragon

import (
	"github.com/Kyagara/equinox/internal"
)

// Data Dragon endpoint URLs.
const (
	VersionsURL  = "/api/versions.json"
	RealmURL     = "/realms/%s.json"
	LanguagesURL = "/cdn/languages.json"
)

type DataDragonClient struct {
	internalClient *internal.InternalClient
}

// Returns a new DataDragonClient using the InternalClient provided.
func NewDataDragonClient(client *internal.InternalClient) *DataDragonClient {
	return &DataDragonClient{
		internalClient: client,
	}
}
