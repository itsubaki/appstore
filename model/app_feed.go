package model

import (
	"encoding/json"
	"fmt"

	"github.com/itsubaki/appstore/format"
)

type AppFeed struct {
	AppList
}

type AppList []App

func (l AppList) JSON(pretty bool) string {
	return format.JSON(l, pretty)
}

func (l AppList) Select(keyword string) AppList {
	list := AppList{}
	for i := 0; i < len(l); i++ {
		if l[i].Contains(keyword) {
			list = append(list, l[i])
		}
	}
	return list
}

func (f *AppFeed) Select(keyword string) AppList {
	list := AppList{}
	for i := 0; i < len(f.AppList); i++ {
		if f.AppList[i].Contains(keyword) {
			list = append(list, f.AppList[i])
		}
	}
	return list
}

func NewAppFeed(b []byte) *AppFeed {
	var content interface{}
	err := json.Unmarshal(b, &content)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	feed := content.(map[string]interface{})["feed"]
	result := feed.(map[string]interface{})["results"]
	rlist := result.([]interface{})

	list := AppList{}
	for i := 0; i < len(rlist); i++ {
		app := NewApp(rlist[i], i+1)
		list = append(list, app)
	}

	return &AppFeed{list}
}
