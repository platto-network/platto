package keeper

import (
	"github.com/platto-network/platto/x/platto/types"
)

var _ types.QueryServer = Keeper{}
