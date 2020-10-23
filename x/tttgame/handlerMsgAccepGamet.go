package tttgame

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/r24zeng/tttgame/x/tttgame/keeper"
	"github.com/r24zeng/tttgame/x/tttgame/types"
)

func handleMsgAcceptGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgAcceptGame) (*sdk.Result, error) {
	// player account has been created or verified in client/cliCommand
	// msg has been verified in NewHandler, then route to the corresponding handler

	if !k.GameExist(ctx, msg.GameID) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "the game does not exist, accept fail")
	}

	if k.GetGameState(ctx, msg.GameID) != "open games" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "this game is not open for joining, accept fail")
	}

	error := k.ActiveGame(ctx, msg.PlayerID, msg.GameID) // when inviter = opponent, accept fail
	if error != nil {
		return inl, error
	}

	return &sdk.Result{}, nil
}
