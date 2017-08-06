package main

import (
	"github.com/sakemi/go-training/ch10/ex02/archive"
	_ "github.com/sakemi/go-training/ch10/ex02/archive/tar"
	_ "github.com/sakemi/go-training/ch10/ex02/archive/zip"
)

func main() {
	z, err := archive.NewReader("test.zip")
	t, err := archive.NewReader("test.tar")
}
