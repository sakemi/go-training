package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/get", db.get)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		fmt.Fprintf(w, "%s already exists", item)
		return
	}
	price := req.URL.Query().Get("price")
	f, err := strconv.ParseFloat(price, 32)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	db[item] = dollars(f)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		fmt.Fprintf(w, "%s does not exist", item)
		return
	}
	price := req.URL.Query().Get("price")
	f, err := strconv.ParseFloat(price, 32)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	db[item] = dollars(f)
}

func (db database) get(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		fmt.Fprintf(w, "%s does not exist", item)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, db[item])
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		fmt.Fprintf(w, "%s does not exist", item)
		return
	}
	delete(db, item)
}
