package keeper_test

import (
	"github.com/gotabit/gotabit/x/inbox/keeper"
	"github.com/gotabit/gotabit/x/inbox/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) SendMsg(sender sdk.AccAddress, to, topics, message string) *types.MsgSendResponse {
	msgServer := keeper.NewMsgServerImpl(&suite.app.InboxKeeper)
	resp, err := msgServer.Msg(sdk.WrapSDKContext(suite.ctx), types.NewMsgSend(
		sender.String(), to, topics, message,
	))
	suite.Require().NoError(err)
	return resp
}

func (suite *KeeperTestSuite) TestMsgServerSendMsg() {
	tests := []struct {
		testCase      string
		to            string
		topics        string
		message       string
		expectPass    bool
		expectedMsgId uint64
	}{
		{
			"send message",
			"to",
			"topic",
			"test message",
			true,
			1,
		},
		{
			"send empty message",
			"to",
			"topic",
			"",
			false,
			0,
		},
	}

	for _, tc := range tests {
		sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

		msgServer := keeper.NewMsgServerImpl(&suite.app.InboxKeeper)
		resp, err := msgServer.Msg(sdk.WrapSDKContext(suite.ctx), types.NewMsgSend(
			sender.String(), tc.to, tc.topics, tc.message,
		))
		if tc.expectPass {
			suite.Require().NoError(err)

			// test response is correct
			suite.Require().Equal(resp.Id, tc.expectedMsgId)

			// test lastMsgId is updated correctly
			lastMsgId := suite.app.InboxKeeper.GetLastMsgId(suite.ctx)
			suite.Require().Equal(lastMsgId, tc.expectedMsgId)

			// test metadataId and nftId to set correctly
			msg, err := suite.app.InboxKeeper.GetMsgById(suite.ctx, resp.Id)
			suite.Require().NoError(err)
			suite.Require().Equal(msg.Id, tc.expectedMsgId)
		} else {
			suite.Require().Error(err)
		}
	}
}
