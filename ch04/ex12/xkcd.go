package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	latest    = 1835
	xkcd      = "https://xkcd.com"
	comicInfo = "info.0.json"
)

type info struct {
	Title string `json:"title"`
	URL   string
}

func main() {
	index := createIndex()
	var t string
	for {
		fmt.Printf("Input title:")
		fmt.Scan(&t)
		result := search(t, index)
		for _, v := range result {
			fmt.Printf("%s\t%s\n", v.Title, v.URL)
		}
	}
}

func createIndex() []info {
	fmt.Println("Creating index...")
	index := make([]info, latest)
	url := []string{xkcd, "", comicInfo}
	for i := 1; i < latest; i++ {
		url[1] = strconv.Itoa(i)
		u := strings.Join(url, "/")

		resp, err := http.Get(u)
		if err != nil {
			fmt.Printf("failed to GET: %s\n", u)
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("failed to GET: %s\t%s\n", u, resp.Status)
			continue
		}

		var result info
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			fmt.Printf("failed to decode json: %s\n", u)
			continue
		}
		result.URL = strings.Join(url[0:2], "/")
		index = append(index, result)
	}
	return index
}

func search(key string, index []info) []info {
	result := []info{}
	for _, v := range index {
		if strings.Contains(v.Title, key) {
			result = append(result, v)
		}
	}
	return result
}
