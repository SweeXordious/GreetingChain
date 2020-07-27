package keeper

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sweexordious/helloworld/x/helloworld/types"
)

// Keeper of the helloworld store
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

// NewKeeper creates a helloworld keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		storeKey: key,
		cdc:      cdc,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Get returns the hello struct from a hello msg
func (k Keeper) GetMsg(ctx sdk.Context, helloMsg string) (types.Hello, error) {
	store := ctx.KVStore(k.storeKey)
	var hello types.Hello
	byteKey := []byte(helloMsg)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &hello)
	if err != nil {
		return hello, err
	}
	return hello, nil
}

func (k Keeper) SetMsg(ctx sdk.Context, helloStruct types.Hello) error {
	store := ctx.KVStore(k.storeKey)
	byteVal := k.cdc.MustMarshalBinaryBare(helloStruct)
	if store.Get([]byte(helloStruct.Msg)) != nil {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "The greeting you are trying to add already exists. Try buying it if it is in sale.")
	}
	store.Set([]byte(helloStruct.Msg), byteVal)
	return nil
}

func (k Keeper) delete(ctx sdk.Context, helloMsg string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(helloMsg))
}

func (k Keeper) GetHelloIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}
