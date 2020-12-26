package client_test

import (
	"testing"

	"github.com/itsubaki/appstore/pkg/client"
)

func TestClientRanking(t *testing.T) {
	b := client.Ranking("10", "all", "top-grossing", "jp")
	if b == nil {
		t.Error("http get failed.")
	}
}

func TestClientReview(t *testing.T) {
	b := client.Review("1094591345", "jp")
	if b == nil {
		t.Error("http get failed.")
	}
}
