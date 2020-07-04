package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Describe your actions, these will implment the interface of `sdk.Msg`

// verify interface at compile time
var _ sdk.Msg = &MsgHelloCosmos{}

// Msg<Action> - struct for unjailing jailed validator
type MsgHelloCosmos struct {
	Sender    sdk.AccAddress `json:"sender" yaml:"sender"`     // address of the account "sending" the greeting
	Recipient sdk.AccAddress `json:"receiver" yaml:"receiver"` // address of the account "receiving" the greeting
	Body      string         `json:"body" yaml:"body"`         // string body of the greeting
}

// NewMsg<Action> creates a new Msg<Action> instance
func NewMsgHelloCosmos(sender sdk.AccAddress, body string, receiver sdk.AccAddress) MsgHelloCosmos {
	return MsgHelloCosmos{
		Recipient: receiver,
		Sender:    sender,
		Body:      body,
	}
}

const HelloCosmosConst = "HelloCosmos"

// nolint
func (msg MsgHelloCosmos) Route() string { return RouterKey }
func (msg MsgHelloCosmos) Type() string  { return "hello" }
func (msg MsgHelloCosmos) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgHelloCosmos) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgHelloCosmos) ValidateBasic() error {
	if msg.Recipient.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Recipient.String())
	}
	if len(msg.Sender) == 0 || len(msg.Body) == 0 || len(msg.Recipient) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Sender, Recipient and/or Body cannot be empty")
	}
	return nil
}
