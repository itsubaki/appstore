package top

import (
	"fmt"
	"github.com/itsubaki/apst/client"
	cli "gopkg.in/urfave/cli.v1"
	"strconv"
)

func Action(c *cli.Context) {
	limit, _ := strconv.Atoi(c.String("limit"))
	b := client.Get(limit)
	feed := NewFeed(b)

	keyword := ""
	if len(c.Args()) > 0 {
		keyword = c.Args().Get(0)
	}

	applist := feed.AppList(keyword)

	for _, app := range applist {
		fmt.Println(app.String())
	}
}
