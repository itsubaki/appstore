package review

import (
	"fmt"
	"os"
	"time"

	"github.com/itsubaki/apst/client"
	"github.com/itsubaki/apst/genre"
	"github.com/itsubaki/apst/ranking"
	"github.com/itsubaki/apst/util"
	"gopkg.in/urfave/cli.v1"
)

func Action(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Println("See: apst review -h")
		os.Exit(1)
	}

	fmt.Println(time.Now())
	country := c.String("country")
	b := client.Ranking(
		util.Limit(c.String("limit")),
		genre.ID(c.String("genre")),
		c.String("feed"),
		country,
	)

	list := ranking.NewAppFeed(b).AppList
	for i := 0; i < len(c.Args()); i++ {
		list = list.Select(c.Args().Get(i))
	}

	for i, app := range list {
		b := client.Review(app.ID, country)
		f := NewReviewFeed(b)

		switch c.String("output") {
		case "json":
			fmt.Println(f.Json())
		case "jsonp":
			fmt.Println(f.JsonPretty())
		default:
			fmt.Println(app)
			for _, r := range f.ReviewList {
				util.ColorPrintln(r.Rating, r.String())
			}

			if c.Bool("stats") {
				fmt.Println(f.Stats())
			}

			if i != (len(list) - 1) {
				fmt.Println("")
			}
		}
	}

}
