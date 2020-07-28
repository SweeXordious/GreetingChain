package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Done: Describe your actions, these will implment the interface of `sdk.Msg`

// verify interface at compile time
var _ sdk.Msg = &MsgBuy{}

// MsgBuy - struct for unjailing jailed validator
type MsgBuy struct {
	ValidatorAddr sdk.AccAddress `json:"address" yaml:"address"` // address of the validator operator
	Hello         string         `json:"hello" yaml:"hello"`
	Price         sdk.Coins      `json:"price" yaml:"price"`
}

// NewMsgBuy creates a new MsgBuy instance
func NewMsgBuy(validatorAddr sdk.AccAddress, hello string, price sdk.Coins) MsgBuy {
	return MsgBuy{
		ValidatorAddr: validatorAddr,
		Hello:         hello,
		Price:         price,
	}
}

const BuyConst = "Buy"

// nolint
func (msg MsgBuy) Route() string { return RouterKey }
func (msg MsgBuy) Type() string  { return BuyConst }
func (msg MsgBuy) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.ValidatorAddr)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgBuy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgBuy) ValidateBasic() error {
	if msg.Price.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}
	if msg.Hello == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing hello message")
	}
	return nil
}
