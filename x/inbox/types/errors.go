package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// module sentinel errors
var (
	ErrMsgDoesNotExist = sdkerrors.Register(ModuleName, 1, "Msg does not exist")
)
