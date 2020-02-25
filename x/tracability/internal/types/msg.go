package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgAddFarm Farm

func NewMsgAddFarm(id, name, location string, owner sdk.AccAddress) MsgAddFarm {
	return MsgAddFarm(NewFarm(id, name, location, owner))
}

func (msg MsgAddFarm) Route() string { return RouterKey }
func (msg MsgAddFarm) Type() string { return "msg_add_farm" }

func (msg MsgAddFarm) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if msg.Id == "" || msg.Location == "" || msg.Name == "" {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Id, Location and Name must not be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAddFarm) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgAddFarm) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

