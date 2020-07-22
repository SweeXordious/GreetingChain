package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Done: Describe your actions, these will implment the interface of `sdk.Msg`
// verify interface at compile time
var _ sdk.Msg = &MsgGet{}

// MsgGet - struct for unjailing jailed validator
type MsgGet struct {
	ValidatorAddr sdk.ValAddress `json:"address" yaml:"address"` // address of the validator operator
	Hello         string         `json:"hello" yaml:"hello"`
}

// NewMsgGet creates a new MsgGet instance
func NewMsgGet(validatorAddr sdk.ValAddress, hello string) MsgGet {
	return MsgGet{
		ValidatorAddr: validatorAddr,
		Hello:         hello,
	}
}

const GetConst = "Get"

// nolint
func (msg MsgGet) Route() string { return RouterKey }
func (msg MsgGet) Type() string  { return GetConst }
func (msg MsgGet) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.ValidatorAddr)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgGet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgGet) ValidateBasic() error {
	if msg.Hello == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}
	return nil
}
