//go:build !race

package state

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"cosmossdk.io/math"
	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	"github.com/sunriselayer/sunrise/app"
	"github.com/sunriselayer/sunrise/test/util/genesis"
	"github.com/sunriselayer/sunrise/test/util/testnode"
	blobtypes "github.com/sunriselayer/sunrise/x/blob/types"

	"github.com/sunrise-zone/sunrise-node/blob"
	"github.com/sunrise-zone/sunrise-node/share"
)

func TestSubmitPayForBlob(t *testing.T) {
	accounts := genesis.NewAccounts(1000000000, "jimy", "rob")
	tmCfg := testnode.DefaultTendermintConfig()
	tmCfg.Consensus.TimeoutCommit = time.Millisecond * 1
	appConf := testnode.DefaultAppConfig()
	appConf.API.Enable = true
	appConf.MinGasPrices = fmt.Sprintf("0.002%s", app.BondDenom)

	config := testnode.DefaultConfig().WithTendermintConfig(tmCfg).WithAppConfig(appConf)
	config.Genesis = config.Genesis.WithAccounts(accounts...)
	cctx, rpcAddr, grpcAddr := testnode.NewNetwork(t, config)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signer := blobtypes.NewKeyringSigner(cctx.Keyring, accounts[0].Name, cctx.ChainID)
	ca := NewCoreAccessor(signer, nil, "127.0.0.1", extractPort(rpcAddr), extractPort(grpcAddr))
	// start the accessor
	err := ca.Start(ctx)
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = ca.Stop(ctx)
	})

	ns, err := share.NewBlobNamespaceV0([]byte("namespace"))
	require.NoError(t, err)
	blobbyTheBlob, err := blob.NewBlobV0(ns, []byte("data"))
	require.NoError(t, err)

	minGas, err := ca.queryMinimumGasPrice(ctx)
	require.NoError(t, err)
	require.Equal(t, float64(0), minGas)

	testcases := []struct {
		name   string
		blobs  []*blob.Blob
		fee    math.Int
		gasLim uint64
		expErr error
	}{
		{
			name:   "empty blobs",
			blobs:  []*blob.Blob{},
			fee:    sdkmath.ZeroInt(),
			gasLim: 0,
			expErr: errors.New("state: no blobs provided"),
		},
		{
			name:   "good blob with user provided gas and fees",
			blobs:  []*blob.Blob{blobbyTheBlob},
			fee:    sdkmath.NewInt(10_000), // roughly 0.12 utia per gas (should be good)
			gasLim: blobtypes.DefaultEstimateGas([]uint32{uint32(len(blobbyTheBlob.Data))}),
			expErr: nil,
		},
		// TODO: add more test cases. The problem right now is that the celestia-app doesn't
		// correctly construct the node (doesn't pass the min gas price) hence the price on
		// everything is zero and we can't actually test the correct behavior
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := ca.SubmitPayForBlob(ctx, tc.fee, tc.gasLim, tc.blobs)
			require.Equal(t, tc.expErr, err)
			if err == nil {
				require.EqualValues(t, 0, resp.Code)
			}
		})
	}

}

func extractPort(addr string) string {
	splitStr := strings.Split(addr, ":")
	return splitStr[len(splitStr)-1]
}
