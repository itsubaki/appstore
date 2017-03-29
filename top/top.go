package top

import (
	"fmt"
	"github.com/itsubaki/apst/client"

	cli "gopkg.in/urfave/cli.v1"
)

func Action(c *cli.Context) {
	b := client.Get(100)
	feed := NewFeed(b)
	applist := feed.GetAppList("keyword")

	for i := 0; i < len(applist); i++ {
		fmt.Println(applist[i].ToString())
	}
}
