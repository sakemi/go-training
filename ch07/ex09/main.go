//ex08ができなかったのでstable sortできてない。とりあえず普通のソートするサーバーだけ
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/tracks", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
