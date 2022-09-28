package main

import (
	"os"

	"github.com/gotabit/gotabit/app"
	"github.com/gotabit/gotabit/cmd/cosmoscmd"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	cmdOptions := GetWasmCmdOptions()

	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.CoinType,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
		cmdOptions...,
	)
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
