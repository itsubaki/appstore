package review

import (
	"fmt"

	"github.com/itsubaki/apst/client"
	cli "gopkg.in/urfave/cli.v1"
)

func Action(c *cli.Context) {
	id := c.Args().Get(0)

	b := client.Review(id)
	f := NewFeed(b)

	for _, r := range f.Review {
		fmt.Println(r)
	}
}
