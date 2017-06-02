package main

import (
	"fmt"
	"io"
	"os"
)

type Writer struct {
	wr    io.Writer
	count *int64
}

func main() {
	w, p := CountingWriter(os.Stdout)
	w.Write([]byte("hello world"))
	fmt.Println("")
	fmt.Println(*p)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	writer := Writer{w, new(int64)}
	return writer, writer.count
}

func (w Writer) Write(p []byte) (int, error) {
	b, err := w.wr.Write(p)
	*(w.count) += int64(b)
	return b, err
}
