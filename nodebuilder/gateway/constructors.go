package gateway

import (
	"github.com/sunrise-zone/sunrise-node/api/gateway"
	"github.com/sunrise-zone/sunrise-node/das"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/header"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/share"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/state"
)

// Handler constructs a new RPC Handler from the given services.
func Handler(
	state state.Module,
	share share.Module,
	header header.Module,
	daser *das.DASer,
	serv *gateway.Server,
) {
	handler := gateway.NewHandler(state, share, header, daser)
	handler.RegisterEndpoints(serv)
	handler.RegisterMiddleware(serv)
}

func server(cfg *Config) *gateway.Server {
	return gateway.NewServer(cfg.Address, cfg.Port)
}
