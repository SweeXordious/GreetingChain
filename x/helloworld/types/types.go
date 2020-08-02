package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Hello struct {
	Owner sdk.AccAddress `json:"owner" yaml:"owner"` // address of the hello sender
	Msg   string         `json:"msg" yaml:"msg"`
	Price sdk.Coins      `json:"price" yaml:"price"`
	Sale  bool           `json:"forSale" yaml:"forSale"`
}

// implement fmt.Stringer
func (h Hello) String() string {
	return fmt.Sprintf(`Owner: %s    Msg: %s    Price: %s     For sale: %s`,
		h.Owner,
		h.Msg,
		h.Price,
		h.Sale,
	)
}
