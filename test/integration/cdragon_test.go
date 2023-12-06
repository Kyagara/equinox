//go:build integration
// +build integration

package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCDragonChampionByName(t *testing.T) {
	ctx := context.Background()
	version, err := client.CDragon.Version.Latest(ctx)
	require.NoError(t, err)
	champion, err := client.CDragon.Champion.ByName(ctx, version, "Lux")
	require.NoError(t, err)
	require.NotEmpty(t, champion, "expecting non-nil champion")
	require.Equal(t, "Lux", champion.Name, "expecting champion name to be Lux")
}

func TestCDragonChampionByID(t *testing.T) {
	ctx := context.Background()
	version, err := client.CDragon.Version.Latest(ctx)
	require.NoError(t, err)
	champion, err := client.CDragon.Champion.ByID(ctx, version, 223)
	require.NoError(t, err)
	require.NotEmpty(t, champion, "expecting non-nil champion")
	require.Equal(t, "Tahm Kench", champion.Name, "expecting champion name to be Tahm Kench")
}
