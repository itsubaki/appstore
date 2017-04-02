package review

import (
	"strconv"
	"strings"
)

type Review struct {
	Rating  int
	Title   string
	Content string
	Author  string
}

func (r *Review) String() string {
	rating := strconv.Itoa(r.Rating)

	return "[" + rating + "][" + r.Title + "] " + r.Content + " /" + r.Author
}

func NewReview(content interface{}) *Review {

	rating := content.(map[string]interface{})["im:rating"]
	ratinglabel := rating.(map[string]interface{})["label"]

	c := content.(map[string]interface{})["content"]
	clabel := c.(map[string]interface{})["label"]

	t := content.(map[string]interface{})["title"]
	tlabel := t.(map[string]interface{})["label"]

	a := content.(map[string]interface{})["author"]
	n := a.(map[string]interface{})["name"]
	nlabel := n.(map[string]interface{})["label"]

	r := &Review{}
	r.Rating, _ = strconv.Atoi(ratinglabel.(string))
	r.Content = strings.Replace(clabel.(string), "\n", " ", -1)
	r.Title = strings.Replace(tlabel.(string), "\n", " ", -1)
	r.Author = strings.Replace(nlabel.(string), "\n", " ", -1)
	return r
}
