package main

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	params := &lissajousParam{}
	initParam(params)

	query := r.URL.RawQuery
	m, _ := url.ParseQuery(query)

	if v, ok := m["cylces"]; ok {
		params.cycles, _ = strconv.ParseFloat(v[0], 64)
	}
	if v, ok := m["res"]; ok {
		params.res, _ = strconv.ParseFloat(v[0], 64)
	}
	if v, ok := m["size"]; ok {
		params.size, _ = strconv.Atoi(v[0])
	}
	if v, ok := m["cylnframesces"]; ok {
		params.nframes, _ = strconv.Atoi(v[0])
	}
	if v, ok := m["delay"]; ok {
		params.delay, _ = strconv.Atoi(v[0])
	}

	lissajous(w, params)
}

func initParam(p *lissajousParam) {
	p.cycles = 5
	p.res = 0.001
	p.size = 100
	p.nframes = 64
	p.delay = 8
}
