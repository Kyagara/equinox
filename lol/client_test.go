package lol_test

import (
	"testing"

	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
	"github.com/stretchr/testify/assert"
)

func TestLOLClient(t *testing.T) {
	client := lol.NewLOLClient(&internal.InternalClient{})

	assert.NotNil(t, client, "expecting non-nil LOLClient")
}
