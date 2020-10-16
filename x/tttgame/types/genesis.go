package types

import "fmt"

// GenesisState - all tttgame state that must be provided at genesis
type GenesisState struct {
	// TODO: Fill out what is needed by the module for genesis
	PlayerRecords []Player   `json: "player_records"`
	GameRecords []Game    `json: "game_records"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState( /* TODO: Fill out with what is needed for genesis state */ ) GenesisState {
	return GenesisState{
		// TODO: Fill out according to your genesis state
		PlayerRecords: nil,
		GameRecords: nil,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
		// TODO: Fill out according to your genesis state, these values will be initialized but empty
		PlayerRecords: []Player{},
		GameRecords: []Game{},
	}
}

// ValidateGenesis validates the tttgame genesis parameters
func ValidateGenesis(data GenesisState) error {
	// TODO: Create a sanity check to make sure the state conforms to the modules needs
	for _, record := range data.PlayerRecords {
		if record.Ox == nil && record.GameID != nil {
			return fmt.Errorf("Invalid PlayerRecord: %s. Error: Missing Ox", record.ID)
		}
		if record.GameID == nil && record.Ox != nil {
			return fmt.Errorf("Invalid PlayerRecord: %s. Error: Missing GameID", record.ID)
		}
	}		
	for _, record := range data.GameRecords {
		if record.CurrTurn == nil {
			return fmt.Errorf("Invalid GameRecord: %s. Error: Missing CurrTurn", record.ID)
		}
		if record.State == "open games" {
			if !(record.Player[0] != "" && record.Player[1] == "") {
				return fmt.Errorf("Invalid GameRecord: %s. Error: 'Open state' doesn't match players", record.ID)
		}	
		if record.State == "games currently in progress" {
			if record.Player[0] == "" && record.Player[1] == "" {
				return fmt.Errorf("Invalid GameRecord: %s. Error: 'In progress state' doesn't match players", record.ID)
		}	
		if record.State == "complete games" {
			if record.Player[0] != "" && record.Player[1] != "" {
				return fmt.Errorf("Invalid GameRecord: %s. Error: 'Open state' doesn't match players", record.ID)
			}
		}
	}	
	return nil
}
