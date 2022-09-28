package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/gotabit/gotabit/x/mint/types"
)

// InitGenesis new mint genesis.
func (k Keeper) InitGenesis(ctx sdk.Context, data *types.GenesisState) {
	if data == nil {
		panic("nil mint genesis state")
	}

	data.Minter.EpochProvisions = data.Params.GenesisEpochProvisions
	k.SetMinter(ctx, data.Minter)
	k.SetParams(ctx, data.Params)

	// The call to GetModuleAccount creates a module account if it does not exist.
	k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)

	k.SetLastReductionEpochNum(ctx, data.ReductionStartedEpoch)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	minter := k.GetMinter(ctx)
	params := k.GetParams(ctx)

	lastHalvenEpoch := k.GetLastReductionEpochNum(ctx)
	return types.NewGenesisState(minter, params, lastHalvenEpoch)
}
