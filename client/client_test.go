package client

import "testing"

func TestClientRanking(t *testing.T) {
	b := Ranking("10", "all", "top-grossing", "jp")
	if b == nil {
		t.Error("http get failed.")
	}
}

func TestClientReview(t *testing.T) {
	b := Review("1094591345", "jp")
	if b == nil {
		t.Error("http get failed.")
	}
}
