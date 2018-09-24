package cmd

import "testing"

func TestAction1Func(t *testing.T) {

	var args = []string{"string1", "string2"}
	err := Action1Func(args)

	if err != nil {
		t.Errorf("Action1Func failed with err: %s", err.Error())
	}
}