package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"golang.org/x/net/html"
)

func main() {
	b, err := ioutil.ReadFile("test.html")
	if err != nil {
		panic(err)
	}
	doc, err := html.Parse(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	elem := ElementsByTagName(doc, "a", "title")

	for _, v := range elem {
		fmt.Println(v.Data)
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node

	if doc.Type == html.ElementNode {
		for _, n := range name {
			if doc.Data == n {
				nodes = append(nodes, doc)
			}
		}
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, ElementsByTagName(c, name...)...)
	}
	return nodes
}
