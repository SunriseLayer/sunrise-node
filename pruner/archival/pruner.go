package archival

import (
	"context"

	"github.com/sunriselayer/sunrise-da/header"
)

// Pruner is a noop implementation of the pruner.Factory interface
// that allows archival nodes to sync and retain historical data
// that is out of the availability window.
type Pruner struct{}

func NewPruner() *Pruner {
	return &Pruner{}
}

func (p *Pruner) Prune(context.Context, ...*header.ExtendedHeader) error {
	return nil
}
