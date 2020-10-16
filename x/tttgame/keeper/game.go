package keeper

import (
	"strconv"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/r24zeng/tttgame/x/tttgame/types"
)

//---------------- player related keeper functions ----------------//
// get a player
func (k Keeper) GetPlayer(ctx sdk.Context, playerID string) (types.Player, err) {
	sotre := ctx.KVStore(k.storeKey)
	var player types.Player
	byteKey := []byte(types.PlayerPrefix + playerID)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &player)
	if err != nil {
		return player, err
	}
	return player, nil
}

// get player.GameID
func (k Keeper) GetPlayerGameID(ctx sdk.Context, playerID string) string {
	player, _ := k.GetPlayer(ctx, playerID)
	return player.GameID
}

// get player.OX
func (k Keeper) GetPlayerGameID(ctx sdk.Context, playerID string) string {
	player, _ := k.GetPlayer(ctx, playerID)
	return player.Ox
}

// set a player in KVStore
func (k Keeper) SetPlayer(ctx sdk.Context, player types.Player) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(player)
	key := []byte(types.PlayerPrefix + player.ID)
	store.Set(key, bz)
}
// set a player's GameID
func (k Keeper) SetPlayerGameID(ctx sdk.Contex, playerID string, gameID string) {
	player, _ := k.GetPlayer(ctx, playerID)
	player.GameID = gameID
	k.SetPlayer(ctx, player)	
}

// set a player's Ox
func (k Keeper) SetPlayerOX(ctx sdk.Contex, playerID string, Ox string) {
	player, _ := k.GetPlayer(ctx, playerID)
	player.Ox = Ox
	k.SetPlayer(ctx, player)	
}

// creates a player in store
func (k Keeper) CreatePlayer(ctx sdk.Context, playerID string) {
	var player types.Player
	player = types.NewPlayer(playerID)
	k.SetPlayer(ctx, player)
}

// clear player.gameID when complete a game
func (k Keeper) ClearPlayer(ctx sdk.Context, playerID string) {
	player, _ := k.GetPlayer(ctx, playerID)
	player.gameID = ""
	player.Ox = ""
	k.SetPlayer(ctx, player)
}

// deletes a player when quit the game
func (k Keeper) DeletePlayer(ctx sdk.Context, playerID string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.GamePrefix + playerID))
}

// if this player exist
func (k Keeper) PlayerExist(ctx sdk.Context, playerID string) {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.PlayerPrefix + playerID))	
}

//---------------- game related keeper functions ----------------//
// GetGame returns the game information, 
func (k Keeper) GetGame(ctx sdk.Context, gameID string) (types.Game, error) {
	store := ctx.KVStore(k.storeKey)
	var game types.Game
	byteKey := []byte(types.GamePrefix + gameID)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &game)
	if err != nil {
		return game, err
	}
	return game, nil
}

// get game.State
func (k Keeper) GetGameState(ctx skd.Context, gameID string) string {
	game, _ := k.GetGame(ctx, gameID)
	return game.State
}

// get game.CurrTurn
func (k Keeper) GetGameCurrTurn(ctx skd.Context, gameID string) string {
	game, _ := k.GetGame(ctx, gameID)
	return game.CurrTurn
}

// get the ith game.Players
func (k Keeper) GetGamePlayers(ctx sdk.Context, gameID string,  i int) string {
	game, _ := k.GetGame(ctx, gameID)
	return game.Player[i]
}

// get the game.Board[i][j]
func (k Keeper) GetGamePlayers(ctx sdk.Context, gameID string,  i int, j int) string {
	game, _ := k.GetGame(ctx, gameID)
	return game.Board[i][j]
}

// create a new game, no players
func (k Keeper) CreateGame(ctx sdk.Context, gameID string) {
	var game types.Game
	game = types.NewGame(gameID)
	k.SetGame(ctx, game)
}

// set game.Board[i][j]
func (k Keeper) SetGameBoard(ctx sdk.Context, gameID string, i int, j int, ox string) {
	game, _ := k.GetGame(ctx, gameID)
	game.Board[i][j] = ox
	k.SetGame(ctx, game)
}

// set game.CurrTurn
func (k Keeper) SetGameCurrTurn(ctx sdk.Context, gameID string, ox string) {
	game, _ := k.GetGame(ctx, gameID)
	game.CurrTurn = ox
	k.SetGame(ctx, game)
}

// open a game, only one player
func (k Keeper) OpenGame(ctx sdk.Context, playerID string, gameID string) {
	game, _ := k.GetGame(ctx, gameID)
	game.State = "open games"
	game.Players[0] = playerID
	game.CurrTurn = "O"
	k.SetGame(ctx, game)
}

// activate the game, two players
func (k Keeper) ActiveGame(ctx sdk.Context, playerID string, gameID string) {
	game, _ := k.GetGame(ctx, gameID)
	game.State = "games currently in progress"
	game.Players[1] = playerID
	k.SetGame(ctx, game)
}

// complte the game
func (k Keeper) CompleteGame(ctx sdk.Context, gameID string) {
	game, _ := k.GetGame(ctx, gameID)
	game.State = "completed games"
	game.Players[0] = ""
	game.Players[1] = ""
	k.SetGame(ctx, game)
}


// set game in KVStore
func (k Keeper) SetGame(ctx sdk.Context, game type.Game) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(game)
	key := []byte(types.GamePrefix + game.ID)
	store.Set(key, bz)
}

// Check if the key exists in the store
func (k Keeper) GameExists(ctx sdk.Context, gameID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.GamePrefix + gameID))
}


//---------------- play helper functions ----------------//
// check vertical of board
func (k Keeper) isVerticalWin(board [3][3]string, ox string) {
	res := true
	for i := 0; i < 3; i ++ {
		res = true
		for j := 0; j < 3; j ++ {
			if board[i][j] != ox { 
				res = false
				break 
			} 
		}
	}
	return res	
}

// check horizontal of board
func (k Keeper) isHorizontalWin(board [3][3]string, ox string) {
	res := true
	for i := 0; i < 3; i ++ {
		res = true
		for j := 0; j < 3; j ++ {
			if board[j][i] != ox { 
				res = false
				break 
			} 
		}
	}
	return res	
}

// check diagonal of board
func (k Keeper) isDiagonalWin(board [3][3]string, ox string) {
	res := true
	for i := 0; i < 3; i ++ {
		for j := 0; j < 3; j ++ {
			if i == j && board[i][j] != ox { 
				res := false
				break
			}
		}
	}
	if res == true { return res }
	res = true
	for i := 0; i < 3; i ++ {
		for j := 0; j < 3; j ++ {
			if i == n-j-1 && board[i][j] != ox { 
				res := false
				break
			}
		}
	}
	return res	
}

// return if the player is win
func (k Keeper) IsWin(board [3][3]string, ox string) {
	return k.isVerticalWin(board, ox) || k.isHorizontalWin(board, ox) || k.isDiagonalWin(board, ox)
}


//---------------- Functions used by querier ----------------//

func getGame(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	game, err := k.GetGame(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, game)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func getPlayer(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	game, err := k.GetPlayer(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, player)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func getGameBoard(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	game, err := k.GetGame(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, game.Board)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}