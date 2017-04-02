package ranking

import (
	"encoding/json"
	"fmt"
)

type Feed struct {
	list AppList
}

type AppList [](*App)

func (f *Feed) Select(keyword string) AppList {
	list := AppList{}
	for i := 0; i < len(f.list); i++ {
		if f.list[i].Contains(keyword) {
			list = append(list, f.list[i])
		}
	}
	return list
}

func (f *Feed) App() AppList {
	return f.list
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

	list := AppList{}
	for i := 0; i < len(entrylist); i++ {
		app := NewApp(entrylist[i], i+1)
		list = append(list, app)
	}

	return &Feed{list}
}
