package main

import (
	"bytes"
	"testing"
)

type test struct {
	str      string
	expected bool
}

func TestIsPalindrome(t *testing.T) {
	testCase := []test{
		test{"a", true},
		test{"aba", true},
		test{"a b c b a", true},
		test{"トマト", true},
		test{"abb", false},
		test{" aba", false},
	}
	for _, tc := range testCase {
		p := bytes.Runes([]byte(tc.str))
		if IsPalindrome(PalindromeCandidate(p)) != tc.expected {
			t.Errorf("%s", tc.str)
		}
	}
}
