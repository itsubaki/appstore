package test

import (
	"github.com/itsubaki/apst/client"
	"testing"
)

func TestClientGet(t *testing.T) {
	b := client.Get(10)
	if b == nil {
		t.Error("http get failed.")
	}
}
