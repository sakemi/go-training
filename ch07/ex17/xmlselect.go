package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string
	var attrStack [][]xml.Attr
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)
			attrStack = append(attrStack, tok.Attr)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
			attrStack = attrStack[:len(attrStack)-1]
		case xml.CharData:
			if elemAndAttr, isContain := containsAll(attrStack, stack, os.Args[1:]); isContain {
				fmt.Printf("%s: %s\n", strings.Join(elemAndAttr, " "), tok)
			}
		}
	}
}

func containsAll(attr [][]xml.Attr, x, y []string) ([]string, bool) {
	var elemAndAttr []string
	for i, v := range x {
		elemAndAttr = append(elemAndAttr, v)
		for _, v := range attr[i] {
			elemAndAttr = append(elemAndAttr, v.Name.Local, v.Value)
		}
	}
	tmp := elemAndAttr
	for len(y) <= len(tmp) {
		if len(y) == 0 {
			return elemAndAttr, true
		}
		if tmp[0] == y[0] {
			y = y[1:]
		}
		tmp = tmp[1:]
	}
	return nil, false
}
