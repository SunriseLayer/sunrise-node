package blobtest

import (
	tmrand "github.com/cometbft/cometbft/libs/rand"
	"github.com/cometbft/cometbft/types"

	"github.com/sunrise-zone/sunrise-app/pkg/appconsts"
	"github.com/sunrise-zone/sunrise-app/pkg/shares"
	"github.com/sunrise-zone/sunrise-app/test/util/testfactory"

	"github.com/sunrise-zone/sunrise-node/share"
)

// GenerateV0Blobs is a test utility producing v0 share formatted blobs with the
// requested size and random namespaces.
func GenerateV0Blobs(sizes []int, sameNamespace bool) ([]types.Blob, error) {
	blobs := make([]types.Blob, 0, len(sizes))

	for _, size := range sizes {
		size := rawBlobSize(appconsts.FirstSparseShareContentSize * size)
		appBlob := testfactory.GenerateRandomBlob(size)
		if !sameNamespace {
			nid, err := share.NewBlobNamespaceV0(tmrand.Bytes(7))
			if err != nil {
				return nil, err
			}
			appBlob.NamespaceVersion = nid[0]
			appBlob.NamespaceID = nid[1:]
		}

		blobs = append(blobs, appBlob)
	}
	return blobs, nil
}

func rawBlobSize(totalSize int) int {
	return totalSize - shares.DelimLen(uint64(totalSize))
}
