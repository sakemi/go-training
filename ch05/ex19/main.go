package main

import "fmt"

func main() {
	fmt.Println(panicRecover())
}

func panicRecover() (i interface{}) {
	defer func() {
		i = recover()
	}()
	panic(1)
}
