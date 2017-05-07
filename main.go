package main

import (
	"os"

	"gopkg.in/urfave/cli.v1"

	"github.com/itsubaki/apst/genre"
	"github.com/itsubaki/apst/ranking"
	"github.com/itsubaki/apst/review"
)

func main() {
	app := cli.NewApp()
	app.Name = "app store command line tool."
	app.Version = "0.0.3"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "limit, l",
			Value: "20",
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
		cli.StringFlag{
			Name:  "output, o",
			Value: "plain",
			Usage: "plain, json, jsonp",
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

	genre := cli.Command{
		Name:    "genre",
		Aliases: []string{"g"},
		Usage:   "Show app genre",
		Action:  genre.Action,
	}

	app.Commands = []cli.Command{
		rank,
		review,
		genre,
	}

	app.Run(os.Args)
}
