package model

import (
	"strconv"
	"strings"
)

type Review struct {
	Rating  int
	ID      string
	Title   string
	Content string
	Author  string
}

func (r Review) String() string {
	return "[" + strconv.Itoa(r.Rating) + "][" + r.Title + "] " + r.Content + " / " + r.Author
}

func NewReview(content interface{}) Review {
	return Review{
		Rating:  rating(content),
		ID:      rid(content),
		Content: contents(content),
		Title:   title(content),
		Author:  author(content),
	}
}

func rid(content interface{}) string {
	id := content.(map[string]interface{})["id"]
	label := id.(map[string]interface{})["label"]
	return label.(string)
}

func rating(content interface{}) int {
	rating := content.(map[string]interface{})["im:rating"]
	ratinglabel := rating.(map[string]interface{})["label"]
	r, _ := strconv.Atoi(ratinglabel.(string))
	return r
}

func contents(content interface{}) string {
	c := content.(map[string]interface{})["content"]
	clabel := c.(map[string]interface{})["label"]
	return strings.Replace(clabel.(string), "\n", " ", -1)
}

func title(content interface{}) string {
	t := content.(map[string]interface{})["title"]
	tlabel := t.(map[string]interface{})["label"]
	return strings.Replace(tlabel.(string), "\n", " ", -1)
}

func author(content interface{}) string {
	a := content.(map[string]interface{})["author"]
	n := a.(map[string]interface{})["name"]
	nlabel := n.(map[string]interface{})["label"]
	return strings.Replace(nlabel.(string), "\n", " ", -1)
}
