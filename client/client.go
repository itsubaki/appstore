package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Get(limit int) []byte {
	var slimit = strconv.Itoa(limit)
	var url = "https://itunes.apple.com/jp/rss/topgrossingapplications/limit=" + slimit + "/json"
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
