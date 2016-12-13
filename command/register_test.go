package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestRegisterCommand_implement(t *testing.T) {
	var _ cli.Command = &RegisterCommand{}
}
