package riot_test

import (
	"testing"

	"github.com/Kyagara/equinox/clients/riot"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/assert"
)

func TestRiotClient(t *testing.T) {
	client := riot.NewRiotClient(&internal.InternalClient{})

	assert.NotNil(t, client, "expecting non-nil RiotClient")
}
