package keeper

import (
	"context"

	"github.com/cosmosregistry/example"
)

func (k *Keeper) InitGenesis(ctx context.Context, data *example.GenesisState) {

}

func (k *Keeper) ExportGenesis(ctx context.Context) *example.GenesisState {
	return &example.GenesisState{}
}
