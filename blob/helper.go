package blob

import (
	"bytes"
	"sort"

	pkgblob "github.com/sunrise-zone/sunrise-app/pkg/blob"
	"github.com/sunrise-zone/sunrise-app/pkg/shares"

	"github.com/sunrise-zone/sunrise-node/share"
)

// SharesToBlobs takes raw shares and converts them to the blobs.
func SharesToBlobs(rawShares []share.Share) ([]*Blob, error) {
	if len(rawShares) == 0 {
		return nil, ErrBlobNotFound
	}

	appShares, err := toAppShares(rawShares...)
	if err != nil {
		return nil, err
	}
	return parseShares(appShares)
}

// parseShares takes shares and converts them to the []*Blob.
func parseShares(appShrs []shares.Share) ([]*Blob, error) {
	shareSequences, err := shares.ParseShares(appShrs, true)
	if err != nil {
		return nil, err
	}

	// ensure that sequence length is not 0
	if len(shareSequences) == 0 {
		return nil, ErrBlobNotFound
	}

	blobs := make([]*Blob, len(shareSequences))
	for i, sequence := range shareSequences {
		data, err := sequence.RawData()
		if err != nil {
			return nil, err
		}
		if len(data) == 0 {
			continue
		}

		shareVersion, err := sequence.Shares[0].Version()
		if err != nil {
			return nil, err
		}

		blob, err := NewBlob(shareVersion, sequence.Namespace.Bytes(), data)
		if err != nil {
			return nil, err
		}
		blobs[i] = blob
	}
	return blobs, nil
}

// BlobsToShares accepts blobs and convert them to the Shares.
func BlobsToShares(blobs ...*Blob) ([]share.Share, error) {
	b := make([]*pkgblob.Blob, len(blobs))
	for i, blob := range blobs {
		namespace := blob.Namespace()
		b[i] = &pkgblob.Blob{
			NamespaceVersion: uint32(namespace[0]),
			NamespaceId:      namespace[1:],
			Data:             blob.Data,
			ShareVersion:     uint32(blob.ShareVersion),
		}
	}

	sort.Slice(b, func(i, j int) bool {
		val := bytes.Compare(b[i].NamespaceId, b[j].NamespaceId)
		return val <= 0
	})

	rawShares, err := shares.SplitBlobs(b...)
	if err != nil {
		return nil, err
	}
	return shares.ToBytes(rawShares), nil
}
