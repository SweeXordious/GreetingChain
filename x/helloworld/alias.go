package helloworld

import (
	"github.com/sweexordious/helloworld/x/helloworld/keeper"
	"github.com/sweexordious/helloworld/x/helloworld/types"
)

const (
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	// QueryParams       = types.QueryParams
	QuerierRoute = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	// variable aliases
	ModuleCdc = types.ModuleCdc

	NewMsgGet     = types.NewMsgGet
	NewMsgSet     = types.NewMsgSet
	NewMsgSell    = types.NewMsgSell
	NewMsgPropose = types.NewMsgPropose
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params

	MsgGet     = types.MsgGet
	MsgSet     = types.MsgSet
	MsgPropose = types.MsgPropose
	MsgSell    = types.MsgSell
)
