package keeper

import (
	"fmt"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/gotabit/gotabit/x/epochs/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		storeKey storetypes.StoreKey
		hooks    types.EpochHooks
	}
)

// NewKeeper returns a new keeper by codec and storeKey inputs.
func NewKeeper(cdc codec.Codec, storeKey storetypes.StoreKey) *Keeper {
	return &Keeper{
		storeKey: storeKey,
	}
}

// Set the gamm hooks.
func (k *Keeper) SetHooks(eh types.EpochHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set epochs hooks twice")
	}

	k.hooks = eh

	return k
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
