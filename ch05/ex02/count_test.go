package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"golang.org/x/net/html"
)

func TestCount(t *testing.T) {
	b, err := ioutil.ReadFile("test.html")
	if err != nil {
		panic(err)
	}
	doc, err := html.Parse(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	expected := map[string]int{"html": 1, "head": 1, "script": 1, "meta": 3, "title": 1, "body": 1, "a": 2, "link": 1}
	m := map[string]int{}
	count(m, doc)
	if !equals(m, expected) {
		t.Errorf("Expected:\n%v\nActual:\n%v\n", expected, m)
	}
}

func equals(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}

	return true
}
