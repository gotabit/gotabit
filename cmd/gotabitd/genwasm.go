package main

import (
	"github.com/gotabit/gotabit/cmd/cosmoscmd"
	"github.com/spf13/cobra"

	"github.com/CosmWasm/wasmd/x/wasm"
)

func GetWasmCmdOptions() []cosmoscmd.Option {
	var options []cosmoscmd.Option

	options = append(options,
		cosmoscmd.CustomizeStartCmd(func(startCmd *cobra.Command) {
			wasm.AddModuleInitFlags(startCmd)
		}),
	)

	return options
}
