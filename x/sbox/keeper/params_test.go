package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "SBox/testutil/keeper"
	"SBox/x/sbox/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.SboxKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
