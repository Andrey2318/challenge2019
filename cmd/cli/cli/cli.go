package cli

import (
	"challenge2019/internal/container"
	Cli "challenge2019/internal/interfaces/cli/v1"
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	c := Cli.New(container.TrafficApplication())
	return &cli.Command{
		Name:  "cli",
		Usage: "cli interface",
		Subcommands: []*cli.Command{
			c.Statement1(),
			c.Statement2(),
		},
	}
}
