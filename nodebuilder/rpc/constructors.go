package rpc

import (
	"github.com/cristalhq/jwt"

	"github.com/sunrise-zone/sunrise-node/api/rpc"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/blob"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/da"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/das"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/fraud"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/header"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/node"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/p2p"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/share"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/state"
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
