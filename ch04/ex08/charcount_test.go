package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestCountCharType(t *testing.T) {
	s := "1\na b2c 34"
	expected := map[string]int{"letter": 3, "digit": 4, "space": 3}
	r := strings.NewReader(s)
	actual, _ := countCharType(bufio.NewReader(r))
	if len(expected) != len(actual) {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}
	for k, v := range expected {
		if actual[k] != v {
			t.Errorf("expected:%v, actual:%v", expected, actual)
		}
	}
}
