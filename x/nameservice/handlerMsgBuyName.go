package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/palacios95/nameservice/x/nameservice/keeper"
	"github.com/palacios95/nameservice/x/nameservice/types"
)

func handleMsgBuyName(ctx sdk.Context, k keeper.Keeper, msg types.MsgBuyName) (*sdk.Result, error) {
	whois, _ := k.GetWhois(ctx, msg.Name)

	if whois.Price.IsAllGT(msg.Bid) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid must be higher than the current price")
	}

	if whois.Creator.Empty() {
		k.CoinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid)
	} else {
		k.CoinKeeper.SendCoins(ctx, msg.Buyer, whois.Creator, msg.Bid)
	}

	k.SetWhois(ctx, msg.Name, types.Whois{Creator: msg.Buyer, ID: whois.ID, Value: whois.Value, Price: msg.Bid})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
