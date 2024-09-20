package intf_test

import (
	"testing"

	"main.go/intf"
)

func TestQuack(t *testing.T) {
	duck := intf.NewDuck()
	if x := duck.Quack(); len(x) < 1 {
		t.Errorf("duck doesn't quack")
	}
}
