package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgPlayGame defines a PlayGame message
type MsgPlayGame struct {
	Player sdk.AccAddress `json:"Player" yaml:"Player"`
	GameID string         `json:"GameID" yaml:"GameID"`
	X      int            `json:"X" yaml:"X"`
	Y      int            `json:"Y" yaml:"Y"`
}

// NewMsgPlayGame is a constructor function for MsgPlayGame
func NewMsgPlayGame(player sdk.AccAddress, game string, x int, y int) MsgPlayGame {
	return MsgPlayGame{
		Player: player,
		GameID: game,
		X:      x,
		Y:      y,
	}
}

// Route should return the name of the module
func (msg MsgPlayGame) Route() string { return RouterKey }

// Type should return the action
func (msg MsgPlayGame) Type() string { return "Play_game" }

// ValidateBasic runs stateless checks on the message
func (msg MsgPlayGame) ValidateBasic() error {
	if len(msg.GameID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "GameID cannot be empty")
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
	return []sdk.AccAddress{msg.Player}
}
