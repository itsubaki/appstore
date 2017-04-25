package genre

import (
	"fmt"

	"github.com/itsubaki/apst/model"

	"gopkg.in/urfave/cli.v1"
)

func Action(c *cli.Context) {
	for k := range model.Genre() {
		fmt.Println(k)
	}
}
