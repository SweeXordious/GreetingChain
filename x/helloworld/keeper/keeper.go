package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/sweexordious/helloworld/x/helloworld/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/log"
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
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Problem when sending money to the module.")
	}
	return nil
}

func (k Keeper) ProposeMsg(ctx sdk.Context, helloStruct types.Hello) error {
	store := ctx.KVStore(k.storeKey)
	msgByteKey := []byte(types.GreetingPrefix + helloStruct.Msg)
	existentMsgBytes := store.Get(msgByteKey)
	if existentMsgBytes == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The greeting you are trying to buy does not exist. Try creating it.")
	}
	byteVal := k.cdc.MustMarshalBinaryBare(helloStruct)
	proposalByteKey := []byte(types.ProposalPrefix + helloStruct.Msg + "-" + helloStruct.Owner.String())
	store.Set(proposalByteKey, byteVal)
	return nil
}

func (k Keeper) SellMsg(ctx sdk.Context, helloStruct types.Hello) error {
	store := ctx.KVStore(k.storeKey)
	msgByteKey := []byte(types.GreetingPrefix + helloStruct.Msg)
	existentMsgBytes := store.Get(msgByteKey)
	if existentMsgBytes == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The greeting you are trying to sell does not exist.")
	}
	var existentMsg types.Hello
	codec.Cdc.MustUnmarshalBinaryBare(existentMsgBytes, &existentMsg)

	proposalsIterator := k.GetProposalSpecificHelloIterator(ctx, helloStruct.Msg)
	var bestPrice = sdk.NewInt(0)
	var bestProposal types.Hello
	for ; proposalsIterator.Valid(); proposalsIterator.Next() {
		var currentProposal types.Hello
		codec.Cdc.MustUnmarshalBinaryBare(proposalsIterator.Value(), &currentProposal)
		if currentProposal.Price.AmountOf(types.GreetingCoinDenom).GTE(bestPrice) {
			bestPrice = currentProposal.Price.AmountOf(types.GreetingCoinDenom)
			bestProposal = currentProposal
		}
	}
	if bestPrice.Int64() == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrPanic, "No proposals exists. Try again in the future")
	}

	sdkError := k.CoinKeeper.SendCoins(ctx, bestProposal.Owner, existentMsg.Owner, bestProposal.Price)
	if sdkError != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Problem happened when sending the money.")
	}
	existentMsg.Price = bestProposal.Price
	existentMsg.Owner = bestProposal.Owner
	byteVal := k.cdc.MustMarshalBinaryBare(existentMsg)
	key := []byte(types.GreetingPrefix + helloStruct.Msg)
	store.Set(key, byteVal)
	return nil
}

func (k Keeper) delete(ctx sdk.Context, helloMsg string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(helloMsg))
}

func (k Keeper) GetAllHelloIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

func (k Keeper) GetMsgHelloIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.GreetingPrefix))
}

func (k Keeper) GetProposalHelloIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.ProposalPrefix))
}

func (k Keeper) GetProposalSpecificHelloIterator(ctx sdk.Context, msg string) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.ProposalPrefix+msg))
}

// Add specific message proposals list
