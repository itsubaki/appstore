package review

type Feed struct {
	Review [](*Review)
}

func NewFeed(content interface{}) *Feed {
	f := &Feed{}

	return f
}
