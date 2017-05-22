package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"golang.org/x/net/html"
)

func TestElementByID(t *testing.T) {
	b, err := ioutil.ReadFile("test.html")
	if err != nil {
		panic(err)
	}
	doc, err := html.Parse(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	expectedData := "div"
	expectedAttr := map[string]string{"class": "c1 c2", "id": "i1"}
	actual := ElementByID(doc, "i1")

	if actual.Data != expectedData {
		t.Errorf("Expected data:\n%v\nActual data:\n%v\n", expectedData, actual.Data)
	}
	for _, attr := range actual.Attr {
		if attr.Val != expectedAttr[attr.Key] {
			t.Errorf("Expected attr:\n%v\nActual attr:\n%v\n", expectedAttr, actual.Attr)
		}
	}
}
