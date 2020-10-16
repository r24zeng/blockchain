package tttgame

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/r24zeng/tttgame/x/tttgame/keeper"
	"github.com/r24zeng/tttgame/x/tttgame/types"
)

// Handle a message to delete name
func handleMsgPlayGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgPlayGame) (*sdk.Result, error) {

	return &sdk.Result{}, nil
}
