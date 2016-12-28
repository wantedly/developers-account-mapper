package command

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"strings"

	"github.com/wantedly/developers-account-mapper/store"
)

type SetenvCommand struct {
	Meta
}

func (c *SetenvCommand) Run(args []string) int {
	loginName := args[0]
	if args[1] != "exec" {
		log.Println("$ developers-account-mapper setenv <user> exec")
		return 1
	}

	s := store.NewDynamoDB()

	user, err := s.GetUserByLoginName(loginName)
	if err != nil {
		log.Println(err)
		return 1
	}

	envs := os.Environ()
	envs = append(envs, fmt.Sprintf("%s=%s", "GITHUB_USERNAME", user.GitHubUsername))

	execCmd := exec.Command(args[1], args[2:]...)
	execCmd.Env = envs
	execCmd.Stderr = os.Stderr
	execCmd.Stdout = os.Stdout
	err = execCmd.Run()

	if execCmd.Process == nil {
		log.Println(err)
	}

	os.Exit(execCmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus())

	return 0
}

func (c *SetenvCommand) Synopsis() string {
	return "Set account information as env vars and exec commands"
}

func (c *SetenvCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
