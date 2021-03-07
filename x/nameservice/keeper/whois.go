package keeper

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/palacios95/nameservice/x/nameservice/types"
)

//
// Keeper CRUD operations
//

// GetWhoisCount get the total number of whois
func (k Keeper) GetWhoisCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.WhoisCountPrefix)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetWhoisCount set the total number of whois
func (k Keeper) SetWhoisCount(ctx sdk.Context, count int64) {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.WhoisCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// GetWhois returns the whois information
func (k Keeper) GetWhois(ctx sdk.Context, name string) (types.Whois, error) {
	store := ctx.KVStore(k.storeKey)
	var whois types.Whois
	byteKey := []byte(types.WhoisPrefix + name)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &whois)
	if err != nil {
		return whois, err
	}
	return whois, nil
}

// SetWhois sets a whois
func (k Keeper) SetWhois(ctx sdk.Context, name string, whois types.Whois) {
	whoisKey := name
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(whois)
	key := []byte(types.WhoisPrefix + whoisKey)
	store.Set(key, bz)
}

// DeleteWhois deletes a whois
func (k Keeper) DeleteWhois(ctx sdk.Context, name string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.WhoisPrefix + name))
}

//GetName gets the value of the name
func (k Keeper) GetName(ctx sdk.Context, key string) (string, error) {
	whois, err := k.GetWhois(ctx, key)
	if err != nil {
		return "", err
	}
	return whois.Value, nil
}

// SetName sets the value corresponding to the name.
func (k Keeper) SetName(ctx sdk.Context, name string, value string) {
	whois, _ := k.GetWhois(ctx, name)
	whois.Value = value
	k.SetWhois(ctx, name, whois)
}

//GetOwner gets the owner of the item
func (k Keeper) GetOwner(ctx sdk.Context, key string) (sdk.AccAddress, error) {
	whois, err := k.GetWhois(ctx, key)
	if err != nil {
		return nil, err
	}
	return whois.Owner, nil
}

//SetOwner sets the Owner of a whois
func (k Keeper) SetOwner(ctx sdk.Context, name string, owner sdk.AccAddress) {
	whois, _ := k.GetWhois(ctx, name)
	whoisKey := name
	whois.Owner = owner
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(whois)
	key := []byte(types.WhoisPrefix + whoisKey)
	store.Set(key, bz)
}

//GetPrice gets the owner of the item
func (k Keeper) GetPrice(ctx sdk.Context, key string) (sdk.Coins, error) {
	whois, err := k.GetWhois(ctx, key)
	if err != nil {
		return nil, err
	}
	return whois.Price, nil
}

//SetPrice sets the Price of a whois
func (k Keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {
	whois, _ := k.GetWhois(ctx, name)
	whoisKey := name
	whois.Price = price
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(whois)
	key := []byte(types.WhoisPrefix + whoisKey)
	store.Set(key, bz)
}

// WhoisExists ...
func (k Keeper) WhoisExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.WhoisPrefix + key))
}

//
// Functions used by querier
//

func listWhois(ctx sdk.Context, k Keeper) ([]byte, error) {
	var whoisList []types.Whois
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.WhoisPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var whois types.Whois
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &whois)
		whoisList = append(whoisList, whois)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, whoisList)
	return res, nil
}

func getWhois(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	whois, err := k.GetWhois(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, whois)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
