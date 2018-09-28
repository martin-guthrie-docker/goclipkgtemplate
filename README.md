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
    * GOCLIP_ONE is accepted (export GOCLIP_ONE=igotthis)
  
Note that currently yaml and environment variables are hard coded, meaning
that the name of the variable to import is coded into th program before it
can be used.  Its possible that arbitrary variables can be used...

## Examples

```
$ ./goclipkgtemplate foo hello world
Hello! My name is world
And my host is world
```

```
$ ./goclipkgtemplate --verbose foo hello world
[0000]  INFO       command.go: 852:cobra.(*Command).ExecuteC     | Start
[0000]  INFO       command.go: 766:cobra.(*Command).execute      | Hello! My name is world
Hello! My name is world
[0000]  INFO       command.go: 766:cobra.(*Command).execute      | And my host is world
And my host is world
[0000] DEBUG       command.go: 852:cobra.(*Command).ExecuteC     | Done

```

```bash
$ export GOCLIP_ONE=hello
$ cat ./.goclipkgtemplate.yaml 
Manufacturer: Hyundai
Model: Elantra
Year: 2018
Options:
  leather
  sun roof
# add verbosity 0-Panic, 1-fatal, 2-error, 3-warn(default), 4-Info, 5-Debug
verbosity: 3
$ ./goclipkgtemplate --config ./.goclipkgtemplate.yaml state
CarData: Hyundai Elantra 2018 [leather sun roof]
ENV GOCLIP_ONE: hello

```