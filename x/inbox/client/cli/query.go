package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/gotabit/gotabit/x/inbox/types"
)

// GetQueryCmd returns the query commands for the inbox module.
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                types.ModuleName,
		Short:              "Querying commands for the Inbox module",
		DisableFlagParsing: true,
	}

	queryCmd.AddCommand(
		GetCmdQuerySentMessages(),
		GetCmdQueryReceivedMessages(),
	)

	return queryCmd
}

// GetCmdQuerySentMessages returns sent messsages
func GetCmdQuerySentMessages() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "sent [flags]",
		Long: "Query messages by sender",
		Example: fmt.Sprintf(
			`$ %s query inbox sent
				--sender="gio1yx06xsqreefnhwmtu8ypd6vlatwxfqs9c2h2cq"`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			sender, err := cmd.Flags().GetString(FlagSender)
			if err != nil {
				return err
			}

			res, err := queryClient.SentMessages(context.Background(), &types.SentMessagesRequest{
				Address: sender,
			})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().AddFlagSet(FlagQuerySentMessages())
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryReceivedMessages returns received messsages
func GetCmdQueryReceivedMessages() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "received [flags]",
		Long: "Query messages by sender",
		Example: fmt.Sprintf(
			`$ %s query inbox received
				--receiver="gio1yx06xsqreefnhwmtu8ypd6vlatwxfqs9c2h2cq" --topics="topic1"`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			receiver, err := cmd.Flags().GetString(FlagReceiver)
			if err != nil {
				return err
			}

			topics, err := cmd.Flags().GetString(FlagTopics)
			if err != nil {
				return err
			}

			res, err := queryClient.ReceivedMessages(context.Background(), &types.ReceivedMessagesRequest{
				Address: receiver,
				Topics:  topics,
			})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().AddFlagSet(FlagQueryReceivedMessages())
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
