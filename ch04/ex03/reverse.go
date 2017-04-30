package main

import "fmt"

func main() {
	s := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("before: %v\n", s)
	reverse(&s)
	fmt.Printf("after: %v\n", s)
}

func reverse(s *[5]int) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}
