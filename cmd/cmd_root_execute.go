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
// FIXME: Does this need to be global scoped?  Better to use a struct at least?
var log *logrus.Logger = logrus.New()  // init global logger
var logLevel int = 5

var cfgEnvVarsPrefix = "GOCLIP"  // vars in format GOCLIP_<key>
var cfgFileName string = ".goclipkgtemplate"
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
	rootCmd.PersistentFlags().IntVar(&logLevel, "logLevel", 5, "Set the logging level [0=panic, 3=warning, 5=debug]")

	configUsage := fmt.Sprintf("config file (default is $HOME/%s, ./%s)", cfgFileName, cfgFileName)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", configUsage)

	// root command
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

	// FIXME: this doesn't work because logrus hasn't parsed flags yet!
	switch logLevel {
	case 0: log.Level = logrus.PanicLevel
	case 1: log.Level = logrus.FatalLevel
	case 2: log.Level = logrus.ErrorLevel
	case 3: log.Level = logrus.WarnLevel
	case 4: log.Level = logrus.InfoLevel  // this is default by the flag
	case 5: log.Level = logrus.DebugLevel
	}

	filenameHook := filename.NewHook()
	filenameHook.Field = "src"
	log.AddHook(filenameHook)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func logConfig() {
	log.Info("Using config file:", viper.ConfigFileUsed())
	log.Info("Manufacturer: ", viper.Get("Manufacturer"))
	log.Info("Options: ", viper.Get("Options"))

	// TODO: how can the viper vars be accessed elsewhere in the program?
	//       should the fields be populated in a global struct? or is Viper global?
}

// got this from https://github.com/spf13/cobra.... not sure how/what it does yet
// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Since the user specified a config file, throw error, exit if not found
		log.Debug("target config file: [", cfgFile, "]")
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)  // includes full path

		if err := viper.ReadInConfig(); err != nil {
			log.Error("config file: ", err.Error())
			fmt.Println(err)
			os.Exit(1)
		}
		logConfig()

	} else {
		// Search config in home and current directory with name "cfgFileName"
		home, err := homedir.Dir()
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
		log.Debug("target config file: [", home, "/", cfgFileName, "]")
		// see https://github.com/spf13/viper/issues/390
		viper.SetConfigType("yaml")
		viper.AddConfigPath(home)
		viper.SetConfigName(cfgFileName)

		if err := viper.ReadInConfig(); err != nil {
			log.Warn("config file: ", err.Error())
			// NOTE: if this program needs external config, then error and exit here
		} else {
			logConfig()
		}
	}

	// get environment vars
	viper.SetEnvPrefix(cfgEnvVarsPrefix)
	viper.AutomaticEnv() // read in environment variables that match

	one := viper.Get("one")  // without the prefix, vars in format <prefix>_<key>
	log.Info("Environment var one: ", one)

}
