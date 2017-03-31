package review

type Review struct {
	Rating  int
	Title   string
	Content string
	Author  string
}

func NewReview(content interface{}) *Review {
	r := &Review{}

	return r
}
