package tttgame

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/r24zeng/tttgame/x/tttgame/keeper"
	"github.com/r24zeng/tttgame/x/tttgame/types"
)

// Handle a message to delete name
func handleMsgPlayGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgPlayGame) (*sdk.Result, error) {
	// player account has been created or verified in client/cliCommand
	// msg has been verified in NewHandler, then route to the corresponding handler
	
	if !k.GameExist(ctx, msg.GameID) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "this game doesn't exist")
	}

	if k.GetGameState(ctx, msg.GameID) != "games currently in progress" { 
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "this game is not in progress, invalid play")
	}

	if k.GetGameCurrPlayer(ctx, msg.GameID) != msg.PlayerID {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "wrong player or game, invalid play"
	}

	// play game, if win then complete the game
	k.PlayGame(ctx, msg.GameID, msg.X, msg.Y)
	if k.GetGameState(ctx, msg.GameID) == "completed games" {
		fmt.Print("current player win, game %ID is completed\n", msg.GameID)
	}
	
	return &sdk.Result{}, nil
}
