package keeper

import (
	// this line is used by starport scaffolding # 1
	"github.com/r24zeng/tttgame/x/tttgame/types"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for tttgame clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// this line is used by starport scaffolding # 2
		case types.QueryGetGame:
			return getGame(ctx, path[1:], k)
		case types.QueryGetPlayer:
			return listGame(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown tttgame query endpoint")
		}
	}
}
