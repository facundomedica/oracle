package example

import "cosmossdk.io/collections"

const (
	ModuleName = "example"

	// StoreKey is the default store key for mint
	StoreKey = ModuleName
)

var (
	ParamsKey  = collections.NewPrefix(0)
	CounterKey = collections.NewPrefix(1)
)
