package main

import (
	"github.com/itsubaki/apst/top"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "app store command line tool."
	app.Version = "0.0.1"

	top := cli.Command{
		Name:    "top",
		Aliases: []string{"t"},
		Usage:   "app store top-sales",
		Action:  top.Action,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "limit",
				Value: "200",
			},
		},
	}

	app.Commands = []cli.Command{
		top,
	}

	app.Run(os.Args)
}
