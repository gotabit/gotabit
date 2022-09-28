package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	// this line is used by starport scaffolding # stargate/app/moduleImport

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	ica "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts"
	icacontrollertypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"

	store "github.com/cosmos/cosmos-sdk/store/types"

	msgtypes "github.com/gotabit/gotabit/x/inbox/types"
)

// RegisterUpgradeHandlers returns upgrade handlers
func (app *App) RegisterUpgradeHandlers(cfg module.Configurator) {
	// left here for hysterical raisins - we should probably remove
	// bankBaseKeeper, _ := app.BankKeeper.(bankkeeper.BaseKeeper)
	// app.UpgradeKeeper.SetUpgradeHandler(veritas.UpgradeName, veritas.CreateUpgradeHandler(app.mm, cfg, &app.StakingKeeper, &bankBaseKeeper))
	app.UpgradeKeeper.SetUpgradeHandler("multiverse", func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		vm[icatypes.ModuleName] = app.mm.Modules[icatypes.ModuleName].ConsensusVersion()

		// create ICS27 Controller submodule params, controller module not enabled.
		controllerParams := icacontrollertypes.Params{}

		// create ICS27 Host submodule params
		hostParams := icahosttypes.Params{
			HostEnabled: true,
			AllowMessages: []string{
				sdk.MsgTypeURL(&banktypes.MsgSend{}),
				sdk.MsgTypeURL(&stakingtypes.MsgDelegate{}),
				sdk.MsgTypeURL(&stakingtypes.MsgBeginRedelegate{}),
				sdk.MsgTypeURL(&stakingtypes.MsgCreateValidator{}),
				sdk.MsgTypeURL(&stakingtypes.MsgEditValidator{}),
				sdk.MsgTypeURL(&distrtypes.MsgWithdrawDelegatorReward{}),
				sdk.MsgTypeURL(&distrtypes.MsgSetWithdrawAddress{}),
				sdk.MsgTypeURL(&distrtypes.MsgWithdrawValidatorCommission{}),
				sdk.MsgTypeURL(&distrtypes.MsgFundCommunityPool{}),
				sdk.MsgTypeURL(&govtypes.MsgVote{}),
				sdk.MsgTypeURL(&authz.MsgExec{}),
				sdk.MsgTypeURL(&authz.MsgGrant{}),
				sdk.MsgTypeURL(&authz.MsgRevoke{}),
				// wasm msgs here
				// note we only support these three for now
				sdk.MsgTypeURL(&wasmtypes.MsgStoreCode{}),
				sdk.MsgTypeURL(&wasmtypes.MsgInstantiateContract{}),
				sdk.MsgTypeURL(&wasmtypes.MsgExecuteContract{}),
				// msg module msgs
				sdk.MsgTypeURL(&msgtypes.MsgSend{}),
			},
		}

		// initialize ICS27 module
		icamodule, correctTypecast := app.mm.Modules[icatypes.ModuleName].(ica.AppModule)
		if !correctTypecast {
			panic("mm.Modules[icatypes.ModuleName] is not of type ica.AppModule")
		}
		icamodule.InitModule(ctx, controllerParams, hostParams)

		return app.mm.RunMigrations(ctx, cfg, vm)
	})

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if upgradeInfo.Name == "multiverse" && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := store.StoreUpgrades{
			Added: []string{icacontrollertypes.StoreKey, icahosttypes.StoreKey},
		}

		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}
