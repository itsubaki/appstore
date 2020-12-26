package model_test

import (
	"testing"

	"github.com/itsubaki/appstore/pkg/client"
	"github.com/itsubaki/appstore/pkg/model"
)

func TestFeed(t *testing.T) {
	b := client.Ranking("10", "all", "top-grossing", "jp")
	if b == nil {
		t.Error("http get failed.")
	}

	f := model.NewAppFeed(b)
	if f == nil {
		t.Error("feed unmarshal failed.")
	}
}
