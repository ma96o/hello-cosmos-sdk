package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "hello-cosmos-sdk/testutil/keeper"
	"hello-cosmos-sdk/x/hellocosmossdk/keeper"
	"hello-cosmos-sdk/x/hellocosmossdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.HellocosmossdkKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
