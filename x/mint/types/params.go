package types

import (
	"errors"
	"fmt"
	"strings"

	yaml "gopkg.in/yaml.v2"

	epochtypes "github.com/gotabit/gotabit/x/epochs/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys.
var (
	KeyMintDenom                            = []byte("MintDenom")
	KeyGenesisEpochProvisions               = []byte("GenesisEpochProvisions")
	KeyEpochIdentifier                      = []byte("EpochIdentifier")
	KeyReductionPeriodInEpochs              = []byte("ReductionPeriodInEpochs")
	KeyReductionFactor                      = []byte("ReductionFactor")
	KeyPoolAllocationRatio                  = []byte("PoolAllocationRatio")
	KeyMintingRewardsDistributionStartEpoch = []byte("MintingRewardsDistributionStartEpoch")
	KeyEcoFundPoolAddress                   = []byte("EcoFundPoolAddress")
	KeyDeveloperFundPoolAddress             = []byte("DeveloperFundPoolAddress")
)

// ParamTable for minting module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams returns new mint module parameters initialized to the given values.
func NewParams(
	mintDenom string, genesisEpochProvisions sdk.Dec, epochIdentifier string,
	ReductionFactor sdk.Dec, reductionPeriodInEpochs int64, distrProportions DistributionProportions,
	mintingRewardsDistributionStartEpoch int64, ecoFundPoolAddress, developerFundPoolAddress string,
) Params {
	return Params{
		MintDenom:                            mintDenom,
		GenesisEpochProvisions:               genesisEpochProvisions,
		EpochIdentifier:                      epochIdentifier,
		ReductionPeriodInEpochs:              reductionPeriodInEpochs,
		ReductionFactor:                      ReductionFactor,
		DistributionProportions:              distrProportions,
		MintingRewardsDistributionStartEpoch: mintingRewardsDistributionStartEpoch,
		EcoFundPoolAddress:                   ecoFundPoolAddress,
		DeveloperFundPoolAddress:             developerFundPoolAddress,
	}
}

// DefaultParams returns the default minting module parameters.
func DefaultParams() Params {
	return Params{
		MintDenom:               sdk.DefaultBondDenom,
		GenesisEpochProvisions:  sdk.NewDec(1000000000),
		EpochIdentifier:         "hour",                                     // 1 hour
		ReductionPeriodInEpochs: 6,                                          // 6 hours
		ReductionFactor:         sdk.NewDecWithPrec(666666666666666666, 18), // 2/3
		DistributionProportions: DistributionProportions{
			Staking:           sdk.NewDecWithPrec(4, 1), // 0.4
			EcoFundPool:       sdk.NewDecWithPrec(3, 1), // 0.3
			DeveloperFundPool: sdk.NewDecWithPrec(2, 1), // 0.2
			CommunityPool:     sdk.NewDecWithPrec(1, 1), // 0.1
		},
		MintingRewardsDistributionStartEpoch: 1,
		EcoFundPoolAddress:                   "",
		DeveloperFundPoolAddress:             "",
	}
}

// Validate validates mint module parameters. Returns nil if valid,
// error otherwise
func (p Params) Validate() error {
	if err := validateMintDenom(p.MintDenom); err != nil {
		return err
	}
	if err := validateGenesisEpochProvisions(p.GenesisEpochProvisions); err != nil {
		return err
	}
	if err := epochtypes.ValidateEpochIdentifierInterface(p.EpochIdentifier); err != nil {
		return err
	}
	if err := validateReductionPeriodInEpochs(p.ReductionPeriodInEpochs); err != nil {
		return err
	}
	if err := validateReductionFactor(p.ReductionFactor); err != nil {
		return err
	}
	if err := validateDistributionProportions(p.DistributionProportions); err != nil {
		return err
	}
	if err := validateMintingRewardsDistributionStartEpoch(p.MintingRewardsDistributionStartEpoch); err != nil {
		return err
	}
	if err := validateBech32Address(p.EcoFundPoolAddress); err != nil {
		return err
	}
	if err := validateBech32Address(p.DeveloperFundPoolAddress); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// Implements params.ParamSet.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintDenom, &p.MintDenom, validateMintDenom),
		paramtypes.NewParamSetPair(KeyGenesisEpochProvisions, &p.GenesisEpochProvisions, validateGenesisEpochProvisions),
		paramtypes.NewParamSetPair(KeyEpochIdentifier, &p.EpochIdentifier, epochtypes.ValidateEpochIdentifierInterface),
		paramtypes.NewParamSetPair(KeyReductionPeriodInEpochs, &p.ReductionPeriodInEpochs, validateReductionPeriodInEpochs),
		paramtypes.NewParamSetPair(KeyReductionFactor, &p.ReductionFactor, validateReductionFactor),
		paramtypes.NewParamSetPair(KeyPoolAllocationRatio, &p.DistributionProportions, validateDistributionProportions),
		paramtypes.NewParamSetPair(KeyMintingRewardsDistributionStartEpoch, &p.MintingRewardsDistributionStartEpoch, validateMintingRewardsDistributionStartEpoch),
		paramtypes.NewParamSetPair(KeyEcoFundPoolAddress, &p.EcoFundPoolAddress, validateBech32Address),
		paramtypes.NewParamSetPair(KeyDeveloperFundPoolAddress, &p.DeveloperFundPoolAddress, validateBech32Address),
	}
}

func validateMintDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return errors.New("mint denom cannot be blank")
	}
	if err := sdk.ValidateDenom(v); err != nil {
		return err
	}

	return nil
}

func validateGenesisEpochProvisions(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.LT(sdk.ZeroDec()) {
		return fmt.Errorf("genesis epoch provision must be non-negative")
	}

	return nil
}

func validateReductionPeriodInEpochs(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("max validators must be positive: %d", v)
	}

	return nil
}

func validateReductionFactor(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.GT(sdk.NewDec(1)) {
		return fmt.Errorf("reduction factor cannot be greater than 1")
	}

	if v.IsNegative() {
		return fmt.Errorf("reduction factor cannot be negative")
	}

	return nil
}

func validateDistributionProportions(i interface{}) error {
	v, ok := i.(DistributionProportions)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.Staking.IsNegative() {
		return errors.New("staking distribution ratio should not be negative")
	}

	if v.EcoFundPool.IsNegative() {
		return errors.New("pool incentives distribution ratio should not be negative")
	}

	if v.DeveloperFundPool.IsNegative() {
		return errors.New("developer fund pool distribution ratio should not be negative")
	}

	if v.CommunityPool.IsNegative() {
		return errors.New("community pool distribution ratio should not be negative")
	}

	totalProportions := v.Staking.Add(v.EcoFundPool).Add(v.DeveloperFundPool).Add(v.CommunityPool)

	if !totalProportions.Equal(sdk.NewDec(1)) {
		return errors.New("total distributions ratio should be 1")
	}

	return nil
}

func validateMintingRewardsDistributionStartEpoch(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return fmt.Errorf("start epoch must be non-negative")
	}

	return nil
}

func validateBech32Address(i interface{}) error {
	address, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if address != "" {
		_, err := sdk.AccAddressFromBech32(address)
		if err != nil {
			return fmt.Errorf("invalid address at %v", i)
		}
	}

	return nil
}
