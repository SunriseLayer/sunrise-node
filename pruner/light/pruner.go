package light

import (
	"context"

	"github.com/sunrise-zone/sunrise-node/header"
)

type Pruner struct{}

func NewPruner() *Pruner {
	return &Pruner{}
}

func (p *Pruner) Prune(context.Context, ...*header.ExtendedHeader) error {
	return nil
}
