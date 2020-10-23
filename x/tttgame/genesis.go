package tttgame

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/r24zeng/tttgame/x/tttgame/keeper"
	"github.com/r24zeng/tttgame/x/tttgame/types"
	// abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k keeper.Keeper /* TODO: Define what keepers the module needs */, data types.GenesisState) {
	// TODO: Define logic for when you would like to initalize a new genesis
	for _, record := range data.PlayerRecords {
		keeper.SetPlayer(ctx, record)
	}
	for _, record := range data.GameRecords {
		keeper.SetGame(ctx, record)
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) (data types.GenesisState) {
	// TODO: Define logic for exporting state
	var records []types.Player
	iterator := k.GetPlayerIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		ID := string(iterator.Key())
		player, _ := k.GetPlayer(ctx, ID)
		records = append(records, player)

	}

	var records []types.Game
	iterator := k.GetGameIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		ID := string(iterator.Key())
		game, _ := k.GetGame(ctx, ID)
		records = append(records, game)

	}
	return types.GenesisState{PlayerRecords: records}
}
