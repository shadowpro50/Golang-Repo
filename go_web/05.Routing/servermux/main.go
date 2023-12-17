package main

import (
	"io"
	"net/http"
)

func dg(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "gau")
}

func ct(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "meow")
}

func main() {
	http.HandleFunc("/dog/", dg)
	http.HandleFunc("/cat", ct)
	http.ListenAndServe(":8080", nil)
}
