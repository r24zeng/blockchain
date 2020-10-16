package tttgame

import (
	"crypto/md5"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/r24zeng/tttgame/x/tttgame/keeper"
	"github.com/r24zeng/tttgame/x/tttgame/types"
)

func handleMsgAcceptGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgAcceptGame) (*sdk.Result, error) {
	var player types.Player
	var game types.Game

	if !k.PlayerExist(ctx, msg.playerID) { // the player doesn't exist, create a new player
		player, _ := k.CreatePlayer(ctx, playerID)
	} else {
		player = k.GetPlayer(ctx, playerID)
	}

	if player.ID != "" { // the player is in antoher game progress, fail
		return nil, types.ErrInvalidPlay
	}

	if !k.GameExist(ctx, msg.GameID) { // if the game doesn't exist, the accept is invalid, fail
		return nil, types.ErrGameNotExist
	}

	game = k.GetGame(ctx, msg.GameID)

	if game.State != "open games" { // game is not waitting for another player to join
		return nil, types.ErrGameInProgress
	}

	// decide who play first, default "O" plays first
	// a hash function is applied on the concatenation of the user's public keys.
	// if the first bit of the output is 0, then the game's initiator (whoever posted the invitation) plays "O"
	// and the second player plays "X" and vice versa.

	k.SetPlayerGameID(ctx, msg.GameID)
	k.ActiveGame(ctx, msg.playerID, game)

	hash_value = md5.Sum([]byte(game.Players[0] + game.Players[1]))
	if hash_value[0] == 0 {
		k.SetPlayerOx(ctx, game.Player[0], "O")
	} else {
		k.SetPlayerOx(ctx, game.Player[0], "X")
	}

	return &sdk.Result{}, nil
}
