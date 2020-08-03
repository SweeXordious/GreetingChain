package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// verify interface at compile time
var _ sdk.Msg = &MsgSell{}

// MsgSell - struct for unjailing jailed validator
type MsgSell struct {
	Sender sdk.AccAddress `json:"sender" yaml:"sender"` // address of the validator operator
	Hello  string         `json:"hello" yaml:"hello"`
}

// NewMsgSell creates a new MsgSell instance
func NewMsgSell(validatorAddr sdk.AccAddress, hello string) MsgSell {
	return MsgSell{
		Sender: validatorAddr,
		Hello:  hello,
	}
}

const SellConst = "Buy"

// nolint
func (msg MsgSell) Route() string { return RouterKey }
func (msg MsgSell) Type() string  { return SellConst }
func (msg MsgSell) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgSell) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgSell) ValidateBasic() error {
	if msg.Hello == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing hello message")
	}
	return nil
}
