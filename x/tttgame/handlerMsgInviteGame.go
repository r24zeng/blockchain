package tttgame

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/r24zeng/tttgame/x/tttgame/keeper"
	"github.com/r24zeng/tttgame/x/tttgame/types"
)

func handleMsgInviteGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgInviteGame) (*sdk.Result, error) {
	var player types.Player
	var game types.Game

	if !k.PlayerExist(ctx, msg.playerID) { // the player doesn't exist, then create a new player
		player = k.CreatePlayer(ctx, msg.playerID)
	} else {
		player, _ := k.GetPlayer(ctx, msg.playerID)
	}

	if player.GameID != "" { // the player is in another game progress, fail
		return nil, types.ErrInvalidPlay
	}

	if !k.GameExist(ctx, msg.GameID) { // if the game doesn't exist, initial a game
		game = k.CreateGame(ctx, msg.gameID)
	} else {
		game, _ := k.GetGame(ctx, msg.gameID)
	}

	if game.State != "completed games" { // the game is in progress, fail
		return nil, types.ErrGameInProgress
	}

	k.SetPlayerGameID(ctx, msg.GameID)
	k.OpenGame(ctx, msg.PlayerID, game)

	return &sdk.Result{}, nil
}
