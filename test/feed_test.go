package test

import (
	"github.com/itsubaki/apst/client"
	"github.com/itsubaki/apst/top"
	"testing"
)

func TestGetAppList(t *testing.T) {
	b := client.Get(1)
	if b == nil {
		t.Error("http get failed.")
	}

	feed := top.NewFeed(b)
	if b == nil {
		t.Error("feed unmarshal failed.")
	}

	applist := feed.GetAppList("")
	t.Log(applist)

}

func TestFeed(t *testing.T) {
	b := client.Get(1)
	if b == nil {
		t.Error("http get failed.")
	}

	feed := top.NewFeed(b)
	if b == nil {
		t.Error("feed unmarshal failed.")
	}
	t.Log(feed)

}
