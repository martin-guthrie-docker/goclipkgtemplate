package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

// TODO: these come from the makefile (or goxc?) - figure this out
// GoCLIPkgTemplateVersion is the release TAG
var GoCLIPkgTemplateVersion string
// GoCLIPkgTemplateBuild is the current GIT commit
var GoCLIPkgTemplateBuild string

//LogPointer to have same logging in pkg and cmds
// FIXME: Does this need to be global scoped?  Better to use a struct at least?
var log *logrus.Logger = logrus.New()  // init global logger
var logLevel int

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
	rootCmd.PersistentFlags().IntVar(&logLevel, "logLevel", 4, "Set the logging level [0=panic, 3=warning, 5=debug]")

	configUsage := fmt.Sprintf("config file (default is $HOME/%s, ./%s)", cfgFileName, cfgFileName)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", configUsage)

	// root command
	rootCmd.AddCommand(cmdVersion)
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
		log.Debug("target config file: [", cfgFile, "]")
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}

		// Search config in home and current directory with name "cfgFileName"
		log.Debug("target config file: [", home, "/", cfgFileName, "]")
		log.Debug("target config file: [./", cfgFileName, "]")
		viper.AddConfigPath(home)
		//viper.AddConfigPath(".")
		viper.SetConfigName(cfgFileName)
	}

	viper.SetEnvPrefix(cfgEnvVarsPrefix)
	viper.AutomaticEnv() // read in environment variables that match

	one := viper.Get("one")  // without the prefix, vars in format <prefix>_<key>
	log.Info("Environment var one: ", one)

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file:", viper.ConfigFileUsed())
		log.Info("Manufacturer: ", viper.Get("Manufacturer"))
	} else{
		log.Warn("Error config file: ", err.Error())
	}


}

