package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type movie struct {
	Poster string `json:"Poster"`
}

const omdb = "http://www.omdbapi.com/?t="

func main() {
	title := os.Args[1:]
	url, err := getPosterURL(title)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	if err := fetchPoster(url, strings.Join(title, "_")+".jpg"); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func getPosterURL(title []string) (string, error) {
	q := url.QueryEscape(strings.Join(title, "+"))
	url := omdb + q

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result movie
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Poster, nil
}

func fetchPoster(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write(data)
	return nil
}
