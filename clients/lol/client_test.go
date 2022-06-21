package lol_test

import (
	"testing"

	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/assert"
)

func TestLOLClient(t *testing.T) {
	client := lol.NewLOLClient(&internal.InternalClient{})

	assert.NotNil(t, client, "expecting non-nil LOLClient")
}
