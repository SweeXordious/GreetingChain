package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Done: Describe your actions, these will implment the interface of `sdk.Msg`

// verify interface at compile time
var _ sdk.Msg = &MsgPropose{}

// MsgPropose - struct for unjailing jailed validator
type MsgPropose struct {
	Sender sdk.AccAddress `json:"sender" yaml:"sender"` // address of the validator operator
	Hello  string         `json:"hello" yaml:"hello"`
	Price  sdk.Coins      `json:"price" yaml:"price"`
}

// NewMsgPropose creates a new MsgPropose instance
func NewMsgPropose(validatorAddr sdk.AccAddress, hello string, price sdk.Coins) MsgPropose {
	return MsgPropose{
		Sender: validatorAddr,
		Hello:  hello,
		Price:  price,
	}
}

const BuyConst = "Buy"

// nolint
func (msg MsgPropose) Route() string { return RouterKey }
func (msg MsgPropose) Type() string  { return BuyConst }
func (msg MsgPropose) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgPropose) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgPropose) ValidateBasic() error {
	if msg.Price.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}
	if msg.Hello == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing hello message")
	}
	return nil
}
