package core

import (
	"go.uber.org/fx"

	"github.com/sunriselayer/sunrise-da/core"
	"github.com/sunriselayer/sunrise-da/header"
	"github.com/sunriselayer/sunrise-da/libs/fxutil"
)

// WithClient sets custom client for core process
func WithClient(client core.Client) fx.Option {
	return fxutil.ReplaceAs(client, new(core.Client))
}

// WithHeaderConstructFn sets custom func that creates extended header
func WithHeaderConstructFn(construct header.ConstructFn) fx.Option {
	return fx.Replace(construct)
}
