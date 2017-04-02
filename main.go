package main

import (
	"os"

	"github.com/itsubaki/apst/ranking"
	"github.com/itsubaki/apst/review"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "app store command line tool."
	app.Version = "0.0.2"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "limit, l",
			Value: "30",
		},
		cli.StringFlag{
			Name:  "genre, g",
			Value: "",
		},
		cli.StringFlag{
			Name:  "country, c",
			Value: "jp",
		},
		cli.StringFlag{
			Name:  "feed, f",
			Value: "grossing",
			Usage: "grossing, free, paid.",
		},
	}

	rank := cli.Command{
		Name:    "ranking",
		Aliases: []string{"r"},
		Usage:   "Show app store ranking",
		Action:  ranking.Action,
		Flags:   flags,
	}

	review := cli.Command{
		Name:    "review",
		Aliases: []string{"rv"},
		Usage:   "Show app store review",
		Action:  review.Action,
		Flags: append(flags,
			cli.BoolFlag{
				Name:  "stats, s",
				Usage: "Show stats.",
			}),
	}

	app.Commands = []cli.Command{
		rank,
		review,
	}

	app.Run(os.Args)
}
