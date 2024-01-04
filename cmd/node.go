package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/sunrise-zone/sunrise-node/nodebuilder/core"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/gateway"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/header"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/node"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/p2p"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/rpc"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/state"
)

func NewBridge(options ...func(*cobra.Command, []*pflag.FlagSet)) *cobra.Command {
	flags := []*pflag.FlagSet{
		NodeFlags(),
		p2p.Flags(),
		MiscFlags(),
		core.Flags(),
		rpc.Flags(),
		gateway.Flags(),
		state.Flags(),
	}
	cmd := &cobra.Command{
		Use:   "bridge [subcommand]",
		Args:  cobra.NoArgs,
		Short: "Manage your Bridge node",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return PersistentPreRunEnv(cmd, node.Bridge, args)
		},
	}
	for _, option := range options {
		option(cmd, flags)
	}
	return cmd
}

func NewLight(options ...func(*cobra.Command, []*pflag.FlagSet)) *cobra.Command {
	flags := []*pflag.FlagSet{
		NodeFlags(),
		p2p.Flags(),
		header.Flags(),
		MiscFlags(),
		core.Flags(),
		rpc.Flags(),
		gateway.Flags(),
		state.Flags(),
	}
	cmd := &cobra.Command{
		Use:   "light [subcommand]",
		Args:  cobra.NoArgs,
		Short: "Manage your Light node",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return PersistentPreRunEnv(cmd, node.Light, args)
		},
	}
	for _, option := range options {
		option(cmd, flags)
	}
	return cmd
}

func NewFull(options ...func(*cobra.Command, []*pflag.FlagSet)) *cobra.Command {
	flags := []*pflag.FlagSet{
		NodeFlags(),
		p2p.Flags(),
		header.Flags(),
		MiscFlags(),
		core.Flags(),
		rpc.Flags(),
		gateway.Flags(),
		state.Flags(),
	}
	cmd := &cobra.Command{
		Use:   "full [subcommand]",
		Args:  cobra.NoArgs,
		Short: "Manage your Full node",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return PersistentPreRunEnv(cmd, node.Full, args)
		},
	}
	for _, option := range options {
		option(cmd, flags)
	}
	return cmd
}
