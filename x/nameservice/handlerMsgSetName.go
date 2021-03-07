package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/palacios95/nameservice/x/nameservice/keeper"
	"github.com/palacios95/nameservice/x/nameservice/types"
)

func handleMsgSetName(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetName) (*sdk.Result, error) {

	owner, _ := k.GetOwner(ctx, msg.Name)

	if !msg.Owner.Equals(owner) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetName(ctx, msg.Name, msg.Value)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
