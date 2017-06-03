package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type Result struct {
	Tracks []*Track
}

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.RawQuery
	m, _ := url.ParseQuery(q)
	sortKey := ""

	if v, ok := m["sort"]; ok {
		sortKey = v[0]
	}
	log.Print(sortKey)
	sortBy(sortKey)
	printTracks(tracks)
	result := Result{tracks}

	t := template.Must(template.ParseFiles("track.html.tpl"))
	if err := t.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}
