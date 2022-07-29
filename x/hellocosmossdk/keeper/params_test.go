package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "hello-cosmos-sdk/testutil/keeper"
	"hello-cosmos-sdk/x/hellocosmossdk/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HellocosmossdkKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
