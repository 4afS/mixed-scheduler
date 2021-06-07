package command

import "testing"

func TestCommand(t *testing.T) {
	command := Command()
	if command != "command" {
		t.Error("test failed")
	}
}
