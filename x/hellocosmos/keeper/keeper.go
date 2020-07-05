package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/bank"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sweexordious/hellocosmos/x/hellocosmos/types"
)

// Keeper of the hellocosmos store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
	CoinKeeper bank.Keeper
}

// NewKeeper creates a hellocosmos keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, coinKeeper bank.Keeper) Keeper {
	keeper := Keeper{
		storeKey:   key,
		cdc:        cdc,
		CoinKeeper: coinKeeper,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Get returns the pubkey from the adddress-pubkey relation
func (k Keeper) GetGreet(ctx sdk.Context, key string) (types.MsgHelloCosmos, error) {
	store := ctx.KVStore(k.storeKey)
	var hello types.MsgHelloCosmos
	byteKey := []byte(key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &hello)
	if err != nil {
		return hello, err
	}
	return hello, nil
}

func (k Keeper) setGreet(ctx sdk.Context, key string, msg types.MsgHelloCosmos) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(msg)
	store.Set([]byte(key), bz)
}

func (k Keeper) delete(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(key))
}

func (k Keeper) GetHelloIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(""))
}
