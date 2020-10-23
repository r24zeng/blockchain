package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// define a PlayGame message
type MsgPlayGame struct {
	PlayerID sdk.AccAddress `json:"PlayerID"`
	GameID   string         `json:"GameID"`
	X        int            `[0 - 2]`
	Y        int            `[0 - 2]`
}

// constructor function for MsgPlayGame
func NewMsgPlayGame(player sdk.AccAddress, game string, x int, y int) MsgInviteGame {
	return MsgInviteGame{
		PlayerID: player,
		GameID:   game,
		X:        x,
		Y:        y,
	}
}

//------------ Msg Interface ------------//
// Route should return the name of the module
func (msg MsgPlayGame) Route() string { return RouterKey }

// Type should return the action
func (msg MsgPlayGame) Type() string { return "Play_game" }

// ValidateBasic runs stateless checks on the message
func (msg MsgPlayGame) ValidateBasic() error {
	if len(msg.PlayerID) == 0 || len(msg.GameID) == 0 || msg.X.Empty() || msg.Y.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "PlayerID and/or coordinate cannot be empty")
	}
	if msg.X < 0 || msg.X >= 3 || msg.Y < 0 || msg.Y >= 3 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "coordination is invalid")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgPlayGame) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgPlayGame) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.PlayerID}
}
