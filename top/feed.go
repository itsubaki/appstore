package top

import (
	"encoding/json"
	"fmt"
)

type Feed struct {
	Feed  interface{}
	Entry interface{}
}

func NewFeed(b []byte) *Feed {
	var content interface{}
	err := json.Unmarshal(b, &content)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	feed := &Feed{}
	feed.Feed = content.(map[string]interface{})["feed"]
	feed.Entry = feed.Feed.(map[string]interface{})["entry"]
	return feed
}

func (feed *Feed) AppList(keyword string) [](*App) {
	applist := [](*App){}
	//	entry := feed.Entry.([]interface{})

	/*
		for i := 0; i < len(entry.([]interface{})); i++ {
			app := NewApp(entry.([]interface{})[i], i)

			if keyword == "" {
				applist = append(applist, app)
				continue
			}

			if app.Contains(keyword) {
				applist = append(applist, app)
			}
		}*/

	return applist
}
