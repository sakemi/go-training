package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	mode := flag.String("m", "256", "hash mode: 256/384/512")
	flag.Parse()
	var in string
	fmt.Scan(&in)
	switch *mode {
	case "256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(in)))
	case "384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(in)))
	case "512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(in)))
	default:
		fmt.Printf("Undefined mode: %v\n", *mode)
	}
}
