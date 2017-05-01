package main

import "testing"

func TestDup(t *testing.T) {
	s := [][]string{{"a", "b", "b", "bb", "abc", "a", "a"}, {"", "あああ", "あああ", "a", "", ""}}
	expected := [][]string{{"a", "b", "bb", "abc", "a"}, {"", "あああ", "a", ""}}
	var actual []string
	for i := range s {
		actual = dup(s[i])
		if !equals(actual, expected[i]) {
			t.Errorf("input: %v, expected:%v, actual:%v\n", s[i], expected[i], actual)
		}
	}
}

func equals(s1, s2 []string) bool {
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
