package main

import (
	"github.com/mitchellh/cli"
	"github.com/wantedly/developers-account-mapper/command"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"register": func() (cli.Command, error) {
			return &command.RegisterCommand{
				Meta: *meta,
			}, nil
		},
		"list": func() (cli.Command, error) {
			return &command.ListCommand{
				Meta: *meta,
			}, nil
		},
		"setenv": func() (cli.Command, error) {
			return &command.ExecCommand{
				Meta: *meta,
			}, nil
		},
		"to-github-name": func() (cli.Command, error) {
			return &command.ToGithubNameCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
