package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"golang.org/x/net/html"
)

func TestPrintHTML(t *testing.T) {
	doc := fetchMocked()
	forEachNode(doc, startElementMocked, endElementMocked)
	actual := buf.Bytes()

	expected, err := ioutil.ReadFile("expected.html")
	if err != nil {
		t.Error(err)
	}
	if string(actual) != string(expected) {
		t.Errorf("\n**********Actual**********\n%s\n**********Expected**********\n%s\n", string(actual), string(expected))
	}
}

func fetchMocked() *html.Node {
	b, err := ioutil.ReadFile("test.html")
	if err != nil {
		panic(err)
	}
	doc, err := html.Parse(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	return doc
}

var b []byte
var buf *bytes.Buffer = bytes.NewBuffer(b)

func startElementMocked(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild == nil {
			noChild = true
			fPrintNodeWithAttribute(buf, n)
		} else {
			fPrintNodeWithAttribute(buf, n)
			depth++
		}
	}
	if n.Type == html.TextNode || n.Type == html.DocumentNode || n.Type == html.CommentNode {
		fPrintWithIndent(buf, n)
	}
}

func endElementMocked(n *html.Node) {
	if n.Type == html.ElementNode && !noChild {
		depth--
		fmt.Fprintf(buf, "%*s</%s>\n", depth*2, "", n.Data)
	}
	noChild = false
}
