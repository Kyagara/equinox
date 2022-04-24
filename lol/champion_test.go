package lol_test

import (
	"os"
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/lol"
	"github.com/stretchr/testify/assert"
)

func TestFreeChampionsRotation(t *testing.T) {
	client := equinox.NewClient(os.Getenv("RIOT_API_KEY"))

	res, err := client.LOL.Champion.FreeRotation(lol.BR1)

	assert.Nil(t, err, "expecting nil error")

	assert.NotNil(t, res, "expecting non-nil response")
}
