package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys.
var (
	KeyId      = []byte("Id")
	KeySender  = []byte("Sender")
	KeyTo      = []byte("To")
	KeyTopics  = []byte("Topics")
	KeyMessage = []byte("Message")
)

// ParamTable for minting module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Msg{})
}

// Implements params.ParamSet.
func (p *Msg) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyId, &p.Id, validate),
		paramtypes.NewParamSetPair(KeySender, &p.Sender, validate),
		paramtypes.NewParamSetPair(KeyTo, &p.To, validate),
		paramtypes.NewParamSetPair(KeyTopics, &p.Topics, validate),
		paramtypes.NewParamSetPair(KeyMessage, &p.Message, validate),
	}
}

func validate(i interface{}) error {
	return nil
}
