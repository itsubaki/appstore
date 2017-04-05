package util

import (
	"fmt"
	"strconv"
)

func Genre(genre string) string {
	switch genre {
	case "business":
		return "6000"
	case "weather":
		return "6001"
	case "utilities":
		return "6002"
	case "travel":
		return "6003"
	case "sports":
		return "6004"
	case "socialnetworking":
		return "6005"
	case "reference":
		return "6006"
	case "productivity":
		return "6007"
	case "photo&video":
		return "6008"
	case "news":
		return "6009"
	case "navigation":
		return "6010"
	case "music":
		return "6011"
	case "lifestyle":
		return "6012"
	case "health&fitness":
		return "6013"
	case "games":
		return "6014"
	case "finance":
		return "6015"
	case "entertainment":
		return "6016"
	case "education":
		return "6017"
	case "books":
		return "6018"
	case "medical":
		return "6020"
	case "newsstand":
		return "6021"
	case "catalogs":
		return "6022"
	case "food&drink":
		return "6023"
	default:
		return ""
	}

}

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

