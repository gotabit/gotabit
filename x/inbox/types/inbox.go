package types

import (
	"github.com/gogo/protobuf/proto"
	"gopkg.in/yaml.v2"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ proto.Message = &Msg{}
)

// NewMsg constructs a new Msg instance
func NewMsg(sender, to, topics, message string) *Msg {
	return &Msg{
		Sender:  sender,
		To:      to,
		Topics:  topics,
		Message: message,
	}
}

// GetFrom implements exported.MsgI
func (sm Msg) GetId() uint64 {
	return sm.Id
}

// GetSender implements exported.MsgI
func (sm Msg) GetSender() string {
	return sm.Sender
}

// GetDenom implements exported.MsgI
func (sm Msg) GetTo() string {
	return sm.To
}

// GetTopics implements exported.MsgI
func (sm Msg) GetTopics() string {
	return sm.Topics
}

// GetMessage implements exported.MsgI
func (sm Msg) GetMessage() string {
	return sm.Message
}

func (sm Msg) String() string {
	bz, _ := yaml.Marshal(sm)
	return string(bz)
}

func (sm Msg) Validate() error {
	if len(sm.To) == 0 {
		return sdkerrors.Wrapf(sdkerrors.Error{}, "missing from")
	}
	if len(sm.To) > 64 {
		return sdkerrors.Wrapf(sdkerrors.Error{}, "from too long")
	}

	if len(sm.Topics) == 0 {
		return sdkerrors.Wrapf(sdkerrors.Error{}, "missing to")
	}
	if len(sm.Topics) > 64 {
		return sdkerrors.Wrapf(sdkerrors.Error{}, "to too long")
	}

	if len(sm.Message) == 0 {
		return sdkerrors.Wrapf(sdkerrors.Error{}, "missing message")
	}
	if len(sm.Message) > 512 {
		return sdkerrors.Wrapf(sdkerrors.Error{}, "message too long")
	}

	return nil
}
