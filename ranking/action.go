package ranking

import (
	"fmt"
	"time"

	"github.com/itsubaki/apst/client"
	"github.com/itsubaki/apst/genre"
	"github.com/itsubaki/apst/util"
	"gopkg.in/urfave/cli.v1"
)

func Action(c *cli.Context) {
	fmt.Println(time.Now())
	b := client.Ranking(
		util.Limit(c.String("limit")),
		genre.ID(c.String("genre")),
		c.String("feed"),
		c.String("country"),
	)

	list := NewAppFeed(b).AppList
	for i := 0; i < len(c.Args()); i++ {
		list = list.Select(c.Args().Get(i))
	}

	switch c.String("output") {
	case "json":
		fmt.Println(list.Json())
	case "jsonp":
		fmt.Println(list.JsonPretty())
	default:
		for _, app := range list {
			fmt.Println(app)
		}
	}

}
