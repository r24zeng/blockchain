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
	if !k.PlayerExist(ctx, msg.PlayerID) { // this player doesn't exist
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "the player doesn't exist, invalid play")
	}

	if k.GetPlayerGameID(ctx, msg.PlayerID) == "" { // this player is not in game progress
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "the player is not in a gam, invalid play")
	}

	if k.GetGameState(ctx, msg.GameID) != "games currently in progress" { // this game is not in progress, can't play
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "this game is not in progress, invalid play")
	}

	// it is not the player's turn, or the coordinate is out of board
	if k.GetPlayerOx(ctx, msg.PlayerID) != k.GetGameCurrTurn(ctx, msg.GameID) {
		reutrn nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "it is not the player's turn, invalid play")
	} 
	if msg.X >= 3 || msg.X < 0 || msg.Y >= 3 || msg.Y < 0 || k.GetGameBoard(ctx, msg.GameID, X, Y) != "_" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "coordinate is invalid, invalid play"
	}

	k.SetGameBoard(ctx, msg.GameID, msg.X, msg.Y, k.GetPlayerOx(ctx, msg.PlayerID))
	if k.GetPlayerOx(ctx, msg.PlayerID) == "O" {
		k.SetGameCurrTurn(ctx, msg.GameID, "X")
	} else {
		k.SetGameCurrTurn(ctx, msg.GameID, "O")
	}

	if k.IsWin(game.board, player.Ox) {
		fmt.Printf("Player %v is win, game %ID is completed\n", msg.PlayerID, msg.GameID)
		k.CompleteGame(ctx, msg.GameID)
		k.ClearPlayer(ctx, msg.PlayerID)
	}

	return &sdk.Result{}, nil
}
