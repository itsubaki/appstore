package review

import (
	"fmt"
	"os"

	"github.com/itsubaki/appstore/client"
	"github.com/itsubaki/appstore/format"
	"github.com/itsubaki/appstore/model"
	"github.com/urfave/cli"
)

func Action(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Println("See: appstore review -h")
		os.Exit(1)
	}

	country := c.String("country")
	b := client.Ranking(
		c.String("limit"),
		c.String("genre"),
		c.String("feed"),
		country,
	)

	list := model.NewAppFeed(b).AppList
	for i := 0; i < len(c.Args()); i++ {
		list = list.Select(c.Args().Get(i))
	}

	for i, app := range list {
		b := client.Review(app.ID, country)
		f := model.NewReviewFeed(app, b)

		switch c.String("format") {
		case "json":
			fmt.Println(f.JSON(c.Bool("pretty")))
		default:
			fmt.Println(app)
			for _, r := range f.ReviewList {
				format.ColorPrintln(r.Rating, r.String())
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
