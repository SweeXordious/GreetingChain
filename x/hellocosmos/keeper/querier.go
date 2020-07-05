package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sweexordious/hellocosmos/x/hellocosmos"
	types3 "github.com/sweexordious/hellocosmos/x/hellocosmos/types"
	types2 "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier creates a new querier for hellocosmos clients.
//func NewQuerier(k Keeper) sdk.Querier {
//	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
//		switch path[0] {
//		case types.QueryParams:
//			return queryParams(ctx, k)
//			// TODO: Put the modules query routes
//		default:
//			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown hellocosmos query endpoint")
//		}
//	}
//}
//
//func queryParams(ctx sdk.Context, k Keeper) ([]byte, error) {
//	params := k.GetParams(ctx)
//
//	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params)
//	if err != nil {
//		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
//	}
//
//	return res, nil
//}

// TODO: Add the modules query functions
// They will be similar to the above one: queryParams()

func NewQuerier(k hellocosmos.Keeper) types.Querier {
	return func(ctx types.Context, path []string, req types2.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types3.QueryListHello:
			return ListHellos(ctx, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown scavenge query endpoint")
		}
	}
}

func ListHellos(ctx types.Context, k Keeper) ([]byte, error) {
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
