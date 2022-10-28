package cmd

import "github.com/urfave/cli/v2"

//commands
func Commands() []*cli.Command {
	cmd := []*cli.Command{
		{
			Name:   "new",
			Usage:  "micro new",
			Action: newService,
		},
		{
			Name:   "build",
			Usage:  "micro build .",
			Action: buildService,
		},
		{
			Name:   "lint",
			Usage:  "micro lint .",
			Action: lintService,
		},
		{
			Name:   "deploy",
			Usage:  "micro deploy .",
			Action: deployService,
		},
	}
	return cmd
}
