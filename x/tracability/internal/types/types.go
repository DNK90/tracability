package types

import "github.com/cosmos/cosmos-sdk/types"

type (
	Farm struct {
		Id    string               `json:"FarmId"` // It might be an uuid4
		Owner types.AccAddress     `json:"Owner"` //address of farm's owner
		Name  string               `json:"Name"`
		QCs   map[string]struct{}  `json:"QCs"`
		Location string            `json:"Location"`
	}
)

func NewFarm(id, name, location string, owner types.AccAddress) Farm {
	return Farm {
		Id: id,
		Name: name,
		QCs: make(map[string]struct{}),
		Location: location,
		Owner: owner,
	}
}

