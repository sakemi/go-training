package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		if n.Data == "a" || n.Data == "link" {
			for _, e := range n.Attr {
				if e.Key == "href" {
					links = append(links, e.Val)
				}
			}
		} else if n.Data == "img" || n.Data == "script" {
			for _, e := range n.Attr {
				if e.Key == "src" {
					links = append(links, e.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
