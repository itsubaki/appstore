package ranking

import (
	"fmt"
	"github.com/itsubaki/apst/client"
	cli "gopkg.in/urfave/cli.v1"
	"strconv"
)

func limit(input string) int {
	limit := 30
	if tmp, _ := strconv.Atoi(input); tmp < 200 {
		limit = tmp
	}
	if limit < 10 {
		limit = 10
	}
	return limit
}

func keyword(args []string) string {
	keyword := ""
	if len(args) > 0 {
		keyword = args[0]
	}

	return keyword
}

func Action(c *cli.Context) {
	b := client.Ranking(limit(c.String("limit")), "grossing")
	f := NewFeed(b)
	kw := keyword(c.Args())

	for _, app := range f.Applist {
		if app.Contains(kw) {
			fmt.Println(app)
		}
	}
}
