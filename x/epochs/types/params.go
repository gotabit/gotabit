package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys.
var (
	KeyIdentifier              = []byte("Identifier")
	KeyStartTime               = []byte("StartTime")
	KeyDuration                = []byte("Duration")
	KeyCurrentEpoch            = []byte("CurrentEpoch")
	KeyCurrentEpochStartTime   = []byte("CurrentEpochStartTime")
	KeyEpochCountingStarted    = []byte("EpochCountingStarted")
	KeyCurrentEpochStartHeight = []byte("CurrentEpochStartHeight")
)

// ParamTable for minting module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&EpochInfo{})
}

// Implements params.ParamSet.
func (p *EpochInfo) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyIdentifier, &p.Identifier, validate),
		paramtypes.NewParamSetPair(KeyStartTime, &p.StartTime, validate),
		paramtypes.NewParamSetPair(KeyDuration, &p.Duration, validate),
		paramtypes.NewParamSetPair(KeyCurrentEpochStartTime, &p.CurrentEpochStartTime, validate),
		paramtypes.NewParamSetPair(KeyEpochCountingStarted, &p.EpochCountingStarted, validate),
		paramtypes.NewParamSetPair(KeyCurrentEpochStartHeight, &p.CurrentEpochStartHeight, validate),
	}

}

func validate(i interface{}) error {
	return nil
}
