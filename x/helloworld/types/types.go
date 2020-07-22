package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

type Hello struct {
	Sender sdk.AccAddress `json:"sender" yaml:"sender"` // address of the hello sender
	Msg    string         `json:"msg" yaml:"msg"`
}

// implement fmt.Stringer
func (h Hello) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Sender: %s
	Msg: %s`,
		h.Sender,
		h.Msg,
	))
}
