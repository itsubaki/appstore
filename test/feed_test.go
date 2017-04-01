package test

import (
	"testing"

	"github.com/itsubaki/apst/client"
	"github.com/itsubaki/apst/ranking"
)

func TestFeed(t *testing.T) {
	b := client.Ranking(10, "", "grossing", "jp")
	if b == nil {
		t.Error("http get failed.")
	}

	f := ranking.NewFeed(b)
	if f == nil {
		t.Error("feed unmarshal failed.")
	}

}
