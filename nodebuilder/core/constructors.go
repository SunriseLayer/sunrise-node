package core

import (
	"github.com/sunriselayer/sunrise-da/core"
)

func remote(cfg Config) (core.Client, error) {
	return core.NewRemote(cfg.IP, cfg.RPCPort)
}
