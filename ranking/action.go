package ranking

import (
	"fmt"

	"github.com/itsubaki/apst/client"
	"github.com/itsubaki/apst/util"
	"gopkg.in/urfave/cli.v1"
)

func Action(c *cli.Context) {
	b := client.Ranking(
		util.Limit(c.String("limit")),
		util.Genre(c.String("genre")),
		c.String("feed"),
		c.String("country"),
	)

	list := NewFeed(b).Select("")
	for i := 0; i < len(c.Args()); i++ {
		list = list.Select(c.Args().Get(i))
	}

	for _, app := range list {
		fmt.Println(app)
	}
}
