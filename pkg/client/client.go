package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Review output customer reviews
func Review(id, country string) []byte {
	var url = "https://itunes.apple.com/" + country + "/rss/customerreviews/id=" + id + "/sortBy=mostRecent/json"
	return get(url)
}

// Ranking output ranking
func Ranking(limit, genre, feed, country string) []byte {
	var url = "https://rss.itunes.apple.com/api/v1/" + country + "/ios-apps/" + feed + "/" + genre + "/" + limit + "/explicit.json"
	return get(url)
}

func get(url string) []byte {
	rsp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rsp.Body.Close()

	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return b
}
