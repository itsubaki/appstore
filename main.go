package main

import (
	"github.com/itsubaki/apst/ranking"
	"github.com/itsubaki/apst/review"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "app store command line tool."
	app.Version = "0.0.2"

	rank := cli.Command{
		Name:    "ranking",
		Aliases: []string{"r"},
		Usage:   "app store ranking",
		Action:  ranking.Action,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "limit, l",
				Value: "200",
			}, cli.StringFlag{
				Name:  "genre, g",
				Value: "",
			},
			cli.StringFlag{
				Name:  "feed, f",
				Value: "grossing",
				Usage: "grossing, free, paid",
			},
		},
	}

	review := cli.Command{
		Name:    "review",
		Aliases: []string{"rv"},
		Usage:   "app store review",
		Action:  review.Action,
	}

	app.Commands = []cli.Command{
		rank,
		review,
	}

	app.Run(os.Args)
}
