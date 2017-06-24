package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

type workingState struct {
	workList []string
	depth    int
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	var limit int
	flag.IntVar(&limit, "depth", 3, "depth to crawl.")
	flag.Parse()

	state := make(chan workingState)
	unseenLinks := make(chan workingState)

	go func() {
		state <- workingState{os.Args[1:], 1}
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				link := link
				foundLinks := crawl(link.workList[0])
				go func() {
					state <- workingState{foundLinks, link.depth + 1}
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	for s := range state {
		if s.depth > limit {
			continue
		}
		for _, link := range s.workList {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- workingState{[]string{link}, s.depth}
			}
		}
	}
}
