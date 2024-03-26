package main

import (
	"github.com/sunrise-zone/sunrise-node/cmd"
	blob "github.com/sunrise-zone/sunrise-node/nodebuilder/blob/cmd"
	das "github.com/sunrise-zone/sunrise-node/nodebuilder/das/cmd"
	header "github.com/sunrise-zone/sunrise-node/nodebuilder/header/cmd"
	node "github.com/sunrise-zone/sunrise-node/nodebuilder/node/cmd"
	p2p "github.com/sunrise-zone/sunrise-node/nodebuilder/p2p/cmd"
	share "github.com/sunrise-zone/sunrise-node/nodebuilder/share/cmd"
	state "github.com/sunrise-zone/sunrise-node/nodebuilder/state/cmd"
)

func init() {
	blob.Cmd.PersistentFlags().AddFlagSet(cmd.RPCFlags())
	das.Cmd.PersistentFlags().AddFlagSet(cmd.RPCFlags())
	header.Cmd.PersistentFlags().AddFlagSet(cmd.RPCFlags())
	p2p.Cmd.PersistentFlags().AddFlagSet(cmd.RPCFlags())
	share.Cmd.PersistentFlags().AddFlagSet(cmd.RPCFlags())
	state.Cmd.PersistentFlags().AddFlagSet(cmd.RPCFlags())
	node.Cmd.PersistentFlags().AddFlagSet(cmd.RPCFlags())

	rootCmd.AddCommand(
		blob.Cmd,
		das.Cmd,
		header.Cmd,
		p2p.Cmd,
		share.Cmd,
		state.Cmd,
		node.Cmd,
	)
}
