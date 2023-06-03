package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/gotabit/gotabit/cmd/cosmoscmd"
)

func main() {
	rootCmd := cosmoscmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "gotabit", "gotabit"); err != nil {
		os.Exit(1)
	}
}
