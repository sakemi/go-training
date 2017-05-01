package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4}
	fmt.Printf("input: %v, rotate3: %v", s, rotate(s, 3))
}

func rotate(s []int, n int) []int {
	return append(s[n%len(s):], s[:n%len(s)]...)
}
