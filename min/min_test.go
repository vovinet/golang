package main

import "testing"

func TestMin(t *testing.T) {
	x := []int{76, 92, 7, 34, 2, 56}
	res := min(x)
	if res != 2 {
		t.Error("expected 2, got ", res)
	}
}
