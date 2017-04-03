package review

import (
	"fmt"
	"os"

	"github.com/itsubaki/apst/client"
	"github.com/itsubaki/apst/ranking"
	"github.com/itsubaki/apst/util"
	cli "gopkg.in/urfave/cli.v1"
)

func Action(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Println("See: apst review -h")
		os.Exit(1)
	}

	country := c.String("country")
	b := client.Ranking(
		util.Limit(c.String("limit")),
		c.String("genre"),
		c.String("feed"),
		country,
	)

	kw := c.Args().Get(0)
	list := ranking.NewFeed(b).Select(kw)

	for i, app := range list {
		fmt.Println(app)

		b := client.Review(app.ID, country)
		f := NewFeed(b)

		for _, r := range f.ReviewList {
			r.ColorPrintln()
		}

		if c.Bool("stats") {
			fmt.Println(f.Stats())
		}

		if i != (len(list) - 1) {
			fmt.Println("")
		}

	}

}
