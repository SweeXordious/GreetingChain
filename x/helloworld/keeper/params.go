package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

/*
// TODO: Define if your module needs Parameters, if not this can be deleted

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sweexordious/./x/helloworld/types"
)

// GetParams returns the total set of helloworld parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramspace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the helloworld parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramspace.SetParamSet(ctx, &params)
}
*/
