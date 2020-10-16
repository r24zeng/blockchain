package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const RouterKey = ModuleName // this was defined in key.go file

// define a InviteGame message
type MsgAcceptGame struct {
	PlayerID string `json:"PlayerID"`
	GameID   string `json:"value"`
}

// constructor function for MsgInviteGame
func NewMsgAcceptGame(player string, game string) MsgInviteGame {
	return MsgAcceptGame{
		PlayerID: player,
		GameID:   game,
	}
}

//------------ Msg Interface ------------//
// Route should return the name of the module
func (msg MsgAcceptGame) Route() string { return RouterKey }

// Type should return the action
func (msg MsgAcceptGame) Type() string { return "Accept_game" }

// ValidateBasic runs stateless checks on the message
func (msg MsgAcceptGame) ValidateBasic() error {
	if len(msg.PlayerID) == 0 || len(msg.GameID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "PlayerID and/or GameID cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAcceptGame) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgAcceptGame) GetSigners() string {
	return msg.PlayerID
}
