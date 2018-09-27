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
	Short: "print state to the console/log",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		stateFunc(args)
	},
}

func stateFunc(args []string) error {
	log.Term.Info("State:")
	// CmdConfig is from cmd_root_exec.go, global
	GlobalConfig.Dump(true, false, "")
	log.Term.Info("State: ---")
	return nil
}