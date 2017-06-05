package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"testing"
)

func TestGetChildren(t *testing.T) {
	b, err := ioutil.ReadFile("test.html")
	if err != nil {
		t.Error(err)
	}
	dec := xml.NewDecoder(bytes.NewReader(b))
	tok, err := dec.Token()
	if err != nil {
		t.Error(err)
	}
	root := new(Element)
	switch tok := tok.(type) {
	case xml.StartElement:
		root.Attr = tok.Attr
		root.Type = tok.Name
		getChildren(root, dec)
	default:
		t.Fatal("root is not start element")
	}

	firstChild := root.Children[0]
	switch firstChild := firstChild.(type) {
	case *Element:
		if actual := firstChild.Type.Local; actual != "head" {
			t.Errorf("first child is %s", actual)
		}
		firstGrandChild := firstChild.Children[0]
		switch firstGrandChild := firstGrandChild.(type) {
		case *Element:
			if actual := firstGrandChild.Type.Local; actual != "script" {
				t.Errorf("first child is %s", actual)
			}
			if actual := firstGrandChild.Attr[0].Name.Local; actual != "src" {
				t.Errorf("first grandchild's attr name is %s", actual)
			}
			if actual := firstGrandChild.Attr[0].Value; actual != "source" {
				t.Errorf("first grandchild's attr value is %s", actual)
			}
		default:
			t.Fatal("first child is not start element")
		}
	default:
		t.Fatal("first child is not start element")
	}

	secondChild := root.Children[1]
	switch secondChild := secondChild.(type) {
	case *Element:
		if actual := secondChild.Type.Local; actual != "body" {
			t.Errorf("first child is %s", actual)
		}
	default:
		t.Fatal("first child is not start element")
	}
}
