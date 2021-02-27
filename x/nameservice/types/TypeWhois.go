package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

"Struct of value mapping for names"
type Whois struct {
	Owner sdk.AccAddress `json:"creator" yaml:"creator"`
	Value string         `json:"value" yaml:"value"`
	Price string         `json:"price" yaml:"price"`
}
