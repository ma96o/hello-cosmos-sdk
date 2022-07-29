package keeper

import (
	"hello-cosmos-sdk/x/hellocosmossdk/types"
)

var _ types.QueryServer = Keeper{}
