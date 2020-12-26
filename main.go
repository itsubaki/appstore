package main

import (
	"os"

	"github.com/itsubaki/appstore/pkg/ranking"
	"github.com/itsubaki/appstore/pkg/review"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "appstore"
	app.Usage = "app store command line tool."
	app.Version = "0.0.3"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "limit, l",
			Value: "20",
		},
		cli.StringFlag{
			Name:  "genre, g",
			Value: "all",
			Usage: "all, games",
		},
		cli.StringFlag{
			Name:  "country, c",
			Value: "jp",
		},
		cli.StringFlag{
			Name:  "feed, f",
			Value: "top-grossing",
			Usage: "top-grossing, top-free, top-paid",
		},
		cli.StringFlag{
			Name:  "format",
			Value: "plain",
			Usage: "plain, json",
		},
		cli.BoolFlag{
			Name:  "pretty, p",
			Usage: "Print pretty json.",
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
