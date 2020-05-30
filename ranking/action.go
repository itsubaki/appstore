package ranking

import (
	"fmt"

	"github.com/itsubaki/appstore/client"
	"github.com/itsubaki/appstore/model"
	"github.com/urfave/cli"
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
		fmt.Println(list.JSON(c.Bool("pretty")))
	default:
		for _, app := range list {
			fmt.Println(app)
		}
	}

}
