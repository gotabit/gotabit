# Mint

The `mint` module is responsible for creating tokens in a flexible way to reward 
validators, ecosystem fund, provide funds for Gotabit governance,
and pay developers to maintain and improve Gotabit.

The module is also responsible for reducing the token creation and distribution by a set period
until it reaches its maximum supply (see `reduction_factor` and `reduction_period_in_epochs`)

The module uses time basis epochs supported by the `epochs` module.

## Contents

1. **[Concept](#concepts)**
2. **[State](#state)**
3. **[Begin Epoch](#begin-epoch)**
4. **[Parameters](#network-parameters)**
5. **[Events](#events)**
6. **[Transactions](#transaction)**
7. **[Queries](#queries)**
    
## Concepts

The `x/mint` module is designed to handle the regular printing of new
tokens within a chain. The design taken within Gotabit is to

- Mint new tokens once per epoch (default one week)
- To have a "Reductioning factor" every period, which reduces the number of
    rewards per epoch. (default: period is 3 years, where a
    year is 52 epochs. The next period's rewards are 2/3 of the prior
    period's rewards)

### Reduction factor

This is a generalization over the Bitcoin-style halvenings. Every year, the number
 of rewards issued per week will reduce by a governance-specified 
factor, instead of a fixed `1/2`. So
`RewardsPerEpochNextPeriod = ReductionFactor * CurrentRewardsPerEpoch)`.
When `ReductionFactor = 1/2`, the Bitcoin halvenings are recreated. We
default to having a reduction factor of `2/3` and thus reduce rewards
at the end of every year by `33%`.

The implication of this is that the total supply is finite, according to
the following formula:

`Total Supply = InitialSupply + EpochsPerPeriod * { {InitialRewardsPerEpoch} / {1 - ReductionFactor} }`

## State

### Minter

The [`Minter`](https://github.com/hjcore/gotabit/blob/main/x/mint/types/mint.pb.go#L31) is an abstraction for holding current rewards information.

```go
type Minter struct {
    EpochProvisions sdk.Dec   // Rewards for the current epoch
}
```

### Params

Minting [`Params`](https://github.com/hjcore/gotabit/blob/main/x/mint/types/mint.pb.go#L119) are held in the global params store.

### LastReductionEpoch

Last reduction epoch stores the epoch number when the last reduction of
coin mint amount per epoch has happened.

## Begin-Epoch

Minting parameters are recalculated and inflation is paid at the beginning
of each epoch. An epoch is signaled by x/epochs

### NextEpochProvisions

The target epoch provision is recalculated on each reduction period
(default 3 years). At the time of the reduction, the current provision is
multiplied by the reduction factor (default `2/3`), to calculate the
provisions for the next epoch. Consequently, the rewards of the next
period will be lowered by a `1` - reduction factor.

### EpochProvision

Calculate the provisions generated for each epoch based on current epoch
provisions. The provisions are then minted by the `mint` module's
`ModuleMinterAccount`. These rewards are transferred to a
`FeeCollector`, which handles distributing the rewards per the chain's needs.
This fee collector is specified as the `auth` module's `FeeCollector` `ModuleAccount`.

## Network Parameters

The minting module contains the following parameters:

| Key                                          | Type        | Example           |
|----------------------------------------------|-------------|-------------------|
| mint_denom                                   | string      | "ugio"            |
| genesis_epoch_provisions                     | string (dec) | "500000000"       |
| epoch_identifier                             | string      | "weekly"          |
| reduction_period_in_epochs                   | int64       | 156               |
| reduction_factor                             | string (dec) | "0.6666666666666" |
| distribution_proportions.staking             | string (dec) | "0.4"             |
| distribution_proportions.eco_fund_pool       | string (dec) | "0.3"             |
| distribution_proportions.developer_fund_pool | string (dec) | "0.2"             |
| distribution_proportions.community_pool      | string (dec) | "0.1"             |
| minting_rewards_distribution_start_epoch     | int64       | 10                |
| eco_fund_pool_address                        | string      |                   |
| developer_fund_pool_address                  | string      |                   |



Below are all the network parameters for the ```mint``` module:

- **```mint_denom```** - Token type being minted
- **```genesis_epoch_provisions```** - Amount of tokens generated at the epoch to the distribution categories (see distribution_proportions)
- **```epoch_identifier```** - Type of epoch that triggers token issuance (day, week, etc.)
- **```reduction_period_in_epochs```** - How many epochs must occur before implementing the reduction factor
- **```reduction_factor```** - What the total token issuance factor will reduce by after the reduction period passes (if set to 66.66%, token issuance will reduce by 1/3)
- **```distribution_proportions```** - Categories in which the specified proportion of newly released tokens are distributed to
  - **```staking```** - Proportion of minted funds to incentivize staking GIO
  - **```eco_fund_pool```** - Proportion of minted funds to pay ecosystem
  - **```developer_fund_pool```** - Proportion of minted funds to pay developers for their past and future work
  - **```community_pool```** - Proportion of minted funds to be set aside for the community pool
- **```minting_rewards_distribution_start_epoch```** - What epoch will start the rewards distribution to the aforementioned distribution categories
- **```eco_fund_pool_address```** - Addresses that ecological foundation rewards will go to.
- **```developer_fund_pool_address```** - Addresses that developer foundation rewards will go to.

**Notes**

1. `mint_denom` defines denom for minting token - ugio
2. `genesis_epoch_provisions` provides minting tokens per epoch at genesis.
3. `epoch_identifier` defines the epoch identifier to be used for the mint module e.g. "weekly"
4. `reduction_period_in_epochs` defines the number of epochs to pass to reduce the mint amount
5. `reduction_factor` defines the reduction factor of tokens at every `reduction_period_in_epochs`
6. `distribution_proportions` defines distribution rules for minted tokens, when the developer 
    rewards address is empty, it distributes tokens to the community pool.
7. `minting_rewards_distribution_start_epoch` defines the start epoch of minting to make sure
    minting start after initial pools are set
8. `eco_fund_pool_address` defines the contract addresses that ecosystem fund will go to
9. `developer_fund_pool_address` defines the contract addresses that developer fund will go to

## Events

The minting module emits the following events:

### End of Epoch

|  Type | Attribute Key      | Attribute Value    |
|------ |--------------------|--------------------|
|  mint | epoch\_number      | {epochNumber}      |
|  mint | epoch\_provisions  | {epochProvisions}  |
|  mint | amount             | {amount}           |

</br>
</br>

## Queries

### params

Query all the current mint parameter values

```sh
query mint params
``` 

::: details Example

List all current min parameters in json format by:

```bash
gotabitd query mint params -o json | jq
```

An example of the output:

```json
{
  "mint_denom": "ugio",
  "genesis_epoch_provisions": "1000000.000000000000000000",
  "epoch_identifier": "day",
  "reduction_period_in_epochs": "365",
  "reduction_factor": "0.666666666666666666",
  "distribution_proportions": {
    "staking": "0.250000000000000000",
    "eco_fund_pool": "0.450000000000000000",
    "developer_fund_pool": "0.250000000000000000",
    "community_pool": "0.050000000000000000"
  },
  "minting_rewards_distribution_start_epoch": "1",
  "eco_fund_pool_address": "gio1d4ysgq9ljs2k6hhafhuy9mnstuunzeqd800vlk75jkemdrn4r0ks7ezpjk",
  "developer_fund_pool_address": "gio14zd54ruj9zu9lyz8puv3rp4tnts8xezjntq9et"
}
```
:::


### epoch-provisions

Query the current epoch provisions

```sh
query mint epoch-provisions
```

::: details Example

List the current epoch provisions:

```bash
gotabitd query mint epoch-provisions
```
As of this writing, this number will be equal to the ```genesis-epoch-provisions```. Once the ```reduction_period_in_epochs``` is reached, the ```reduction_factor``` will be initiated and reduce the amount of GIO minted per epoch.
:::

## License

This is modify by [osmosis mint](https://github.com/osmosis-labs/osmosis/tree/main/x/mint)

This software is licensed under the Apache 2.0 license.

© 2022 GotaBit Limited
