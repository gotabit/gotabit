package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pkg/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/gotabit/gotabit/x/inbox/types"
)

var _ types.QueryServer = Keeper{}

// SentMessages returns messages sent from user
func (k Keeper) SentMessages(c context.Context, req *types.SentMessagesRequest) (*types.SentMessagesResponse, error) {
	if req == nil || len(req.Address) == 0 {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid request")
	}

	msgs, err := k.GetMsgsBySender(sdk.UnwrapSDKContext(c), req.Address)
	if err != nil {
		return nil, err
	}

	return &types.SentMessagesResponse{Messages: msgs}, nil
}

// ReceivedMessages returns messages received by user
func (k Keeper) ReceivedMessages(c context.Context, req *types.ReceivedMessagesRequest) (*types.ReceivedMessagesResponse, error) {
	if req == nil || len(req.Address) == 0 {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid request")
	}

	msgs, err := k.GetMsgsByReceiver(sdk.UnwrapSDKContext(c), req.Address, req.Topics)
	if err != nil {
		return nil, err
	}

	return &types.ReceivedMessagesResponse{Messages: msgs}, nil
}
