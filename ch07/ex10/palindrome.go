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
	return sort.IsSorted(sort.Reverse(s))
}

func (p PalindromeCandidate) Len() int {
	return len(p)
}

func (p PalindromeCandidate) Less(i, j int) bool {
	//TODO
	return p[i] < p[j]
}

func (p PalindromeCandidate) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
