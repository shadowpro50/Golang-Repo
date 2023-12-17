package main

import (
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
