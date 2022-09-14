package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "valoper/testutil/keeper"
	"valoper/x/valoper/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ValoperKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
