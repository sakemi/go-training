package main

import (
	"bytes"
	"fmt"
	"testing"
)

type test struct {
	str string
	num int64
}

func TestCount(t *testing.T) {
	testCase := []test{
		{"a", 1},
		{"hello world", 11},
		{"", 0},
		{"あいう", 9},
		{"a\nbcd  e", 8},
	}
	for _, tc := range testCase {
		buf := new(bytes.Buffer)
		w, c := CountingWriter(buf)
		fmt.Fprintf(w, "%s", tc.str)
		if *c != tc.num {
			t.Errorf("Input:%s, Expected:%d, Actual%d", tc.str, tc.num, *c)
		}
		if tc.str != string(buf.Bytes()) {
			t.Errorf("Expected:%s, Actual%s", tc.str, string(buf.Bytes()))
		}
	}
}
