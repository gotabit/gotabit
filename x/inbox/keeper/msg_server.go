package keeper

import (
	"context"

	"github.com/gotabit/gotabit/x/inbox/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	*Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the token MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// Send sends a new message
func (m msgServer) Send(goCtx context.Context, msg *types.MsgSend) (*types.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id, err := m.Keeper.Send(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendResponse{
		Id:      id,
		Sender:  msg.Sender,
		To:      msg.To,
		Topics:  msg.Topics,
		Message: msg.Message,
	}, nil
}
