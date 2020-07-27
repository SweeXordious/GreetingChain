package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Hello struct {
	Sender sdk.AccAddress `json:"sender" yaml:"sender"` // address of the hello sender
	Msg    string         `json:"msg" yaml:"msg"`
	Price  sdk.Coins      `json:"price" yaml:"price"`
}

// implement fmt.Stringer
func (h Hello) String() string {
	return fmt.Sprintf(`Sender: %s    Msg: %s    Price: %s`,
		h.Sender,
		h.Msg,
		h.Price,
	)
}
