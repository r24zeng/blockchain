package types

import (
	"fmt"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Player struct {
	ID      sdk.AccAddress      `json: "ID"  yaml: "ID"`
	PubKey  cryptotypes.PubKey  `json: "PubKey"  yaml: "PubKey"`
	PrivKey cryptotypes.PrivKey `json: "PrivKey"  yaml: "PrivKey"`
}

type Game struct {
	ID string `json: "ID" yaml: "ID"`
	// no player -> "completed games"; one player -> "open games"; two players -> "games currently in progress"
	// Players[0] always play "O", Players[1] plays "X"
	Players []sdk.AccAddress
	// 0 represent "O", 1 represent "X", -1 represent "_"
	Board [3][3]int `json: "0" or "1" or "-1" yaml: "0" or "1" or "-1" `
	// player who holds "O" always play first;
	// CurrTurn is the index of current player in Players[2]
	CurrTurn int `json: "0" or "1" yaml: "0" or "1"`
}

// implement fmt.Stringer
func (g Game) String() string {
	return fmt.Sprintf("The current game %v is: ", g.ID)
}
