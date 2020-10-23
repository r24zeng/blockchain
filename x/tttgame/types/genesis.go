package types

import "fmt"

// GenesisState - all tttgame state that must be provided at genesis
type GenesisState struct {
	// TODO: Fill out what is needed by the module for genesis
	GameRecords []Game `json: "game_records"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState( /* TODO: Fill out with what is needed for genesis state */ ) GenesisState {
	return GenesisState{
		// TODO: Fill out according to your genesis state
		GameRecords: nil,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
		// TODO: Fill out according to your genesis state, these values will be initialized but empty
		GameRecords: []Game{},
	}
}

// ValidateGenesis validates the tttgame genesis parameters
func ValidateGenesis(data GenesisState) error {
	// TODO: Create a sanity check to make sure the state conforms to the modules needs
	for _, record := range data.GameRecords {
		if record.GameID == nil {
			return fmt.Errorf("Invalid GameRecord: %s. Error: Missing GameID", record.ID)
		}
		if record.Board == nil {
			return fmt.Errorf("Invalid GameRecord: %s. Error: Missing Board", record.ID)
		}
		if record.CurrTurn == nil {
			return fmt.Errorf("Invalid GameRecord: %s. Error: Missing CurrTurn", record.ID)
		}
	}
	return nil
}
