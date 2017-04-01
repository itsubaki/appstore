package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Review(id string) []byte {
	var url = "https://itunes.apple.com/jp/rss/customerreviews/id=" + id + "/sortBy=mostRecent/json"
	return get(url)
}

func Ranking(limit int, feed string) []byte {
	var slimit = strconv.Itoa(limit)
	var url = "https://itunes.apple.com/jp/rss/top" + feed + "applications/limit=" + slimit + "/json"
	return get(url)
}

func get(url string) []byte {
	rsp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	defer rsp.Body.Close()

	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return b
}
