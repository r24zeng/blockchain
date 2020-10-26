package tttgame

import (
	"fmt"

	crkeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/r24zeng/tttgame/x/tttgame/keeper"
	"github.com/r24zeng/tttgame/x/tttgame/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, tx auth.StdTx, key crkeys.Keybase) (*sdk.Result, error) {
		// verify signture
		err := tx.ValidateBasic()
		if err != nill {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "invalid stdTx")
		}
		msg := tx.GetMsgs()[0]
		sig := tx.GetSignaures()[0]
		signer := msg.GetSigner() // get signer from msg
		keyInfo, err := key.GetByAddress(signer)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "signer doesn't exsited")
		}
		if !keyInfo.GetPubKey().VerifySignature(msg, sig) {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "verify signature fail")
		}

		// switch to other handler
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
