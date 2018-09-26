package cmd

import (
	"github.com/martin-guthrie-docker/goclipkgtemplate/pkg/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdState)
}

var cmdState = &cobra.Command{
	Use:   "state",
	Short: "print your name to the console",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		stateFunc(args)
	},
}

func stateFunc(args []string) error {
	log.Term.Info("Start")

	// CmdConfig is from cmd_root_exec.go, global
	GlobalConfig.Dump(false, "")

	log.Term.Debug("Done")
	return nil
}