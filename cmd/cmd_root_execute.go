package cmd

import (
	"bytes"
	"fmt"
	"github.com/martin-guthrie-docker/goclipkgtemplate/pkg/goclipkgtemplate"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/martin-guthrie-docker/goclipkgtemplate/pkg/log"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// TODO: these come from the makefile (or goxc?) - figure this out
// GoCLIPkgTemplateVersion is the release TAG
var GoCLIPkgTemplateVersion string
// GoCLIPkgTemplateBuild is the current GIT commit
var GoCLIPkgTemplateBuild string

var verboseFlag bool = false

var cfgEnvVarsPrefix = "GOCLIP"  // vars in format GOCLIP_<key>
var cfgFileName string = ".goclipkgtemplate"
var cfgFile string

var GlobalConfig goclipkgtemplate.ConfigClass

var rootCmd = &cobra.Command{
	Use:   "goclipkgtemplate",
	Short: "This tool is a template for creating CLI tools with a PKG option",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("verbFlag") {
			v := int(logrus.DebugLevel)
			setLoggingLevel(v)
		}
	},
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
	// Global flag across all subcommands
	rootCmd.PersistentFlags().BoolVar(&verboseFlag, "verbose", false,"Set debug level")

	configUsage := fmt.Sprintf("config file (default is $HOME/%s, ./%s.yaml)", cfgFileName, cfgFileName)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", configUsage)

	// root command
	rootCmd.AddCommand(cmdVersion)

	viper.BindPFlags(rootCmd.PersistentFlags())
	viper.BindPFlag("verbFlag", rootCmd.PersistentFlags().Lookup("verbose"))

	cobra.OnInitialize(initConfig)

	GlobalConfig, err := goclipkgtemplate.NewConfigClass(viper.GetViper(),
										 			  goclipkgtemplate.ConfigClassCfg{ Log: log.Term } )
	if err != nil {
		log.Term.Error(err)
		os.Exit(1)
	}
	GlobalConfig.Open()
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
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func setLoggingLevel(level int) {
	switch level {
	case 0:
		log.Term.Level = logrus.PanicLevel
	case 1:
		log.Term.Level = logrus.FatalLevel
	case 2:
		log.Term.Level = logrus.ErrorLevel
	case 3:
		log.Term.Level = logrus.WarnLevel
	case 4:
		log.Term.Level = logrus.InfoLevel
	case 5:
		log.Term.Level = logrus.DebugLevel
	}
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// from https://github.com/spf13/cobra

	// uncomment in order to log this func. This occurs before loglevel is set.
	// setLoggingLevel(5)

	if cfgFile != "" {
		// Since the user specified a config file, throw error, exit if not found
		log.Term.Debug("target config file: [", cfgFile, "]")
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)  // includes full path

		if err := viper.ReadInConfig(); err != nil {
			log.Term.Error("config file: ", err.Error())
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		// Search config in home and current directory with name "cfgFileName"
		home, err := homedir.Dir()
		if err != nil {
			log.Term.Error(err)
			os.Exit(1)
		}
		log.Term.Debug("target config file: [", home, "/", cfgFileName, "]")
		// see https://github.com/spf13/viper/issues/390
		viper.SetConfigType("yaml")
		viper.AddConfigPath(home)
		viper.SetConfigName(cfgFileName)

		if err := viper.ReadInConfig(); err != nil {
			log.Term.Warn("config file: ", err.Error())
			// NOTE: if this program needs external config, then error and exit here
		}
	}

	if viper.Get("verbosity") != nil {
		v := viper.Get("verbosity").(int)
		setLoggingLevel(v)
	}

	// get environment vars
	viper.SetEnvPrefix(cfgEnvVarsPrefix)
	viper.AutomaticEnv() // read in environment variables that match
}
