package main

import "testing"

type test struct {
	a        []string
	sep      string
	expected string
}

func TestJoin(t *testing.T) {
	testCase := []test{
		test{[]string{"foo", "bar", "baz"}, "/", "foo/bar/baz"},
		test{[]string{"あ", "い", "う", "えお"}, " ", "あ い う えお"},
		test{[]string{}, "/", ""},
	}
	for _, tc := range testCase {
		if actual := join(tc.sep, tc.a...); actual != tc.expected {
			t.Errorf("join(%s, %v) = %s", tc.sep, tc.a, actual)
		}
	}
}
