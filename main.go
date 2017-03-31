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
				Name:  "limit, l",
				Value: "200",
			},
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "genre, g",
				Value: "",
			},
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "feed, f",
				Value: "grossing",
               Usage: "grossing, free, paid",
			},
		},
	}

	app.Commands = []cli.Command{
		top,
	}

	app.Run(os.Args)
}
