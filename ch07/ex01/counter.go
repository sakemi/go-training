package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int
type LineCounter int

func main() {
	hello := []byte("hello world\nhello world again")
	var wc WordCounter
	wc.Write(hello)
	fmt.Println(wc)

	var lc LineCounter
	lc.Write(hello)
	fmt.Println(lc)
}

func (c *WordCounter) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	l := 0
	for scanner.Scan() {
		*c++
		if err := scanner.Err(); err != nil {
			return l, err
		}
		l += len(scanner.Bytes())
	}
	return len(p), nil
}

func (c *LineCounter) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	l := 0
	for scanner.Scan() {
		*c++
		if err := scanner.Err(); err != nil {
			return l, err
		}
		l += len(scanner.Bytes())
	}
	return len(p), nil
}
