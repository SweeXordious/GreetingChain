package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(k Keeper) types.Querier {
	return func(ctx types.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case "hellos":
			return QueryListHello(ctx, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown query something")
		}
	}
}

func QueryListHello(ctx types.Context, k Keeper) ([]byte, error) {
	var helloList []string
	iterator := k.GetHelloIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		hello := iterator.Key()
		helloList = append(helloList, string(hello))
	}

	res, err := codec.MarshalJSONIndent(k.cdc, helloList)
	if err != nil {
		return res, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}
