package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"golang.org/x/net/html"
)

func TestElementsByTagName(t *testing.T) {
	testCase := [][]string{
		{"head"},
		{"a"},
		{"script", "title"},
	}

	b, err := ioutil.ReadFile("test.html")
	if err != nil {
		panic(err)
	}
	doc, err := html.Parse(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	test := func(name []string) {
		elem := ElementsByTagName(doc, name...)
		checked := 0
		for _, e := range elem {
			ok := false
			for _, n := range name {
				if n == e.Data {
					ok = true
					checked++
					break
				}
			}
			if !ok {
				t.Errorf("ElementsByTagName(doc, %v)=%v", name, elem)
			}
		}
		if checked != len(elem) {
			t.Errorf("ElementsByTagName(doc, %v)=%v", name, elem)
		}
	}

	for _, tc := range testCase {
		test(tc)
	}
}
