package test

import (
	"github.com/itsubaki/apst/client"
	"github.com/itsubaki/apst/top"
	"testing"
)

func TestFeed(t *testing.T) {
	b := client.Ranking(10, "grossing")
	if b == nil {
		t.Error("http get failed.")
	}

	f := top.NewFeed(b)
	if f == nil {
		t.Error("feed unmarshal failed.")
	}

}
