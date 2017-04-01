package test

import (
	"github.com/itsubaki/apst/client"
	"github.com/itsubaki/apst/review"
	"testing"
)

func TestClientRanking(t *testing.T) {
	b := client.Ranking(10, "grossing")
	if b == nil {
		t.Error("http get failed.")
	}
}

func TestClientReview(t *testing.T) {
	b := client.Review("658511662")
	if b == nil {
		t.Error("http get failed.")
	}

	f := review.NewFeed(b)
	if f == nil {
		t.Error("feed error.")
	}
}
