package keeper_test

import (
	"github.com/gotabit/gotabit/x/inbox/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestGRPCSentMsgs() {
	// send messages
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	msg1 := suite.SendMsg(sender, "bob", "topic1", "message1")
	msg2 := suite.SendMsg(sender, "carol", "topic1", "message2")
	msg3 := suite.SendMsg(sender, "bob", "topic2", "message3")

	tests := []struct {
		testCase     string
		sender       string
		expectedMsgs []*types.MsgSendResponse
	}{
		{
			"query alice's sent messages",
			sender.String(),
			[]*types.MsgSendResponse{
				{
					Id:      1,
					Sender:  msg1.Sender,
					To:      msg1.To,
					Topics:  msg1.Topics,
					Message: msg1.Message,
				},
				{
					Id:      2,
					Sender:  msg2.Sender,
					To:      msg2.To,
					Topics:  msg2.Topics,
					Message: msg2.Message,
				},
				{
					Id:      3,
					Sender:  msg3.Sender,
					To:      msg3.To,
					Topics:  msg3.Topics,
					Message: msg3.Message,
				},
			},
		},
		{
			"query bob's sent messages",
			"bob",
			[]*types.MsgSendResponse{},
		},
	}

	for _, tc := range tests {
		resp := suite.app.InboxKeeper.GetMsgsBySender(suite.ctx, tc.sender)
		suite.Require().Equal(len(resp), len(tc.expectedMsgs))
		for i, msg := range resp {
			suite.Require().Equal(msg.Id, tc.expectedMsgs[i].Id)
			suite.Require().Equal(msg.Sender, tc.expectedMsgs[i].Sender)
			suite.Require().Equal(msg.To, tc.expectedMsgs[i].To)
			suite.Require().Equal(msg.Topics, tc.expectedMsgs[i].Topics)
			suite.Require().Equal(msg.Message, tc.expectedMsgs[i].Message)
		}
	}
}

func (suite *KeeperTestSuite) TestGRPCReceivedMsgs() {
	// send messages
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	msg1 := suite.SendMsg(sender, "bob", "topic1", "message1")
	msg2 := suite.SendMsg(sender, "carol", "topic1", "message2")
	msg3 := suite.SendMsg(sender, "bob", "topic2", "message3")

	tests := []struct {
		testCase     string
		receiver     string
		topics       string
		expectedMsgs []*types.MsgSendResponse
	}{
		{
			"query bob's received messages",
			"bob",
			"",
			[]*types.MsgSendResponse{
				{
					Id:      1,
					Sender:  msg1.Sender,
					To:      msg1.To,
					Topics:  msg1.Topics,
					Message: msg1.Message,
				},
				{
					Id:      3,
					Sender:  msg3.Sender,
					To:      msg3.To,
					Topics:  msg3.Topics,
					Message: msg3.Message,
				},
			},
		},
		{
			"query bob's received messages",
			"bob",
			"topic1",
			[]*types.MsgSendResponse{
				{
					Id:      1,
					Sender:  msg1.Sender,
					To:      msg1.To,
					Topics:  msg1.Topics,
					Message: msg1.Message,
				},
			},
		},
		{
			"query carol's received messages",
			"carol",
			"",
			[]*types.MsgSendResponse{
				{
					Id:      2,
					Sender:  msg2.Sender,
					To:      msg2.To,
					Topics:  msg2.Topics,
					Message: msg2.Message,
				},
			},
		},
	}

	for _, tc := range tests {
		resp := suite.app.InboxKeeper.GetMsgsByReceiver(suite.ctx, tc.receiver, tc.topics)
		suite.Require().Equal(len(resp), len(tc.expectedMsgs))
		for i, msg := range resp {
			suite.Require().Equal(msg.Id, tc.expectedMsgs[i].Id)
			suite.Require().Equal(msg.Sender, tc.expectedMsgs[i].Sender)
			suite.Require().Equal(msg.To, tc.expectedMsgs[i].To)
			suite.Require().Equal(msg.Topics, tc.expectedMsgs[i].Topics)
			suite.Require().Equal(msg.Message, tc.expectedMsgs[i].Message)
		}
	}
}
