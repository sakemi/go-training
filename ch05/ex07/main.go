package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		printHTML(url)
	}
}

func printHTML(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int
var noChild bool

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild == nil {
			noChild = true
			fPrintNodeWithAttribute(os.Stdout, n)
		} else {
			fPrintNodeWithAttribute(os.Stdout, n)
			depth++
		}
	}
	if n.Type == html.TextNode || n.Type == html.DocumentNode || n.Type == html.CommentNode {
		fPrintWithIndent(os.Stdout, n)
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode && !noChild {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
	noChild = false
}
func fPrintNodeWithAttribute(w io.Writer, n *html.Node) {
	fmt.Fprintf(w, "%*s<%s", depth*2, "", n.Data)

	if n.Attr != nil {
		for _, attr := range n.Attr {
			fmt.Fprintf(w, " %s=\"%s\"", attr.Key, attr.Val)
		}
	}

	if noChild {
		fmt.Fprintln(w, "/>")
	} else {
		fmt.Fprintln(w, ">")
	}
}
func fPrintWithIndent(w io.Writer, n *html.Node) {
	fmt.Fprintf(w, "%*s%s\n", depth*2, "", n.Data)
}
