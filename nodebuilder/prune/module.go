package prune

import (
	"context"

	"go.uber.org/fx"

	"github.com/sunrise-zone/sunrise-node/nodebuilder/node"
	"github.com/sunrise-zone/sunrise-node/pruner"
	"github.com/sunrise-zone/sunrise-node/pruner/archival"
	"github.com/sunrise-zone/sunrise-node/pruner/light"
)

func ConstructModule(tp node.Type) fx.Option {
	baseComponents := fx.Options(
		fx.Provide(fx.Annotate(
			pruner.NewService,
			fx.OnStart(func(ctx context.Context, p *pruner.Service) error {
				return p.Start(ctx)
			}),
			fx.OnStop(func(ctx context.Context, p *pruner.Service) error {
				return p.Stop(ctx)
			}),
		)),
	)

	switch tp {
	case node.Full, node.Bridge:
		return fx.Module("prune",
			baseComponents,
			fx.Provide(func() pruner.Pruner {
				return archival.NewPruner()
			}),
			fx.Supply(archival.Window),
		)
	case node.Light:
		return fx.Module("prune",
			baseComponents,
			fx.Provide(func() pruner.Pruner {
				return light.NewPruner()
			}),
			fx.Supply(light.Window),
		)
	default:
		panic("unknown node type")
	}
}
