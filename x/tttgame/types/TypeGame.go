package types

import (
	"fmt"
)

type Player struct {
	ID     string `json: "id"    yaml: "id"`
	GameID string
	Ox     string `json: "O" or "X"    yaml: "O" or "X"`
}

// NewPlayer constructor
func NewPlayer(id string) Player {
	var player Player
	player.ID = id
	return player
}

// implment Stringer
func (p Player) Sring() string {
	return fmt.Sprint("Player %v is playing game %v with piece of %v", p.ID, p.GameID, p.Ox)
}

type Game struct {
	ID       string
	Players  [2]string
	State    string       `"open games" or "games currently in progress" or "completed games"`
	Board    [3][3]string `"O" or "X" or "_"`
	CurrTurn string       `"O" or "X"`
}

// NewGame contructor
func NewGame(ID string) Game {
	var game Game
	game.ID = ID
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			game.Board[i][j] = "_"
		}
	}
	return game
}

// implement fmt.Stringer
func (g Game) String() string {
	return fmt.Sprintf("The current state of game %v is: %v", g.ID, g.State)
}
