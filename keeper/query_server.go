package keeper

import (
	"context"
	"fmt"

	"github.com/cosmosregistry/example"
)

var _ example.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the module QueryServer.
func NewQueryServerImpl(k Keeper) example.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}

// Counter defines the handler for the Query/Counter RPC method.
func (qs queryServer) Counter(ctx context.Context, req *example.QueryCounterRequest) (*example.QueryCounterResponse, error) {
	if _, err := qs.k.addressCodec.StringToBytes(req.Address); err != nil {
		return nil, fmt.Errorf("invalid sender address: %w", err)
	}

	counter, err := qs.k.Counter.Get(ctx, req.Address)
	if err != nil {
		return nil, err
	}

	return &example.QueryCounterResponse{Counter: counter}, nil
}

// Params defines the handler for the Query/Params RPC method.
func (qs queryServer) Params(ctx context.Context, req *example.QueryParamsRequest) (*example.QueryParamsResponse, error) {
	params, err := qs.k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &example.QueryParamsResponse{Params: params}, nil
}
