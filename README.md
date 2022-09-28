<div align="center">
  <h1> GotaBit Chain </h1>
  <p> The full-node software implementation of the GotaBit blockchain. </p>
  
<p align="center">
  <a href="https://docs.gotabit.io"><strong>Explore the Docs</strong></a>
  <br />
  <br />
  ·
  <a href="https://api.gotabit.dev">Rest API</a>
  ·
  <a href="https://rpc.gotabit.dev">RPC API</a>
  ·
  <a href="https://scan.gotabit.io">GotaBit Scan</a>
  ·
  <a href="https://app.gotabit.io">Chain App</a>
</p>
</div>

## Table of Contents <!-- omit in toc -->
- [What is GotaBit?](#what-is-gotabit)
- [Installation](#installation)
  - [From Binary](#from-binary)
  - [From Source](#from-source)
  - [GotaBitd](#gotabitd)
- [Node Setup](#node-operators)
- [Validators](#validators)
- [Delegators](#delegators)
- [Testnet](#testnet)
- [Talk to us](#talk-to-us)
- [License](#license)

<br >

## What is GotaBit?

**[GotaBit](https://g.io)** is a public, open-source, proof-of-stake, sovereign blockchain in the Cosmos ecosystem. It aims to provide a sandbox environment for the deployment of such inter-operable smart contracts.

The network serves as a decentralized, permissionless, and censorship resistant zone for developers to efficiently and securely launch application specific smart contracts.

The GotaBit is powered by the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk), [CosmWasm](https://github.com/CosmWasm/cosmwasm) and [Tendermint](https://github.com/tendermint/tendermint) BFT consensus.


## Installation

### From Binary

The easiest way to install is to download a pre-built binary. You can find the latest binaries on the [releases](https://github.com/gotabit/node/releases) page.

### From Source

**Step 1: Install Golang**

Go v1.18+ or higher is required for The GotaBit Node.

1. Install [Go 1.18+ from the official site](https://go.dev/dl/). Ensure that your `GOPATH` and `GOBIN` environment variables are properly set up by using the following commands:

   For Linux:

   ```sh
   wget <https://golang.org/dl/go1.18.2.linux-amd64.tar.gz>
   sudo tar -C /usr/local -xzf go1.18.2.linux-amd64.tar.gz
   export PATH=$PATH:/usr/local/go/bin
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

   For Mac:

   ```sh
   export PATH=$PATH:/usr/local/go/bin
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

2. Confirm your Go installation by checking the version:

   ```sh
   go version
   ```


**Step 2: Get source code**

Clone the source code from the [official repo](https://github.com/gotabit/node) and check out the `main` branch for the latest stable release.

```bash
git clone https://github.com/gotabit/node
cd node
git checkout main
```

**Step 3: Build GotaBit node**

Run the following command to install `gotabitd` to your `GOPATH` and build the GotaBit node. `gotabitd` is the node daemon and CLI for interacting with a GotaBit node.

```bash
make install
```

**Step 4: Verify your installation**

Verify your installation with the following command:

```bash
gotabitd version --long
```

A successful installation will return the following:

```bash
name: gotabit
server_name: gotabitd
version: <x.x.x>
commit: <Commit hash>
build_tags: netgo,ledger
go: go version go1.18.2 linux/amd64
build_deps:
...
```

### `Gotabitd`

`gotabitd` is the all-in-one CLI and node daemon for interacting with the GotaBit blockchain. 

To view various subcommands and their expected arguments, use the following command:

``` sh
 gotabitd --help
```

```
Stargate CosmosHub App

Usage:
  gotabitd [command]

Available Commands:
  add-genesis-account      Add a genesis account to genesis.json
  add-wasm-genesis-message Wasm genesis subcommands
  collect-gentxs           Collect genesis txs and output a genesis.json file
  config                   Create or query an application CLI configuration file
  debug                    Tool for helping with debugging your application
  export                   Export state to JSON
  gentx                    Generate a genesis tx carrying a self delegation
  help                     Help about any command
  init                     Initialize private validator, p2p, genesis, and application configuration files
  keys                     Manage your application's keys
  migrate                  Migrate genesis to a specified target version
  query                    Querying subcommands
  rollback                 rollback cosmos-sdk and tendermint state by one height
  start                    Run the full node
  status                   Query remote node for status
  tendermint               Tendermint subcommands
  tx                       Transactions subcommands
  validate-genesis         validates the genesis file at the default location or at the location passed as an arg
  version                  Print the application binary version information

Flags:
  -h, --help                help for gotabitd
      --home string         directory for config and data (default "/root/.gotabit")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors

Use "gotabitd [command] --help" for more information about a command.
```

Visit the [documentation page](https://docs.gotabit.io/how-to) for more info on usage. 

## Node Operators

If you're interested in running a node on the current GotaBit network, check out the docs to [Join the GotaBit Mainnet](https://docs.gotabit.io/).

## Validators

If you want to participate and help secure GotaBit network, check out becoming a validator. Information on what a validator is and how to participate as one can be found at the [Validator FAQ](https://docs.gotabit.io/). If you're running a validator node on the GotaBit network, reach out to a Janitor on the [GotaBit Discord](https://discord.gg/dDgRkVwqD6) to join the `#validators-verified` channel.

## Delegators

If you still want to participate on the GotaBit network, check out becoming a delegator. Information what a delegator is and how to participate as one can be found at the [Delegator FAQ](https://docs.gotabit.io/).

## Testnet

To participate in or utilize the current GotaBit testnet, take a look at the [gotabit/testnets](https://github.com/gotabit/testnets) repository. 

## Talk to us

We have active, helpful communities on Twitter, Discord, and Telegram.

<p>
<a href="https://twitter.com/GotaBitG"><img src="https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white" alt="Tweet" height="30"/></a> 
  &nbsp;
 <a href="https://t.me/gotabit"><img src="https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white" alt="Telegram" height="30"/></a> 
</p>

For updates on the GotaBit team's activities follow us on the [GotaBit Twitter](https://twitter.com/GotaBitG) account.

## License

This software is licensed under the Apache 2.0 license.

© 2022 GotaBit Limited
