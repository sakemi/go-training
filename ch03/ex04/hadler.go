package main

import (
	"net/http"
	"net/url"
)

const (
	width, hight, stroke, fill = "600", "320", "gray", "white"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	params := &svgParam{}
	query := r.URL.RawQuery
	m, _ := url.ParseQuery(query)

	initParam(params)

	if v, ok := m["hight"]; ok {
		params.hight = v[0]
	}
	if v, ok := m["width"]; ok {
		params.width = v[0]
	}
	if v, ok := m["stroke"]; ok {
		params.stroke = v[0]
	}
	if v, ok := m["fill"]; ok {
		params.fill = v[0]
	}

	writeSVG(w, params)
}

func initParam(p *svgParam) {
	p.hight = hight
	p.width = width
	p.stroke = stroke
	p.fill = fill
}
