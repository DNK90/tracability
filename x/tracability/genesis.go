package tracability

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.Farms {
		err := k.SetFarm(ctx, record)
		if err != nil {
			panic(err)
		}
	}
	return []abci.ValidatorUpdate{}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k Keeper) (data GenesisState) {
	var (
		records []Farm
		farm Farm
		err error
	)
	iterator := k.IterateAllKey(ctx)
	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		if farm, err = k.GetFarm(ctx, key); err != nil {
			panic(err)
		}
		records = append(records, farm)
	}
	return NewGenesisState(records)
}
