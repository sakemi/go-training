package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "+1234.567890"
	fmt.Printf("%v -> %v\n", s, comma(s))
}

func comma(s string) string {
	var buf bytes.Buffer

	b := []byte(s)
	if sign := s[:1]; sign == "+" || sign == "-" {
		buf.WriteString(sign)
		b = b[1:]
	}

	var decimal []byte
	if d := bytes.Index(b, []byte(".")); d != -1 {
		decimal = b[d:]
		b = b[:d]
	}

	l := len(b)
	if l <= 3 {
		buf.Write(b)
		buf.Write(decimal)
		return buf.String()
	}
	head := l % 3
	if head != 0 {
		buf.Write(b[:head])
		buf.WriteByte(',')
	}
	for i := head; i < l; i += 3 {
		buf.Write(b[i : i+3])
		if i+3 != l {
			buf.WriteByte(',')
		}
	}
	buf.Write(decimal)
	return buf.String()
}
