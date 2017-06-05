package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	tok, err := dec.Token()
	if err == io.EOF {
		fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
		os.Exit(0)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
		os.Exit(1)
	}
	for {
		switch tok := tok.(type) {
		case xml.StartElement:
			root := new(Element)
			root.Attr = tok.Attr
			root.Type = tok.Name
			getChildren(root, dec)
			break
		default:
			tok, err = dec.Token()
			if err == io.EOF {
				fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
				os.Exit(0)
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
				os.Exit(1)
			}
		}
	}
}

func getChildren(n Node, dec *xml.Decoder) {
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(0)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			elem := new(Element)
			elem.Attr = tok.Attr
			elem.Type = tok.Name
			n, ok := n.(*Element)
			if ok {
				n.Children = append(n.Children, elem)
			}
			getChildren(elem, dec)
		case xml.CharData:
			char := CharData(tok)
			n, ok := n.(*Element)
			if ok {
				if char != "\n" {
					n.Children = append(n.Children, char)
				}
			}
		case xml.EndElement:
			return
		}
	}
}
