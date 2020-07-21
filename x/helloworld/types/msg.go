package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Describe your actions, these will implment the interface of `sdk.Msg`

// verify interface at compile time
var _ sdk.Msg = &MsgHello{}

// MsgHello - struct for unjailing jailed validator
type MsgHello struct {
	Greeting      string         `json:"greeting" yaml:"greeting"`
	ValidatorAddr sdk.ValAddress `json:"address" yaml:"address"` // address of the validator operator
}

// NewMsgHello creates a new MsgHello instance
func NewMsgHello(validatorAddr sdk.ValAddress, greet string) MsgHello {
	return MsgHello{
		Greeting:      greet,
		ValidatorAddr: validatorAddr,
	}
}

const HelloConst = "Hello"

// nolint
func (msg MsgHello) Route() string { return RouterKey }
func (msg MsgHello) Type() string  { return HelloConst }
func (msg MsgHello) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.ValidatorAddr)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgHello) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgHello) ValidateBasic() error {
	if msg.ValidatorAddr.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing validator address")
	}
	if msg.Greeting == "" {
		return sdkerrors.Wrap(sdkerrors.Error{}, "Greeting cannot be empty")
	}
	return nil
}
