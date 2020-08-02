package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	GreetingCoinDenom = "msgcoin"
	BaseGreetingPrice = 100
	GreetingPrefix    = "greeting-"
	ProposalPrefix    = "proposal-"
)

var BaseGreetingCoin = sdk.Coins{
	sdk.Coin{
		Denom:  GreetingCoinDenom,
		Amount: sdk.NewInt(BaseGreetingPrice),
	},
}
