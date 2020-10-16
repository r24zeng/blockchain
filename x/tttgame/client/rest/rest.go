package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers tttgame-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding # 1
		r.HandleFunc("/tttgame/game", createGameHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/tttgame/game", listGameHandler(cliCtx, "tttgame")).Methods("GET")
		r.HandleFunc("/tttgame/game/{key}", getGameHandler(cliCtx, "tttgame")).Methods("GET")
		r.HandleFunc("/tttgame/game", setGameHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/tttgame/game", deleteGameHandler(cliCtx)).Methods("DELETE")

		
}
