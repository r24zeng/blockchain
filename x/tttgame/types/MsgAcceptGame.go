package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// const RouterKey = ModuleName // this was defined in key.go file

// MsgAcceptGame define a InviteGame message
type MsgAcceptGame struct {
	Player sdk.AccAddress `json:"Player"`
	GameID string         `json:"value"`
}

// NewMsgAcceptGame is a constructor function for MsgAcceptGame
func NewMsgAcceptGame(player sdk.AccAddress, game string) MsgAcceptGame {
	return MsgAcceptGame{
		Player: player,
		GameID: game,
	}
}

// Route should return the name of the module
func (msg MsgAcceptGame) Route() string { return RouterKey }

// Type should return the action
func (msg MsgAcceptGame) Type() string { return "Accept_game" }

// ValidateBasic runs stateless checks on the message
func (msg MsgAcceptGame) ValidateBasic() error {
	if len(msg.GameID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "GameID cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAcceptGame) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgAcceptGame) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Player}
}
