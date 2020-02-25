package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dnk90/tracability/x/tracability/internal/types"
)

// Keeper of the tracability store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
}

// NewKeeper creates a tracability keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		storeKey:   key,
		cdc:        cdc,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetFarm(ctx sdk.Context, id string) (types.Farm, error) {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(id)) {
		return types.Farm{}, fmt.Errorf("farmId is not found")
	}
	farmByte := store.Get([]byte(id))
	var farm types.Farm
	k.cdc.MustUnmarshalBinaryBare(farmByte, &farm)
	return farm, nil
}

func (k Keeper) SetFarm(ctx sdk.Context, farm types.Farm) error {
	store := ctx.KVStore(k.storeKey)
	if store.Has([]byte(farm.Id)) {
		return fmt.Errorf("farmId existed")
	}
	store.Set([]byte(farm.Id), k.cdc.MustMarshalBinaryBare(farm))
	return nil
}

func (k Keeper) AddQC(ctx sdk.Context, id, qc string) error {
	store := ctx.KVStore(k.storeKey)
	farm, err := k.GetFarm(ctx, id)
	if err != nil {
		return err
	}
	if _, hasKey := farm.QCs[qc]; hasKey {
		return fmt.Errorf("qc existed")
	}
	farm.QCs[qc] = struct{}{}
	store.Set([]byte(farm.Id), k.cdc.MustMarshalBinaryBare(farm))
	return nil
}

func (k Keeper) delete(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(key))
}

func (k Keeper) IterateAllKey(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}
