package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "1234567890"
	fmt.Printf("%v -> %v\n", s, comma(s))
}

func comma(s string) string {
	var buf bytes.Buffer
	l := len(s)
	if l <= 3 {
		return s
	}
	head := l % 3
	if head != 0 {
		buf.WriteString(s[:head])
		buf.WriteByte(',')
	}
	for i := head; i < l; i += 3 {
		buf.WriteString(s[i : i+3])
		if i+3 != l {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}
