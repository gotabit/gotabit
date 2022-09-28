package testutil

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"

	msgcli "github.com/gotabit/gotabit/x/inbox/client/cli"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
)

func SendMsg(clientCtx client.Context, sender, from, to, message string, bondDenom string) (testutil.BufferWriter, error) {
	cmd := msgcli.GetCmdMsg()

	return clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		from,
		to,
		message,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, sender),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(bondDenom, sdk.NewInt(100))).String()),
	})
}
