package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println(countHashDiff("x", "x"))
}

func countHashDiff(s1, s2 string) int {
	c1 := sha256.Sum256([]byte(s1))
	c2 := sha256.Sum256([]byte(s2))

	count := 0
	for i, v := range c1 {
		diff := v ^ c2[i]
		count += PopCount(diff)
	}

	return count
}
