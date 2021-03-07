package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgBuyName struct {
	Name  string
	Buyer sdk.AccAddress
	Bid   sdk.Coins
}

func NewMsgBuyName(owner sdk.AccAddress, name string, buyer sdk.AccAddress, bid sdk.Coins) MsgBuyName {
	return MsgBuyName{
		Name:  name,
		Buyer: buyer,
		Bid:   bid,
	}
}

func (msg MsgBuyName) Route() string {
	return RouterKey
}

func (msg MsgBuyName) Type() string {
	return "BuyName"
}

func (msg MsgBuyName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer.Bytes()}
}

func (msg MsgBuyName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgBuyName) ValidateBasic() error {
	if msg.Buyer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Buyer cannot be empty")
	}
	if msg.Bid.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Bid must be a valid number")
	}
	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name cannot be empty")
	}
	return nil
}
