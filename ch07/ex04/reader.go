package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type StringReader struct {
	data []byte
}

func main() {
	h := "<html><head></head><body><a></a></body></html>"
	sr := NewReader(h)
	doc, err := html.Parse(sr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func NewReader(s string) *StringReader {
	sr := StringReader{[]byte(s)}
	return &sr
}

func (sr *StringReader) Read(p []byte) (n int, err error) {
	if len(sr.data) < len(p) {
		err = io.EOF
	} else {
		err = nil
	}
	n = copy(p, sr.data)
	return
}
