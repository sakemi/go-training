package main

import "testing"

func TestSameStr(t *testing.T) {
	s1 := []string{"a", "xyz"}
	s2 := []string{"a", "xyz"}
	expected := []int{0, 0}
	for i, v := range s1 {
		if actual := countHashDiff(v, s2[i]); actual != expected[i] {
			t.Errorf("expected:%v actual:%v\n", expected[i], actual)
		}
	}
}

func TestDifferentStr(t *testing.T) {
	s1 := "x"
	s2 := "X"
	if actual := countHashDiff(s1, s2); actual == 0 {
		t.Errorf("%v,%v %v\n", s1, s2, actual)
	}
}
