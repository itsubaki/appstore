package top

import (
	"fmt"
	"github.com/itsubaki/apst/client"
	cli "gopkg.in/urfave/cli.v1"
	"strconv"
)

func Action(c *cli.Context) {
	limit := 200
	climit, _ := strconv.Atoi(c.String("limit"))
	if climit < 200 {
		limit = climit
	}
	if limit < 10 {
		limit = 10
	}

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
