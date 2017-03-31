package ranking

import (
	"encoding/json"
	"fmt"
)

type Feed struct {
	Applist [](*App)
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

	applist := [](*App){}
	for i := 0; i < len(entrylist); i++ {
		app := NewApp(entrylist[i], i+1)
		applist = append(applist, app)
	}

	return &Feed{applist}
}
