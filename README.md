# goclipkgtemplate
A template for starting a Go CLI/PKG tool with nice(er) logging and unit test.

Go dependancies
* "github.com/spf13/cobra"
* "github.com/sirupsen/logrus"
* "github.com/onrik/logrus/filename"
* "github.com/x-cray/logrus-prefixed-formatter"
* "github.com/spf13/viper"

The motivation is to create a sample directory structure and naming format for
creating a CLI tool that can also be used as a package (pkg).

This example has two pkg modules,

* foo
  * Prints command line strings to the log (console)
* config
  * transfers configuration data to a go struct, presumably for more consistant
    access by other parts of the program, rather than doing viper.get()

    
## Configuration

The program can be configured two different ways,

* yaml file
  * a file of name .goclipkgtemplate.yaml
  * by default the program looks for this file in your home directory
  * using '--config <path_to_file>' a different file can be specified
* environment variables 
  * variables with prefix 'GOCLIP_' are imported
  
Note that currently yaml and environment variables are hard coded, meaning
that the name of the variable to import is coded into th program before it
can be used.  Its possible that arbitrary variables can be used...


