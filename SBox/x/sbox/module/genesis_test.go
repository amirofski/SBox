package sbox_test

import (
	"testing"

	keepertest "SBox/testutil/keeper"
	"SBox/testutil/nullify"
	sbox "SBox/x/sbox/module"
	"SBox/x/sbox/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SboxKeeper(t)
	sbox.InitGenesis(ctx, k, genesisState)
	got := sbox.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
