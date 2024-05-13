package light

import (
	"context"

	"github.com/sunriselayer/sunrise-da/header"
)

type Pruner struct{}

func NewPruner() *Pruner {
	return &Pruner{}
}

func (p *Pruner) Prune(context.Context, ...*header.ExtendedHeader) error {
	return nil
}
