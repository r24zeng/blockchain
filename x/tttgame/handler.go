package tttgame

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/r24zeng/tttgame/x/tttgame/keeper"
	"github.com/r24zeng/tttgame/x/tttgame/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, tx sdk.Tx) (*sdk.Result, error) {
		msg := tx.GetMsg()
		sig := tx.GetSignaure()
		if sig.Empty() {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "signature is missing")
		}
		signer := msg.GetSigner()
		pub := k.GetPlayerPubKey(signer)
		if !pub.VerifySignature(msg, sig) {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "verify signature fail")
		}

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case types.MsgInviteGame:
			return handleMsgInviteGame(ctx, k, msg)
		case types.MsgAcceptGame:
			return handleMsgAcceptGame(ctx, k, msg)
		case types.MsgPlayGame:
			return handleMsgPlayGame(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
