package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

const (
	digit  = "digit"
	letter = "letter"
	space  = "space"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	counts, invalid := countCharType(in)

	fmt.Printf("type\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func countCharType(in *bufio.Reader) (map[string]int, int) {
	counts := make(map[string]int)
	invalid := 0

	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsDigit(r) {
			counts[digit]++
		}
		if unicode.IsLetter(r) {
			counts[letter]++
		}
		if unicode.IsSpace(r) {
			counts[space]++
		}
	}
	return counts, invalid
}
