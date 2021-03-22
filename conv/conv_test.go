package main

import "testing"

func TestConv(t *testing.T) {
	var v float64
	v = conv(4)
	if v != 13.123359580052492 {
		t.Error("Expected 13.123359580052492, got ", v)
	}
}
