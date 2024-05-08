package main

import (
	"github.com/sunriselayer/sunrise-da/cmd"
	blob "github.com/sunriselayer/sunrise-da/nodebuilder/blob/cmd"
	das "github.com/sunriselayer/sunrise-da/nodebuilder/das/cmd"
	header "github.com/sunriselayer/sunrise-da/nodebuilder/header/cmd"
	node "github.com/sunriselayer/sunrise-da/nodebuilder/node/cmd"
	p2p "github.com/sunriselayer/sunrise-da/nodebuilder/p2p/cmd"
	share "github.com/sunriselayer/sunrise-da/nodebuilder/share/cmd"
	state "github.com/sunriselayer/sunrise-da/nodebuilder/state/cmd"
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
