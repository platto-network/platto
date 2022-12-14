package platto_test

import (
	"testing"

	keepertest "github.com/platto-network/platto/testutil/keeper"
	"github.com/platto-network/platto/testutil/nullify"
	"github.com/platto-network/platto/x/platto"
	"github.com/platto-network/platto/x/platto/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PlattoKeeper(t)
	platto.InitGenesis(ctx, *k, genesisState)
	got := platto.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
