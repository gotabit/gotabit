package inbox

import (
	"github.com/gotabit/gotabit/x/inbox/keeper"
	"github.com/gotabit/gotabit/x/inbox/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func DefaultGenesisState() *types.GenesisState {
	return &types.GenesisState{
		Messages: []types.Msg{},
	}
}

// InitGenesis stores the genesis state
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
}

// ExportGenesis outputs the genesis state
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Messages: k.GetAllMsgs(ctx),
	}
}
