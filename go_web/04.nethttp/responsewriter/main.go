package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Zeds-Key", "this is from zeds")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>hello world</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
