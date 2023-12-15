//go:build integration
// +build integration

package integration

import (
	"context"
	"testing"

	"github.com/Kyagara/equinox/clients/ddragon"
	"github.com/stretchr/testify/require"
)

// DDragon

func TestDDragonRealmByName(t *testing.T) {
	ctx := context.Background()
	realm, err := client.DDragon.Realm.ByName(ctx, ddragon.BR)
	require.NoError(t, err)
	require.NotEmpty(t, realm, "expecting non-nil realm")
	require.Equal(t, "pt_BR", realm.L, "expecting realm language to be pt_BR")
}

func TestDDragonChampionAllChampions(t *testing.T) {
	ctx := context.Background()
	version, err := client.DDragon.Version.Latest(ctx)
	require.NoError(t, err)
	champions, err := client.DDragon.Champion.AllChampions(ctx, version, ddragon.PtBR)
	require.NoError(t, err)
	require.NotEmpty(t, champions, "expecting non-nil champions")
	require.Equal(t, true, len(champions) > 1, "expecting list to have more than one champions")
	require.Equal(t, "Jarvan IV", champions["JarvanIV"].Name, "expecting champion name to be Jarvan IV")
}

func TestDDragonChampionByName(t *testing.T) {
	ctx := context.Background()
	version, err := client.DDragon.Version.Latest(ctx)
	require.NoError(t, err)
	champion, err := client.DDragon.Champion.ByName(ctx, version, ddragon.PtBR, "Lux")
	require.NoError(t, err)
	require.NotEmpty(t, champion, "expecting non-nil champion")
	require.Equal(t, "Lux", champion.Name, "expecting champion name to be Lux")
}

// CDragon

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
