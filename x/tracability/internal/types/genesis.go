package types

import "fmt"

type GenesisState struct {
	Farms []Farm   `json:"farms"`
}

func NewGenesisState(farms []Farm) GenesisState {
	return GenesisState{Farms: farms}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Farms: []Farm{},
	}
}

// ValidateGenesis validates the tracability genesis parameters
func ValidateGenesis(data GenesisState) error {
	for _, farm := range data.Farms {
		if farm.Owner == nil {
			return fmt.Errorf("invalid Farmer: Missing Owner")
		}
		if farm.Id == "" {
			return fmt.Errorf("invalid Farmer: Missing Id")
		}
		if farm.Name == "" {
			return fmt.Errorf("invalid Farmer: Missing Name")
		}
		if farm.Location == "" {
			return fmt.Errorf("invalid Farmer: Missing Location")
		}
	}
	return nil
}
