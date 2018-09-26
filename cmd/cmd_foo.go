package cmd

import (
	"github.com/spf13/cobra"

	"github.com/martin-guthrie-docker/goclipkgtemplate/pkg/goclipkgtemplate"
	log "github.com/martin-guthrie-docker/goclipkgtemplate/pkg/log"

)

func init() {
	rootCmd.AddCommand(cmdFoo)
}

var cmdFoo = &cobra.Command{
	Use:   "foo <string1> <string2>",
	Short: "print <string1> <string2> to the console",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fooFunc(args)
	},
}

func fooFunc(args []string) error {
	log.Term.Info("Start")

	foo, err := goclipkgtemplate.NewFooClass(
		args[0],
		goclipkgtemplate.FooClassCfg{
			Name: args[1],
			Log:  log.Term,
		})

	if err != nil {
		log.Term.Fatalf("NewFooClass failed")
		return err
	}

	foo.Open()
	foo.Close()
	log.Term.Debug("Done")
	return nil
}
