package main

import "testing"

func TestPopcount(t *testing.T) {
	x := []uint8{0, 255, 15, 8}
	expected := []int{0, 8, 4, 1}
	for i, v := range x {
		if actual := PopCount(v); actual != expected[i] {
			t.Errorf("expected:%v actual:%v\n", expected[i], actual)
		}
	}
}
