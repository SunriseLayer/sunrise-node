package blob

import (
	"context"

	"go.uber.org/fx"

	"github.com/sunrise-zone/sunrise-node/blob"
	"github.com/sunrise-zone/sunrise-node/header"
	headerService "github.com/sunrise-zone/sunrise-node/nodebuilder/header"
	"github.com/sunrise-zone/sunrise-node/share"
	"github.com/sunrise-zone/sunrise-node/state"
)

func ConstructModule() fx.Option {
	return fx.Module("blob",
		fx.Provide(
			func(service headerService.Module) func(context.Context, uint64) (*header.ExtendedHeader, error) {
				return service.GetByHeight
			}),
		fx.Provide(func(
			state *state.CoreAccessor,
			sGetter share.Getter,
			getByHeightFn func(context.Context, uint64) (*header.ExtendedHeader, error),
		) Module {
			return blob.NewService(state, sGetter, getByHeightFn)
		}))
}
