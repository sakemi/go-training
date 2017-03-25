package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sep := ":"
	for i, arg := range os.Args[:] {
		fmt.Println(strconv.Itoa(i) + sep + arg)
	}
}
