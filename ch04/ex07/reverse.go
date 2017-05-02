package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("abcde")
	fmt.Println(string(b))
	reverse(b)
	fmt.Println(string(b))
}

func reverse(b []byte) {
	if utf8.RuneCount(b) <= 1 {
		return
	}
	head, headLen := utf8.DecodeRune(b)
	tail, tailLen := utf8.DecodeLastRune(b)
	if headLen == tailLen {
		for i := 0; i < headLen; i++ {
			b[i], b[len(b)-tailLen+i] = b[len(b)-tailLen+i], b[i]
		}
	} else {
		h := []byte(string(head))
		t := []byte(string(tail))
		copy(b[tailLen:len(b)-headLen], b[headLen:len(b)-tailLen])
		copy(b[:tailLen], t)
		copy(b[len(b)-headLen:], h)
	}
	reverse(b[tailLen : len(b)-headLen])
}
