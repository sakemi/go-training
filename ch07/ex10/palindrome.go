package main

import (
	"bytes"
	"fmt"
	"sort"
)

type PalindromeCandidate []rune

func main() {
	str := "abcbba"
	p := bytes.Runes([]byte(str))
	b := IsPalindrome(PalindromeCandidate(p))
	fmt.Println(str, b)
}

func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len(); i++ {
		j := s.Len() - i - 1
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

func (p PalindromeCandidate) Len() int {
	return len(p)
}

func (p PalindromeCandidate) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p PalindromeCandidate) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
