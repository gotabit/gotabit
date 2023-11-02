package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pkg/errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/gotabit/gotabit/x/inbox/types"
)

var _ types.QueryServer = Keeper{}

// SentMessages returns messages sent from user
func (k Keeper) SentMessages(c context.Context, req *types.SentMessagesRequest) (*types.SentMessagesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if len(req.Address) == 0 {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "empty address")
	}

	msgs := k.GetMsgsBySender(ctx, req.Address)

	return &types.SentMessagesResponse{Messages: msgs}, nil
}

// ReceivedMessages returns messages received by user
func (k Keeper) ReceivedMessages(c context.Context, req *types.ReceivedMessagesRequest) (*types.ReceivedMessagesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if len(req.Address) == 0 {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "empty address")
	}

	msgs := k.GetMsgsByReceiver(ctx, req.Address, req.Topics)

	return &types.ReceivedMessagesResponse{Messages: msgs}, nil
}
