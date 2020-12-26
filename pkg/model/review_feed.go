package model

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/itsubaki/appstore/pkg/format"
)

type ReviewFeed struct {
	App
	ReviewList
}

type ReviewList []Review

func (f *ReviewFeed) JSON(pretty bool) string {
	return format.JSON(f, pretty)
}

func (f *ReviewFeed) Stats() string {
	str := "stats: "
	for i := 5; i > 0; i-- {
		str = str + f.ratio(i)
		if i != 1 {
			str = str + ", "
		}
	}
	return str
}

func (f *ReviewFeed) ratio(rating int) string {
	rat, count, total := f.Ratio(rating)
	r := strconv.Itoa(rating)
	c := strconv.Itoa(count)
	t := strconv.Itoa(total)
	s := fmt.Sprintf("%2.0f", rat)
	return "[" + r + "]" + s + "%(" + c + "/" + t + ")"
}

func (f *ReviewFeed) Ratio(rating int) (ratio float64, count, total int) {
	r := len(f.Select(rating))
	l := len(f.ReviewList)
	return (float64(r) / float64(l)) * 100, r, l
}

func (f *ReviewFeed) Select(rating int) ReviewList {
	list := ReviewList{}
	for _, r := range f.ReviewList {
		if r.Rating == rating {
			list = append(list, r)
		}
	}
	return list
}

func NewReviewFeed(app App, b []byte) *ReviewFeed {
	var content interface{}
	err := json.Unmarshal(b, &content)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	feed := content.(map[string]interface{})["feed"]
	entry := feed.(map[string]interface{})["entry"]
	entrylist := entry.([]interface{})

	list := ReviewList{}
	for i := 1; i < len(entrylist); i++ {
		list = append(list, NewReview(entrylist[i]))
	}

	return &ReviewFeed{app, list}
}
