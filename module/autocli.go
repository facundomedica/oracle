package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	examplev1 "github.com/cosmosregistry/example/api/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: examplev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "ListExample",
					Use:       "list",
					Short:     "List all sent example messages",
				},
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Get the current module parameters",
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: examplev1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod:      "CreateExample",
					Use:            "create [data]",
					Short:          "Create a new example message",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "data"}},
				},
				// The params tx is purposely left empty, as the only tx is MsgUpdateParams which is gov gated.
			},
		},
	}
}
