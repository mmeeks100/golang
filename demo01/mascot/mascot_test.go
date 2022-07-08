package mascot_test

import (
	"demo01/mascot"
	"testing"
)

func TestBestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong Mascot")
	}
}
