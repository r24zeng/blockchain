package tttgame

import (
	"crypto/md5"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/r24zeng/tttgame/x/tttgame/keeper"
	"github.com/r24zeng/tttgame/x/tttgame/types"
)

func handleMsgAcceptGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgAcceptGame) (*sdk.Result, error) {
	if !k.PlayerExist(ctx, msg.PlayerID) { // the player doesn't exist, create a new player
		k.CreatePlayer(ctx, msg.PlayerID)
	}

	if k.GetPlayerGameID(ctx, msg.PlayerID) != "" { // the player is in antoher game progress, fail
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "the player is in another game, accept fail")
	}

	if !k.GameExist(ctx, msg.GameID) { // if the game doesn't exist, the accept is invalid, fail
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "the game does not exist, accept fail")
	}

	if k.GetGameState(ctx, msg.GameID) != "open games" { // game is not waitting for another player to join
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "this game is not open for joining, accept fail")
	}

	// decide who play first, default "O" plays first
	// a hash function is applied on the concatenation of the user's public keys.
	// if the first bit of the output is 0, then the game's initiator (whoever posted the invitation) plays "O"
	// and the second player plays "X" and vice versa.

	k.SetPlayerGameID(ctx, msg.GameID)
	k.ActiveGame(ctx, msg.PlayerID, game)

	hash_value = md5.Sum(k.GetGamePlayers(ctx, msg.GameID, 0) + k.GetGamePlayers(ctx, msg.GameID, 1))
	if hash_value[0] == 0 {
		k.SetPlayerOx(ctx, k.GetGamePlayers(ctx, msg.GameID, 0), "O")
	} else {
		k.SetPlayerOx(ctx, game.Player[0], "X")
	}

	return &sdk.Result{}, nil
}
