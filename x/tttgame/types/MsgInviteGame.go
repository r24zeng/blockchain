package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgInviteGame define a InviteGame message
type MsgInviteGame struct {
	Player sdk.AccAddress `json:"Player"`
	GameID string         `json:"value"`
}

// NewMsgInviteGame is constructor function for MsgInviteGame
func NewMsgInviteGame(player sdk.AccAddress, game string) MsgInviteGame {
	return MsgInviteGame{
		Player: player,
		GameID: game,
	}
}

// Route should return the name of the module
func (msg MsgInviteGame) Route() string { return RouterKey }

// Type should return the action
func (msg MsgInviteGame) Type() string { return "Invite_game" }

// ValidateBasic runs stateless checks on the message
func (msg MsgInviteGame) ValidateBasic() error {
	if len(msg.GameID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "GameID cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgInviteGame) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgInviteGame) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Player}
}
