package types

const QueryGetGame = "get-game"
const QueryGetPlayer = "get-player"
const QueryGetGameBoard = "get-game-board"

// QueryResNames Queries Result Payload for a names query
type QueryResBoard [3][3]string

// implement fmt.Stringer
func (n QueryResBoard) String() string {
	var tmp string
	for i := 0; i < 3; i++ {
		tmp = tmp + strings.join(n[i][:], "\n")
	}
	return tmp
}
