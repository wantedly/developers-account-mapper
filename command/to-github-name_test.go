package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestToGithubNameCommand_implement(t *testing.T) {
	var _ cli.Command = &ToGithubNameCommand{}
}
