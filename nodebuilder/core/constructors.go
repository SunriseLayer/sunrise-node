package core

import (
	"github.com/sunrise-zone/sunrise-node/core"
)

func remote(cfg Config) (core.Client, error) {
	return core.NewRemote(cfg.IP, cfg.RPCPort)
}
