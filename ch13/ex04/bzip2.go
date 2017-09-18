package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	file, err := os.Create("out.bz2")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	w := NewWriter(file)
	b, err := w.Write(data)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	fmt.Printf("%d bytes", b)
}

type writer struct {
	w io.Writer
}

func NewWriter(out io.Writer) io.Writer {
	return writer{out}
}

func (w writer) Write(data []byte) (int, error) {
	cmd := exec.Command("bzip2")
	in, err := cmd.StdinPipe()
	if err != nil {
		return 0, err
	}
	defer in.Close()

	b, err := in.Write(data)
	if err != nil {
		return b, err
	}

	out, err := cmd.Output()
	if err != nil {
		return b, err
	}

	b, err = w.w.Write(out)
	if err != nil {
		return b, err
	}

	return b, nil
}
