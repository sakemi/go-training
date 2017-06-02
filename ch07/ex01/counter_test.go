package main

import "testing"

type test struct {
	str string
	num int
}

func TestCountWord(t *testing.T) {
	testCase := []test{
		{"a", 1},
		{"hello world", 2},
		{"", 0},
		{"あいう　え　お", 3},
		{"a\nbcd  e", 3},
	}
	var wc WordCounter
	for _, tc := range testCase {
		wc = 0
		wc.Write([]byte(tc.str))
		if int(wc) != tc.num {
			t.Errorf("Input:%s, Expected:%d, Actual%d", tc.str, tc.num, int(wc))
		}
	}
}

func TestCountLine(t *testing.T) {
	testCase := []test{
		{"a", 1},
		{"hello world", 1},
		{"", 0},
		{"あ\nいう　\nえ　お", 3},
		{"a\nbcd \n\n e", 4},
	}
	var wc LineCounter
	for _, tc := range testCase {
		wc = 0
		wc.Write([]byte(tc.str))
		if int(wc) != tc.num {
			t.Errorf("Input:%s, Expected:%d, Actual%d", tc.str, tc.num, int(wc))
		}
	}
}
