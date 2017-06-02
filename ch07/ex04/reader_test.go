package main

import (
	"io"
	"testing"
)

func TestRead(t *testing.T) {
	str := "123456789"
	sr := StringReader{[]byte(str)}
	buf := 10
	p := make([]byte, buf)
	n, err := sr.Read(p)
	if err != io.EOF {
		t.Error(err)
	}
	if n != len(str) {
		t.Errorf("Read %d byte", n)
	}
	if str != string(p[:len(str)]) {
		t.Errorf("Expected:%s, Actual:%s", str, string(p[:len(str)]))
	}

	buf = 3
	p = make([]byte, buf)
	n, err = sr.Read(p)
	if err != nil {
		t.Error(err)
	}
	if n != buf {
		t.Errorf("Read %d byte", n)
	}
	if "123" != string(p) {
		t.Errorf("Expected:%s, Actual:%s", "123", string(p[:len(str)]))
	}
}
