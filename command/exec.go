package command

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/wantedly/developers-account-mapper/store"
)

type ExecCommand struct {
	Meta
}

func (c *ExecCommand) Run(args []string) int {
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

	execCmd := exec.Command(args[2], args[3:]...)
	execCmd.Env = envs
	execCmd.Stderr = os.Stderr
	execCmd.Stdout = os.Stdout
	err = execCmd.Run()
	if err != nil {
		log.Println(err)
		return 1
	}

	return 0
}

func (c *ExecCommand) Synopsis() string {
	return "Set account information as env vars and exec commands"
}

func (c *ExecCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
