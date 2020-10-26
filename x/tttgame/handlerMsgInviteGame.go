package tttgame

import (
	crkeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/r24zeng/tttgame/x/tttgame/keeper"
	"github.com/r24zeng/tttgame/x/tttgame/types"
)

func handleMsgInviteGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgInviteGame, key crkeys.Keybase) (*sdk.Result, error) {
	// player account has been created or verified in client/cliCommand
	// msg has been verified in NewHandler, then route to the corresponding handler

	if !k.GameExist(ctx, msg.GameID) { // if the game doesn't exist, initialize a game
		k.CreateGame(ctx, msg.gameID)
	}

	if k.GetGameState(ctx, msg.GameID) != "completed games" { // the game is in progress, fail
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "This game is occupied, invite fail")
	}

	pubKey := key.GetByAddress(msg.Player)
	k.OpenGame(ctx, pubKey, msg.gameID)

	return &sdk.Result{}, nil
}
