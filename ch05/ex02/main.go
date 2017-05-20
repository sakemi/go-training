package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

//!+
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	m := map[string]int{}
	count(m, doc)
	fmt.Println(m)
	for k, v := range m {
		fmt.Printf("%s : %d\n", k, v)
	}
}

func count(m map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		m[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count(m, c)
	}

	return m
}
