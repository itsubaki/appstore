package review

import (
	"fmt"
	"os"
	"strconv"

	"github.com/itsubaki/apst/client"
	"github.com/itsubaki/apst/ranking"
	cli "gopkg.in/urfave/cli.v1"
)

func Action(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Println("See: apst review -h")
		os.Exit(1)
	}

	kw := c.Args().Get(0)
	country := c.String("country")
	rating, _ := strconv.Atoi(c.String("rating"))

	b := client.Ranking(ranking.Limit(c.String("limit")), c.String("feed"), country)
	f := ranking.NewFeed(b)

	for _, app := range f.Applist {
		if !app.Contains(kw) {
			continue
		}

		if rating == -1 {
			PrintAll(app, country)
			continue
		}

		Print(app, country, rating)
	}

}

func PrintAll(app *ranking.App, country string) {
	fmt.Println(app)

	b := client.Review(app.ID, country)
	f := NewFeed(b)

	for _, r := range f.Review {
		colorPrint(r.Rating, r.String())
	}
}

func Print(app *ranking.App, country string, rating int) {
	fmt.Println(app)

	b := client.Review(app.ID, country)
	f := NewFeed(b)

	hit := 0
	for _, r := range f.Review {
		if r.Rating == rating {
			hit++
			colorPrint(r.Rating, r.String())
		}
	}

	length := len(f.Review)
	shit := strconv.Itoa(hit)
	total := strconv.Itoa(length)

	rate := 0.0
	if hit > 0 {
		rate = (float64(hit) / float64(length)) * 100
	}

	fmt.Println(fmt.Sprintf("%3.0f", rate) + "% (" + shit + "/" + total + ")")
}

func colorPrint(rating int, review string) {
	color := "\x1b[30m%s\x1b[0m"
	switch rating {
	case 5:
		color = "\x1b[32m%s\x1b[0m"
	case 4:
		color = "\x1b[34m%s\x1b[0m"
	case 2:
		color = "\x1b[33m%s\x1b[0m"
	case 1:
		color = "\x1b[31m%s\x1b[0m"
	}

	fmt.Printf(color, review)
	fmt.Println("")
}
