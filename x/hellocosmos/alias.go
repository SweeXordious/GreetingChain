package hellocosmos

import (
	"github.com/sweexordious/hellocosmos/x/hellocosmos/keeper"
	"github.com/sweexordious/hellocosmos/x/hellocosmos/types"
)

const (
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	// QueryParams = types.QueryParams
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

	NewMsgHelloCosmos = types.NewMsgHelloCosmos
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params

	MsgHelloCosmos = types.MsgHelloCosmos
)
