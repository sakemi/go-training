package main

import "testing"

func TestReverse(t *testing.T) {
	s := [][5]int{{1, 2, 3, 4, 5}, {1, -1, 2, 2, -3}}
	expected := [][5]int{{5, 4, 3, 2, 1}, {-3, 2, 2, -1, 1}}
	for i, v := range s {
		reverse(&v)
		if v != expected[i] {
			t.Errorf("expected:%v, actual:%v\n", expected[i], v)
		}
	}
}
