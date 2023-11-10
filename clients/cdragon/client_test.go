package cdragon_test

import (
	"testing"

	"github.com/Kyagara/equinox/clients/cdragon"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/require"
)

func TestDataDragonClient(t *testing.T) {
	c := &internal.InternalClient{}
	client := cdragon.NewCDragonClient(c)
	require.NotNil(t, client, "expecting non-nil CDragonClient")
}
