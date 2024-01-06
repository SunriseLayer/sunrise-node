package gateway

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sunrise-zone/sunrise-app/pkg/appconsts"
	"github.com/sunrise-zone/sunrise-app/pkg/blob"
	"github.com/sunrise-zone/sunrise-app/pkg/shares"

	"github.com/sunrise-zone/sunrise-node/share/sharetest"
)

func Test_dataFromShares(t *testing.T) {
	testData := [][]byte{
		[]byte("beep"),
		[]byte("beeap"),
		[]byte("BEEEEAHP"),
	}

	ns := sharetest.RandV0Namespace()
	sss := shares.NewSparseShareSplitter()
	for _, data := range testData {
		b := &blob.Blob{
			Data:             data,
			NamespaceId:      ns.ID(),
			NamespaceVersion: uint32(ns.Version()),
			ShareVersion:     uint32(appconsts.ShareVersionZero),
		}
		err := sss.Write(b)
		require.NoError(t, err)
	}

	sssShares := sss.Export()

	rawSSSShares := make([][]byte, len(sssShares))
	for i := 0; i < len(sssShares); i++ {
		d := sssShares[i].ToBytes()
		rawSSSShares[i] = d
	}

	parsedSSSShares, err := dataFromShares(rawSSSShares)
	require.NoError(t, err)

	require.Equal(t, testData, parsedSSSShares)
}
