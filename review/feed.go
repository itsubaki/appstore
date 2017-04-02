package review

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Feed struct {
	list ReviewList
}

type ReviewList [](*Review)

func (f *Feed) Review() ReviewList {
	return f.list
}

func (f *Feed) Stats() string {
	str := "stats: "
	for i := 5; i > 0; i-- {
		str = str + f.ratio(i)
		if i != 1 {
			str = str + ", "
		}
	}
	return str
}

func (f *Feed) ratio(rating int) string {
	rat, count, total := f.Ratio(rating)
	r := strconv.Itoa(rating)
	c := strconv.Itoa(count)
	t := strconv.Itoa(total)
	s := fmt.Sprintf("%2.0f", rat)
	return "[" + r + "]" + s + "%(" + c + "/" + t + ")"
}

func (f *Feed) Ratio(rating int) (ratio float64, count, total int) {
	l := len(f.Review())
	r := len(f.Select(rating))
	return (float64(r) / float64(l)) * 100, r, l
}

func (f *Feed) Select(rating int) ReviewList {
	list := ReviewList{}
	for _, r := range f.Review() {
		if r.Rating == rating {
			list = append(list, r)
		}
	}
	return list
}

func NewFeed(b []byte) *Feed {
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

	return &Feed{list: list}
}
