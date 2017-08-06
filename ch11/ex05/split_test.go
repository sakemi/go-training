package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		s    string
		sep  string
		want int
	}{
		{"", "", 0},
		{"", ":", 1},
		{"abc", "", 3},
		{"abc", ":", 1},
		{"a:b:cd", ":", 3},
		{"::", ":", 3},
		{"日 本 語", " ", 3},
		{"a.b..c...d....e", "..", 5},
	}

	for _, test := range tests {
		if got := len(strings.Split(test.s, test.sep)); got != test.want {
			t.Errorf("Split(%q, %q) returned %d word, want %d", test.s, test.sep, got, test.want)
		}
	}
}
