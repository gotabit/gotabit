package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/gotabit/gotabit/x/inbox/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/version"
)

// NewTxCmd returns the transaction commands for the Inbox module.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Inbox transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetCmdSend(),
	)

	return txCmd
}

// GetCmdSend sends new message
func GetCmdSend() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "send TO TOPICS MESSAGE [flags]",
		Long: "Send message to other user with topics",
		Example: fmt.Sprintf(
			`$ %s tx inbox send gio13m350fvnk3s6y5n8ugxhmka277r0t7cw48ru47 "greeting" "Hello there!"`,
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			if len(args) != 3 {
				return sdkerrors.Wrapf(sdkerrors.Error{}, "invalid args length")
			}

			to := args[0]
			if err != nil {
				return err
			}
			topics := args[1]
			if err != nil {
				return err
			}
			message := args[2]
			if err != nil {
				return err
			}

			msg := types.NewMsgSend(clientCtx.GetFromAddress().String(), to, topics, message)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagMsg())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
