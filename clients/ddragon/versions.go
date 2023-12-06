package ddragon

import (
	"context"

	"github.com/Kyagara/equinox/internal"
)

type VersionEndpoint struct {
	internal *internal.Client
}

func (e *VersionEndpoint) Latest(ctx context.Context) (string, error) {
	version, err := e.internal.GetDDragonLOLVersions(ctx, "DDragon_Version_Latest")
	if err != nil {
		return "", err
	}
	return version[0], nil
}

func (e *VersionEndpoint) List(ctx context.Context) ([]string, error) {
	return e.internal.GetDDragonLOLVersions(ctx, "DDragon_Version_List")
}
