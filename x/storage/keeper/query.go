package keeper

import (
	"SBox/x/storage/types"
)

var _ types.QueryServer = Keeper{}
