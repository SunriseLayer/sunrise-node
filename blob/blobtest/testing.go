package blobtest

import (
	tmrand "github.com/cometbft/cometbft/libs/rand"

	"github.com/sunriselayer/sunrise/pkg/appconsts"
	"github.com/sunriselayer/sunrise/pkg/blob"
	"github.com/sunriselayer/sunrise/pkg/shares"
	"github.com/sunriselayer/sunrise/test/util/testfactory"

	"github.com/sunrise-zone/sunrise-node/share"
)

// GenerateV0Blobs is a test utility producing v0 share formatted blobs with the
// requested size and random namespaces.
func GenerateV0Blobs(sizes []int, sameNamespace bool) ([]blob.Blob, error) {
	blobs := make([]blob.Blob, 0, len(sizes))

	for _, size := range sizes {
		size := rawBlobSize(appconsts.FirstSparseShareContentSize * size)
		appBlob := testfactory.GenerateRandomBlob(size)
		if !sameNamespace {
			nid, err := share.NewBlobNamespaceV0(tmrand.Bytes(7))
			if err != nil {
				return nil, err
			}
			appBlob.NamespaceVersion = uint32(nid[0])
			appBlob.NamespaceId = nid[1:]
		}

		blobs = append(blobs, *appBlob)
	}
	return blobs, nil
}

func rawBlobSize(totalSize int) int {
	return totalSize - shares.DelimLen(uint64(totalSize))
}
