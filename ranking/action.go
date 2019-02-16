package ranking

import (
	"fmt"

	"github.com/itsubaki/appstore/client"
	"github.com/itsubaki/appstore/model"
	"gopkg.in/urfave/cli.v1"
)

func Action(c *cli.Context) {
	b := client.Ranking(
		c.String("limit"),
		c.String("genre"),
		c.String("feed"),
		c.String("country"),
	)

	list := model.NewAppFeed(b).AppList
	for i := 0; i < len(c.Args()); i++ {
		list = list.Select(c.Args().Get(i))
	}

	switch c.String("format") {
	case "json":
		fmt.Println(list.Json(c.Bool("pretty")))
	default:
		for _, app := range list {
			fmt.Println(app)
		}
	}

}
