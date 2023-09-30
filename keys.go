package oracle

import "cosmossdk.io/collections"

const ModuleName = "oracle"

var (
	ParamsKey  = collections.NewPrefix(0)
	CounterKey = collections.NewPrefix(1)
)
