package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Whois struct of value mapping for names
type Whois struct {
	Owner sdk.AccAddress `json:"owner" yaml:"owner"`
	ID    string         `json:"id" yaml:"id"`
	Value string         `json:"value" yaml:"value"`
	Price sdk.Coins      `json:"price" yaml:"price"`
}
