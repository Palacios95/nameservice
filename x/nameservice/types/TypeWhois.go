package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Whois struct of value mapping for names
type Whois struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
	Value   string         `json:"value" yaml:"value"`
	Price   sdk.Coins      `json:"price" yaml:"price"`
}
