package main

import (
	"net/http"
	"net/url"
)

const (
	x, y, scale = "0", "0", "1"
)

func handler(w http.ResponseWriter, r *http.Request) {

	params := &fractalParam{}
	query := r.URL.RawQuery
	m, _ := url.ParseQuery(query)

	initParam(params)

	if v, ok := m["x"]; ok {
		params.x = v[0]
	}
	if v, ok := m["y"]; ok {
		params.y = v[0]
	}
	if v, ok := m["scale"]; ok {
		params.scale = v[0]
	}

	renderMandelbrot(w, params)
}

func initParam(p *fractalParam) {
	p.x = x
	p.y = y
	p.scale = scale
}
