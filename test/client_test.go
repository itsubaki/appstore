package test

import (
	"testing"

	"github.com/itsubaki/apst/client"
	"github.com/itsubaki/apst/review"
)

func TestClientRanking(t *testing.T) {
	b := client.Ranking(10, "", "grossing", "jp")
	if b == nil {
		t.Error("http get failed.")
	}
	t.Log(string(b))
}

func TestClientReview(t *testing.T) {
	b := client.Review("658511662", "jp")
	if b == nil {
		t.Error("http get failed.")
	}
	t.Log(string(b))

	f := review.NewFeed(b)
	if f == nil {
		t.Error("feed error.")
	}
}
