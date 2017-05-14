package main

import "fmt"

type operation int

const (
	create operation = iota
	read
	update
)

func main() {
	for {
		fmt.Printf("Select Operation\n%d:create %d:read %d:update\n", create, read, update)
		var op operation
		_, err := fmt.Scan(&op)
		if err != nil {
			fmt.Println("Undefined Operation")
			continue
		}
		switch op {
		case create:
			if err := createIssue(); err != nil {
				fmt.Printf("%v\n", err)
				continue
			}
		case read:
			if err := readIssue(); err != nil {
				fmt.Printf("%v\n", err)
				continue
			}
		case update:
			if err := updateIssue(); err != nil {
				fmt.Printf("%v\n", err)
				continue
			}
		default:
			fmt.Println("Undefined Operation")
		}
	}
}
