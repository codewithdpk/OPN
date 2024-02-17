package opn_test

import (
	"testing"

	keepertest "opn/testutil/keeper"
	"opn/testutil/nullify"
	opn "opn/x/opn/module"
	"opn/x/opn/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OpnKeeper(t)
	opn.InitGenesis(ctx, k, genesisState)
	got := opn.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
