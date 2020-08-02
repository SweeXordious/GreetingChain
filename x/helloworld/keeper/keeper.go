package keeper

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sweexordious/helloworld/x/helloworld/types"
)

// Keeper of the helloworld store
type Keeper struct {
	CoinKeeper bank.Keeper
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
}

// NewKeeper creates a helloworld keeper
func NewKeeper(coinKeeper bank.Keeper, cdc *codec.Codec, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		CoinKeeper: coinKeeper,
		storeKey:   key,
		cdc:        cdc,
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
	byteKey := []byte(types.GreetingPrefix + helloMsg)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &hello)
	if err != nil {
		return hello, err
	}
	return hello, nil
}

func (k Keeper) SetMsg(ctx sdk.Context, helloStruct types.Hello) error {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.GreetingPrefix + helloStruct.Msg)
	byteVal := k.cdc.MustMarshalBinaryBare(helloStruct)
	if store.Get(byteKey) != nil {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "The greeting you are trying to add already exists. Try buying it if it is in sale.")
	}
	store.Set(byteKey, byteVal)
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	sdkError := k.CoinKeeper.SendCoins(ctx, helloStruct.Owner, moduleAcct, helloStruct.Price)
	if sdkError != nil {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Random when sending money to the module.")
	}
	return nil
}

func (k Keeper) BuyMsg(ctx sdk.Context, helloStruct types.Hello) error {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.GreetingPrefix + helloStruct.Msg)
	existentMsgBytes := store.Get(byteKey)
	if existentMsgBytes == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The greeting you are trying to buy does not exist. Try creating it.")
	}
	var existentMsg types.Hello
	codec.Cdc.MustUnmarshalBinaryBare(existentMsgBytes, &existentMsg)
	if existentMsg.Price.AmountOf(types.GreetingCoinDenom).GT(helloStruct.Price.AmountOf(types.GreetingCoinDenom)) {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "The greeting you are trying to buy is more expensive ! Send the right amount.")
	}
	sdkError := k.CoinKeeper.SendCoins(ctx, helloStruct.Owner, existentMsg.Owner, helloStruct.Price)
	if sdkError != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Problem happened when sending the money.")
	}
	//helloStruct.Msg = helloStruct.Msg + "hi"
	byteVal := k.cdc.MustMarshalBinaryBare(helloStruct)
	//store.Delete(byteKey)
	store.Set(byteKey, byteVal)
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
