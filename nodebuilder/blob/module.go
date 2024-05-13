package blob

import (
	"context"

	"go.uber.org/fx"

	"github.com/sunriselayer/sunrise-da/blob"
	"github.com/sunriselayer/sunrise-da/header"
	headerService "github.com/sunriselayer/sunrise-da/nodebuilder/header"
	"github.com/sunriselayer/sunrise-da/share"
	"github.com/sunriselayer/sunrise-da/state"
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
