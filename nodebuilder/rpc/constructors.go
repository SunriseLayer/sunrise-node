package rpc

import (
	"github.com/cristalhq/jwt"

	"github.com/sunriselayer/sunrise-da/api/rpc"
	"github.com/sunriselayer/sunrise-da/nodebuilder/blob"
	"github.com/sunriselayer/sunrise-da/nodebuilder/da"
	"github.com/sunriselayer/sunrise-da/nodebuilder/das"
	"github.com/sunriselayer/sunrise-da/nodebuilder/fraud"
	"github.com/sunriselayer/sunrise-da/nodebuilder/header"
	"github.com/sunriselayer/sunrise-da/nodebuilder/node"
	"github.com/sunriselayer/sunrise-da/nodebuilder/p2p"
	"github.com/sunriselayer/sunrise-da/nodebuilder/share"
	"github.com/sunriselayer/sunrise-da/nodebuilder/state"
)

// registerEndpoints registers the given services on the rpc.
func registerEndpoints(
	stateMod state.Module,
	shareMod share.Module,
	fraudMod fraud.Module,
	headerMod header.Module,
	daserMod das.Module,
	p2pMod p2p.Module,
	nodeMod node.Module,
	blobMod blob.Module,
	daMod da.Module,
	serv *rpc.Server,
) {
	serv.RegisterService("fraud", fraudMod, &fraud.API{})
	serv.RegisterService("das", daserMod, &das.API{})
	serv.RegisterService("header", headerMod, &header.API{})
	serv.RegisterService("state", stateMod, &state.API{})
	serv.RegisterService("share", shareMod, &share.API{})
	serv.RegisterService("p2p", p2pMod, &p2p.API{})
	serv.RegisterService("node", nodeMod, &node.API{})
	serv.RegisterService("blob", blobMod, &blob.API{})
	serv.RegisterService("da", daMod, &da.API{})
}

func server(cfg *Config, auth jwt.Signer) *rpc.Server {
	return rpc.NewServer(cfg.Address, cfg.Port, cfg.SkipAuth, auth)
}
