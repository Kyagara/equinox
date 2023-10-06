package lol_test

import (
	"testing"

	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/require"
)

func TestLOLClient(t *testing.T) {
	c := &internal.InternalClient{}
	require.Equal(t, false, c.IsDataDragonOnly, "expecting non-nil LOLClient")
	client := lol.NewLOLClient(c)

	require.NotNil(t, client, "expecting non-nil LOLClient")

	c = &internal.InternalClient{IsDataDragonOnly: true}
	require.Equal(t, true, c.IsDataDragonOnly, "expecting non-nil LOLClient")
	client = lol.NewLOLClient(c)

	require.Nil(t, client, "expecting nil LOLClient")
}
