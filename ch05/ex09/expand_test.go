package main

import (
	"strings"
	"testing"
)

type testcase struct {
	s        string
	f        func(string) string
	expected string
}

var (
	toUpper = func(s string) string {
		return strings.ToUpper(s)
	}

	toA = func(s string) string {
		return "A"
	}

	addQuestion = func(s string) string {
		return s + "?"
	}
)

func TestExpand(t *testing.T) {
	s := "abc $def あいう $えお $$aaa aa$aa"
	test := []testcase{
		{s, toUpper, "abc DEF あいう えお $AAA aa$aa"},
		{s, toA, "abc A あいう A A aa$aa"},
		{s, addQuestion, "abc def? あいう えお? $aaa? aa$aa"},
	}
	for _, v := range test {
		if actual := expand(v.s, v.f); actual != v.expected {
			t.Errorf("Input\t%v\n Expected\t%v\n Actual\t%v\n", v.s, v.expected, actual)
		}
	}
}
