package review

import (
	"encoding/json"
	"fmt"
)

type Feed struct {
	Review [](*Review)
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

	rlist := [](*Review){}
	for i := 1; i < len(entrylist); i++ {
		rlist = append(rlist, NewReview(entrylist[i]))
	}

	return &Feed{rlist}
}
