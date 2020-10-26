package types

import (
	"fmt"

	crypto "github.com/tendermint/tendermint/crypto"
)

// Game store in KVStore
type Game struct {
	ID string `json:"ID" yaml:"ID"`
	// no player -> "completed games"; one player -> "open games"; two players -> "games currently in progress"
	// Players[0] always play "O", Players[1] plays "X"
	Players []crypto.PubKey
	// 0 represent "O", 1 represent "X", -1 represent "_"
	Board [3][3]int `json:"Board" yaml:"Board" `
	// player who holds "O" always play first;
	// CurrTurn is the index of current player in Players[2]
	CurrTurn int `json:"CurrTurn" yaml:"CurrTurn"`
}

// implement fmt.Stringer
func (g Game) String() string {
	return fmt.Sprintf("The current game %v is: ", g.ID)
}
