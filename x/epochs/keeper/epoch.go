package keeper

import (
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"

	"github.com/gotabit/gotabit/x/epochs/types"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetEpochInfo returns epoch info by identifier.
func (k Keeper) GetEpochInfo(ctx sdk.Context, identifier string) (*types.EpochInfo, error) {
	epoch := new(types.EpochInfo)

	value := ctx.KVStore(k.storeKey).Get(append(types.KeyPrefixEpoch, []byte(identifier)...))
	if value == nil {
		return nil, fmt.Errorf("epoch info not exist . identifier: %s", identifier)
	}

	if err := proto.Unmarshal(value, epoch); err != nil {
		return nil, err
	}

	return epoch, nil
}

// AddEpochInfo adds a new epoch info. Will return an error if the epoch fails validation,
// or re-uses an existing identifier.
// This method also sets the start time if left unset, and sets the epoch start height.
func (k Keeper) AddEpochInfo(ctx sdk.Context, epoch types.EpochInfo) error {
	if err := epoch.Validate(); err != nil {
		return err
	}
	// Check if identifier already exists
	if v, _ := k.GetEpochInfo(ctx, epoch.Identifier); v != nil {
		return fmt.Errorf("epoch with identifier %s already exists", epoch.Identifier)
	}

	// Initialize empty and default epoch values
	if epoch.StartTime.Equal(time.Time{}) {
		epoch.StartTime = ctx.BlockTime()
	}
	epoch.CurrentEpochStartHeight = ctx.BlockHeight()
	k.setEpochInfo(ctx, epoch)
	return nil
}

// setEpochInfo set epoch info.
func (k Keeper) setEpochInfo(ctx sdk.Context, epoch types.EpochInfo) {
	store := ctx.KVStore(k.storeKey)
	value, err := proto.Marshal(&epoch)
	if err != nil {
		panic(err)
	}

	store.Set(append(types.KeyPrefixEpoch, []byte(epoch.Identifier)...), value)
}

// DeleteEpochInfo delete epoch info.
func (k Keeper) DeleteEpochInfo(ctx sdk.Context, identifier string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(append(types.KeyPrefixEpoch, []byte(identifier)...))
}

// IterateEpochInfo iterate through epochs.
func (k Keeper) IterateEpochInfo(ctx sdk.Context, fn func(index int64, epochInfo types.EpochInfo) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := storetypes.KVStorePrefixIterator(store, types.KeyPrefixEpoch)
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		epoch := types.EpochInfo{}
		err := proto.Unmarshal(iterator.Value(), &epoch)
		if err != nil {
			panic(err)
		}
		stop := fn(i, epoch)

		if stop {
			break
		}
		i++
	}
}

// AllEpochInfos iterate through epochs to return all epochs info.
func (k Keeper) AllEpochInfos(ctx sdk.Context) []types.EpochInfo {
	epochs := []types.EpochInfo{}
	k.IterateEpochInfo(ctx, func(index int64, epochInfo types.EpochInfo) (stop bool) {
		epochs = append(epochs, epochInfo)
		return false
	})
	return epochs
}

// NumBlocksSinceEpochStart returns the number of blocks since the epoch started.
// if the epoch started on block N, then calling this during block N (after BeforeEpochStart)
// would return 0.
// Calling it any point in block N+1 (assuming the epoch doesn't increment) would return 1.
func (k Keeper) NumBlocksSinceEpochStart(ctx sdk.Context, identifier string) (int64, error) {
	epoch, err := k.GetEpochInfo(ctx, identifier)
	if err != nil {
		return 0, err
	}

	if epoch == nil {
		return 0, fmt.Errorf("epoch with identifier %s not found", identifier)
	}

	return ctx.BlockHeight() - epoch.CurrentEpochStartHeight, nil
}
