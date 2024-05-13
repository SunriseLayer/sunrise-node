package nodebuilder

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sunriselayer/sunrise-da/core"
	coremodule "github.com/sunriselayer/sunrise-da/nodebuilder/core"
	"github.com/sunriselayer/sunrise-da/nodebuilder/node"
	"github.com/sunriselayer/sunrise-da/nodebuilder/p2p"
)

func TestBridge_WithMockedCoreClient(t *testing.T) {
	t.Skip("skipping") // consult https://github.com/celestiaorg/celestia-core/issues/667 for reasoning
	repo := MockStore(t, DefaultConfig(node.Bridge))

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	client := core.StartTestNode(t).RpcClient
	node, err := New(node.Bridge, p2p.Private, repo, coremodule.WithClient(client))
	require.NoError(t, err)
	require.NotNil(t, node)
	err = node.Start(ctx)
	require.NoError(t, err)

	err = node.Stop(ctx)
	require.NoError(t, err)
}
