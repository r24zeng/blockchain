package tttgame

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/r24zeng/tttgame/x/tttgame/keeper"
	"github.com/r24zeng/tttgame/x/tttgame/types"
)

func handleMsgInviteGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgInviteGame) (*sdk.Result, error) {

	if !k.PlayerExist(ctx, msg.PlayerID) { // the player doesn't exist, then create a new player
		k.CreatePlayer(ctx, msg.PlayerID)
	}

	if k.GetPlayerGameID(ctx, msg.PlayerID) != "" { // the player is in another game progress, fail
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "the player is in another game progress, invite fail")
	}

	if !k.GameExist(ctx, msg.GameID) { // if the game doesn't exist, initialize a game
		k.CreateGame(ctx, msg.gameID)
	}

	if k.GetGameState(ctx, msg.GameID) != "completed games" { // the game is in progress, fail
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "This game is waiting for joining or in progress, invite fail")
	}

	k.SetPlayerGameID(ctx, msg.GameID)
	k.OpenGame(ctx, msg.PlayerID, game)

	return &sdk.Result{}, nil
}
