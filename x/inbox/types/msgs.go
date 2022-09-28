package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// MsgRoute identifies transaction types
	MsgRoute = "inbox"

	TypeMsgSend = "inbox"
)

var (
	_ sdk.Msg = &MsgSend{}
)

// NewMsg - construct token issue msg.
func NewMsgSend(sender, to, topics, message string) *MsgSend {
	return &MsgSend{
		Sender:  sender,
		To:      to,
		Topics:  topics,
		Message: message,
	}
}

// Route Implements Msg.
func (mm MsgSend) Route() string { return MsgRoute }

// Type Implements Msg.
func (mm MsgSend) Type() string { return TypeMsgSend }

// ValidateBasic Implements Msg.
func (mm MsgSend) ValidateBasic() error {
	msg := &Msg{
		Sender:  mm.Sender,
		To:      mm.To,
		Topics:  mm.Topics,
		Message: mm.Message,
	}

	return msg.Validate()
}

// GetSignBytes Implements Msg.
func (mm MsgSend) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&mm)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (mm MsgSend) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(mm.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
