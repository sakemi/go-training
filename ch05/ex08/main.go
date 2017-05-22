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

	printNodeWithAttribute(ElementByID(doc, "i1"))
	printNodeWithAttribute(ElementByID(doc, "i2"))
}

func ElementByID(doc *html.Node, id string) *html.Node {
	startElement := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}
		if n.Attr == nil {
			return true
		}

		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return false
			}
		}

		return true
	}
	return forEachNode(doc, startElement, nil)
}

var cont bool = true
var result *html.Node

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		cont = pre(n)
	}
	if !cont {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result = forEachNode(c, pre, post)
		if result != nil {
			return result
		}
	}

	if post != nil {
		cont = post(n)
	}
	return nil
}

func printNodeWithAttribute(n *html.Node) {
	if n == nil {
		fmt.Println("nil")
		return
	}
	fmt.Printf("<%s", n.Data)

	if n.Attr != nil {
		for _, attr := range n.Attr {
			fmt.Printf(" %s=\"%s\"", attr.Key, attr.Val)
		}
	}

	fmt.Println(">")
}
