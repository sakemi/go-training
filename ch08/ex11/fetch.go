package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func fetch(url string, response chan *http.Response, cancel <-chan struct{}) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return
	}
	req.Cancel = cancel

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return
	}
	response <- resp
}

func main() {
	done := make(chan struct{})
	response := make(chan *http.Response)
	for _, url := range os.Args[1:] {
		go fetch(url, response, done)
	}
	resp := <-response
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading: %v\n", err)
		os.Exit(1)
	}
	close(done)
	fmt.Printf("%s", b)
}

// func cancelled() bool {
// 	select {
// 	case <-done:
// 		return true
// 	default:
// 		return false
// 	}
// }
