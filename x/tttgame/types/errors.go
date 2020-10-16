package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidPlay       = sdkerrors.Register(ModuleName, 1, "This play is invalid")
)
