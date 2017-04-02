package util

import "strconv"

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
