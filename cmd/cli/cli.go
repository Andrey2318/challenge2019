package cli

import (
	cliInterface "challenge2019/cmd/cli/cli"
	"github.com/urfave/cli/v2"
)

func App() *cli.App {
	return &cli.App{
		Name:        "challenge",
		Description: "Program",
		Commands: []*cli.Command{
			cliInterface.Command(),
		},
	}
}
