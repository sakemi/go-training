package main

import (
	"fmt"
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
	q := r.URL.RawQuery
	m, _ := url.ParseQuery(q)
	input := ""

	if v, ok := m["expr"]; ok {
		input = v[0]
	}

	expr, err := Parse(input)
	if err != nil {
		fmt.Fprintln(w, err)
		log.Println("Failed to parse")
		return
	}
	env := Env{}
	for k, _ := range vars {
		if v, ok := m[k]; ok {
			f, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				fmt.Fprintln(w, err)
				log.Printf("Failed to parse var: %s = %s\n.", k, v[0])
				return
			}
			env[Var(k)] = f
		} else {
			fmt.Fprintln(w, err)
			log.Printf("Missing var: %s\n", k)
			return
		}
	}
	for k, v := range env {
		fmt.Fprintf(w, "%s = %v\n", k, v)
	}
	fmt.Fprintf(w, "%s = %v", input, expr.Eval(env))
}
