package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidPlay       = sdkerrors.Register(ModuleName, 1, "This play is invalid")
	ErrGameNotExist      = sdkerrors.Register(ModuleName, 1, "This game ID does not exist")
	ErrGameExist         = sdkerrors.Register(ModuleName, 1, "This game ID exists, can't create a new game")
	ErrGameNotInProgress = sdkerrors.Register(ModuleName, 1, "This game is not in progress")
	ErrGameInProgress    = sdkerrors.Register(ModuleName, 1, "This game is in progress")
)
