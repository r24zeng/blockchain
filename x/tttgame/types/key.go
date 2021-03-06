package types

const (
	// ModuleName is the name of the module
	ModuleName = "tttgame"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)

// GamePrefix is for storing game in KVStore
const (
	GamePrefix = "game-"
)
