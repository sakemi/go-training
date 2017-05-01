package main

import "testing"

func TestRotate(t *testing.T) {
	s := [][]int{{1, 2, 3, 4, 5}, {1, 2, 3}, {1, 2, 3, 4, 5}}
	n := []int{3, 4, 0}
	expected := [][]int{{4, 5, 1, 2, 3}, {2, 3, 1}, {1, 2, 3, 4, 5}}
	var actual []int
	for i := range s {
		actual = rotate(s[i], n[i])
		if !equals(actual, expected[i]) {
			t.Errorf("input: %v, rotate:%d, expected:%v, actual:%v\n", s[i], n[i], expected[i], actual)
		}
	}
}

func equals(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
