package keeper

import (
	"fmt"
	"crypto/md5"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	crypto "github.com/tendermint/tendermint/crypto"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/r24zeng/tttgame/x/tttgame/types"
)

//---------------- game related keeper functions ----------------//
// CreateGame create a game which does not exist, no players
func (k Keeper) CreateGame(ctx sdk.Context, gameID string) {
	var game types.Game
	game.ID = gameID
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			game.Board[i][j] = -1
		}
	}
	game.Players = []
	game.CurrTurn = 0
	k.SetGame(ctx, game)
}

// GetGame returns the game information
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
func (k Keeper) GetGameState(ctx sdk.Context, gameID string) string {
	game, _ := k.GetGame(ctx, gameID)
	if game.Players == nil { 
		return "completed games" 
	}
	if len(game.Players) == 1 {
		return "open games"
	} 
	return "games currently in progress"
}

// get current player
func (k Keeper) GetGameCurrPlayer(ctx sdk.Context, gameID string) sdk.AccAddress {
	game, _ := k.GetGame(ctx, gameID)
	return game.Players[game.CurrTurn]
}

// open a game(the game exists), only one player, if the player doesn't exist, create a new player
func (k Keeper) OpenGame(ctx sdk.Context, pub crypto.PubKey, gameID string) {
	game, _ := k.GetGame(ctx, gameID)
	game.Players.append(pub)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			game.Board[i][j] = -1
		}
	}
	game.CurrTurn = 0
	k.SetGame(ctx, game)
}

// activate the game, two players, if the player doesn't exist, create a new player 
func (k Keeper) ActiveGame(ctx sdk.Context, pub crypto.PubKey, gameID string) error {
	game, _ := k.GetGame(ctx, gameID)
	
	// if pub = inviter, then invalid
	if game.Players[0] == pub {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Inviter can't be opponent, accept fail")
	}

	game.Players.append(pub)
	hash_value = md5.Sum(game.Players[0] + game.Players[1])
	if hash_value[0]/16 != 0 {
		tmp := game.Player[0]
		game.Player[1] = game.Player[0]
		game.Player[0] = tmp
	}
	k.SetGame(ctx, game)
	return nil
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

// place the piece to game board and set the current turn
func (k Keeper) PlayGame(ctx sdk.Context, gameID string, i int, j int) {
	game, _ := k.GetGame(ctx, gameID)
	game.Board[i][j] = game.CurrTurn
	if IsWin(ctx, game.Board, i, j, game.CurrTurn) {
		game.Players = nil	
	} else {
		game.CurrTurn = game.CurrTurn == 0? 1: 0
	}
	k.SetGame(ctx, game)
}

func (k Keeper) IsWin(ctx sdk.Context, board int, i int, j int, curr int) bool {
	isWin := false
	if game.Board[i][0] == curr & game.Board[i][1] == curr & game.Board[i][2] == curr {
		isWin = true
	}
	if game.Board[0][j] == curr & game.Board[1][j] == curr & game.Board[2][j] == curr {
		isWin = true
	}
	if i == j & game.Board{
		if game.Board[0][0] == curr & game.Board[1][1] == curr & game.Board[2][2] == curr {
			isWin = true
		}
	} else {
		if game.Board[0][2] == curr & game.Board[1][1] == curr & game.Board[2][0] == curr {
			isWin = true
		}
	}
	return isWin	
}

// Get an iterator over all GameIDs in which the keys are the gameIDs and the values are the Games
func (k Keeper) GetGameIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.GamePrefix))
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

func listGame(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	var gameList []types.Game
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.GamePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var game types.Game
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &game)
		gameList = append(gameList, game)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, gameList)
	return res, nil
}