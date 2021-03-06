package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteName{}

type MsgDeleteName struct {
	ID    string         `json:"id" yaml:"id"`
	Owner sdk.AccAddress `json:"owner" yaml:"owner"`
}

func NewMsgDeleteName(id string, owner sdk.AccAddress) MsgDeleteName {
	return MsgDeleteName{
		ID:    id,
		Owner: owner,
	}
}

func (msg MsgDeleteName) Route() string {
	return RouterKey
}

func (msg MsgDeleteName) Type() string {
	return "DeleteName"
}

func (msg MsgDeleteName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg MsgDeleteName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteName) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Owner can't be empty")
	}
	return nil
}
