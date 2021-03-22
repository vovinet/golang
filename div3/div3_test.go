package main

import "testing"

func TestDiv3(t *testing.T) {
	var res []int
	var check = [3]int{3, 6, 9}
	start := 1
	stop := 11
	res = div3(start, stop)
	if len(res) != len(check) {
		t.Error("resulting array length mismath, ", len(res), " != ", len(check))
	} else {
		for i := 0; i < len(res); i++ {
			if res[i] != check[i] {
				t.Error("resulting array element mismath, ", res[i], " != ", check[i])
			}
		}
	}
}
