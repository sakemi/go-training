package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	const httpPrefix = "http://"
	for _, url := range os.Args[1:] {
		url = insertPrefixIfNeeded(url, httpPrefix)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func insertPrefixIfNeeded(s string, prefix string) string {
	if !(strings.HasPrefix(s, prefix)) {
		return prefix + s
	}
	return s
}
