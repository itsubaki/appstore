package util

import (
	"fmt"
	"strconv"

	"github.com/itsubaki/apst/genre"
)

func ColorPrintln(rating int, a ...interface{}) {
	ColorPrint(rating, a)
	fmt.Println("")
}

func ColorPrint(rating int, a ...interface{}) {
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

	fmt.Printf(color, a)
}

func Limit(input string) int {
	limit, _ := strconv.Atoi(input)
	if limit > 201 {
		limit = 200
	}

	if limit < 2 {
		limit = 2
	}
	return limit
}

func Genre(name string) string {
	v, ok := genre.Genre()[name]
	if ok {
		return v
	}
	return ""
}
