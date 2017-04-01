package ranking

import (
	"fmt"
	"strconv"

	"github.com/itsubaki/apst/client"
	cli "gopkg.in/urfave/cli.v1"
)

func Limit(input string) int {
	limit, _ := strconv.Atoi(input)
	return limit
}

func Keyword(args []string) string {
	keyword := ""
	if len(args) > 0 {
		keyword = args[0]
	}

	return keyword
}

func Action(c *cli.Context) {
	b := client.Ranking(Limit(c.String("limit")), c.String("feed"), c.String("country"))
	f := NewFeed(b)
	k := Keyword(c.Args())

	for _, app := range f.Applist {

		if app.Contains(k) {
			fmt.Println(app)
		}

	}
}
