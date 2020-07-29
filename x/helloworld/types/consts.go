package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	GreetingCoinDenom = "msgcoin"
	BaseGreetingPrice = 100
)

var BaseGreetingCoin = sdk.Coins{
	sdk.Coin{
		Denom:  GreetingCoinDenom,
		Amount: sdk.NewInt(BaseGreetingPrice),
	},
}
