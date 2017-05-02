package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	s := "a b  c　d"
	fmt.Println(s)
	b := []byte(s)
	s = string(bundleSpace(b))
	fmt.Println(s)
}

func bundleSpace(b []byte) []byte {
	runes := []rune(bytes.Runes(b))
	isPrevSpace := false
	i := 0
	for _, r := range runes {
		if unicode.IsSpace(r) {
			if isPrevSpace {
				continue
			}
			isPrevSpace = true
			runes[i] = ' '
			i++
		} else {
			isPrevSpace = false
			runes[i] = r
			i++
		}
	}
	return []byte(string(runes[:i]))
}
