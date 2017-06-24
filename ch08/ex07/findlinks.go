package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"gopl.io/ch5/links"
)

const (
	result   = "pages"
	fileName = "out.html"
)

var origin string

func crawl(url string) []string {
	fmt.Println(url)

	s := strings.Index(url, "://") + len("://")
	dom := strings.Split(url[s:], "/")[0]
	if len(origin) == 0 {
		origin = dom
	}
	if dom == origin {
		skip := false
		resp, err := http.Get(url)
		if err != nil {
			skip = true
		}
		defer resp.Body.Close()

		var p string
		if !skip {
			p = strings.Join([]string{result, url[s:]}, "/")
			if err := os.MkdirAll(p, 0777); err != nil {
				skip = true
			}
		}

		if !skip {
			var name string
			if strings.HasSuffix(p, "/") {
				name = p + fileName
			} else {
				name = p + "/" + fileName
			}
			fp, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0644)
			if err == nil {
				b, err := ioutil.ReadAll(resp.Body)
				if err == nil {
					cur, _ := os.Getwd()
					path := strings.Join([]string{"href=\"", cur, "/", result, "/", dom}, "")
					replaced := bytes.Replace(b, []byte("href=\""), []byte(path), -1)
					replaced = bytes.Replace(replaced, []byte("/\">"), []byte("/"+fileName+"\">"), -1)
					io.Copy(fp, bytes.NewReader(replaced))
				} else {
					io.Copy(fp, resp.Body)
				}
				fp.Close()
			}
		}
	}

	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

type workingState struct {
	workList []string
	depth    int
}

func main() {
	if err := os.Mkdir(result, 0777); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create folder to store result.")
	}

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
