package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	QueryFarm = "farm"
)

// NewQuerier creates a new querier for tracability clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case QueryFarm:
			return queryFarm(ctx, path[1:], req, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown tracability query endpoint")
		}
	}
}

func queryFarm(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	id := path[0]
	if id == "" {
		return []byte{}, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "could not resolve id")
	}
	farm, err := k.GetFarm(ctx, id)
	if err != nil {
		return nil, err
	}
	res, err := codec.MarshalJSONIndent(k.cdc, farm)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

// TODO: Add the modules query functions
// They will be similar to the above one: queryParams()
