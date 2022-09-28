package app

import (
	"encoding/json"

	"github.com/gotabit/gotabit/cmd/cosmoscmd"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp"
)

// Setup initializes a new App
func Setup(isCheckTx bool) *App {
	db := dbm.NewMemDB()
	encoding := cosmoscmd.MakeEncodingConfig(ModuleBasics)
	app := NewGotabitApp(log.NewNopLogger(), db, nil, true, map[int64]bool{}, DefaultNodeHome, 5, encoding, simapp.EmptyAppOptions{})
	if !isCheckTx {
		interfaceRegistry := types.NewInterfaceRegistry()
		marshaler := codec.NewProtoCodec(interfaceRegistry)
		genesisState := NewDefaultGenesisState(marshaler)
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}

		app.InitChain(
			abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: simapp.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return app
}
