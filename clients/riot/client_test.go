package riot_test

import (
	"testing"

	"github.com/Kyagara/equinox/clients/riot"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/require"
)

func TestRiotClient(t *testing.T) {
	c := &internal.InternalClient{}
	require.Equal(t, false, c.IsDataDragonOnly, "expecting non-nil RiotClient")
	client := riot.NewRiotClient(c)

	require.NotNil(t, client, "expecting non-nil RiotClient")

	c = &internal.InternalClient{IsDataDragonOnly: true}
	require.Equal(t, true, c.IsDataDragonOnly, "expecting non-nil RiotClient")
	client = riot.NewRiotClient(c)

	require.Nil(t, client, "expecting nil RiotClient")
}
