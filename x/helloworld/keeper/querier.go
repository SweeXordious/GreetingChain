package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sweexordious/helloworld/x/helloworld/types"
)

// NewQuerier creates a new querier for helloworld clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case "list":
			return listHellos(ctx, k)
		case "get":
			return getHello(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown helloworld query endpoint")
		}
	}
}

func getHello(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	msg, err := k.GetMsg(ctx, path[0])
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil

}

func listHellos(ctx sdk.Context, k Keeper) ([]byte, error) {
	hellosList := make(map[string]string)
	iterator := k.GetHelloIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		hellosList[string(iterator.Key())] = string(iterator.Value())
	}
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, hellosList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
