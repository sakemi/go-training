package main

import (
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	str := "123456789"
	limit := int64(10)
	buf := 1024
	sr := strings.NewReader(str)
	lr := LimitReader(sr, limit)

	p := make([]byte, buf)
	n, _ := lr.Read(p)
	if n != len(str) {
		t.Errorf("Read %d byte", n)
	}
	if str != string(p[:len(str)]) {
		t.Errorf("Expected:%s, Actual:%s", str, string(p[:len(str)]))
	}

	limit = 3
	sr = strings.NewReader(str)
	lr = LimitReader(sr, limit)
	p = make([]byte, buf)
	n, _ = lr.Read(p)
	if int64(n) != limit {
		t.Errorf("Read %d byte", n)
	}
	if "123" != string(p[:n]) {
		t.Errorf("Expected:%s, Actual:%s", "123", string(p[:n]))
	}
}
