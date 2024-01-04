package nodebuilder

import (
	"github.com/sunrise-zone/sunrise-node/nodebuilder/blob"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/das"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/fraud"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/header"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/node"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/p2p"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/share"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/state"
)

// PackageToAPI maps a package to its API struct. Currently only used for
// method discovery for openrpc spec generation
var PackageToAPI = map[string]interface{}{
	"fraud":  &fraud.API{},
	"state":  &state.API{},
	"share":  &share.API{},
	"header": &header.API{},
	"das":    &das.API{},
	"p2p":    &p2p.API{},
	"blob":   &blob.API{},
	"node":   &node.API{},
}
