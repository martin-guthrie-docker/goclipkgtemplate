package cmd

import (
	"github.com/spf13/cobra"

	"github.com/martin-guthrie-docker/goclipkgtemplate/pkg/goclipkgtemplate"
)

func init() {
	rootCmd.AddCommand(cmdAction1)
}

var cmdAction1 = &cobra.Command{
	Use:   "action1 <string1> <string2>",
	Short: "print your name to the console",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		Action1Func(args)
	},
}

// functions are pulled out of the cobra struct for
// purpose of unit testing
func Action1Func (args []string) error {

	foo, err := goclipkgtemplate.NewFooClass(
		args[0],
		goclipkgtemplate.FooClassCfg{
			Name: args[1],
			Log: log,
		})

	if err != nil {
		log.Fatalf("NewFooClass failed")
		return err
	}

	foo.Open()
	foo.Close()
	return nil
}
