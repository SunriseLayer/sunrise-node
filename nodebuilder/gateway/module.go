package gateway

import (
	"context"

	logging "github.com/ipfs/go-log/v2"
	"go.uber.org/fx"

	"github.com/sunriselayer/sunrise-da/api/gateway"
	headerServ "github.com/sunriselayer/sunrise-da/nodebuilder/header"
	"github.com/sunriselayer/sunrise-da/nodebuilder/node"
	shareServ "github.com/sunriselayer/sunrise-da/nodebuilder/share"
	stateServ "github.com/sunriselayer/sunrise-da/nodebuilder/state"
)

var log = logging.Logger("module/gateway")

func ConstructModule(tp node.Type, cfg *Config) fx.Option {
	// sanitize config values before constructing module
	cfgErr := cfg.Validate()
	if !cfg.Enabled {
		return fx.Options()
	}

	baseComponents := fx.Options(
		fx.Supply(cfg),
		fx.Error(cfgErr),
		fx.Provide(fx.Annotate(
			server,
			fx.OnStart(func(ctx context.Context, server *gateway.Server) error {
				return server.Start(ctx)
			}),
			fx.OnStop(func(ctx context.Context, server *gateway.Server) error {
				return server.Stop(ctx)
			}),
		)),
	)

	switch tp {
	case node.Light, node.Full:
		return fx.Module(
			"gateway",
			baseComponents,
			fx.Invoke(Handler),
		)
	case node.Bridge:
		return fx.Module(
			"gateway",
			baseComponents,
			fx.Invoke(func(
				state stateServ.Module,
				share shareServ.Module,
				header headerServ.Module,
				serv *gateway.Server,
			) {
				Handler(state, share, header, nil, serv)
			}),
		)
	default:
		panic("invalid node type")
	}
}
