package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosregistry/example"
)

type msgServer struct {
	Keeper
}

var _ example.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) example.MsgServer {
	return &msgServer{Keeper: keeper}
}

// CreateExample defines the handler for the MsgCreateExample message.
func (srv msgServer) CreateExample(ctx context.Context, msg *example.MsgExample) (*example.MsgExampleResponse, error) {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return nil, err
	}

	if msg.Data == "" {
		return nil, fmt.Errorf("data cannot be empty")
	}

	// TODO: implement the module's logic here

	return &example.MsgExampleResponse{}, nil
}

// UpdateParams params is defining the handler for the MsgUpdateParams message.
func (srv msgServer) UpdateParams(ctx context.Context, msg *example.MsgUpdateParams) (*example.MsgUpdateParamsResponse, error) {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return nil, err
	}

	if strings.EqualFold(msg.Authority, srv.Keeper.authority) {
		return nil, fmt.Errorf("unauthorized, authority does not match the module's authority")
	}

	// TODO: implement the module's logic here

	return &example.MsgUpdateParamsResponse{}, nil
}
