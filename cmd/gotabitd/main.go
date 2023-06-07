package main

import (
	"fmt"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/gotabit/gotabit/app"
	"github.com/gotabit/gotabit/cmd/cosmoscmd"
)

func main() {
	rootCmd := cosmoscmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, fmt.Sprint(".", app.Name), app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
