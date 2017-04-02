package ranking

import (
	"fmt"

	"github.com/itsubaki/apst/client"
	"github.com/itsubaki/apst/util"
	cli "gopkg.in/urfave/cli.v1"
)

func Action(c *cli.Context) {
	b := client.Ranking(
		util.Limit(c.String("limit")),
		c.String("genre"),
		c.String("feed"),
		c.String("country"),
	)

	k := util.Keyword(c.Args())
	f := NewFeed(b)
	for _, app := range f.Applist {
		if app.Contains(k) {
			fmt.Println(app)
		}
	}
}
