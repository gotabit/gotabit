package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagSender   = "sender"
	FlagReceiver = "receiver"
	FlagTopics   = "topics"
)

func FlagMsg() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	return fs
}

func FlagQuerySentMessages() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagSender, "", "Message sender address")

	return fs
}

func FlagQueryReceivedMessages() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagReceiver, "", "Message receiver address")
	fs.String(FlagTopics, "", "Message topics")

	return fs
}
