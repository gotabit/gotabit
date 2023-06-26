package keeper_test

import (
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	simapp "github.com/gotabit/gotabit/app"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx sdk.Context
	app *simapp.App
}
