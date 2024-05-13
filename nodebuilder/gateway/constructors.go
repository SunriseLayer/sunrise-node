package gateway

import (
	"github.com/sunriselayer/sunrise-da/api/gateway"
	"github.com/sunriselayer/sunrise-da/das"
	"github.com/sunriselayer/sunrise-da/nodebuilder/header"
	"github.com/sunriselayer/sunrise-da/nodebuilder/share"
	"github.com/sunriselayer/sunrise-da/nodebuilder/state"
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
