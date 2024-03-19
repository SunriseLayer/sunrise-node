package pruner

import (
	"context"

	"github.com/sunrise-zone/sunrise-node/header"
)

// Pruner contains methods necessary to prune data
// from the node's datastore.
type Pruner interface {
	Prune(context.Context, ...*header.ExtendedHeader) error
}
