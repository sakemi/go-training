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

	expectedWords := 12
	expectedImages := 1
	actualWords, actualImages := countWordsAndImages(doc)
	if expectedWords != actualWords {
		t.Errorf("Expected words:\n%v\nActual words:\n%v\n", expectedWords, actualWords)
	}
	if expectedImages != actualImages {
		t.Errorf("Expected images:\n%v\nActual images:\n%v\n", expectedImages, actualImages)
	}
}
