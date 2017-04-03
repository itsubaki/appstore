package util

import (
	"fmt"
	"strconv"
)

func ColorPrintln(rating int, message string) {
    ColorPrint(rating, message)
    fmt.Println("")
}

func ColorPrint(rating int, message string) {
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

	fmt.Printf(color, message)
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

func Keyword(args []string) string {
	keyword := ""
	if len(args) > 0 {
		keyword = args[0]
	}

	return keyword
}
