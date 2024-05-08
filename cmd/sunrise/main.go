package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	cmdnode "github.com/sunriselayer/sunrise-da/cmd"
)

func WithSubcommands() func(*cobra.Command, []*pflag.FlagSet) {
	return func(c *cobra.Command, flags []*pflag.FlagSet) {
		c.AddCommand(
			cmdnode.Init(flags...),
			cmdnode.Start(cmdnode.WithFlagSet(flags)),
			cmdnode.AuthCmd(flags...),
			cmdnode.ResetStore(flags...),
			cmdnode.RemoveConfigCmd(flags...),
			cmdnode.UpdateConfigCmd(flags...),
		)
	}
}

func init() {
	bridgeCmd := cmdnode.NewBridge(WithSubcommands())
	lightCmd := cmdnode.NewLight(WithSubcommands())
	fullCmd := cmdnode.NewFull(WithSubcommands())
	rootCmd.AddCommand(
		bridgeCmd,
		lightCmd,
		fullCmd,
		versionCmd,
	)
	rootCmd.SetHelpCommand(&cobra.Command{})
}

func main() {
	err := run()
	if err != nil {
		os.Exit(1)
	}
}

func run() error {
	return rootCmd.ExecuteContext(context.Background())
}

var rootCmd = &cobra.Command{
	Use: "sunrise-da [  bridge  ||  full ||  light  ] [subcommand]",
	Short: `
	   _____                      _
	  / ____|                    (_)
	 | (___   _   _  _ __   _ __  _  ___   ___
	  \___ \ | | | || '_ \ | '__|| |/ __| / _ \
	  ____) || |_| || | | || |   | |\__ \|  __/
	 |_____/  \__,_||_| |_||_|   |_||___/ \___|
	`,
	Args: cobra.NoArgs,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: false,
	},
}
