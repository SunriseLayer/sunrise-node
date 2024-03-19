package nodebuilder

import (
	"context"

	"go.uber.org/fx"

	"github.com/sunrise-zone/sunrise-node/header"
	"github.com/sunrise-zone/sunrise-node/libs/fxutil"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/blob"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/core"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/da"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/das"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/fraud"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/gateway"
	modhead "github.com/sunrise-zone/sunrise-node/nodebuilder/header"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/node"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/p2p"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/prune"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/rpc"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/share"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/state"
)

func ConstructModule(tp node.Type, network p2p.Network, cfg *Config, store Store) fx.Option {
	log.Infow("Accessing keyring...")
	ks, err := store.Keystore()
	if err != nil {
		fx.Error(err)
	}
	signer, err := state.KeyringSigner(cfg.State, ks, network)
	if err != nil {
		fx.Error(err)
	}

	baseComponents := fx.Options(
		fx.Supply(tp),
		fx.Supply(network),
		fx.Provide(p2p.BootstrappersFor),
		fx.Provide(func(lc fx.Lifecycle) context.Context {
			return fxutil.WithLifecycle(context.Background(), lc)
		}),
		fx.Supply(cfg),
		fx.Supply(store.Config),
		fx.Provide(store.Datastore),
		fx.Provide(store.Keystore),
		fx.Supply(node.StorePath(store.Path())),
		fx.Supply(signer),
		// modules provided by the node
		p2p.ConstructModule(tp, &cfg.P2P),
		state.ConstructModule(tp, &cfg.State, &cfg.Core),
		modhead.ConstructModule[*header.ExtendedHeader](tp, &cfg.Header),
		share.ConstructModule(tp, &cfg.Share),
		gateway.ConstructModule(tp, &cfg.Gateway),
		core.ConstructModule(tp, &cfg.Core),
		das.ConstructModule(tp, &cfg.DASer),
		fraud.ConstructModule(tp),
		blob.ConstructModule(),
		da.ConstructModule(),
		node.ConstructModule(tp),
		prune.ConstructModule(tp),
		rpc.ConstructModule(tp, &cfg.RPC),
	)

	return fx.Module(
		"node",
		baseComponents,
	)
}
