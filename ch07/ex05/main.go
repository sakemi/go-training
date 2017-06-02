package main

import (
	"fmt"
	"io"
	"os"
)

type Reader struct {
	reader io.Reader
	limit  int64
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &Reader{r, n}
}

func (lr *Reader) Read(p []byte) (n int, err error) {
	if lr.limit < int64(len(p)) {
		p = p[:lr.limit]
	}
	n, err = lr.reader.Read(p)
	lr.limit -= int64(n)
	return
}

func main() {
	lr := LimitReader(os.Stdin, 3)
	p := make([]byte, 10)
	lr.Read(p)
	fmt.Println(string(p))
}
