package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// define a InviteGame message
type MsgInviteGame struct {
	PlayerID string `json:"PlayerID"`
	GameID   string `json:"value"`
}

// constructor function for MsgInviteGame
func NewMsgInviteGame(player string, game string) MsgInviteGame {
	return MsgInviteGame{
		PlayerID: player,
		GameID:   game,
	}
}

//------------ Msg Interface ------------//
// Route should return the name of the module
func (msg MsgInviteGame) Route() string { return RouterKey }

// Type should return the action
func (msg MsgInviteGame) Type() string { return "Invite_game" }

// ValidateBasic runs stateless checks on the message
func (msg MsgInviteGame) ValidateBasic() error {
	if len(msg.PlayerID) == 0 || len(msg.GameID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "PlayerID and/or GameID cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgInviteGame) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgInviteGame) GetSigners() string {
	return msg.PlayerID
}
