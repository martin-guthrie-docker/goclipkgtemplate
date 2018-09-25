package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-cray/logrus-prefixed-formatter"
)

// TODO: these come from the makefile (or goxc?) - figure this out
// GoCLIPkgTemplateVersion is the release TAG
var GoCLIPkgTemplateVersion string

// GoCLIPkgTemplateBuild is the current GIT commit
var GoCLIPkgTemplateBuild string

//LogPointer to have same logging in pkg and cmds
// FIXME: Does this need to be global scoped?  Better to use a stuct at least?
var log = logrus.New()
var logLevel int

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "goclipkgtemplate",
	Short: "This tool is a template for creating CLI tools with a PKG option",
}

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Version and Release information about the goclipkgtemplate",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("goclipkgtemplate, a CLI tool example\n")
		fmt.Printf("Version:  %s\n", GoCLIPkgTemplateVersion)
		fmt.Printf("Build:    %s\n", GoCLIPkgTemplateBuild)
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flag across all subcommands
	rootCmd.PersistentFlags().IntVar(&logLevel, "logLevel", 4, "Set the logging level [0=panic, 3=warning, 5=debug]")
	rootCmd.AddCommand(cmdVersion)
}

// ExecuteCommand executes commands, intended for testing
func ExecuteCommand(args ...string) (output string, err error) {
	_, output, err = executeCommandC(rootCmd, args...)
	return output, err
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOutput(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}

// Execute - starts the command parsing process
func Execute() {
	log.Formatter = new(prefixed.TextFormatter)
	log.Level = logrus.DebugLevel

	filenameHook := filename.NewHook()
	filenameHook.Field = "src"
	log.AddHook(filenameHook)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// got this from https://github.com/spf13/cobra.... not sure how/what it does yet
// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra-example" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra-example")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
