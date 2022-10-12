package keeper_test

import (
	"testing"

	testkeeper "github.com/platto-network/platto/testutil/keeper"
	"github.com/platto-network/platto/x/platto/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PlattoKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
