package tracability

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler creates an sdk.Handler for all the tracability type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgAddFarm:
			return handleMsgAddFarm(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName,  msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// Handle a message to set name
func handleMsgAddFarm(ctx sdk.Context, keeper Keeper, msg MsgAddFarm) (*sdk.Result, error) {
	err := keeper.SetFarm(ctx, Farm(msg))
	if err != nil {
		return nil, err
	}
	return &sdk.Result{}, nil
}
