package types

const (
	// ModuleName is the name of the module
	ModuleName = "inbox"

	// StoreKey is the string store representation
	StoreKey string = ModuleName

	// QuerierRoute is the querier route for the module
	QuerierRoute string = ModuleName

	// RouterKey is the inbox router key for the module
	RouterKey string = ModuleName
)

var (
	PrefixMsg                    = []byte{0x01}
	PrefixMsgBySender            = []byte{0x02}
	PrefixMsgByReceiver          = []byte{0x03}
	PrefixMsgByReceiverAndTopics = []byte{0x04}
	KeyLastMsgId                 = []byte{0x05}
)
