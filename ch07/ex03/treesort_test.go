package main

import (
	"bytes"
	"fmt"
	"testing"
)

type test struct {
	vals []int
	str  string
}

func TestString(t *testing.T) {
	testCase := []test{
		{[]int{1}, "1"},
		{[]int{3, 2, 1}, "1 2 3"},
		{[]int{0, -1, 10, -10, 1}, "-10 -1 0 1 10"},
	}
	for _, tc := range testCase {
		buf := new(bytes.Buffer)
		tr := &tree{tc.vals[0], nil, nil}
		for _, v := range tc.vals[1:] {
			tr = add(tr, v)
		}
		fmt.Fprintf(buf, "%s", tr)
		if actual := string(buf.Bytes()); actual != tc.str {
			t.Errorf("Expected:%s, Actual:%s", tc.str, actual)
		}
	}
}
