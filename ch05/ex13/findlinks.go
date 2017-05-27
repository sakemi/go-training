package main

import (
	"fmt"
	"io"
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

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

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
			fp, err := os.OpenFile(p+fileName, os.O_RDWR|os.O_CREATE, 0644)
			if err == nil {
				io.Copy(fp, resp.Body)
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

func main() {
	if err := os.Mkdir(result, 0777); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create folder to store result.")
	}
	breadthFirst(crawl, os.Args[1:])
}
