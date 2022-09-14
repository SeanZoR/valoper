package valoper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "valoper/testutil/keeper"
	"valoper/testutil/nullify"
	"valoper/x/valoper"
	"valoper/x/valoper/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ValoperKeeper(t)
	valoper.InitGenesis(ctx, *k, genesisState)
	got := valoper.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
