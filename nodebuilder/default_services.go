package nodebuilder

import (
	"github.com/sunriselayer/sunrise-da/nodebuilder/blob"
	"github.com/sunriselayer/sunrise-da/nodebuilder/das"
	"github.com/sunriselayer/sunrise-da/nodebuilder/fraud"
	"github.com/sunriselayer/sunrise-da/nodebuilder/header"
	"github.com/sunriselayer/sunrise-da/nodebuilder/node"
	"github.com/sunriselayer/sunrise-da/nodebuilder/p2p"
	"github.com/sunriselayer/sunrise-da/nodebuilder/share"
	"github.com/sunriselayer/sunrise-da/nodebuilder/state"
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
