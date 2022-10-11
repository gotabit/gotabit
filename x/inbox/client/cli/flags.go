package cli

import (
	flag "github.com/spf13/pflag"
)

// flags for inbox module tx/query commands
const (
	FlagSender   = "sender"
	FlagReceiver = "receiver"
	FlagTopics   = "topics"
)

// FlagMsg returns flags for sending msg
func FlagMsg() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	return fs
}

// FlagQuerySentMessages returns flags for querying sent messages
func FlagQuerySentMessages() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagSender, "", "Message sender address")

	return fs
}

// FlagQueryReceivedMessages returns flags for querying received messages
func FlagQueryReceivedMessages() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagReceiver, "", "Message receiver address")
	fs.String(FlagTopics, "", "Message topics")

	return fs
}
