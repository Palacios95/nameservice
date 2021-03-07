package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/palacios95/nameservice/x/nameservice/keeper"
	"github.com/palacios95/nameservice/x/nameservice/types"
)

// Handle a message to delete name
func handleMsgDeleteName(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteName) (*sdk.Result, error) {
	if !k.WhoisExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	owner, err := k.GetOwner(ctx, msg.ID)

	if err != nil {
		return nil, sdkerrors.Wrap(err, "An error has ocurred retrieving the Owner of the whois")
	}
	if !msg.Owner.Equals(owner) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteWhois(ctx, msg.ID)
	return &sdk.Result{}, nil
}
