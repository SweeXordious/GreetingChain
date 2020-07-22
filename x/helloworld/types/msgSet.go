package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Done: Describe your actions, these will implment the interface of `sdk.Msg`

// verify interface at compile time
var _ sdk.Msg = &MsgSet{}

// MsgSet - struct for unjailing jailed validator
type MsgSet struct {
	ValidatorAddr sdk.ValAddress `json:"address" yaml:"address"` // address of the validator operator
	Sender        sdk.AccAddress `json:"sender" yaml:"sender"`
	Hello         string         `json:"hello" yaml:"hello"`
}

// NewMsgSet creates a new MsgSet instance
func NewMsgSet(sender sdk.AccAddress, hello string) MsgSet {
	return MsgSet{
		Sender: sender,
		Hello:  hello,
	}
}

const SetConst = "Set"

// nolint
func (msg MsgSet) Route() string { return RouterKey }
func (msg MsgSet) Type() string  { return SetConst }
func (msg MsgSet) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.ValidatorAddr)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgSet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgSet) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}
	if msg.Hello == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing hello message")
	}
	return nil
}
