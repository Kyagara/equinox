package ddragon_test

import (
	"testing"

	"github.com/Kyagara/equinox/clients/ddragon"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/require"
)

func TestDDragonClient(t *testing.T) {
	c := &internal.InternalClient{}
	require.Equal(t, false, c.IsDataDragonOnly, "expecting non-nil DDragonClient")
	client := ddragon.NewDDragonClient(c)

	require.NotNil(t, client, "expecting non-nil DDragonClient")

	c = &internal.InternalClient{IsDataDragonOnly: true}
	require.Equal(t, true, c.IsDataDragonOnly, "expecting non-nil DDragonClient")
	client = ddragon.NewDDragonClient(c)

	require.NotNil(t, client, "expecting non-nil DDragonClient")
}
