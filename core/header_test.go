package core

import (
	"context"
	"testing"

	"github.com/cometbft/cometbft/libs/rand"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sunriselayer/sunrise-da/header"
	"github.com/sunriselayer/sunrise-da/header/headertest"
	"github.com/sunriselayer/sunrise-da/share"
)

func TestMakeExtendedHeaderForEmptyBlock(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	client := StartTestNode(t).RpcClient
	fetcher := NewBlockFetcher(client)

	sub, err := fetcher.SubscribeNewBlockEvent(ctx)
	require.NoError(t, err)
	<-sub

	height := int64(1)
	b, err := fetcher.GetBlock(ctx, &height)
	require.NoError(t, err)

	comm, val, err := fetcher.GetBlockInfo(ctx, &height)
	require.NoError(t, err)

	eds, err := extendBlock(b.Data, b.Header.Version.App)
	require.NoError(t, err)

	headerExt, err := header.MakeExtendedHeader(&b.Header, comm, val, eds)
	require.NoError(t, err)

	assert.Equal(t, share.EmptyRoot(), headerExt.DAH)
}

func TestMismatchedDataHash_ComputedRoot(t *testing.T) {
	header := headertest.RandExtendedHeader(t)
	header.DataHash = rand.Bytes(32)

	err := header.Validate()
	assert.Contains(t, err.Error(), "mismatch between data hash commitment from"+
		" core header and computed data root")
}
