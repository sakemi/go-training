package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type result struct {
	counts    int
	fileNames string
}

func main() {
	results := make(map[string]*result)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "Stdin", results)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, results)
			f.Close()
		}
	}
	for line, r := range results {
		if r.counts > 1 {
			fmt.Printf("%d\t%s\t%s\n", r.counts, line, r.fileNames)
		}
	}
}

func countLines(f *os.File, fileName string, results map[string]*result) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if results[input.Text()] == nil {
			results[input.Text()] = &result{}
		}
		results[input.Text()].counts++
		if !strings.Contains(results[input.Text()].fileNames, fileName) {
			results[input.Text()].fileNames += " " + fileName
		}
	}
}
