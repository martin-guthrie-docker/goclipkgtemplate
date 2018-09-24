package main

import (
	"github.com/martin-guthrie-docker/goclipkgtemplate/cmd"
)

// TODO: check if goxc is the way to assign these, or use make file?
// Version is populated from the Makefile and is tied to the release TAG
var Version string = "0.0.1"
// Build is the last GIT commit
var Build string = "DEADBEEF"

func init() {
	cmd.GoCLIPkgTemplateVersion = Version
	cmd.GoCLIPkgTemplateBuild = Build
}

func main() {

	cmd.Execute()
}
