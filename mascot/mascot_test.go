package mascot_test

import (
	"fluxyl-go-demo-1/mascot"
	"testing"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong mascot :(")
	}
}
