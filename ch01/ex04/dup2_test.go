package main

import (
	"fmt"
	"os"
	"testing"
)

func TestCountLines(t *testing.T) {
	//setup
	results := make(map[string]*result)
	fileName1 := "friends.txt"
	f1, err := os.Open(fileName1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
		t.Fatal("Cannot Open " + fileName1)
	}
	defer f1.Close()
	fileName2 := "friends2.txt"
	f2, err := os.Open(fileName2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
		t.Fatal("Cannot Open " + fileName2)
	}
	defer f2.Close()

	//execute
	countLines(f1, fileName1, results)
	countLines(f2, fileName2, results)

	//test
	actual := results["サーバル"]
	expect1 := 5
	if actual.counts != expect1 {
		t.Error("サーバル", expect1, actual.counts)
	}
	expect2 := " " + fileName1 + " " + fileName2
	if actual.fileNames != expect2 {
		t.Error("サーバル", expect2, actual.fileNames)
	}
}
