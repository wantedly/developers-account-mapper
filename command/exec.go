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
