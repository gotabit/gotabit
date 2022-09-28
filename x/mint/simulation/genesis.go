package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/gotabit/gotabit/x/mint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

// Simulation parameter constants.
const (
	epochProvisionsKey         = "genesis_epoch_provisions"
	reductionFactorKey         = "reduction_factor"
	reductionPeriodInEpochsKey = "reduction_period_in_epochs"

	mintingRewardsDistributionStartEpochKey = "minting_rewards_distribution_start_epoch"

	epochIdentifier = "day"
	maxInt64        = int(^uint(0) >> 1)

	ecoFundPoolAddress       = "gio1d4ysgq9ljs2k6hhafhuy9mnstuunzeqd800vlk75jkemdrn4r0ks7ezpjk"
	developerFundPoolAddress = "gio14zd54ruj9zu9lyz8puv3rp4tnts8xezjntq9et"
)

var distributionProportions = types.DistributionProportions{
	Staking:           sdk.NewDecWithPrec(25, 2),
	EcoFundPool:       sdk.NewDecWithPrec(45, 2),
	DeveloperFundPool: sdk.NewDecWithPrec(25, 2),
	CommunityPool:     sdk.NewDecWithPrec(0o5, 2),
}

// RandomizedGenState generates a random GenesisState for mint.
func RandomizedGenState(simState *module.SimulationState) {
	var epochProvisions sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, epochProvisionsKey, &epochProvisions, simState.Rand,
		func(r *rand.Rand) { epochProvisions = genEpochProvisions(r) },
	)

	var reductionFactor sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, reductionFactorKey, &reductionFactor, simState.Rand,
		func(r *rand.Rand) { reductionFactor = genReductionFactor(r) },
	)

	var reductionPeriodInEpochs int64
	simState.AppParams.GetOrGenerate(
		simState.Cdc, reductionPeriodInEpochsKey, &reductionPeriodInEpochs, simState.Rand,
		func(r *rand.Rand) { reductionPeriodInEpochs = genReductionPeriodInEpochs(r) },
	)

	var mintintRewardsDistributionStartEpoch int64
	simState.AppParams.GetOrGenerate(
		simState.Cdc, mintingRewardsDistributionStartEpochKey, &mintintRewardsDistributionStartEpoch, simState.Rand,
		func(r *rand.Rand) { mintintRewardsDistributionStartEpoch = genMintintRewardsDistributionStartEpoch(r) },
	)

	reductionStartedEpoch := genReductionStartedEpoch(simState.Rand)

	mintDenom := sdk.DefaultBondDenom
	params := types.NewParams(
		mintDenom,
		epochProvisions,
		epochIdentifier,
		reductionFactor,
		reductionPeriodInEpochs,
		distributionProportions,
		mintintRewardsDistributionStartEpoch,
		ecoFundPoolAddress, developerFundPoolAddress)

	minter := types.NewMinter(epochProvisions)

	mintGenesis := types.NewGenesisState(minter, params, reductionStartedEpoch)

	bz, err := json.MarshalIndent(&mintGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected pseudo-randomly generated minting parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(mintGenesis)
}

func genEpochProvisions(r *rand.Rand) sdk.Dec {
	return sdk.NewDec(int64(r.Intn(maxInt64)))
}

func genReductionFactor(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(r.Intn(10)), 1)
}

func genReductionPeriodInEpochs(r *rand.Rand) int64 {
	return int64(r.Intn(maxInt64))
}

func genMintintRewardsDistributionStartEpoch(r *rand.Rand) int64 {
	return int64(r.Intn(maxInt64))
}

func genReductionStartedEpoch(r *rand.Rand) int64 {
	return int64(r.Intn(maxInt64))
}
