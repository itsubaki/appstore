package ranking

import (
	"fmt"
	"strconv"

	"github.com/itsubaki/apst/client"
	cli "gopkg.in/urfave/cli.v1"
)

func limit(input string) int {
	limit := 30
	if tmp, _ := strconv.Atoi(input); tmp < 201 {
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
	b := client.Ranking(limit(c.String("limit")), c.String("feed"))
	f := NewFeed(b)
	k := keyword(c.Args())

	for _, app := range f.Applist {
		if app.Contains(k) {
			fmt.Println(app)
		}
	}
}
