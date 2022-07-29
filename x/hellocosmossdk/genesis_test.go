package hellocosmossdk_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "hello-cosmos-sdk/testutil/keeper"
	"hello-cosmos-sdk/testutil/nullify"
	"hello-cosmos-sdk/x/hellocosmossdk"
	"hello-cosmos-sdk/x/hellocosmossdk/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HellocosmossdkKeeper(t)
	hellocosmossdk.InitGenesis(ctx, *k, genesisState)
	got := hellocosmossdk.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
