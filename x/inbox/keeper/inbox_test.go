package keeper_test

import (
	"github.com/gotabit/gotabit/x/inbox/types"
)

func (suite *KeeperTestSuite) TestLastMsgIdGetSet() {
	// get default last msg id
	lastMsgId := suite.app.InboxKeeper.GetLastMsgId(suite.ctx)
	suite.Require().Equal(lastMsgId, uint64(0))

	// set last msg id to new value
	newMsgId := uint64(2)
	suite.app.InboxKeeper.SetLastMsgId(suite.ctx, newMsgId)

	// check last msg id update
	lastMsgId = suite.app.InboxKeeper.GetLastMsgId(suite.ctx)
	suite.Require().Equal(lastMsgId, newMsgId)
}

func (suite *KeeperTestSuite) TestMsgGetSet() {
	// get msg by not available id
	_, err := suite.app.InboxKeeper.GetMsgById(suite.ctx, 0)
	suite.Require().Error(err)

	// get sent msgs when not available
	sentMsgs := suite.app.InboxKeeper.GetMsgsBySender(suite.ctx, "gio13m350fvnk3s6y5n8ugxhmka277r0t7cw48ru47")
	suite.Require().Len(sentMsgs, 0)

	// get received msgs when not available
	receivedMsgs := suite.app.InboxKeeper.GetMsgsByReceiver(suite.ctx, "gio13m350fvnk3s6y5n8ugxhmka277r0t7cw48ru47", "")
	suite.Require().Len(receivedMsgs, 0)

	// create a new msg
	msgs := []*types.Msg{
		{
			Id:      1,
			Sender:  "gio13m350fvnk3s6y5n8ugxhmka277r0t7cw48ru47",
			To:      "gio1pwhmp2d53crcqervmv5c6h4chdnctkvjaya9vs",
			Topics:  "topic1",
			Message: "message 1",
		},
		{
			Id:      2,
			Sender:  "gio13m350fvnk3s6y5n8ugxhmka277r0t7cw48ru47",
			To:      "gio1daxjpnra6jpahzjg8e6c865hmtt7469n249ln2",
			Topics:  "topic2",
			Message: "message 2",
		},
		{
			Id:      3,
			Sender:  "gio13m350fvnk3s6y5n8ugxhmka277r0t7cw48ru47",
			To:      "gio1daxjpnra6jpahzjg8e6c865hmtt7469n249ln2",
			Topics:  "topic3",
			Message: "message 3",
		},
	}

	for _, msg := range msgs {
		suite.app.InboxKeeper.SetMsg(suite.ctx, msg)
	}

	for _, msg := range msgs {
		c, err := suite.app.InboxKeeper.GetMsgById(suite.ctx, msg.Id)
		suite.Require().NoError(err)
		suite.Require().Equal(*msg, *c)
	}

	sentMsgs = suite.app.InboxKeeper.GetMsgsBySender(suite.ctx, "gio13m350fvnk3s6y5n8ugxhmka277r0t7cw48ru47")
	suite.Require().Len(sentMsgs, 3)
	for i, msg := range msgs {
		suite.Require().Equal(*msg, *sentMsgs[i])
	}

	receivedMsgs = suite.app.InboxKeeper.GetMsgsByReceiver(suite.ctx, "gio1daxjpnra6jpahzjg8e6c865hmtt7469n249ln2", "")
	suite.Require().Len(receivedMsgs, 2)

	receivedMsgs = suite.app.InboxKeeper.GetMsgsByReceiver(suite.ctx, "gio1daxjpnra6jpahzjg8e6c865hmtt7469n249ln2", "topic2")
	suite.Require().Len(receivedMsgs, 1)
}
