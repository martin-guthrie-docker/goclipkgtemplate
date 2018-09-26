package cmd_test

import (
	"testing"

	"github.com/martin-guthrie-docker/goclipkgtemplate/cmd"
)

func TestAction1Func(t *testing.T) {

	_, err := cmd.ExecuteCommand("foo", "string1", "string2")

	if err != nil {
		t.Errorf("fooFunc failed with err: %s", err.Error())
	}
}
